package client

import (
	"testing"
	"net/http"
	"time"
)

type TestRequest struct {
	Format     string
	RequestURL string
}

func (b *TestRequest)Sign(cert *Credentials) {
}

func (b *TestRequest)HttpRequestInstance() (*http.Request, error) {
	httpReq, err := http.NewRequest("GET", b.RequestURL, nil)
	return httpReq, err
}

func (b *TestRequest)ResponseFormat() string {
	return b.Format
}

// A Timeout of zero means no timeout.
func (b *TestRequest)DeadLine() time.Duration {
	return 0
}
func (b *TestRequest)String() string {
	return ""
}

type youdao struct {
	errorCode int  `xml:"errorCode"`
	query     string  `xml:"query"`
}

func TestClient(t *testing.T) {
	end := make(chan int)
	client := NewClient(&Credentials{"214564", "46546"})
	client.SetDebug(true)
	go func() {
		resp := make(map[string]interface{})
		req := &TestRequest{Format:JSONResponseFormat, RequestURL:"http://www.weather.com.cn/data/cityinfo/101190408.html"}
		err := client.Query(req, &resp)
		if (nil != err) {
			t.Fatal("JSON", err)
		}
		end <- 1
	}()
	respxml := youdao{}
	req := &TestRequest{Format:XMLResponseFormat, RequestURL:"http://fanyi.youdao.com/openapi.do?keyfrom=cbping&key=1366735279&type=data&doctype=xml&version=1.1&q=%E8%A6%81%E7%BF%BB%E8%AF%91%E7%9A%84%E6%96%87%E6%9C%AC"}
	err := client.Query(req, &respxml)
	if (nil != err) {
		t.Fatal("XML", err)
	}
	<-end
}