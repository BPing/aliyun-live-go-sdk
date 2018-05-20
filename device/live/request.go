package live

import (
	"github.com/BPing/aliyun-live-go-sdk/aliyun"
)

// Request 直播请求信息
type Request struct {
	*aliyun.BaseRequest
	DomainName string
	AppName    string
}

func (l *Request) ToArgs() {
	l.Args.Set("DomainName", l.DomainName)
	if "" != l.AppName {
		l.Args.Set("AppName", l.AppName)
	}
}

func (l *Request) Sign(cert *aliyun.Credentials) {
	l.ToArgs()
	l.BaseRequest.Sign(cert)
}

func (l *Request) Clone() interface{} {
	new_obj := *l
	new_obj.BaseRequest = l.BaseRequest.Clone().(*aliyun.BaseRequest)
	return &new_obj
}

func NewLiveRequest(action, domainName, appname string) (l *Request) {
	l = &Request{
		BaseRequest: aliyun.NewBaseRequest(action),
		DomainName:  domainName,
		AppName:     appname,
	}
	l.Host = aliyun.APILiveHost
	l.Version = aliyun.APILiveVersion
	return
}
