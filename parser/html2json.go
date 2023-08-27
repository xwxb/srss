package parser

import (
	"encoding/json"
	"log"
)

type NewsItem struct {
	Title   string
	Content string
}

func CrawlHTML(url string) string {

	resp, err := GetHtmlResp(url)
	if err != nil {
		log.Println(err)
	}

	news, err := parseResp(resp)
	if err != nil {
		log.Println(err)
	}

	// 返回JSON格式的字符串
	bytes, _ := json.Marshal(news)
	return string(bytes)
}
