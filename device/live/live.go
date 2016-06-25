//Copyright cbping
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License

//
//  阿里云直播API
//  文档信息：https://help.aliyun.com/document_detail/27191.html?spm=0.0.0.0.60u2Ny
//  @author cbping
//
package live

import (
	"aliyun-live-go-sdk/util"
	"aliyun-live-go-sdk/client"
	"time"
)

const (
//action
	DescribeLiveStreamsPublishListAction = "DescribeLiveStreamsPublishList"
	DescribeLiveStreamsOnlineListAction = "DescribeLiveStreamsOnlineList"
	DescribeLiveStreamsBlockListAction = "DescribeLiveStreamsBlockList"
	DescribeLiveStreamsControlHistoryAction = "DescribeLiveStreamsControlHistory"
	DescribeLiveStreamOnlineUserNumAction = "DescribeLiveStreamOnlineUserNum"
	ForbidLiveStreamAction = "ForbidLiveStream"
	ResumeLiveStreamAction = "ResumeLiveStream"
	SetLiveStreamsNotifyUrlConfigAction = "SetLiveStreamsNotifyUrlConfig"
)


//
//
// 直播
// @author cbping
// @describe 方法名以"WithApp"结尾代表可以更改请求中  "应用名字（AppName）"，否则按默认的(初始化时设置的AppName)。如果为空，代表忽略参数AppName
type Live struct {
	Rpc        *client.Client
	LiveReq    *LiveRequest

	//鉴权凭证
	//如果为nil，则代表不开启直播流推流鉴权
	streamCert *StreamCredentials

	debug      bool
}

//@param cert  请求凭证
//@param domainName 加速域名
//@param appname    应用名字
//@param streamCert  直播流推流凭证
func NewLive(cert *client.Credentials, domainName, appname string, streamCert *StreamCredentials) *Live {
	return &Live{
		Rpc:client.NewClient(cert),
		LiveReq:NewLiveRequest("", domainName, appname),
		debug:false,
		streamCert:streamCert,
	}
}

//获取直播流
func (l *Live)GetStream(streamName string) *Stream {
	if ("" == streamName) {
		return nil
	}

	var credentials *StreamCredentials
	if (nil != l.streamCert) {
		credentials = l.streamCert.Clone()
	}

	return &Stream{
		DomainName:l.LiveReq.DomainName,
		AppName:l.LiveReq.AppName,
		StreamName:streamName,
		credentials:credentials,
		signOn:(nil != l.streamCert),
		live:l,
	}
}

func (l *Live) StreamsPublishList(startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.StreamsPublishListWithApp(l.LiveReq.AppName, startTime, endTime, resp)
	return
}

//@appname 应用名 为空时，忽略此参数
//@startTime 开始时间
//@endTime   结束时间
func (l *Live) StreamsPublishListWithApp(appname string, startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.SetAction(DescribeLiveStreamsPublishListAction).LiveReq.Clone().(*LiveRequest)
	req.AppName = appname
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.Rpc.Query(req, resp)
	return
}

//获取在线流
func (l *Live) StreamsOnlineList(resp interface{}) (err error) {
	err = l.StreamsOnlineListWithApp(l.LiveReq.AppName, resp)
	return
}

//@appname 应用名 为空时，忽略此参数
func (l *Live) StreamsOnlineListWithApp(appname string, resp interface{}) (err error) {
	req := l.SetAction(DescribeLiveStreamsOnlineListAction).LiveReq.Clone().(*LiveRequest)
	req.AppName = appname
	err = l.Rpc.Query(req, resp)
	return
}


//获取黑名单
func (l *Live) StreamsBlockList(resp interface{}) (err error) {
	req := l.SetAction(DescribeLiveStreamsBlockListAction).LiveReq.Clone().(*LiveRequest)
	req.AppName = ""
	err = l.Rpc.Query(req, resp)
	return
}

//获取控制历史
func (l *Live) StreamsControlHistory(startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.StreamsControlHistoryWithApp(l.LiveReq.AppName, startTime, endTime, resp)
	return
}

//@appname 应用名 为空时，忽略此参数
func (l *Live) StreamsControlHistoryWithApp(appname string, startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.SetAction(DescribeLiveStreamsControlHistoryAction).LiveReq.Clone().(*LiveRequest)
	req.AppName = appname
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.Rpc.Query(req, resp)
	return
}

//或者在线人数
func (l *Live) StreamOnlineUserNum(streamName string, resp interface{}) (err error) {
	err = l.StreamOnlineUserNumWithApp(l.LiveReq.AppName, streamName, resp)
	return
}

//@appname 应用名 为空时，忽略此参数
func (l *Live) StreamOnlineUserNumWithApp(appname string, streamName string, resp interface{}) (err error) {
	req := l.SetAction(DescribeLiveStreamOnlineUserNumAction).LiveReq.Clone().(*LiveRequest)
	req.AppName = appname
	if ("" != streamName) {
		req.SetArgs("StreamName", streamName)
	}
	err = l.Rpc.Query(req, resp)
	return
}


// 禁止流
//StreamName	String	是	流名称
//LiveStreamType	String	是	用于指定主播推流还是客户端拉流, 目前支持"publisher" (主播推送)
//ResumeTime	String	否	恢复流的时间 UTC时间 格式：2015-12-01T17:37:00Z
func (l *Live) ForbidLiveStream(streamName string, liveStreamType string, resumeTime *time.Time, resp interface{}) (err error) {
	req := l.SetAction(ForbidLiveStreamAction).LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("LiveStreamType", liveStreamType)
	if (nil != resumeTime) {
		req.SetArgs("ResumeTime", util.GetISO8601TimeStamp(*resumeTime))
	}
	err = l.Rpc.Query(req, resp)
	return
}

func (l *Live) ForbidLiveStreamWithPublisher(streamName string, resumeTime *time.Time, resp interface{}) (err error) {
	return l.ForbidLiveStream(streamName, "publisher", resumeTime, resp)
}

//恢复流
func (l *Live) ResumeLiveStream(streamName string, liveStreamType string, resp interface{}) (err error) {
	req := l.SetAction(ResumeLiveStreamAction).LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("LiveStreamType", liveStreamType)
	err = l.Rpc.Query(req, resp)
	return
}

func (l *Live) ResumeLiveStreamWithPublisher(streamName string, resp interface{}) (err error) {
	return l.ResumeLiveStream(streamName, "publisher", resp)
}

//设置回调链接
//Action	String	是	操作接口名，系统规定参数，取值：SetLiveStreamsNotifyUrlConfig
//DomainName	String	是	您的加速域名
//NotifyUrl	String	是	设置直播流信息推送到的URL地址，必须以http://开头；
func (l *Live) SetStreamsNotifyUrlConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.SetAction(SetLiveStreamsNotifyUrlConfigAction).LiveReq.Clone().(*LiveRequest)
	req.AppName = ""
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.Rpc.Query(req, resp)
	return
}

// GET 和 SET
// -------------------------------------------------------------------------------

func (l *Live)SetDomainName(domainName string) *Live {
	l.LiveReq.DomainName = domainName
	return l
}

//修改默认或者说全局  domainName（加速域名）
func (l *Live)GetDomainName() (domainName string) {
	domainName = l.LiveReq.DomainName
	return
}

func (l *Live)SetStreamCredentials(streamCert *StreamCredentials) *Live {
	l.streamCert = streamCert
	return l
}

//修改默认或者说全局 appname（应用名）
func (l *Live)SetAppName(appname string) *Live {
	l.LiveReq.AppName = appname
	return l
}

func (l *Live)GetAppName() (appname string) {
	appname = l.LiveReq.AppName
	return
}

func (l *Live)SetAction(action string) *Live {
	l.LiveReq.Action = action
	return l
}
func (l *Live)SetDebug(debug bool) *Live {
	l.debug = debug
	l.Rpc.SetDebug(debug)
	return l
}

