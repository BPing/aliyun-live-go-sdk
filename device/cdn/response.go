package cdn

import (
	"aliyun-live-go-sdk/util/global"
	"aliyun-live-go-sdk/client"
)

// 加速域名的业务类型
type  CdnType string

// 源站类型
type  SourceType string

// 域名查询类型
type  DomainSearchType string

// 加速域名运行状态
type  DomainStatus string

const (
//CdnType
	WebCdnType CdnType = "web" //图片及小文件分发
	DownloadCdnType CdnType = "download" //大文件下载加速
	VideoCdnType CdnType = "video" //视音频点播加速
	LiveStreamCdnType CdnType = "liveStream" //直播流媒体加速
	HttpsDeliveryStreamCdnType CdnType = "httpsDelivery" //Https安全加速
	NullCdnType CdnType = global.EmptyString //

//SourceType
	IpaddrSourceType SourceType = "ipaddr" //IP源站
	DomainSourceType SourceType = "domain" //域名源站
	OSSSourceType SourceType = "OSS" //OSS Bucket为源站
	LiveStreamSourceType SourceType = global.EmptyString // 注：若选择了直播流媒体加速的业务类型，无需填写源站类型和信息

//DomainSearchType 域名查询类型：fuzzy_match 模糊匹配,pre_match 前匹配,suf_match 后匹配,full_match 完全匹配，默认fuzzy_match
	FuzzyMatch DomainSearchType = "fuzzy_match"
	PreMatch DomainSearchType = "pre_match "
	SufMatch DomainSearchType = "suf_match "
	FullMatch DomainSearchType = "full_match"

//DomainStatus
//取值意义：online表示启用；offline表示停用；configuring表示配置中；configure_failed表示配置失败;checking表示正在审核；check_failed表示审核失败
	OnlineDomainStatus DomainStatus = "online"
	OfflineDomainStatus DomainStatus = "offline"
	ConfiguringDomainStatus DomainStatus = "configuring"
	ConfigureFailedDomainStatus DomainStatus = "configure_failed"
	CheckFailedDomainStatus DomainStatus = "check_failed"
	CheckingDomainStatus DomainStatus = "checking"
)

type DomainInfo struct {
									 //添加加速域名用到的信息
	DomainName              string                                   //(接入CDN进行加速的域名)
	Sources                 string      `json:"" xml:"null"`         //(源站信息，域名或IP) 输入信息 相当
	CdnType                 CdnType                                  //(加速域名的业务类型 取值含义：web表示图片及小文件加速；download表示大文件下载加速；video表示视音频点播加速；liveStream表示直播流媒体加速 ；httpsDelivery表示HTTPS安全加速)
	SourceType              SourceType                               //(源站类型 取值含义：ipaddr表示IP源站；domain表示域名源站；oss表示指定OSS Bucket为源站)
	ServerCertificate       string                                   //(如果开启，此处为证书公钥) 如果是HttpsDelivery，需要上传的安全证书。
	SourcePort              int64                                    //可以指定443,80。默认值80。443的话走https回源。oss不支持443
	PrivateKey              string                                   //如果是HttpsDelivery，需要上传的私钥。

									 //其他信息
	Cname                   string                                   //(为加速域名生成的一个CNAME域名，需要在域名解析服务商处将加速域名CNAME解析到该域名)
	HttpsCname              string                                   //(开启https的CNAME域名)
	DomainStatus            DomainStatus                             //(加速域名运行状态 取值意义：online表示启用；offline表示停用；configuring表示配置中；configure_failed表示配置失败;checking表示正在审核；check_failed表示审核失败)
	ServerCertificateStatus string                                   //(是否开启ssl证书 on表示开启；off表示关闭)
	GmtCreated              string                                   //(创建时间)
	GmtModified             string                                   //(最近修改时间)

	SourcesJson             map[string]interface{}  `json:"Sources"` //(Json 返回信息 Sources别名)
	SourcesXml              []string `xml:"Sources>Source"`          //(Xml  返回信息 Sources 别名)
}

type DomainInfoResponse struct {
	client.Response
	GetDomainDetailModel DomainInfo
}