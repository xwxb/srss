package parser

import (
	"encoding/json"
	"testing"
)

func TestCrawlHTML(t *testing.T) {

  // 正常用例
  // urls := [
  //   "http://clxy.bjfu.edu.cn/lmdh/xydt/index.html",
  //   "http://it.bjfu.edu.cn/xyxw/index.html",
  //   "http://jwc.bjfu.edu.cn/ksxx/index.html"
  // ]

  urls := []string{
    "http://clxy.bjfu.edu.cn/lmdh/xydt/index.html",
    "http://it.bjfu.edu.cn/xyxw/index.html",
    "http://jwc.bjfu.edu.cn/ksxx/index.html",
  }

  result := CrawlHTML(urls[0])
  t.Log(result)

  var resJsonObj NewsItem
  err := json.Unmarshal([]byte(result), &resJsonObj)

  // 检查结果长度
  if err != nil {
    t.Error(err)
  } else {
    t.Logf("%v", resJsonObj)
  }


}