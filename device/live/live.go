package live

import (
	"aliyun-live-go-sdk/util"
	"aliyun-live-go-sdk/client"
)

type Live struct {
	Rpc     *client.Client
	LiveReq *LiveRequest
	debug   bool
}

func NewLive(cert *client.Credentials, domainName, appname string) *Live {
	return &Live{client.NewClient(cert), NewLiveRequest("", domainName, appname), false}
}

func (l *Live)SetDomainName(domainName string) *Live {
	l.LiveReq.DomainName = domainName
	return l
}

func (l *Live)SetAppName(appname string) *Live {
	l.LiveReq.AppName = appname
	return l
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

func (l *Live) StreamsPublishList(startTime, endTime util.ISO6801Time, resp interface{}) (err error) {
	req := l.SetAction("DescribeLiveStreamsPublishList").LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StartTime", startTime.String())
	req.SetArgs("EndTime", endTime.String())
	err = l.Rpc.Query(req, &resp)
	return
}

func (l *Live) StreamsOnlineList(resp interface{}) (err error) {
	req := l.SetAction("DescribeLiveStreamsOnlineList").LiveReq.Clone().(*LiveRequest)
	err = l.Rpc.Query(req, &resp)
	return
}

func (l *Live) StreamsBlockList(resp interface{}) (err error) {
	req := l.SetAction("DescribeLiveStreamsBlockList").LiveReq.Clone().(*LiveRequest)
	req.AppName = ""
	err = l.Rpc.Query(req, &resp)
	return
}

func (l *Live) StreamsControlHistory(startTime, endTime util.ISO6801Time, resp interface{}) (err error) {
	req := l.SetAction("DescribeLiveStreamsControlHistory").LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StartTime", startTime.String())
	req.SetArgs("EndTime", endTime.String())
	err = l.Rpc.Query(req, &resp)
	return
}

func (l *Live) StreamOnlineUserNum(streamName string, resp interface{}) (err error) {
	req := l.SetAction("DescribeLiveStreamOnlineUserNum").LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StreamName", streamName)
	err = l.Rpc.Query(req, &resp)
	return
}

//StreamName	String	是	流名称
//LiveStreamType	String	是	用于指定主播推流还是客户端拉流, 目前支持"publisher" (主播推送)
//ResumeTime	String	否	恢复流的时间 UTC时间 格式：2015-12-01T17:37:00Z
func (l *Live) ForbidLiveStream(streamName string, liveStreamType string, resumeTime *util.ISO6801Time, resp interface{}) (err error) {
	req := l.SetAction("ForbidLiveStream").LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("LiveStreamType", liveStreamType)
	if (nil != resumeTime) {
		req.SetArgs("ResumeTime", resumeTime.String())
	}
	err = l.Rpc.Query(req, &resp)
	return
}

func (l *Live) ForbidLiveStreamWithPublisher(streamName string, resumeTime *util.ISO6801Time, resp interface{}) (err error) {
	return l.ForbidLiveStream(streamName, "publisher", resumeTime, resp)
}

func (l *Live) ResumeLiveStream(streamName string, liveStreamType string, resp interface{}) (err error) {
	req := l.SetAction("ResumeLiveStream").LiveReq.Clone().(*LiveRequest)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("LiveStreamType", liveStreamType)
	err = l.Rpc.Query(req, &resp)
	return
}

func (l *Live) ResumeLiveStreamWithPublisher(streamName string, resp interface{}) (err error) {
	return l.ResumeLiveStream(streamName, "publisher", resp)
}

//Action	String	是	操作接口名，系统规定参数，取值：SetLiveStreamsNotifyUrlConfig
//DomaiName	String	是	您的加速域名
//NotifyUrl	String	是	设置直播流信息推送到的URL地址，必须以http://开头；
func (l *Live) SetStreamsNotifyUrlConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.SetAction("SetLiveStreamsNotifyUrlConfig").LiveReq.Clone().(*LiveRequest)
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.Rpc.Query(req, &resp)
	return
}