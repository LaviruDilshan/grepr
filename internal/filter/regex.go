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
		// fmt.Printf("[DEBUG] Compiling pattern: '%s'\n", p)
		r, err := regexp.Compile(p)
		if err != nil {
			fmt.Printf("[!] Skipping invalid regex: %s\n", p)
			continue
		}
		regexes = append(regexes, r)
	}

	for scanner.Scan() {
		line := scanner.Text()
		matched := false
		for _, r := range regexes {
			if r.MatchString(line) {
				// fmt.Printf("[MATCH] Line: %s | Regex: %s\n", line, r.String())
				writer.WriteString(line + "\n")
				matched = true
				break
			}
		}
		if !matched {
			// fmt.Printf("[NO MATCH] Line: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	return writer.Flush()
}

// LoadRegexPatterns combines regexes from file and comma-separated list
func LoadRegexPatterns(filePath, list string) ([]string, error) {
	var patterns []string

	// Load from file
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open regex file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				patterns = append(patterns, line)
			}
		}
	}

	// Load from comma-separated list
	if list != "" {
		split := regexp.MustCompile(`\s*,\s*`).Split(list, -1)
		for _, item := range split {
			if item != "" {
				patterns = append(patterns, item)
			}
		}
	}

	return patterns, nil
}
