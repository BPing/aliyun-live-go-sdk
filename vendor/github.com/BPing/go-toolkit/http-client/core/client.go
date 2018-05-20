package core

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func init() {
	SetDefaultClient("", http.DefaultClient)
}

//----------------------------------------------------------------------------------------------------------------------

//
//  客户端
//  处理http请求
type Client struct {
	// 采用默认&http.Client{}
	*http.Client

	// 存放钩子对象队列数组
	hookList []Hook

	// 用户
	userAgent string

	// 和http.Client.Timeout相关
	timeout time.Duration

	// 失败尝试最大次数
	// 默认2次
	maxBadRetryCount int

	// 版本号
	version string
	// debug
	debug bool

	// 上下文
	ctx Context
}

func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	return c
}

func (c *Client) SetVersion(version string) *Client {
	c.version = version
	return c
}

func (c *Client) SetUserAgent(userAgent string) *Client {
	c.userAgent = userAgent
	return c
}

func (c *Client) AppendHook(hook ...Hook) *Client {
	c.hookList = append(c.hookList, hook...)
	return c
}

func (c *Client) SetTimeOut(timeout time.Duration) *Client {
	c.Timeout = timeout
	return c
}

func (c *Client) SetContext(ctx Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) SetMaxBadRetryCount(retryCount int) *Client {
	if retryCount <= 0 {
		retryCount = 1
	}
	c.maxBadRetryCount = retryCount
	return c
}

// 设置代理
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
//  你也可以通过设置环境变量 HTTP_PROXY 来设置代理，如：
//      os.Setenv("HTTP_PROXY", "http://127.0.0.1:8888")
func (c *Client) SetProxy(proxy func(*http.Request) (*url.URL, error)) {
	//TODO::默认http.Client或者默认http.Transport时，是否值得改变代理（影响其他请求）？
	if nil != c.Client && nil != c.Client.Transport {
		c.Client.Transport.(*http.Transport).Proxy = proxy
	}

	return
}

// 处理请求
//
// 请求将有一定次数的失败重连机会。
// 默认为2次，可以通过SetMaxBadRetryCount()设置失败重连次数.
// 真实尝试的次数会记录在请求实体中。
//
// 记录请求处理时间。
//
func (c *Client) doRequest(req Request) (resp *Response, err error) {
	if nil == c.Client {
		c.Client = http.DefaultClient
	}
	t0 := time.Now()
	req.setReqCount(0)
	httpReq, err := req.HttpRequest()
	if nil != err {
		return nil, err
	}
	//必要头部信息设置
	httpReq.Header.Set("User-Agent", `Bping-Curl-`+c.userAgent+"/"+c.version)
	// 超时时间设置
	c.Client.Timeout = c.Timeout
	var httpResp *http.Response
	// 尝试次数记录
	reqCount := 0
	for ; reqCount < c.maxBadRetryCount; reqCount++ {
		httpResp, err = c.Client.Do(httpReq)
		if nil == err {
			break
		}
	}
	t1 := time.Now()
	req.setReqCount(reqCount)
	req.setReqLongTime(t1.Sub(t0))
	resp = &Response{Response: httpResp}
	req.setResponse(resp)
	return
}

// 请求开始处理之前的操作。
// 钩子将在此执行，其相应的方法会被执行。
func (c *Client) doBefore(req Request) (err error) {
	if err = c.doCtx(c.ctx); err != nil {
		return err
	}
	for _, hook := range c.hookList {
		err = hook.BeforeRequest(req, *c)
		if nil != err {
			break
		}
	}
	return
}

// 请求处理之后的操作。
// 钩子将在此执行，其相应的方法会被执行。
func (c *Client) doAfter(err error, req Request) {
	for _, hook := range c.hookList {
		hook.AfterRequest(err, req, *c)
	}
	return
}

// 处理Context
func (c *Client) doCtx(ctx Context) error {
	select {
	default:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// 处理请求
func (c *Client) DoRequest(req Request) (resp *Response, err error) {
	if err = c.doBefore(req); err != nil {
		return nil, err
	}
	defer func() {
		e := recover()
		if e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
		c.doAfter(err, req)
		if nil != err {
			err = clientError(err)
		}
		if e != nil {
			panic(e)
		}
	}()
	resp, err = c.doRequest(req)
	if nil != err {
		return nil, err
	}
	return
}

//type Setting struct {
//	Version          string
//	UserAgent        string
//	SlowReqLong      time.Duration
//	MaxBadRetryCount int
//}

// NewClient
func NewClient(title string, client *http.Client) *Client {
	return NewClientCtx(BackgroundContext(), title, client)
}

// NewClientCtx
func NewClientCtx(ctx Context, title string, client *http.Client) *Client {
	return &Client{
		Client:           client,
		version:          Version,
		userAgent:        title,
		debug:            false,
		maxBadRetryCount: defaultMaxBadRetryCount,
		ctx:              ctx,
	}
}

func clientError(err error) error {
	if nil == err {
		return nil
	}
	return errors.New("Bping-Http-Client-Failure:" + err.Error())
}

//----------------------------------------------------------------------------------------------------------------------

var DefaultClient *Client

// 设置DefaultClient
func SetDefaultClient(title string, client *http.Client) {
	DefaultClient = NewClient(title, client)
}

// 设置代理
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
// 内部调用DefaultClient
func SetProxy(proxy func(*http.Request) (*url.URL, error)) {
	DefaultClient.SetProxy(proxy)
}

func AppendHook(hook ...Hook) *Client {
	return DefaultClient.AppendHook(hook...)
}

// 设置超时时间
// 内部调用DefaultClient
func SetTimeOut(timeout time.Duration) *Client {
	return DefaultClient.SetTimeOut(timeout)
}

// 设置失败尝试次数
// 内部调用DefaultClient
func SetMaxBadRetryCount(retryCount int) *Client {
	return DefaultClient.SetMaxBadRetryCount(retryCount)
}

func SetVersion(version string) *Client {
	return DefaultClient.SetVersion(version)
}

func SetUserAgent(userAgent string) *Client {
	return DefaultClient.SetUserAgent(userAgent)
}

func SetContext(ctx Context) *Client {
	return DefaultClient.SetContext(ctx)
}

// 处理请求，内部调用DefaultClient
func DoRequest(req Request) (*Response, error) {
	return DefaultClient.DoRequest(req)
}
