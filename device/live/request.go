package live

import (
	"github.com/BPing/aliyun-live-go-sdk/client"
)

// LiveRequest 直播请求信息
type LiveRequest struct {
	*client.CDNRequest
	DomainName string
	AppName    string
}

func (l *LiveRequest) StructToArgs() {
	l.Args.Set("DomainName", l.DomainName)
	if "" != l.AppName {
		l.Args.Set("AppName", l.AppName)
	}
}

func (l *LiveRequest) Sign(cert *client.Credentials) {
	l.StructToArgs()
	l.CDNRequest.Sign(cert)
}

func (l *LiveRequest) Clone() interface{} {
	new_obj := (*l)
	new_obj.CDNRequest = l.CDNRequest.Clone().(*client.CDNRequest)
	return &new_obj
}

func NewLiveRequest(action, domainName, appname string) (l *LiveRequest) {
	l = &LiveRequest{
		CDNRequest: client.NewCDNRequest(action),
		DomainName: domainName,
		AppName:    appname,
	}
	return
}
