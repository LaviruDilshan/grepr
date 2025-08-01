package filter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load keywords from both file and comma-list
func LoadKeywords(file string, list string) ([]string, error) {
	keywords := []string{}

	// From file
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				keywords = append(keywords, line)
			}
		}
	}

	// From comma-list
	if list != "" {
		split := strings.Split(list, ",")
		for _, word := range split {
			clean := strings.TrimSpace(word)
			if clean != "" {
				keywords = append(keywords, clean)
			}
		}
	}

	return keywords, nil
}

// Filter lines by checking if they contain any keyword
func ByKeywords(input string, keywords []string, output string) error {
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
		lower := strings.ToLower(line)

		for _, kw := range keywords {
			if strings.Contains(lower, strings.ToLower(kw)) {
				fmt.Printf("%v\n", line)
				writer.WriteString(line + "\n")
				break
			}
		}
	}

	return writer.Flush()
}
