package cmd

import (
	"fmt"
	"os"
	"strings"

	"grepr/internal/banner"
	"grepr/internal/filter"
	"grepr/internal/soora"

	"github.com/spf13/cobra"
)

var (
	inputFile   string
	fileTypes   string
	outputFile  string
	showBanner  bool
	keywordFile string
	keywordList string
	regexFilter string
	enableSoora bool
)

var rootCmd = &cobra.Command{
	Use:   "Grepr",
	Short: "Grepr - Lightweight URL filter tool for bug bounty hunters",
	Long:  `Grepr filters URLs by file types like .js, .php. Inspired by grep, built for automation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !showBanner {
			banner.Print()
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
			err := filter.ByFileType(inputFile, types, outputFile)
			if err != nil {
				fmt.Printf("Error filtering by file type: %v\n", err)
				os.Exit(1)
			} else {
				fmt.Printf("[✓] Filetype filtered results written to: %s\n", outputFile)
			}
		}

		if keywordFile != "" || keywordList != "" {
			keywords, err := filter.LoadKeywords(keywordFile, keywordList)
			if err != nil {
				fmt.Printf("Keyword loading error: %v\n", err)
				os.Exit(1)
			}

			kwOutput := strings.TrimSuffix(outputFile, ".txt") + "-keywords-Grepr.txt"
			err = filter.ByKeywords(inputFile, keywords, kwOutput)
			if err != nil {
				fmt.Printf("Keyword filtering error: %v\n", err)
				os.Exit(1)
			} else {
				fmt.Printf("[✓] Keyword filtered results written to: %s\n", kwOutput)
			}
		}

		if regexFilter != "" {
			regexOutput := strings.TrimSuffix(outputFile, ".txt") + "-regex-Grepr.txt"
			err := filter.ByRegex(inputFile, regexFilter, regexOutput)
			if err != nil {
				fmt.Printf("Regex filtering error: %v\n", err)
				os.Exit(1)
			} else {
				fmt.Printf("[✓] Regex filtered results saved to: %s\n", regexOutput)
			}
		}

		if fileTypes == "" && keywordFile == "" && keywordList == "" {
			fmt.Println("[!] No filters applied. Use --filetypes, --keywords-file, or --keywords to filter data.")
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path (required)")
	rootCmd.Flags().StringVar(&fileTypes, "filetypes", "", "Filetypes to filter (comma-separated, e.g., js,php)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "output.txt", "Output file path")
	rootCmd.Flags().BoolVarP(&showBanner, "nobanner", "n", false, "Disable banner display")
	rootCmd.Flags().StringVar(&keywordFile, "keywords-file", "", "Keyword file (e.g., sensitive.txt)")
	rootCmd.Flags().StringVar(&keywordList, "keywords", "", "Comma-separated keywords (e.g., admin,dev,secret)")
	rootCmd.Flags().StringVarP(&regexFilter, "regex", "r", "", "Filter lines matching this regex pattern")
	rootCmd.Flags().BoolVarP(&enableSoora, "soora", "s", false, "Enable Soora Super Mode (automated deep filtering)")

	rootCmd.MarkFlagRequired("input")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
