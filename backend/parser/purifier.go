package parser

import (
	"regexp"
)

func purifyTitle(title string) string {
  re := regexp.MustCompile(`<[^>]*>`)
  return re.ReplaceAllString(title, "")
}

func purifyContent(content string) string {
    return content
}