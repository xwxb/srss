package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"strings"
)

func parseResp(resp *http.Response) ([]*NewsItem, error) {
	// 2. 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
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

			// fmt.Println( href)
			// 不知为何有时只出现网址的一部分
			
			if !strings.HasPrefix(href, "http") {
				return 
			}

			innerResp, _ := GetHtmlResp(href)
			contentDoc, _ := goquery.NewDocumentFromReader(innerResp.Body)
			content := purifyContent(contentDoc.Find("p").Text())
			fmt.Println(content)
			
			news = append(news, &NewsItem{Title: title, Content: content})
		})
	})

	return news, nil 
}