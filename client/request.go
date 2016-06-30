package client

import (
	"net/http"
	"time"
	"github.com/BPing/aliyun-live-go-sdk/util"
	"net/url"
	"fmt"
)

//常量
const (
	DefaultSignatureVersion = "1.0"
	DefaultSignatureMethod = "HMAC-SHA1"
	JSONResponseFormat = "JSON"
	XMLResponseFormat = "XML"
	ECSRequestMethod = "GET"
)

//
//  请求接口
//    所有请求对象继承的接口，也是Client接受处理的请求接口
//    签名方式和必要参数信息。
type Request interface {
	//签名
	Sign(*Credentials)
	//返回*http.Request
	HttpRequestInstance() (*http.Request, error)
	//返回值的类型，支持JSON与XML.
	ResponseFormat() string
	//
	String() string
	//
	Clone() interface{}
	////返回请求处理超时限制时长
	//DeadLine() time.Duration
}

//type BaseRequest struct {
//}
//
//func (b *BaseRequest)Sign(accessKeyId, accessKeySecret string) {
//}
//
//func (b *BaseRequest)HttpRequestInstance() (*http.Request, error) {
//	return nil, nil
//}
//
//func (b BaseRequest)ResponseFormat() string {
//	return JSONResponseFormat
//}

//// A Timeout of zero means no timeout.
//func (b BaseRequest)DeadLine() time.Duration {
//	return 0
//}


//
//  cdn 请求对象。实现 Request 接口
//
// 名称	类型	是否必须	描述
// Format	String	否	返回值的类型，支持JSON与XML。默认为XML \n
// Version	String	是	API版本号，为日期形式：YYYY-MM-DD，本版本对应为2014-11-11 \n
// AccessKeyId	String	是	阿里云颁发给用户的访问服务所用的密钥ID \n
// Signature	String	是	签名结果串，关于签名的计算方法，请参见签名机制。\n
// SignatureMethod	String	是	签名方式，目前支持HMAC-SHA1 \n
// Timestamp	String	是	请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。例如，2014-11-11T12:00:00Z（为北京时间2014年11月11日20点0分0秒）\n
// SignatureVersion	String	是	签名算法版本，目前版本是1.0 \n
// SignatureNonce	String	是	唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值 \n
type CDNRequest struct {
	Format           string
	Version          string
	AccessKeyId      string
	Signature        string
	SignatureMethod  string
	Timestamp        util.ISO6801Time
	SignatureVersion string
	SignatureNonce   string
	// ResourceOwnerAccount string
	Action           string

	// http
	Host             string
	Method           string
	Url              string
	Args             url.Values
}

// CDNRequest的必要字段转成参数
func (Cdn *CDNRequest)StructToArgs() {
	Cdn.SignatureNonce = util.CreateRandomString()
	//Cdn.Timestamp = util.NewISO6801Time(time.Now().UTC())
	Cdn.Args.Set("Format", Cdn.Format)
	Cdn.Args.Set("Version", Cdn.Version)
	Cdn.Args.Set("AccessKeyId", Cdn.AccessKeyId)
	Cdn.Args.Set("SignatureMethod", Cdn.SignatureMethod)
	Cdn.Args.Set("Timestamp", Cdn.Timestamp.String())
	Cdn.Args.Set("SignatureVersion", Cdn.SignatureVersion)
	Cdn.Args.Set("SignatureNonce", Cdn.SignatureNonce)
	Cdn.Args.Set("Action", Cdn.Action)
}

// 签名
func (Cdn *CDNRequest)Sign(cert *Credentials) {
	Cdn.AccessKeyId = cert.AccessKeyId
	Cdn.StructToArgs()
	// 生成签名
	Cdn.Signature = util.CreateSignatureForRequest(Cdn.Method, &Cdn.Args, cert.AccessKeySecret + "&")

}

func (Cdn *CDNRequest)HttpRequestInstance() (httpReq *http.Request, err error) {
	// 生成请求url
	Cdn.Url = Cdn.Host + "?" + Cdn.Args.Encode() + "&Signature=" + url.QueryEscape(Cdn.Signature)
	httpReq, err = http.NewRequest(Cdn.Method, Cdn.Url, nil)
	return
}

func (Cdn *CDNRequest)ResponseFormat() string {
	return Cdn.Format
}

// A Timeout of zero means no timeout.
func (Cdn *CDNRequest)DeadLine() time.Duration {
	return 0
}

func (Cdn *CDNRequest)String() string {
	return fmt.Sprintf("Method:%s,Url:%s", Cdn.Method, Cdn.Url)
}

// 克隆
func (l *CDNRequest)Clone() interface{} {
	new_obj := (*l)
	//清空数据
	new_obj.Args = url.Values{}
	return &new_obj
}

func (Cdn *CDNRequest)SetArgs(key, value string) {
	Cdn.Args.Set(key, value)
}

func (Cdn *CDNRequest)DelArgs(key string) {
	Cdn.Args.Del(key)
}

const (
	ApiCDNVersion = "2014-11-11"
	ApiCDNHost = "https://cdn.aliyuncs.com/"
)

// 生成CDNRequest
func NewCDNRequest(action string) *CDNRequest {
	return &CDNRequest{
		Format:JSONResponseFormat,
		Version:ApiCDNVersion,
		SignatureNonce:util.CreateRandomString(),
		SignatureMethod:DefaultSignatureMethod,
		SignatureVersion:DefaultSignatureVersion,
		Timestamp:util.NewISO6801Time(time.Now().UTC()),

		Action:action,

		Host:ApiCDNHost,
		Method:ECSRequestMethod,
		Args:url.Values{},

	}
}
