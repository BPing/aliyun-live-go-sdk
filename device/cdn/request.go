package cdn

import "github.com/BPing/aliyun-live-go-sdk/aliyun"

//  cdn 请求对象。实现 Request 接口
type Request struct {
	*aliyun.BaseRequest
}

func (c *Request) Clone() interface{} {
	new_obj := *c
	new_obj.BaseRequest = c.BaseRequest.Clone().(*aliyun.BaseRequest)
	return &new_obj
}

// 生成CDNRequest
func NewCDNRequest(action string) *Request {
	return &Request{
		BaseRequest: aliyun.NewBaseRequest(action),
	}
}
