package parser

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type NewsItem struct {
	Title   string
	Content string
}

func CrawlHTML(url string) string {
	// 1. 获取HTML页面
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	// 2. 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ""
	}

	news := make([]*NewsItem, 0)

	doc.Find("ul").Each(func(i int, s *goquery.Selection) {
		hasSpan := s.Find("li span").Length() > 0
		if !hasSpan {
			return
		}

		s.Find("li").Each(func(j int, li *goquery.Selection) {

			// 标题
			title := purifyTitle(li.Find("a").Text())

			// 链接和内容
			href, _ := li.Find("a").Attr("href")
			resp, _ := http.Get(href)
			contentDoc, _ := goquery.NewDocumentFromReader(resp.Body)
			content := purifyContent(contentDoc.Find(".content").Text())

			news = append(news, &NewsItem{Title: title, Content: content})
		})
	})

	// 4. 返回JSON格式的字符串
	bytes, _ := json.Marshal(news)
	return string(bytes)
}
