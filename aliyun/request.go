package aliyun

import (
	"fmt"
	"github.com/BPing/aliyun-live-go-sdk/util"
	"github.com/BPing/go-toolkit/http-client/core"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request interface {
	core.Request
	//签名
	Sign(*Credentials)
	//返回值的类型，支持JSON与XML.
	ResponseFormat() string
}

// BaseRequest 基础请求结构
// 名称	类型	是否必须	描述
// Format	String	否	返回值的类型，支持JSON与XML。默认为XML \n
// Version	String	是	API版本号，为日期形式：YYYY-MM-DD，本版本对应为2014-11-11 \n
// AccessKeyId	String	是	阿里云颁发给用户的访问服务所用的密钥ID \n
// Signature	String	是	签名结果串，关于签名的计算方法，请参见签名机制。\n
// SignatureMethod	String	是	签名方式，目前支持HMAC-SHA1 \n
// Timestamp	String	是	请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。例如，2014-11-11T12:00:00Z（为北京时间2014年11月11日20点0分0秒）\n
// SignatureVersion	String	是	签名算法版本，目前版本是1.0 \n
// SignatureNonce	String	是	唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值 \n
type BaseRequest struct {
	core.BaseRequest
	Format           string
	Version          string
	AccessKeyId      string
	Signature        string
	SignatureMethod  string
	Timestamp        util.ISO6801Time
	SignatureVersion string
	SignatureNonce   string
	// ResourceOwnerAccount string
	Action string

	// http
	Host   string
	Method string
	Url    string
	Args   url.Values
}

// 必要字段转成参数
func (base *BaseRequest) ToArgs() {
	base.SignatureNonce = util.CreateRandomString()
	//Cdn.Timestamp = util.NewISO6801Time(time.Now().UTC())
	base.Args.Set("Format", base.Format)
	base.Args.Set("Version", base.Version)
	base.Args.Set("AccessKeyId", base.AccessKeyId)
	base.Args.Set("SignatureMethod", base.SignatureMethod)
	base.Args.Set("Timestamp", base.Timestamp.String())
	base.Args.Set("SignatureVersion", base.SignatureVersion)
	base.Args.Set("SignatureNonce", base.SignatureNonce)
	base.Args.Set("Action", base.Action)
}

// 签名
func (base *BaseRequest) Sign(cert *Credentials) {
	base.AccessKeyId = cert.AccessKeyID
	base.ToArgs()
	// 生成签名
	base.Signature = util.CreateSignatureForRequest(base.Method, &base.Args, cert.AccessKeySecret+"&")
}

func (base *BaseRequest) HttpRequest() (httpReq *http.Request, err error) {
	// 生成请求url
	base.Url = base.Host + "?" + base.Args.Encode() + "&Signature=" + url.QueryEscape(base.Signature)
	httpReq, err = http.NewRequest(base.Method, base.Url, nil)
	httpReq.Header.Set("X-SDK-Client", `AliyunLiveGoSDK/`+Version)
	httpReq.Header.Set("Content-Type", `application/`+strings.ToLower(base.Format))
	return
}

func (base *BaseRequest) ResponseFormat() string {
	return base.Format
}

func (base *BaseRequest) String() string {
	return fmt.Sprintf("Method:%s,Url:%s", base.Method, base.Url)
}

// 克隆
func (base *BaseRequest) Clone() interface{} {
	newObj := *base
	//清空数据
	newObj.Args = url.Values{}
	return &newObj
}

func (base *BaseRequest) SetArgs(key, value string) {
	base.Args.Set(key, value)
}

func (base *BaseRequest) DelArgs(key string) {
	base.Args.Del(key)
}

//
func NewBaseRequest(action string) *BaseRequest {
	return &BaseRequest{
		Format:           JSONResponseFormat,
		Version:          APICDNVersion,
		SignatureNonce:   util.CreateRandomString(),
		SignatureMethod:  DefaultSignatureMethod,
		SignatureVersion: DefaultSignatureVersion,
		Timestamp:        util.NewISO6801Time(time.Now().UTC()),

		Action: action,

		Host:   APICDNHost,
		Method: ECSRequestMethod,
		Args:   url.Values{},
	}
}
