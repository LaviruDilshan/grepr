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

		if fileTypes != "" {
			types := strings.Split(fileTypes, ",")
			ftOutput := strings.TrimSuffix(outputFile, ".txt") + "-filetypes-Grepr.txt"

			err := filter.ByFileType(inputFile, types, ftOutput)
			if err != nil {
				fmt.Printf("Error filtering by file type: %v\n", err)
				os.Exit(1)
			} else {
				lineCount, sizeKB, err := utils.GetFileStats(ftOutput)
				if err != nil {
					fmt.Printf("Error getting file stats: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("[✓] Filetype filtered results written to: %s (%d lines, %.2f KB)\n", ftOutput, lineCount, sizeKB)
			}
		}

		// New multiple-regex filtering (via file + list)
		if regexFile != "" || regexList != "" {
			patterns, err := filter.LoadRegexPatterns(regexFile, regexList)
			if err != nil {
				fmt.Printf("Failed to load regex patterns: %v\n", err)
				os.Exit(1)
			}

			out := strings.TrimSuffix(outputFile, ".txt") + "-regexes-Grepr.txt"
			err = filter.MultiRegex(inputFile, patterns, out)
			if err != nil {
				fmt.Printf("Multi-regex filtering failed: %v\n", err)
				os.Exit(1)
			} else {
				lineCount, sizeKB, err := utils.GetFileStats(out)
				if err != nil {
					fmt.Printf("Error getting file stats: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("[✓] Multi-regex filtered lines written to: %s (%d lines, %.2f KB)\n", out, lineCount, sizeKB)
			}
		}

		if fileTypes == "" && regexFile == "" && regexList == "" {
			fmt.Println("[!] No filters applied. Use --filetypes, --keywords-file, or --keywords to filter data.")
		}
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
