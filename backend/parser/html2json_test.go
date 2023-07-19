package parser

import (
	"encoding/json"
	"testing"
)

func TestCrawlHTML(t *testing.T) {
  result := CrawlHTML(urls[0])
  t.Log(result)

  var resJsonObj NewsItem
  err := json.Unmarshal([]byte(result), &resJsonObj)

  if err != nil {
    t.Error(err)
  } else {
    t.Logf("%v", resJsonObj)
  }


}