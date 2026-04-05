package soora

import (
	"fmt"
	"grepr/internal/filter"
	"grepr/internal/template"
	"os"
	"path/filepath"
	"strings"
)

func Run(input string) error {
	var outputFiles []string
	finalOut := "Final-Grepr.txt"

	// --- Step 1: Original Core Filters ---
	fmt.Println("[i] Running Core Filters...")
	
	// JS filter
	jsOut := "All-Js-Grepr.txt"
	if err := filter.ByFileType(input, []string{"js"}, jsOut); err == nil {
		outputFiles = append(outputFiles, jsOut)
	}

	// TXT filter
	txtOut := "All-Text-Grepr.txt"
	if err := filter.ByFileType(input, []string{"txt"}, txtOut); err == nil {
		outputFiles = append(outputFiles, txtOut)
	}

	// --- Step 2: Special Config Filters (extensions.txt / regex.txt) ---
	
	// Load extensions
	extPath := getConfigPath("extensions.txt")
	if exts, err := filter.LoadLineFile(extPath); err == nil && len(exts) > 0 {
		extOut := "Special-Files-Grepr.txt"
		if err := filter.ByFileType(input, exts, extOut); err == nil {
			outputFiles = append(outputFiles, extOut)
		}
	}

	// Load regex patterns
	regexPath := getConfigPath("regex.txt")
	if regexList, err := filter.LoadLineFile(regexPath); err == nil && len(regexList) > 0 {
		regexOut := "Special-Regex-Grepr.txt"
		if err := filter.MultiRegex(input, regexList, regexOut); err == nil {
			outputFiles = append(outputFiles, regexOut)
		}
	}

	// --- Step 3: Dynamic Template Filters ---
	templates := template.List()
	if len(templates) > 0 {
		fmt.Printf("[i] Applying %d Security Templates...\n", len(templates))
		for _, tName := range templates {
			t, err := template.Load(tName)
			if err != nil {
				continue
			}

			outName := fmt.Sprintf("%s-Grepr.txt", strings.Title(tName))
			err = t.Apply(input, outName)
			if err != nil {
				continue
			}
			outputFiles = append(outputFiles, outName)
		}
	}

	// --- Step 4: Merge + Dedupe ---
	fmt.Println("[i] Merging and deduplicating all results...")
	err := filter.MergeAndDedupe(outputFiles, finalOut)
	if err != nil {
		return fmt.Errorf("merge failed: %v", err)
	}

	// Success Summary
	fmt.Println("\n[✓] Soora mode summary of generated files:")
	for _, file := range outputFiles {
		if _, err := os.Stat(file); err == nil {
			fmt.Printf("    → %s\n", file)
		}
	}
	fmt.Printf("\n[✓] Soora Super Mode Complete: %s generated.\n", finalOut)

	return nil
}

func getConfigPath(filename string) string {
	// Try local config first
	localConfig := filepath.Join("config", filename)
	if _, err := os.Stat(localConfig); err == nil {
		return localConfig
	}

	// Fallback to home config
	home, err := os.UserHomeDir()
	if err == nil {
		return filepath.Join(home, ".config", "grepr", filename)
	}
	return ""
}
