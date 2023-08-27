package parser

import (
	"log"
	"net/http"
)

func GetHtmlResp(url string) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加常见的请求头字段
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	// req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()

	return resp, nil
}
