package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Title       string   `yaml:"title"`
	Author      string   `yaml:"author"`
	Tags        []string `yaml:"tags"`
	Category    []string `yaml:"category"`
	Created     string   `yaml:"created"`
	Published   string   `yaml:"published"`
	Updated     string   `yaml:"updated"`
	StarterRepo string   `yaml:"starter-repo"`
	Demo        string   `yaml:"demo"`
	Video       string   `yaml:"video"`
}

func extractMetadata(content string) (Metadata, string) {
	re := regexp.MustCompile(`(?s)---\n(.*?)\n---`)
	metadataContent := re.FindStringSubmatch(content)

	var meta Metadata
	if len(metadataContent) > 1 {
		yaml.Unmarshal([]byte(metadataContent[1]), &meta)
		content = re.ReplaceAllString(content, "")
	}

	return meta, content
}

func parseMarkdown(content string) string {
	// Split content into lines for easier processing
	lines := strings.Split(content, "\n")
	var parsedLines []string
	inCodeBlock := false

	for _, line := range lines {
		// Headers
		if strings.HasPrefix(line, "#") {
			count := strings.Count(line, "#")
			line = strings.TrimLeft(line, "# ")
			line = fmt.Sprintf("<h%d>%s</h%d>", count, line, count)
		}

		// Bold and Italic
		re := regexp.MustCompile(`\*\*(.*?)\*\*`)
		line = re.ReplaceAllString(line, "<strong>$1</strong>")
		re = regexp.MustCompile(`\*(.*?)\*`)
		line = re.ReplaceAllString(line, "<em>$1</em>")

		// Hyperlinks
		re = regexp.MustCompile(`\[([^\[]+)\]\(([^)]+)\)`)
		line = re.ReplaceAllString(line, `<a href="$2">$1</a>`)

		// Unordered Lists
		if strings.HasPrefix(line, "- ") {
			line = "<li>" + line[2:] + "</li>"
		}

		// Code Blocks
		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
			if inCodeBlock {
				parsedLines = append(parsedLines, "<pre><code>")
			} else {
				parsedLines = append(parsedLines, "</code></pre>")
			}
			continue
		}

		if inCodeBlock {
			parsedLines = append(parsedLines, line)
			continue
		}

		parsedLines = append(parsedLines, line)
	}

	return strings.Join(parsedLines, "\n")
}

func ParseMarkdownFile(filePath string) (Metadata, string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Metadata{}, "", err
	}

	meta, mdContent := extractMetadata(string(content))
	htmlContent := parseMarkdown(mdContent)

	return meta, htmlContent, nil
}
