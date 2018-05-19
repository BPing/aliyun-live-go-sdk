package core

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// 请求接口
type Request interface {
	//返回*http.Request
	HttpRequest() (*http.Request, error)

	// 返回请求服务名
	// 以便归类不同服务的统计数据
	// 如果使用断路器，则必须根据不同服务端返回不同名字。
	// 以 主机+端口 是一个不错的选择
	ServerName() string

	//
	Clone() interface{}
	String() string

	// 尝试次数
	setReqCount(reqCount int)
	ReqCount() (int)

	// 钩子存放获取数据
	SetHookData(key string, data interface{}) (ok bool)
	HookData(key string) (data interface{}, ok bool)

	// 请求响应实体处理
	setResponse(resp *Response)
	Response() (*Response)

	// 请求响应时间处理
	setReqLongTime(long time.Duration)
	ReqLongTime() time.Duration
}

// 请求基类
//    实现请求接口的基类，所有请求对象继承必须继承此基类
type BaseRequest struct {
	reqCount    int
	reqLongTime time.Duration
	Resp        *Response

	// 钩子存放数据Map
	hookData map[string]interface{}
}

//返回*http.Request
func (b *BaseRequest) HttpRequest() (*http.Request, error) {
	return nil, errors.New("implement Interface's Method::HttpRequest")
}

func (b *BaseRequest) ServerName() string {
	return "localhost"
}

func (b *BaseRequest) String() string {
	return fmt.Sprintf("\n ReqCount:%d \n", b.reqCount)
}

func (b *BaseRequest) TimeOut() time.Duration {
	return -1
}

func (b *BaseRequest) Clone() interface{} {
	new_obj := *b
	return &new_obj
}

func (b *BaseRequest) setReqCount(reqCount int) {
	b.reqCount = reqCount
}

func (b *BaseRequest) ReqCount() int {
	return b.reqCount
}

// 设置请求处理时间
func (b *BaseRequest) setReqLongTime(long time.Duration) {
	b.reqLongTime = long
}

func (b *BaseRequest) ReqLongTime() time.Duration {
	return b.reqLongTime
}

// 设置请求返回信息结构
func (b *BaseRequest) setResponse(resp *Response) {
	b.Resp = resp
}

func (b *BaseRequest) Response() *Response {
	return b.Resp
}

func (b *BaseRequest) SetHookData(key string, data interface{}) (ok bool) {
	if nil == b.hookData {
		b.hookData = make(map[string]interface{})
	}
	b.hookData[key] = data
	_, ok = b.hookData[key]
	return
}

func (b *BaseRequest) HookData(key string) (data interface{}, ok bool) {
	if nil == b.hookData {
		return nil, false
	}
	data, ok = b.hookData[key]
	return
}
