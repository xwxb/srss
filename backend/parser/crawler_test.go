package parser

import (
	"testing"
	"io/ioutil"
)

func TestGetHtmlResp(t *testing.T) {
	resp, err := GetHtmlResp(urls[0])

	if err != nil {
		t.Error(err)
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error("读取响应失败:", err)
			return
		}

		t.Log(string(body))
	}

}
