package template

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/LaviruDilshan/grepr/v2/internal/filter"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates
var defaultTemplates embed.FS

// Template represents a search profile like gf
type Template struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	FileTypes   []string `json:"filetypes"`
	Patterns    []string `json:"patterns"`
}

// Load loads a template by name from user directory or embedded defaults
func Load(templateName string) (*Template, error) {
	// 1. Search in ~/.config/grepr/templates/ (User override)
	home, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(home, ".config", "grepr", "templates", templateName+".json")
		if _, err := os.Stat(configPath); err == nil {
			return readTemplateFromFile(configPath)
		}
	}

	// 2. Search in embedded templates
	embeddedPath := "templates/" + templateName + ".json"
	data, err := defaultTemplates.ReadFile(embeddedPath)
	if err == nil {
		return parseTemplate(data)
	}

	return nil, fmt.Errorf("template '%s' not found in embedded defaults or ~/.config/grepr/templates/", templateName)
}

func readTemplateFromFile(path string) (*Template, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseTemplate(data)
}

func parseTemplate(data []byte) (*Template, error) {
	var t Template
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %v", err)
	}
	return &t, nil
}

// Apply applies a template to an input file
func (t *Template) Apply(input string, output string) error {
	var intermediateOut string
	
	// Step 1: Filter by file types if specified
	if len(t.FileTypes) > 0 {
		intermediateOut = "temp-Grepr-template-filetype.txt"
		err := filter.ByFileType(input, t.FileTypes, intermediateOut)
		if err != nil {
			return fmt.Errorf("filetype filter error: %v", err)
		}
	} else {
		intermediateOut = input
	}

	// Step 2: Filter by regex patterns
	if len(t.Patterns) > 0 {
		err := filter.MultiRegex(intermediateOut, t.Patterns, output)
		if intermediateOut != input {
			os.Remove(intermediateOut)
		}
		if err != nil {
			return fmt.Errorf("regex filter error: %v", err)
		}
	} else {
		if intermediateOut != output && output != "" {
			err := os.Rename(intermediateOut, output)
			if err != nil {
				return fmt.Errorf("failed to move results to %s: %v", output, err)
			}
		}
	}

	return nil
}

// List available templates
func List() []string {
	templateMap := make(map[string]bool)
	var templates []string
	
	// 1. Scan embedded templates
	entries, _ := fs.ReadDir(defaultTemplates, "templates")
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".json") {
			name := strings.TrimSuffix(entry.Name(), ".json")
			templateMap[name] = true
		}
	}
	
	// 2. Scan ~/.config/grepr/templates/
	home, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(home, ".config", "grepr", "templates")
		files, _ := os.ReadDir(configPath)
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				name := strings.TrimSuffix(f.Name(), ".json")
				templateMap[name] = true
			}
		}
	}

	for name := range templateMap {
		templates = append(templates, name)
	}
	
	return templates
}
