package parser

import (
	"encoding/json"
	"log"
	"testing"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func TestCrawlHTML(t *testing.T) {

  // CrawlHTML(urls[0])

	// resString := CrawlHTML(urls[0])
  // log.Println(resString)

	// 依次遍历 url 的 1 ~ 3 项，将函数得到的 JSON 字符串反序列化成对象
	// 然后存储到 ../db/test.db 的 sqlite 数据库中的 feed 表中
	for i := 0; i < 3; i++ {
		resString := CrawlHTML(urls[i])
		// 反序列化

		var resJsonObj []*NewsItem
		err := json.Unmarshal([]byte(resString), &resJsonObj)

		if err != nil {
			t.Error(err)
		} else {
			// 把 resJsonObj 存储
			savaToDb(resJsonObj)
      log.Println(resJsonObj[0].Content)
		}
	}

}

func savaToDb(news []*NewsItem) {
  db, err := sql.Open("sqlite3", "../db/test.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // 创建表
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS feed (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, content TEXT)")
  if err != nil {
    log.Fatal(err)
  }

  // 插入数据
  for _, n := range news {
    _, err = db.Exec("INSERT INTO feed(title, content) VALUES(?, ?)", n.Title, n.Content)
    if err != nil {
      log.Fatal(err)
    } else {
      log.Println(n.Content)
    }
  }
}
