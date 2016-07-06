package device

import (
	"testing"
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
	"github.com/BPing/aliyun-live-go-sdk/device/cdn"
)

const (
	AccessKeyId = ""
	AccessKeySecret = ""
	DomainName = "DomainName"
	AppName = "AppName"
	PrivateKey = ""
)

//
func TestDevice(t *testing.T) {

	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	streamCert := live.NewStreamCredentials(PrivateKey, live.DefualtStreamTimeout)

	cdnDev, err := GetDevice(CdnDevice, Config{Credentials:cert, })
	if _, ok := cdnDev.(*cdn.CDN); err != nil || !ok {
		t.Fatal("get cdn device fail", err, ok)
	}

	cdnDev, err = GetDevice(CdnDevice, Config{})
	if cdnDev != nil || err == nil {
		t.Fatal("get cdn device : param error")
	}

	liveDev, err := GetDevice(LiveDevice,
		Config{Credentials:cert,
			StreamCert:streamCert,
			DomainName:DomainName,
			AppName:AppName, })
	if _, ok := liveDev.(*live.Live); err != nil || !ok {
		t.Fatal("get live device fail", err, ok)
	}

	liveDev, err = GetDevice(LiveDevice,
		Config{Credentials:cert,
			StreamCert:streamCert,
			DomainName:"",
			AppName:AppName, })
	if cdnDev != nil || err == nil {
		t.Fatal("get cdn device : param error")
	}

}
