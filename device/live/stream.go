package live

import (
	"time"
	"fmt"
	"aliyun-live-go-sdk/util"
)

const (
	DefualtStreamTimeout = time.Hour * 2
)

//签名凭证
type StreamCredentials struct {
	PrivateKey string //privatekey
	Timeout    time.Duration
}

func (s *StreamCredentials)Clone() *StreamCredentials {
	new_obj := (*s)
	return &new_obj
}

func NewStreamCredentials(privateKey string, timeout time.Duration) *StreamCredentials {
	return &StreamCredentials{
		PrivateKey:privateKey,
		Timeout:timeout,
	}
}

//
// 直播流
//
type Stream struct {
	live            *Live
	DomainName      string
	AppName         string //app-name
	StreamName      string //video-name

	credentials     *StreamCredentials
	signOn          bool   //是否启用签名

	expireTimestamp int64
	authKey         string //auth_key

	rtmpPublishUrl  string
	rtmpLiveUrls    string
	hlsLiveUrls     string
	httpFlvLiveUrls string
}


// -------------------------------------------------------------------------------
func (s *Stream)InitOrUpdate() {
	s.RtmpPublishUrl()
}

//在线状态
// -------------------------------------------------------------------------------
func (s *Stream)Online() bool {
	resp := OnlineInfoResponse{}
	s.live.StreamOnlineUserNum(s.StreamName, &resp)
	if (len(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo) > 0) {
		return true
	}
	return false
}


//获取在线人数
// 如果流不是在线状态，返回0
// -------------------------------------------------------------------------------
func (s *Stream)OnlineUserNum() (num int64) {
	resp := OnlineInfoResponse{}
	s.live.StreamOnlineUserNum(s.StreamName, &resp)
	num = resp.TotalUserNumber
	return
}


// 禁止和恢复推流
// -------------------------------------------------------------------------------
func (s *Stream)ForbidPush() (err error) {
	err = s.live.ForbidLiveStreamWithPublisher(s.StreamName, nil, nil)
	return
}

func (s *Stream)ResumePush() (err error) {
	err = s.live.ResumeLiveStreamWithPublisher(s.StreamName, nil)
	return
}

// -------------------------------------------------------------------------------
func (s *Stream)baseUrl() (url string) {
	url = fmt.Sprintf("%s/%s/%s", s.DomainName, s.AppName, s.StreamName)
	return
}
//
// 启动鉴权才可以使用签名
func (s *Stream)sign() {
	s.authKey, s.expireTimestamp = util.CreateSignatureForStreamUrlWithA("/" + s.AppName + "/" + s.StreamName, "0", "0", s.credentials.PrivateKey, s.credentials.Timeout)
	return
}

func (s *Stream)generateAllUrls() {
	if (s.signOn&&nil != s.credentials) {
		s.sign()
		//rtmp://cdn/app-name/video-name?vhost=cdn&auth_key=timestamp-rand-uid-sign
		s.rtmpPublishUrl = fmt.Sprintf("rtmp://%s?vhost=%s&auth_key=%s", s.baseUrl(), s.DomainName, s.authKey)
		//rtmp://cdn/app-name/video-name?auth_key=timestamp-rand-uid-sign
		s.rtmpLiveUrls = fmt.Sprintf("rtmp://%s?auth_key=%s", s.baseUrl(), s.authKey)
		//http://cdn/app-name/video-name.m3u8?auth_key=timestamp-rand-uid-sign
		s.hlsLiveUrls = fmt.Sprintf("http://%s.m3u8?auth_key=%s", s.baseUrl(), s.authKey)
		//http://cdn/app-name/video-name.flv?auth_key=timestamp-rand-uid-sign
		s.httpFlvLiveUrls = fmt.Sprintf("http://%s.flv?auth_key=%s", s.baseUrl(), s.authKey)
	}else {
		//rtmp://cdn/app-name/video-name?vhost=cdn
		s.rtmpPublishUrl = fmt.Sprintf("rtmp://%s?vhost=%s", s.baseUrl(), s.DomainName)
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
func (s *Stream)RtmpPublishUrl() (url string) {
	if ("" == s.rtmpPublishUrl || "" == s.rtmpLiveUrls || (s.signOn&&s.isExpired())) {
		//RTMP推流地址或者RTMP播放流地址为空或者签名失效
		s.generateAllUrls()
	}
	url = s.rtmpPublishUrl
	return
}

// authKey签名是否过期
// -------------------------------------------------------------------------------
func (s *Stream)isExpired() bool {
	return time.Now().Unix() > s.expireTimestamp
}

// RTMP 直播播放地址
// --------------------------------------------------------------------------------

func (s *Stream)RtmpLiveUrls() (url string) {
	url = s.rtmpLiveUrls
	return
}

// HLS 直播播放地址
// --------------------------------------------------------------------------------

func (s *Stream)HlsLiveUrls() (url string) {
	url = s.hlsLiveUrls
	return
}

// FLV 直播播放地址
// --------------------------------------------------------------------------------

func (s *Stream)HttpFlvLiveUrls() (url string) {
	url = s.httpFlvLiveUrls
	return
}

func (s *Stream)String() (str string) {
	return fmt.Sprintf(
		"RtmpPublishUrl:%s\n RtmpLiveUrls:%s\n  HlsLiveUrls:%s\n  HttpFlvLiveUrls:%s\n  OnlineUserNum:%d \n",
		s.rtmpPublishUrl,
		s.rtmpLiveUrls,
		s.hlsLiveUrls,
		s.httpFlvLiveUrls,
		s.OnlineUserNum(),
	)
}