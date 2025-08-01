package cmd

import (
	"fmt"
	"grepr/internal/banner"
	"grepr/internal/filter"
	"grepr/internal/soora"
	"grepr/internal/utils"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	inputFile   string
	fileTypes   string
	outputFile  string
	showBanner  bool
	enableSoora bool
	regexFile   string
	regexList   string
)

var rootCmd = &cobra.Command{
	Use:   "Grepr",
	Short: "Grepr - Lightweight URL filter tool for bug bounty hunters",
	Long:  `Grepr filters URLs by file types like .js, .php. Inspired by grep, built for automation.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		// Case 1: Filetype + Regex (combo filtering)
		if fileTypes != "" && (regexFile != "" || regexList != "") {
			types := strings.Split(fileTypes, ",")
			ftOutput := strings.TrimSuffix(outputFile, ".txt") + "-Grepr.txt"
			err := filter.ByFileType(inputFile, types, ftOutput)
			if err != nil {
				fmt.Printf("Error filtering by file type: %v\n", err)
				os.Exit(1)
			}

			patterns, err := filter.LoadRegexPatterns(regexFile, regexList)
			if err != nil {
				fmt.Printf("Failed to load regex patterns: %v\n", err)
				os.Exit(1)
			}

			finalOut := strings.TrimSuffix(outputFile, ".txt") + "-Grepr.txt"
			err = filter.MultiRegex(ftOutput, patterns, finalOut)
			if err != nil {
				fmt.Printf("Multi-regex filtering failed: %v\n", err)
				os.Exit(1)
			} else {
				lineCount, sizeKB, _ := utils.GetFileStats(finalOut)
				fmt.Printf("[✓] Filtered results (filetype + regex) → %s (%d lines, %.2f KB)\n", finalOut, lineCount, sizeKB)
			}
			return
		}

		// Case 2: Only Filetypes
		if fileTypes != "" {
			types := strings.Split(fileTypes, ",")
			finalOut := strings.TrimSuffix(outputFile, ".txt") + "-Grepr.txt"
			err := filter.ByFileType(inputFile, types, finalOut)
			if err != nil {
				fmt.Printf("Filetype filtering failed: %v\n", err)
				os.Exit(1)
			} else {
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

			finalOut := strings.TrimSuffix(outputFile, ".txt") + "-Grepr.txt"
			err = filter.MultiRegex(inputFile, patterns, finalOut)
			if err != nil {
				fmt.Printf("Regex filtering failed: %v\n", err)
				os.Exit(1)
			} else {
				lineCount, sizeKB, _ := utils.GetFileStats(finalOut)
				fmt.Printf("[✓] Regex filtered output → %s (%d lines, %.2f KB)\n", finalOut, lineCount, sizeKB)
			}
			return
		}

		// No valid flags used
		fmt.Println("[!] No filters applied. Use --filetypes, --regex-file, or --regex-list to filter data.")
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path (required)")
	rootCmd.Flags().StringVarP(&fileTypes, "filetypes", "f", "", "Filetypes to filter (comma-separated, e.g., js,php)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "output.txt", "Output file path")
	rootCmd.Flags().BoolVarP(&showBanner, "nobanner", "n", false, "Disable banner display")
	rootCmd.Flags().BoolVarP(&enableSoora, "soora", "s", false, "Enable Soora Super Mode (automated deep filtering)")
	rootCmd.Flags().StringVarP(&regexList, "regex-list", "r", "", "Comma-separated regex patterns (e.g., admin.*,login)")
	rootCmd.Flags().StringVar(&regexFile, "regex-file", "", "Path to file containing regex patterns")

	rootCmd.MarkFlagRequired("input")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
