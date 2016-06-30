package cdn

import (
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"github.com/BPing/aliyun-live-go-sdk/client"
)

// 加速域名的业务类型
type  CdnType string

// 源站类型
type  SourceType string

// 域名查询类型
type  DomainSearchType string

// 加速域名运行状态
type  DomainStatus string

// 开通服务的计费类型
type CdnPayType string

// ConfigList参数
type ConfigName string

// 刷新类型
type RefreshObjectType string

// 错误页面类型
type PageType string

// 强制跳转类型, 取值: Off, Http, Https
type RedirectType string

// refer类型,
type ReferType string

// 缓存内容类型,
type CacheType string

// 鉴权类型
type AuthType string

const (
// CdnType
	WebCdnType CdnType = "web" //图片及小文件分发
	DownloadCdnType CdnType = "download" //大文件下载加速
	VideoCdnType CdnType = "video" //视音频点播加速
	LiveStreamCdnType CdnType = "liveStream" //直播流媒体加速
	HttpsDeliveryStreamCdnType CdnType = "httpsDelivery" //Https安全加速
	NullCdnType CdnType = global.EmptyString //

// SourceType
	IpaddrSourceType SourceType = "ipaddr" //IP源站
	DomainSourceType SourceType = "domain" //域名源站
	OSSSourceType SourceType = "OSS" //OSS Bucket为源站
	LiveStreamSourceType SourceType = global.EmptyString // 注：若选择了直播流媒体加速的业务类型，无需填写源站类型和信息

// DomainSearchType 域名查询类型：fuzzy_match 模糊匹配,pre_match 前匹配,suf_match 后匹配,full_match 完全匹配，默认fuzzy_match
	FuzzyMatch DomainSearchType = "fuzzy_match"
	PreMatch DomainSearchType = "pre_match "
	SufMatch DomainSearchType = "suf_match "
	FullMatch DomainSearchType = "full_match"

// DomainStatus
// 取值意义：online表示启用；offline表示停用；configuring表示配置中；configure_failed表示配置失败;checking表示正在审核；check_failed表示审核失败
	OnlineDomainStatus DomainStatus = "online"
	OfflineDomainStatus DomainStatus = "offline"
	ConfiguringDomainStatus DomainStatus = "configuring"
	ConfigureFailedDomainStatus DomainStatus = "configure_failed"
	CheckFailedDomainStatus DomainStatus = "check_failed"
	CheckingDomainStatus DomainStatus = "checking"

// CdnPayType
	PayByTrafficType CdnPayType = "PayByTraffic" //按流量
	PayByBandwidthType CdnPayType = "PayByBandwidth" //按带宽峰值
	PayNullType CdnPayType = "" //空值，默认采用 PayByTrafficType

// ConfigList参数
	CacheExpiredConfig ConfigName = "cache_expired" //目录或文件过期配置
	CcConfig ConfigName = "cc" //cc	String	CC防护功能、IP黑白名单配置
	ErrorPageConfig ConfigName = "error_page" //error_page	String	自定义404错误页面跳转
	HttpHeaderConfig ConfigName = "http_header" //http_header	String	自定义http头
	OptimizeConfig ConfigName = "optimize" //optimize	String	页面优化
	PageCompressConfig ConfigName = "page_compress" //page_compress	String	智能压缩
	IgnoreQueryStringConfig ConfigName = "ignore_query_string" //ignore_query_string	String	过滤参数
	RangeConfig ConfigName = "range" //range	String	range回源功能
	RefererConfig ConfigName = "referer" //referer	String	Refer防盗链功能
	ReqAuthConfig ConfigName = "req_auth" //req_auth	String	访问鉴权配置
	SrcHostConfig ConfigName = "src_host" //src_host	String	回源host
	VideoSeekConfig ConfigName = "video_seek" //video_seek	String	拖拽播放功能
	WafConfig ConfigName = "waf" //waf	String	Waf防护功能
	NotifyRrlConfig ConfigName = "notify_url" //notify_url	String	视频直播notify url
	RedirectTypeConfig ConfigName = "redirect_type" //redirect_type	String	强制访问跳转方式, 取值: Off, Http, Https

// RefreshObjectType
	FileRefreshType RefreshObjectType = "File"
	DirectoryRefreshType RefreshObjectType = "Directory"

// PageType 错误页面类型； 取值：default：默认页面；charity：公益页面；other：自定义页面
	DefaultPageType PageType = "default"
	CharityPageType PageType = "charity"
	OtherPageType PageType = "other"

// RedirectType 强制跳转类型, 取值: Off, Http, Https
	OffRedirectType RedirectType = "Off"
	HttpRedirectType RedirectType = "Http"
	HttpsRedirectType RedirectType = "Https"

// ReferType refer类型，取值：block：黑名单；allow：白名单
	BlockReferType ReferType = "block"
	AllowReferType ReferType = "allow"

// CacheType 缓存内容类型，取值：suffix：文件名后缀；path：路径，支持目录和完整路
	SuffixCacheType CacheType = "suffix"
	PathCacheType CacheType = "path"

// AuthType	String	是	鉴权类型，取值: "no_auth":关闭,"type_a":A方式，"type_b":B方式, "type_c":C方式
	NoAuthType AuthType = "no_auth"
	AAuthType AuthType = "type_a"
	BAuthType AuthType = "type_b"
	CAuthType AuthType = "type_c"
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