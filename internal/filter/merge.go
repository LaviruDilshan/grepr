package filter

import (
	"bufio"
	"os"
)

func LoadLineFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

func MergeAndDedupe(inputs []string, output string) error {
	seen := make(map[string]bool)
	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	for _, file := range inputs {
		f, err := os.Open(file)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if !seen[line] {
				seen[line] = true
				writer.WriteString(line + "\n")
			}
		}
		f.Close()
	}
	return writer.Flush()
}
