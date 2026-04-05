package filter

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ByFileType(input string, types []string, output string) error {
	inFile, err := os.Open(input)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var writer *bufio.Writer
	if output != "" {
		outFile, err := os.Create(output)
		if err != nil {
			return err
		}
		defer outFile.Close()
		writer = bufio.NewWriter(outFile)
	}

	scanner := bufio.NewScanner(inFile)

	// Create regex like \.(js|php)($|\?|#)
	pattern := "\\.(" + strings.Join(types, "|") + ")($|\\?|#)"
	re := regexp.MustCompile("(?i)" + pattern)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			if writer != nil {
				writer.WriteString(line + "\n")
			} else {
				fmt.Println(line)
			}
		}
	}

	if writer != nil {
		return writer.Flush()
	}
	return nil
}
