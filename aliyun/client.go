package aliyun

import (
	"errors"
	"github.com/BPing/go-toolkit/http-client/core"
	"log"
	"net"
	"net/http"
	"time"
)

type Client struct {
	*Credentials
	*core.Client
	//
	debug bool
}

func (c *Client) responseUnmarshal(req Request, respInfo *core.Response, resp interface{}) error {
	if req.ResponseFormat() == XMLResponseFormat {
		return respInfo.ToXML(resp)
	} else {
		return respInfo.ToJSON(resp)
	}
}

// 处理请求
func (c *Client) Query(req Request, resp interface{}) error {
	if nil == req {
		return clientError(errors.New("request is nil"))
	}
	req.Sign(c.Credentials)
	respInfo, err := c.DoRequest(req)
	if nil != err {
		return clientError(err)
	}
	if c.debug {
		log.Printf("http query %s %d (%v) ", req.String(), respInfo.StatusCode, req.ReqLongTime())
	}
	//失败响应处理
	if respInfo.StatusCode >= 400 && respInfo.StatusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = c.responseUnmarshal(req, respInfo, errorResponse)
		errorResponse.StatusCode = respInfo.StatusCode
		return &errorResponse
	}
	err = c.responseUnmarshal(req, respInfo, resp)
	if err != nil {
		return clientError(err)
	}
	return nil
}

func clientError(err error) error {
	if nil == err {
		return nil
	}
	return errors.New("AliyunLiveGoClientFailure:" + err.Error())
}

func NewClientTimeout(cert *Credentials, connectTimeout time.Duration) *Client {
	c := &Client{
		Credentials: cert,
		debug:       false,
	}
	if connectTimeout <= 0 {
		c.Client = core.NewClient("aliyun-client", &http.Client{})
	} else {
		c.Client = core.NewClient("aliyun-client", &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   connectTimeout,
					KeepAlive: 30 * time.Second,
				}).Dial,
				DisableKeepAlives: false,
			},
		})
	}
	return c
}

func NewClient(cert *Credentials) (c *Client) {
	return NewClientTimeout(cert, time.Duration(0))
}
