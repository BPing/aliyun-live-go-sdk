package example

import (
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/device/cdn"
	"fmt"
)

func CDNExample() {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	cdnM := cdn.NewCDN(cert).SetDebug(false)
	resp := make(map[string]interface{})
	cdnM.ReadUserDomains(DomainName, 0, 0, "", "", &resp)
	fmt.Println(resp)

	resp = make(map[string]interface{})
	err := cdnM.ReadUserDomains(DomainName, 0, 0, cdn.OnlineDomainStatus, "", &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = cdnM.AddCdnDomain(cdn.DomainInfo{
		DomainName:DomainName + "test",
		CdnType:cdn.LiveStreamCdnType,
		SourceType:cdn.DomainSourceType,
		Sources:"2",
	}, &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = cdnM.DeleteCdnDomain(DomainName + "test", &resp)
	fmt.Println(err, resp)

	resp1 := cdn.DomainInfoResponse{}
	err = cdnM.CdnDomainDetail(DomainName, &resp1)
	fmt.Println(err, resp1)

	resp = make(map[string]interface{})
	err = cdnM.CdnDomainDetail(DomainName, &resp)
	fmt.Println(err, resp)


	//config
	resp = make(map[string]interface{})
	err = cdnM.DomainConfigs(DomainName, &resp)
	fmt.Println(err, resp)
}
