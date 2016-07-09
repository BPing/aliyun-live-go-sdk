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
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/util"
	"time"
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"errors"
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

	AddLiveAppRecordConfigAction = "AddLiveAppRecordConfig"
	CreateLiveStreamRecordIndexFilesAction = "CreateLiveStreamRecordIndexFiles"
	DeleteLiveAppRecordConfigAction = "DeleteLiveAppRecordConfig"
	DescribeLiveAppRecordConfigAction = "DescribeLiveAppRecordConfig"
	DescribeLiveRecordConfigAction = "DescribeLiveRecordConfig"
	DescribeLiveStreamRecordContentAction = "DescribeLiveStreamRecordContent"
	DescribeLiveStreamRecordIndexFileAction = "DescribeLiveStreamRecordIndexFile"
	DescribeLiveStreamRecordIndexFilesAction = "DescribeLiveStreamRecordIndexFiles"
	DescribeLiveStreamsFrameRateAndBitRateDataAction = "DescribeLiveStreamsFrameRateAndBitRateData"

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
	rpc            *client.Client
	liveReq        *LiveRequest

	//鉴权凭证
	//如果为nil，则代表不开启直播流推流鉴权
	streamCert     *StreamCredentials

	// 推流地址：rtmp://video-center.alivecdn.com/AppName/StreamName?vhost=CDN
	// video-center.alivecdn.com是直播中心服务器，允许自定义，
	// 例如您的域名是live.yourcompany.com，可以设置DNS，将您的域名CNAME指向video-center.alivecdn.com即可；
	// 直播中心服务器或者自定义域名
	videoCenterDns string

	debug          bool
}
// 新建"直播接口控制器"
// @param cert  请求凭证
// @param domainName 加速域名
// @param appname    应用名字
// @param streamCert  直播流推流凭证
func NewLive(cert *client.Credentials, domainName, appName string, streamCert *StreamCredentials) *Live {
	return &Live{
		rpc:        client.NewClient(cert),
		liveReq:    NewLiveRequest("", domainName, appName),
		debug:      false,
		streamCert: streamCert,
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
		domainName:  l.liveReq.DomainName,
		appName:     l.liveReq.AppName,
		StreamName:  streamName,
		videoCenterDns:  l.videoCenterDns,
		streamCert: credentials,
		signOn:      (nil != l.streamCert),
		live:        l,
	}
}

func (l *Live)cloneRequest(action string) (req *LiveRequest) {
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
	if (global.EmptyString == appName) {
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
	if (global.EmptyString == appName) {
		return errors.New("appName should not to be empty")
	}
	req := l.cloneRequest(ResumeLiveStreamAction)
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

// 录制视频
// -------------------------------------------------------------------------------

// AddLiveAppRecordConfig 配置APP录制，输出内容保存到OSS中
//
// https://help.aliyun.com/document_detail/35231.html?spm=5176.doc27193.6.221.xU2Kqb
func (l *Live) AddLiveAppRecordConfigWithApp(appName string, ossInfo OssInfo, resp interface{}) (err error) {
	if (global.EmptyString == appName || ossInfo.OssEndpoint == global.EmptyString || ossInfo.OssBucket == global.EmptyString || ossInfo.OssObjectPrefix == global.EmptyString) {
		return errors.New(" appName|ossEndpoint|ossBucket|ossObjectPrefix should not to be empty")
	}
	req := l.cloneRequest(AddLiveAppRecordConfigAction)
	req.AppName = appName
	req.SetArgs("OssEndpoint", ossInfo.OssEndpoint)
	req.SetArgs("OssBucket", ossInfo.OssBucket)
	req.SetArgs("OssObjectPrefix", ossInfo.OssObjectPrefix)
	err = l.rpc.Query(req, resp)
	return
}

// @see AddLiveAppRecordConfigWithApp
func (l *Live) AddLiveAppRecordConfig(ossInfo OssInfo, resp interface{}) (err error) {
	err = l.AddLiveAppRecordConfigWithApp(l.liveReq.AppName, ossInfo, resp)
	return
}

// DeleteLiveAppRecordConfigWithApp 解除录制配置
//
// https://help.aliyun.com/document_detail/35234.html?spm=5176.doc35239.6.223.4J6IYq
func (l *Live) DeleteLiveAppRecordConfigWithApp(appName string, resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveAppRecordConfigAction)
	req.AppName = appName
	err = l.rpc.Query(req, resp)
	return

}

// @see DeleteLiveAppRecordConfigWithApp
func (l *Live) DeleteLiveAppRecordConfig(resp interface{}) (err error) {
	err = l.DeleteLiveAppRecordConfigWithApp(l.liveReq.AppName, resp)
	return
}

// DescribeLiveAppRecordConfigWithApp 查询域名下指定App录制配置
//
// https://help.aliyun.com/document_detail/35239.html?spm=5176.doc35234.6.224.iCk6RL
func (l *Live) DescribeLiveAppRecordConfigWithApp(appName string, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveAppRecordConfigAction)
	req.AppName = appName
	err = l.rpc.Query(req, resp)
	return

}

// @see DescribeLiveAppRecordConfigWithApp
func (l *Live) DescribeLiveAppRecordConfig(resp interface{}) (err error) {
	err = l.DescribeLiveAppRecordConfigWithApp(l.liveReq.AppName, resp)
	return
}


//AppName	String	是	直播流所属应用名称
//StreamName	String	是	直播流名称
//OssEndpoint	String	否	oss endpoint，如：oss-cn-hangzhou.aliyuncs.com，详细请参照oss相关文档
//OssBucket	String	否	oss存储bucket名称
//OssObject	String	否	oss存储的录制文件名
//StartTime	String	是	开始时间，格式：2015-12-01T17:36:00Z
//EndTime	String	是	结束时间，格式：2015-12-01T17:36:00Z

// CreateLiveStreamRecordIndexFilesWithApp 创建录制索引文件
//
// https://help.aliyun.com/document_detail/35233.html?spm=5176.doc35239.6.225.dvRRZz
func (l *Live) CreateLiveStreamRecordIndexFilesWithApp(appName, streamName string, ossInfo OssInfo, startTime, endTime time.Time, resp interface{}) (err error) {
	if (global.EmptyString == appName || global.EmptyString == streamName ) {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(CreateLiveStreamRecordIndexFilesAction)

	req.AppName = appName
	req.SetArgs("StreamName", streamName)

	if (ossInfo.OssEndpoint != global.EmptyString) {
		req.SetArgs("OssEndpoint", ossInfo.OssEndpoint)
	}
	if (ossInfo.OssBucket != global.EmptyString) {
		req.SetArgs("OssBucket", ossInfo.OssBucket)
	}
	if (ossInfo.OssObject != global.EmptyString) {
		req.SetArgs("OssObject", ossInfo.OssObject)
	}
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// @see CreateLiveStreamRecordIndexFilesWithApp
func (l *Live) CreateLiveStreamRecordIndexFiles(streamName string, ossInfo OssInfo, startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.CreateLiveStreamRecordIndexFilesWithApp(l.liveReq.AppName, streamName, ossInfo, startTime, endTime, resp)
	return
}


// DescribeLiveRecordConfig 查询域名下所有App录制配置
//
// https://help.aliyun.com/document_detail/35235.html?spm=5176.doc35231.6.228.oRPQTW
func (l *Live) DescribeLiveRecordConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveRecordConfigAction)
	err = l.rpc.Query(req, resp)
	return
}


// DescribeLiveStreamRecordContentWithApp 查询某路直播流录制内容
//
// https://help.aliyun.com/document_detail/35236.html?spm=5176.doc35235.6.229.4IXXYR
func (l *Live) DescribeLiveStreamRecordContentWithApp(appName, streamName string, startTime, endTime time.Time, resp interface{}) (err error) {
	if (global.EmptyString == appName || global.EmptyString == streamName ) {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(DescribeLiveStreamRecordContentAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// @see DescribeLiveStreamRecordContentWithApp
func (l *Live) DescribeLiveStreamRecordContent(streamName string, startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.DescribeLiveStreamRecordContentWithApp(l.liveReq.AppName, streamName, startTime, endTime, resp)
	return
}


// DescribeLiveStreamRecordIndexFileWithApp 查询单个录制索引文件
//
// https://help.aliyun.com/document_detail/35237.html?spm=5176.doc35236.6.230.XnsJuD
func (l *Live) DescribeLiveStreamRecordIndexFileWithApp(appName, streamName, recordId string, resp interface{}) (err error) {
	if (global.EmptyString == appName || global.EmptyString == streamName ) {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(DescribeLiveStreamRecordIndexFileAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	if (recordId != global.EmptyString) {
		req.SetArgs("RecordId", recordId)
	}
	err = l.rpc.Query(req, resp)
	return
}

// @see DescribeLiveStreamRecordIndexFileWithApp
func (l *Live) DescribeLiveStreamRecordIndexFile(streamName, recordId string, resp interface{}) (err error) {
	err = l.DescribeLiveStreamRecordIndexFileWithApp(l.liveReq.AppName, streamName, recordId, resp)
	return
}


// DescribeLiveStreamRecordIndexFilesWithApp 查询录制索引文件
//
// https://help.aliyun.com/document_detail/35238.html?spm=5176.doc35237.6.231.L8KuPI
func (l *Live) DescribeLiveStreamRecordIndexFilesWithApp(appName, streamName string, startTime, endTime time.Time, resp interface{}) (err error) {
	if (global.EmptyString == appName || global.EmptyString == streamName ) {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(DescribeLiveStreamRecordIndexFilesAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// @see DescribeLiveStreamRecordIndexFilesWithApp
func (l *Live) DescribeLiveStreamRecordIndexFiles(streamName string, startTime, endTime time.Time, resp interface{}) (err error) {
	err = l.DescribeLiveStreamRecordIndexFilesWithApp(l.liveReq.AppName, streamName, startTime, endTime, resp)
	return
}


// DescribeLiveStreamsFrameRateAndBitRateDataWithApp 获取直播流的帧率和码率，支持基于域名和基于流的查询；
//
// https://help.aliyun.com/document_detail/35362.html?spm=5176.doc35238.6.232.wDsJeH
func (l *Live) DescribeLiveStreamsFrameRateAndBitRateDataWithApp(appName, streamName string, resp interface{}) (err error) {
	if (global.EmptyString == appName || global.EmptyString == streamName ) {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(DescribeLiveStreamsFrameRateAndBitRateDataAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	err = l.rpc.Query(req, resp)
	return
}

// @see DescribeLiveStreamRecordIndexFilesWithApp
func (l *Live) DescribeLiveStreamsFrameRateAndBitRateData(streamName string, resp interface{}) (err error) {
	err = l.DescribeLiveStreamsFrameRateAndBitRateDataWithApp(l.liveReq.AppName, streamName, resp)
	return
}

// GET 和 SET
// -------------------------------------------------------------------------------

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
