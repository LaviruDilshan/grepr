package filter

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Filter lines by regex pattern
func ByRegex(input string, pattern string, output string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	inFile, err := os.Open(input)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			fmt.Printf("%v\n", line)
			writer.WriteString(line + "\n")
		}
	}

	return writer.Flush()
}

func MultiRegex(inputFile string, patterns []string, outputFile string) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	scanner := bufio.NewScanner(input)

	// Compile all regex patterns
	var regexes []*regexp.Regexp
	for _, p := range patterns {
		r, err := regexp.Compile(p)
		if err != nil {
			fmt.Printf("[!] Invalid regex: %s\n", p)
			continue
		}
		regexes = append(regexes, r)
	}

	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range regexes {
			if r.MatchString(line) {
				fmt.Printf("%v\n", line)
				writer.WriteString(line + "\n")
				break
			}
		}
	}

	return writer.Flush()
}
