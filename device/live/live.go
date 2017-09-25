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
package live

import (
	"errors"
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/util"
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"time"
)

const (
	//action
	DescribeLiveStreamsPublishListAction    = "DescribeLiveStreamsPublishList"
	DescribeLiveStreamsOnlineListAction     = "DescribeLiveStreamsOnlineList"
	DescribeLiveStreamsBlockListAction      = "DescribeLiveStreamsBlockList"
	DescribeLiveStreamsControlHistoryAction = "DescribeLiveStreamsControlHistory"
	DescribeLiveStreamOnlineUserNumAction   = "DescribeLiveStreamOnlineUserNum"
	ForbidLiveStreamAction                  = "ForbidLiveStream"
	ResumeLiveStreamAction                  = "ResumeLiveStream"
	SetLiveStreamsNotifyUrlConfigAction     = "SetLiveStreamsNotifyUrlConfig"

	// 录制处理
	AddLiveAppRecordConfigAction                     = "AddLiveAppRecordConfig"
	CreateLiveStreamRecordIndexFilesAction           = "CreateLiveStreamRecordIndexFiles"
	DeleteLiveAppRecordConfigAction                  = "DeleteLiveAppRecordConfig"
	DescribeLiveAppRecordConfigAction                = "DescribeLiveAppRecordConfig"
	DescribeLiveRecordConfigAction                   = "DescribeLiveRecordConfig"
	DescribeLiveStreamRecordContentAction            = "DescribeLiveStreamRecordContent"
	DescribeLiveStreamRecordIndexFileAction          = "DescribeLiveStreamRecordIndexFile"
	DescribeLiveStreamRecordIndexFilesAction         = "DescribeLiveStreamRecordIndexFiles"
	DescribeLiveStreamsFrameRateAndBitRateDataAction = "DescribeLiveStreamsFrameRateAndBitRateData"

	// 截图处理
	AddLiveAppSnapshotConfigAction       = "AddLiveAppSnapshotConfig"
	UpdateLiveAppSnapshotConfigAction    = "UpdateLiveAppSnapshotConfig"
	DeleteLiveAppSnapshotConfigAction    = "DeleteLiveAppSnapshotConfig"
	DescribeLiveSnapshotConfigAction     = "DescribeLiveSnapshotConfig"
	DescribeLiveStreamSnapshotInfoAction = "DescribeLiveStreamSnapshotInfo"

	// 转码处理
	AddLiveStreamTranscodeAction          = "AddLiveStreamTranscode"
	DeleteLiveStreamTranscodeAction       = "DeleteLiveStreamTranscode"
	DescribeLiveStreamTranscodeInfoAction = "DescribeLiveStreamTranscodeInfo"

	// 混流处理
	StartMixStreamsServiceAction = "StartMixStreamsService"
	StopMixStreamsServiceAction  = "StopMixStreamsService"

	// 直播拉流
	AddLivePullStreamInfoConfigcAction   = "AddLivePullStreamInfoConfig"
	DeleteLivePullStreamInfoConfigAction = "DeleteLivePullStreamInfoConfig"
	DescribeLivePullStreamConfigAction   = "DescribeLivePullStreamConfig"

	//直播中心服务器域名
	DefaultVideoCenter = "video-center.alivecdn.com"
)

// Live 直播接口控制器
//      每一个实例都固定对应一个Cdn，并且无法更改。
//
//      方法名以"WithApp"结尾代表可以更改请求中  "应用名字（AppName）"，否则按默认的(初始化时设置的AppName)。
//      如果为空，代表忽略参数AppName
// @author cbping
type Live struct {
	rpc     *client.Client
	liveReq *LiveRequest

	//鉴权凭证
	//如果为nil，则代表不开启直播流推流鉴权
	streamCert *StreamCredentials

	// 推流地址：rtmp://video-center.alivecdn.com/AppName/StreamName?vhost=CDN
	// video-center.alivecdn.com是直播中心服务器，允许自定义，
	// 例如您的域名是live.yourcompany.com，可以设置DNS，将您的域名CNAME指向video-center.alivecdn.com即可；
	// 直播中心服务器或者自定义域名
	videoCenterDns string

	debug bool
}

// 新建"直播接口控制器"
// @param cert  请求凭证
// @param domainName 加速域名
// @param appname    应用名字
// @param streamCert  直播流推流凭证
func NewLive(cert *client.Credentials, domainName, appName string, streamCert *StreamCredentials) *Live {
	return &Live{
		rpc:            client.NewClient(cert),
		liveReq:        NewLiveRequest("", domainName, appName),
		debug:          false,
		streamCert:     streamCert,
		videoCenterDns: DefaultVideoCenter, //默认
	}
}

// GetStream 获取直播流
// @describe 每一次都生成新的流实例，不检查流名的唯一性，并且同一个名字会生成不同的实例的，
//          所以，使用时候，请自行确保流名的唯一性
func (l *Live) GetStream(streamName string) *Stream {
	if "" == streamName {
		return nil
	}

	var credentials *StreamCredentials
	if nil != l.streamCert {
		credentials = l.streamCert.Clone()
	}

	return &Stream{
		domainName:     l.liveReq.DomainName,
		appName:        l.liveReq.AppName,
		StreamName:     streamName,
		videoCenterDns: l.videoCenterDns,
		streamCert:     credentials,
		signOn:         (nil != l.streamCert),
		live:           l,
	}
}

func (l *Live) cloneRequest(action string) (req *LiveRequest) {
	req = l.SetAction(action).liveReq.Clone().(*LiveRequest)
	return
}

// @see StreamsPublishListWithApp
func (l *Live) StreamsPublishList(startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.StreamsPublishListWithApp(l.liveReq.AppName, startTime, endTime, resp)
	return
}

// StreamsPublishListWithApp 获取推流列表
// @appname 应用名 为空时，忽略此参数
// @startTime 开始时间
// @endTime   结束时间
// @link https://help.aliyun.com/document_detail/27191.html?spm=0.0.0.0.Dm58D2
func (l *Live) StreamsPublishListWithApp(appname string, startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamsPublishListAction)
	req.AppName = appname
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// @see StreamsOnlineListWithApp
func (l *Live) StreamsOnlineList(resp interface{}) (err error) {
	err = l.StreamsOnlineListWithApp(l.liveReq.AppName, resp)
	return
}

// StreamsOnlineListWithApp 获取在线流
// @appname 应用名 为空时，忽略此参数
// @link  https://help.aliyun.com/document_detail/27192.html?spm=0.0.0.0.7uWhjM
func (l *Live) StreamsOnlineListWithApp(appname string, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamsOnlineListAction)
	req.AppName = appname
	err = l.rpc.Query(req, resp)
	return
}

// StreamsBlockList 获取黑名单
// @link https://help.aliyun.com/document_detail/27193.html?spm=0.0.0.0.96SCaE
func (l *Live) StreamsBlockList(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamsBlockListAction)
	req.AppName = ""
	err = l.rpc.Query(req, resp)
	return
}

// @see StreamsControlHistoryWithApp
func (l *Live) StreamsControlHistory(startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.StreamsControlHistoryWithApp(l.liveReq.AppName, startTime, endTime, resp)
	return
}

// StreamsControlHistoryWithApp 获取控制历史
// @appname 应用名 为空时，忽略此参数
// @link  https://help.aliyun.com/document_detail/27194.html?spm=0.0.0.0.4DUTT7
func (l *Live) StreamsControlHistoryWithApp(appname string, startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamsControlHistoryAction)
	req.AppName = appname
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// @see StreamOnlineUserNumWithApp
func (l *Live) StreamOnlineUserNum(streamName string, resp interface{}) (err error) {
	err = l.StreamOnlineUserNumWithApp(l.liveReq.AppName, streamName, resp)
	return
}

// StreamOnlineUserNumWithApp 获取在线人数
// @appname 应用名 为空时，忽略此参数
// @link https://help.aliyun.com/document_detail/27195.html?spm=0.0.0.0.n6eAJJ
func (l *Live) StreamOnlineUserNumWithApp(appname string, streamName string, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamOnlineUserNumAction)
	req.AppName = appname
	if "" != streamName {
		req.SetArgs("StreamName", streamName)
	}
	err = l.rpc.Query(req, resp)
	return
}

// ForbidLiveStream 禁止流
// StreamName	String	是	流名称
// LiveStreamType	String	是	用于指定主播推流还是客户端拉流, 目前支持"publisher" (主播推送)
// ResumeTime	String	否	恢复流的时间 UTC时间 格式：2015-12-01T17:37:00Z
func (l *Live) ForbidLiveStream(appName, streamName string, liveStreamType string, resumeTime *time.Time, resp interface{}) (err error) {
	if global.EmptyString == appName {
		return errors.New("appName should not to be empty")
	}
	req := l.cloneRequest(ForbidLiveStreamAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	req.SetArgs("LiveStreamType", liveStreamType)
	if nil != resumeTime {
		req.SetArgs("ResumeTime", util.GetISO8601TimeStamp(*resumeTime))
	}
	err = l.rpc.Query(req, resp)
	return
}

// @see ForbidLiveStream
func (l *Live) ForbidLiveStreamWithPublisher(streamName string, resumeTime *time.Time, resp interface{}) (err error) {
	return l.ForbidLiveStream(l.liveReq.AppName, streamName, "publisher", resumeTime, resp)
}

// @see ForbidLiveStream
func (l *Live) ForbidLiveStreamWithPublisherWithApp(appName, streamName string, resumeTime *time.Time, resp interface{}) (err error) {
	return l.ForbidLiveStream(appName, streamName, "publisher", resumeTime, resp)
}

// ResumeLiveStream 恢复流
func (l *Live) ResumeLiveStream(appName, streamName string, liveStreamType string, resp interface{}) (err error) {
	if global.EmptyString == appName {
		return errors.New("appName should not to be empty")
	}
	req := l.cloneRequest(ResumeLiveStreamAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	req.SetArgs("LiveStreamType", liveStreamType)
	err = l.rpc.Query(req, resp)
	return
}

// @see ResumeLiveStream
func (l *Live) ResumeLiveStreamWithPublisher(streamName string, resp interface{}) (err error) {
	return l.ResumeLiveStream(l.liveReq.AppName, streamName, "publisher", resp)
}

// @see ResumeLiveStream
func (l *Live) ResumeLiveStreamWithPublisherWithApp(appName, streamName string, resp interface{}) (err error) {
	return l.ResumeLiveStream(appName, streamName, "publisher", resp)
}

// SetStreamsNotifyUrlConfig 设置回调链接
// NotifyUrl	String	是	设置直播流信息推送到的URL地址，必须以http://开头；
func (l *Live) SetStreamsNotifyUrlConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.cloneRequest(SetLiveStreamsNotifyUrlConfigAction)
	req.AppName = ""
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.rpc.Query(req, resp)
	return
}

// GET 和 SET
// ---------------------------------------------------------------------------------------------------------------------

/*// 修改默认或者说全局  domainName（加速域名）
func (l *Live) SetDomainName(domainName string) *Live {
	l.liveReq.DomainName = domainName
	return l
}*/

func (l *Live) GetDomainName() (domainName string) {
	domainName = l.liveReq.DomainName
	return
}

//修改默认或者说全局  StreamCredentials（流签名凭证）
func (l *Live) SetStreamCredentials(streamCert *StreamCredentials) *Live {
	l.streamCert = streamCert
	return l
}

// 修改默认或者说全局 appname（应用名）
func (l *Live) SetAppName(appname string) *Live {
	l.liveReq.AppName = appname
	return l
}

func (l *Live) GetAppName() (appname string) {
	appname = l.liveReq.AppName
	return
}

// 修改默认或者说全局 action（操作名称类别）
func (l *Live) SetAction(action string) *Live {
	l.liveReq.Action = action
	return l
}
func (l *Live) SetDebug(debug bool) *Live {
	l.debug = debug
	l.rpc.SetDebug(debug)
	return l
}

// 修改默认或者说全局 videoCenterDns（对应的直播推流域名）
func (l *Live) SetVideoCenter(videoCenterDns string) *Live {
	l.videoCenterDns = videoCenterDns
	return l
}
