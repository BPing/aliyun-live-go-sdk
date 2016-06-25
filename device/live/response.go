package live

import "aliyun-live-go-sdk/client"

// 在线
// -------------------------------------------------------------------------------

type OnlineInfoResponse struct {
	client.Response
	OnlineUserInfo  OnlineUserInfo
	TotalUserNumber int64
}

type OnlineUserInfo struct {
	LiveStreamOnlineUserNumInfo []LiveStreamOnlineInfo
}

type LiveStreamOnlineInfo struct {
	DomainName  string //
	AppName     string //
	StreamName  string //
	PublishTime string //
	PublishUrl  string //
}

//
// -------------------------------------------------------------------------------
type StreamList struct {
	DomainName string   //流所属加速域名
	StreamUrls []string //	流完整URL地址
}

//直播流的操作记录
type LiveStreamControlInfo struct {
	StreamName string //	流的名字
	ClientIP   string //	用户端的IP地址
	Action     string //	执行的操作名称
	TimeStamp  string //	操作执行的时间 UTC时间
}

//RTMP直播流的在线人数
type OnlineUserNum struct {
	TotalUserNumber int64                         //	所有流的用户数总和
	OnlineUserInfo  []LiveStreamOnlineUserNumInfo //每条直播流的用户数信息
}

type LiveStreamOnlineUserNumInfo struct {
	StreamUrl  string //直播流的URL
	UserNumber int64  //	直播流的在线人数
}