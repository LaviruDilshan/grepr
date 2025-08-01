package utils

import (
	"os"
	"strings"
)

func GetFileStats(filePath string) (lines int, sizeKB float64, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, 0, err
	}

	lines = strings.Count(string(data), "\n")
	sizeKB = float64(len(data)) / 1024.0
	return lines, sizeKB, nil
}
