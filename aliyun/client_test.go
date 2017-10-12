package aliyun

import (
	"net/http"
	"testing"
	"github.com/BPing/go-toolkit/http-client/hook"
	"time"
	"fmt"
)

type TestRequest struct {
	BaseRequest
	Format     string
	RequestURL string
}

func (b *TestRequest) ResponseFormat() string {
	return b.Format
}

func (b *TestRequest) Sign(cert *Credentials) {
}

func (b *TestRequest) HttpRequest() (*http.Request, error) {
	httpReq, err := http.NewRequest("GET", b.RequestURL, nil)
	return httpReq, err
}

type youdao struct {
	ErrorCode int    `xml:"errorCode"`
	Query     string `xml:"query"`
}

func TestClient(t *testing.T) {
	end := make(chan int)
	cert := NewCredentials("214564", "46546")
	client := NewClient(cert)
	client.SetDebug(true)
	client.AppendHook(hook.NewLogHook(time.Second*3, func(tag, msg string) {
		fmt.Println(tag, msg)
	}))
	go func() {
		resp := make(map[string]interface{})
		req := &TestRequest{Format: JSONResponseFormat, RequestURL: "http://www.weather.com.cn/data/cityinfo/101190408.html"}
		err := client.Query(req, &resp)
		if nil != err {
			t.Fatal("JSON", err)
		}
		end <- 1
	}()

	respxml := youdao{}
	req := &TestRequest{Format: XMLResponseFormat, RequestURL: "http://fanyi.youdao.com/openapi.do?keyfrom=cbping&key=1366735279&type=data&doctype=xml&version=1.1&q=%E8%A6%81%E7%BF%BB%E8%AF%91%E7%9A%84%E6%96%87%E6%9C%AC"}
	err := client.Query(req, &respxml)
	if nil != err {
		t.Fatal("XML", err)
	}
	<-end
}
