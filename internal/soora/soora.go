package soora

import (
	"fmt"
	"grepr/internal/filter"
	"os"
	"path/filepath"
)

func Run(input string) error {
	// Output file names
	jsOut := "All-Js-Grepr.txt"
	txtOut := "All-Text-Grepr.txt"
	extOut := "Special-Files-Grepr.txt"
	regexOut := "Special-Regex-Grepr.txt"
	finalOut := "Final-Grepr.txt"

	// Step 1: JS filter
	filter.ByFileType(input, []string{"js"}, jsOut)

	// Step 2: TXT filter
	filter.ByFileType(input, []string{"txt"}, txtOut)

	// Step 3: Load extensions
	extPath, err := getConfigPath("extensions.txt")
	if err != nil {
		return fmt.Errorf("extensions path: %v", err)
	}
	exts, err := filter.LoadLineFile(extPath)
	if err != nil {
		return fmt.Errorf("extensions load: %v", err)
	}
	filter.ByFileType(input, exts, extOut)

	// Step 4: Load regex patterns
	regexPath, err := getConfigPath("regex.txt")
	if err != nil {
		return fmt.Errorf("regex path: %v", err)
	}
	regexList, err := filter.LoadLineFile(regexPath)
	if err != nil {
		return fmt.Errorf("regex load: %v", err)
	}
	filter.MultiRegex(input, regexList, regexOut)

	// Step 6: Merge + Dedupe
	allFiles := []string{jsOut, txtOut, extOut, regexOut}
	filter.MergeAndDedupe(allFiles, finalOut)

	// Output success messages
	fmt.Println("[✓] Soora mode complete: All-Js-Grepr.txt generated.")
	fmt.Println("[✓] Soora mode complete: All-Text-Grepr.txt generated.")
	fmt.Println("[✓] Soora mode complete: Special-Files-Grepr.txt generated.")
	fmt.Println("[✓] Soora mode complete: Special-Regex-Grepr.txt generated.")
	fmt.Println("[✓] Soora mode complete: Final-Grepr.txt generated.")

	return nil
}

func getConfigPath(filename string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "grepr", filename), nil
}
