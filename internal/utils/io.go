package utils

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
