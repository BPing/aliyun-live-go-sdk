package live

import (
	"github.com/BPing/aliyun-live-go-sdk/util"
	"fmt"
	"time"
)

const (
	DefualtStreamTimeout = time.Hour * 2
)

// StreamCredentials 流的地址信息的签名凭证
type StreamCredentials struct {
	PrivateKey string //privatekey
	Timeout    time.Duration
}

func (s *StreamCredentials) Clone() *StreamCredentials {
	new_obj := (*s)
	return &new_obj
}

// NewStreamCredentials 新建凭证
func NewStreamCredentials(privateKey string, timeout time.Duration) *StreamCredentials {
	return &StreamCredentials{
		PrivateKey: privateKey,
		Timeout:    timeout,
	}
}

//
// Stream 直播流
// 包括推流地址，流的播放地址，流的在线状态和在线人数等基本信息
//
// 推流地址：rtmp://video-center.alivecdn.com/appName/streamName?vhost=CDN
// video-center.alivecdn.com是直播中心服务器，允许自定义，
// 例如您的域名是live.yourcompany.com，可以设置DNS，将您的域名CNAME指向video-center.alivecdn.com即可；
// AppName是应用名称，支持自定义，可以更改；
// StreamName是流名称，支持自定义，可以更改；
// vhost参数是最终在边缘节点播放的域名，即你的加速域名。
//
// @author cbping
type Stream struct {
	live            *Live
	domainName      string
	videoCenterDns  string
	appName         string //app-name
	StreamName      string //video-name

	streamCert      *StreamCredentials
	signOn          bool   //是否启用签名

	expireTimestamp int64  //过期时间戳
	authKey         string //auth_key

	rtmpPublishUrl  string
	rtmpLiveUrls    string
	hlsLiveUrls     string
	httpFlvLiveUrls string
}

// -------------------------------------------------------------------------------
func (s *Stream) InitOrUpdate() {
	s.RtmpPublishUrl()
}

//在线状态
// -------------------------------------------------------------------------------
func (s *Stream) Online() bool {
	resp := OnlineInfoResponse{}
	s.live.StreamOnlineUserNumWithApp(s.appName, s.StreamName, &resp)
	if len(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo) > 0 {
		return true
	}
	return false
}

// OnlineUserNum 获取在线人数
// 如果流不是在线状态，返回0
// -------------------------------------------------------------------------------
func (s *Stream) OnlineUserNum() (num int64) {
	resp := OnlineInfoResponse{}
	s.live.StreamOnlineUserNumWithApp(s.appName, s.StreamName, &resp)
	num = resp.TotalUserNumber
	return
}
// Blocked 是否在黑名单中
// -------------------------------------------------------------------------------
func (s *Stream) Blocked() (bool) {
	resp := StreamListResponse{}
	s.live.StreamsBlockList(&resp)
	for _, val := range resp.StreamUrls.StreamUrl {
		//遍历判断此流是否存在黑名单中
		if (val == s.baseUrl()) {
			return true
		}
	}
	return false
}

// ForbidPush 禁止推流
// -------------------------------------------------------------------------------
func (s *Stream) ForbidPush() (err error) {
	err = s.live.ForbidLiveStreamWithPublisherWithApp(s.appName, s.StreamName, nil, nil)
	return
}

// ResumePush 恢复推流
// -------------------------------------------------------------------------------
func (s *Stream) ResumePush() (err error) {
	err = s.live.ResumeLiveStreamWithPublisherWithApp(s.appName, s.StreamName, nil)
	return
}

// -------------------------------------------------------------------------------
func (s *Stream) baseUrl() (url string) {
	url = fmt.Sprintf("%s/%s/%s", s.domainName, s.appName, s.StreamName)
	return
}

func (s *Stream) basePushUrl() (url string) {
	url = fmt.Sprintf("%s/%s/%s", s.videoCenterDns, s.appName, s.StreamName)
	return
}

// 启动鉴权才可以使用签名
func (s *Stream) sign() {
	s.authKey, s.expireTimestamp = util.CreateSignatureForStreamUrlWithA("/" + s.appName + "/" + s.StreamName, "0", "0", s.streamCert.PrivateKey, s.streamCert.Timeout)
	return
}

// generateAllUrls 生成所有基本地址链接
func (s *Stream) generateAllUrls() {
	if s.signOn && nil != s.streamCert {
		s.sign()
		//rtmp://cdn/app-name/video-name?vhost=cdn&auth_key=timestamp-rand-uid-sign
		s.rtmpPublishUrl = fmt.Sprintf("rtmp://%s?vhost=%s&auth_key=%s", s.basePushUrl(), s.domainName, s.authKey)
		//rtmp://cdn/app-name/video-name?auth_key=timestamp-rand-uid-sign
		s.rtmpLiveUrls = fmt.Sprintf("rtmp://%s?auth_key=%s", s.baseUrl(), s.authKey)
		//http://cdn/app-name/video-name.m3u8?auth_key=timestamp-rand-uid-sign
		s.hlsLiveUrls = fmt.Sprintf("http://%s.m3u8?auth_key=%s", s.baseUrl(), s.authKey)
		//http://cdn/app-name/video-name.flv?auth_key=timestamp-rand-uid-sign
		s.httpFlvLiveUrls = fmt.Sprintf("http://%s.flv?auth_key=%s", s.baseUrl(), s.authKey)
	} else {
		//rtmp://cdn/app-name/video-name?vhost=cdn
		s.rtmpPublishUrl = fmt.Sprintf("rtmp://%s?vhost=%s", s.baseUrl(), s.domainName)
		//rtmp://cdn/app-name/video-name
		s.rtmpLiveUrls = fmt.Sprintf("rtmp://%s", s.baseUrl())
		//http://cdn/app-name/video-name.m3u8
		s.hlsLiveUrls = fmt.Sprintf("http://%s.m3u8", s.baseUrl())
		//http://cdn/app-name/video-name.flv
		s.httpFlvLiveUrls = fmt.Sprintf("http://%s.flv", s.baseUrl())
	}
	return
}

// RTMP 推流地址
// 如果开启了直播流鉴权，签名失效后，会重新生成新的有效的推流地址
// -------------------------------------------------------------------------------
func (s *Stream) RtmpPublishUrl() (url string) {
	if "" == s.rtmpPublishUrl || "" == s.rtmpLiveUrls || (s.signOn && s.isExpired()) {
		//RTMP推流地址或者RTMP播放流地址为空或者签名失效
		s.generateAllUrls()
	}
	url = s.rtmpPublishUrl
	return
}

// authKey签名是否过期
// -------------------------------------------------------------------------------
func (s *Stream) isExpired() bool {
	return time.Now().Unix() > s.expireTimestamp
}

// RTMP 直播播放地址
// --------------------------------------------------------------------------------
func (s *Stream) RtmpLiveUrls() (url string) {
	url = s.rtmpLiveUrls
	return
}

// HLS 直播播放地址
// --------------------------------------------------------------------------------
func (s *Stream) HlsLiveUrls() (url string) {
	url = s.hlsLiveUrls
	return
}

// FLV 直播播放地址
// --------------------------------------------------------------------------------
func (s *Stream) HttpFlvLiveUrls() (url string) {
	url = s.httpFlvLiveUrls
	return
}

// String 返回流的基本信息
func (s *Stream) String() (str string) {
	return fmt.Sprintf(
		"RtmpPublishUrl:%s\n RtmpLiveUrls:%s\n  HlsLiveUrls:%s\n  HttpFlvLiveUrls:%s\n  OnlineUserNum:%d \n",
		s.rtmpPublishUrl,
		s.rtmpLiveUrls,
		s.hlsLiveUrls,
		s.httpFlvLiveUrls,
		s.OnlineUserNum(),
	)
}


// 录制 ----------------------------------------------------------------------------------------------------------------

// 创建直播流录制索引文件
func (s *Stream) CreateRecordIndexFiles(ossInfo OssInfo, startTime, endTime time.Time) (info RecordInfo, err error) {
	resp := &RecordInfoResponse{}
	err = s.live.CreateLiveStreamRecordIndexFilesWithApp(s.appName, s.StreamName, ossInfo, startTime, endTime, resp)
	if (err == nil) {
		info = resp.RecordInfo
	}
	return
}

// 查询某路直播流录制内容
func (s *Stream) RecordContent(startTime, endTime time.Time) (list RecordContentInfoList, err error) {
	resp := &RecordContentInfoListResponse{}
	err = s.live.DescribeLiveStreamRecordContentWithApp(s.appName, s.StreamName, startTime, endTime, resp)
	if (err == nil) {
		list = resp.RecordContentInfoList
	}
	return
}

// 查询直播流录制索引文件
func (s *Stream) RecordIndexFiles(startTime, endTime time.Time) (list RecordIndexInfoList, err error) {
	resp := &RecordIndexInfoListResponse{}
	err = s.live.DescribeLiveStreamRecordIndexFilesWithApp(s.appName, s.StreamName, startTime, endTime, resp)
	if (err == nil) {
		list = resp.RecordIndexInfoList
	}
	return
}

// 查询直播流单个录制索引文件
func (s *Stream) RecordIndexFile(recordId string) (info RecordIndexInfo, err error) {
	resp := &RecordIndexInfoResponse{}
	err = s.live.DescribeLiveStreamRecordIndexFileWithApp(s.appName, s.StreamName, recordId, resp)
	if (err == nil) {
		info = resp.RecordIndexInfo
	}
	return
}

// 获取直播流的帧率和码率
func (s *Stream) FrameRateAndBitRateData() (info FrameRateAndBitRateInfos, err error) {
	resp := &FrameRateAndBitRateInfosResponse{}
	err = s.live.DescribeLiveStreamsFrameRateAndBitRateDataWithApp(s.appName, s.StreamName, resp)
	if (err == nil) {
		info = resp.FrameRateAndBitRateInfos
	}
	return
}