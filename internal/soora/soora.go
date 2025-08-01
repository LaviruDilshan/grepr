package soora

import (
	"fmt"
	"grepr/internal/filter"
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
	exts, err := filter.LoadLineFile("config/extensions.txt")
	if err != nil {
		return fmt.Errorf("extensions: %v", err)
	}
	filter.ByFileType(input, exts, extOut)

	// Step 4: Regex list
	regexList, err := filter.LoadLineFile("config/regex.txt")
	if err != nil {
		return fmt.Errorf("regex: %v", err)
	}
	filter.MultiRegex(input, regexList, regexOut)

	// Step 6: Merge + Dedupe
	allFiles := []string{jsOut, txtOut, extOut, regexOut}
	filter.MergeAndDedupe(allFiles, finalOut)

	fmt.Println("[✓] Soora mode complete: Final-Grepr.txt generated.")
	return nil
}
