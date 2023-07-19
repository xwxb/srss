package parser

import (
	// "encoding/json"
	"testing"
)

func TestCrawlHTML(t *testing.T) {
  
  resString := CrawlHTML(urls[2])
  t.Log(resString)

  // var resJsonObj NewsItem
  // err := json.Unmarshal([]byte(result), &resJsonObj)

  // if err != nil {
  //   t.Error(err)
  // } else {
  //   t.Logf("%v", resJsonObj)
  // }


}