package cmd

import (
	"fmt"
	"github.com/LaviruDilshan/grepr/v2/internal/banner"
	"github.com/LaviruDilshan/grepr/v2/internal/filter"
	"github.com/LaviruDilshan/grepr/v2/internal/soora"
	"github.com/LaviruDilshan/grepr/v2/internal/template"
	"github.com/LaviruDilshan/grepr/v2/internal/utils"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	inputFile     string
	fileTypes     string
	outputFile    string
	showBanner    bool
	enableSoora   bool
	regexFile     string
	regexList     string
	templateName  string
	listTemplates bool
)

var rootCmd = &cobra.Command{
	Use:   "Grepr",
	Short: "Grepr - Lightweight URL filter tool for bug bounty hunters",
	Long:  `Grepr v2.0.0 filters URLs by file types like .js, .php. Inspired by grep, built for automation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if listTemplates {
			templates := template.List()
			fmt.Println("Available Templates:")
			for _, t := range templates {
				fmt.Printf("- %s\n", t)
			}
			return
		}

		if !showBanner {
			banner.Print()
			time.Sleep(2 * time.Second)
		}

		if enableSoora {
			err := soora.Run(inputFile)
			if err != nil {
				fmt.Printf("Soora mode error: %v\n", err)
				os.Exit(1)
			}
			return
		}

		var finalOut string
		if outputFile != "" {
			finalOut = strings.TrimSuffix(outputFile, ".txt") + "-Grepr.txt"
		}

		// Template Mode
		if templateName != "" {
			t, err := template.Load(templateName)
			if err != nil {
				fmt.Printf("Error loading template: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("[i] Applying template: %s (%s)\n", t.Name, t.Description)
			err = t.Apply(inputFile, finalOut)
			if err != nil {
				fmt.Printf("Error applying template: %v\n", err)
				os.Exit(1)
			} else if finalOut != "" {
				lineCount, sizeKB, _ := utils.GetFileStats(finalOut)
				fmt.Printf("[✓] Template filtered output → %s (%d lines, %.2f KB)\n", finalOut, lineCount, sizeKB)
			}
			return
		}

		// Case 1: Filetype + Regex (combo filtering)
		if fileTypes != "" && (regexFile != "" || regexList != "") {
			types := strings.Split(fileTypes, ",")
			
			// For combo, we need an intermediate result even if printing to console.
			intermediateOut := "temp-Grepr-intermediate.txt"
			err := filter.ByFileType(inputFile, types, intermediateOut)
			if err != nil {
				fmt.Printf("Error filtering by file type: %v\n", err)
				os.Exit(1)
			}

			patterns, err := filter.LoadRegexPatterns(regexFile, regexList)
			if err != nil {
				fmt.Printf("Failed to load regex patterns: %v\n", err)
				os.Remove(intermediateOut)
				os.Exit(1)
			}

			err = filter.MultiRegex(intermediateOut, patterns, finalOut)
			os.Remove(intermediateOut)

			if err != nil {
				fmt.Printf("Multi-regex filtering failed: %v\n", err)
				os.Exit(1)
			} else if finalOut != "" {
				lineCount, sizeKB, _ := utils.GetFileStats(finalOut)
				fmt.Printf("[✓] Filtered results (filetype + regex) → %s (%d lines, %.2f KB)\n", finalOut, lineCount, sizeKB)
			}
			return
		}

		// Case 2: Only Filetypes
		if fileTypes != "" {
			types := strings.Split(fileTypes, ",")
			err := filter.ByFileType(inputFile, types, finalOut)
			if err != nil {
				fmt.Printf("Filetype filtering failed: %v\n", err)
				os.Exit(1)
			} else if finalOut != "" {
				lineCount, sizeKB, _ := utils.GetFileStats(finalOut)
				fmt.Printf("[✓] Filetype filtered output → %s (%d lines, %.2f KB)\n", finalOut, lineCount, sizeKB)
			}
			return
		}

		// Case 3: Only Regexes
		if regexFile != "" || regexList != "" {
			patterns, err := filter.LoadRegexPatterns(regexFile, regexList)
			if err != nil {
				fmt.Printf("Failed to load regex patterns: %v\n", err)
				os.Exit(1)
			}

			err = filter.MultiRegex(inputFile, patterns, finalOut)
			if err != nil {
				fmt.Printf("Regex filtering failed: %v\n", err)
				os.Exit(1)
			} else if finalOut != "" {
				lineCount, sizeKB, _ := utils.GetFileStats(finalOut)
				fmt.Printf("[✓] Regex filtered output → %s (%d lines, %.2f KB)\n", finalOut, lineCount, sizeKB)
			}
			return
		}

		// No valid flags used
		fmt.Println("[!] No filters applied. Use --filetypes, --regex-file, --regex-list, or --template to filter data.")
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path (required)")
	rootCmd.Flags().StringVarP(&fileTypes, "filetypes", "f", "", "Filetypes to filter (comma-separated, e.g., js,php)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (optional, prints to console if omitted)")
	rootCmd.Flags().BoolVarP(&showBanner, "nobanner", "n", false, "Disable banner display")
	rootCmd.Flags().BoolVarP(&enableSoora, "soora", "s", false, "Enable Soora Super Mode (automated deep filtering)")
	rootCmd.Flags().StringVarP(&regexList, "regex-list", "r", "", "Comma-separated regex patterns (e.g., admin.*,login)")
	rootCmd.Flags().StringVar(&regexFile, "regex-file", "", "Path to file containing regex patterns")
	rootCmd.Flags().StringVarP(&templateName, "template", "t", "", "Apply a predefined template (e.g., sqli, xss)")
	rootCmd.Flags().BoolVarP(&listTemplates, "list-templates", "l", false, "List all available templates")

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}
