package aliyun

import (
	"errors"
	"github.com/BPing/go-toolkit/http-client/core"
	"log"
	"net"
	"net/http"
	"time"
)

// Client 统一处理器。
//        底层使用core包的client来处理请求内容。
//   Credentials API请求凭证
//   Client      core.Client
//   debug       是否开启调试
type Client struct {
	*Credentials
	*core.Client
	debug bool
}

func (c *Client) responseUnmarshal(req Request, respInfo *core.Response, resp interface{}) error {
	if req.ResponseFormat() == XMLResponseFormat {
		return respInfo.ToXML(resp)
	}
	return respInfo.ToJSON(resp)
}

func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	c.Client.SetDebug(debug)
	return c
}

// Query 处理请求
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
		errorResponse := &ErrorResponse{}
		err = c.responseUnmarshal(req, respInfo, errorResponse)
		errorResponse.StatusCode = respInfo.StatusCode
		return errorResponse
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

// NewClientTimeout 新建client实例，可设置超时时间
func NewClientTimeoutCtx(ctx core.Context, cert *Credentials, connectTimeout time.Duration) *Client {
	c := &Client{
		Credentials: cert,
	}
	if connectTimeout <= 0 {
		c.Client = core.NewClientCtx(ctx, "aliyun-client", &http.Client{})
	} else {
		c.Client = core.NewClientCtx(ctx, "aliyun-client", &http.Client{
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
	c.SetDebug(false)
	return c
}

// NewClient 新建client实例。
func NewClientCtx(ctx core.Context, cert *Credentials) (c *Client) {
	return NewClientTimeoutCtx(ctx, cert, time.Duration(0))
}
