package parser

import (
	"regexp"
	"strings"
)

func purifyContent(content string) string {
	// 去除乱码字符
	content = removeInvalidCharacters(content)

	// 去除HTML标签
	content = removeHTMLTags(content)

	// 去除空白字符
	content = removeWhitespace(content)

	return content
}

// 去除乱码字符
func removeInvalidCharacters(content string) string {
	// 使用正则表达式匹配无效字符，并替换为空字符串
	re := regexp.MustCompile(`[^\x00-\x7F]+`)
	content = re.ReplaceAllString(content, "")

	return content
}

// 去除HTML标签
func removeHTMLTags(content string) string {
	// 使用正则表达式匹配HTML标签，并替换为空字符串
	re := regexp.MustCompile(`<[^>]+>`)
	content = re.ReplaceAllString(content, "")

	return content
}

// 去除空白字符
func removeWhitespace(content string) string {
	// 使用strings.TrimSpace()函数去除字符串两端的空白字符
	content = strings.TrimSpace(content)

	return content
}

func purifyTitle(title string) string {
  re := regexp.MustCompile(`<[^>]*>`)
  return re.ReplaceAllString(title, "")
}
