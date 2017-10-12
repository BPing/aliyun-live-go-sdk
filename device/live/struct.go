package live

import "github.com/BPing/aliyun-live-go-sdk/aliyun"

// ---------------------------------------------------------------------------------------------------------------------
// 直播基本信息
type LiveBase struct {
	DomainName string `json:"DomainName" xml:"DomainName"` //    加速域名
	AppName    string `json:"AppName" xml:"AppName"`       //	应用名称
}

// 流基本信息
type StreamBase struct {
	LiveBase
	StreamName string //	直播流名称
}

// 在线
// ---------------------------------------------------------------------------------------------------------------------
type OnlineInfoResponse struct {
	aliyun.Response
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

// 黑名单
// ---------------------------------------------------------------------------------------------------------------------
type StreamListResponse struct {
	aliyun.Response
	DomainName string     //流所属加速域名
	StreamUrls StreamUrls //	流完整URL地址
}

type StreamUrls struct {
	StreamUrl []string
}

// ---------------------------------------------------------------------------------------------------------------------

// 直播流的操作记录
type LiveStreamControlInfo struct {
	StreamName string //	流的名字
	ClientIP   string //	用户端的IP地址
	Action     string //	执行的操作名称
	TimeStamp  string //	操作执行的时间 UTC时间
}

// RTMP直播流的在线人数
type OnlineUserNum struct {
	TotalUserNumber int64                         //	所有流的用户数总和
	OnlineUserInfo  []LiveStreamOnlineUserNumInfo //每条直播流的用户数信息
}

type LiveStreamOnlineUserNumInfo struct {
	StreamUrl  string //直播流的URL
	UserNumber int64  //	直播流的在线人数
}

// ---------------------------------------------------------------------------------------------------------------------

// 对象存储oss信息
type OssInfo struct {
	OssEndpoint     string //	oss endpoint，如：oss-cn-hangzhou.aliyuncs.com，详细请参照oss相关文档
	OssBucket       string //	oss存储bucket名称
	OssObject       string //	oss存储的录制文件名
	OssObjectPrefix string //	oss存储文件名，支持变量匹配，包含{AppName}、{StreamName}、{UnixTimestamp}、{Sequence}，如：record/live/{StreamName}/{UnixTimestamp}_{Sequence}
}

// ---------------------------------------------------------------------------------------------------------------------

// 录制配置
type LiveAppRecord struct {
	LiveBase
	OssInfo
	CreateTime string
}

// 录制
// ---------------------------------------------------------------------------------------------------------------------

type FrameRateAndBitRateInfosResponse struct {
	aliyun.Response
	FrameRateAndBitRateInfos FrameRateAndBitRateInfos
}

type FrameRateAndBitRateInfos struct {
	FrameRateAndBitRateInfo []FrameRateAndBitRateInfo
}

// 各直播流的帧率/码率信息
type FrameRateAndBitRateInfo struct {
	StreamUrl      string // 直播流的URL
	VideoFrameRate int    // 直播流的视频帧率
	AudioFrameRate int    // 直播流的音频帧率
	BitRate        int    // 直播流的码率
}

// ---------------------------------------------------------------------------------------------------------------------

// 录制配置列表
type RecordIndexInfoListResponse struct {
	aliyun.Response
	RecordIndexInfoList RecordIndexInfoList
}

type RecordInfoListResponse struct {
	aliyun.Response
	RecordInfoList RecordInfoList
}

// 录制配置单个
type RecordIndexInfoResponse struct {
	aliyun.Response
	RecordIndexInfo RecordIndexInfo
}

type RecordInfoResponse struct {
	aliyun.Response
	RecordInfo RecordInfo
}

type RecordIndexInfoList struct {
	RecordIndexInfo []RecordIndexInfo
}

type RecordInfoList struct {
	RecordIndexInfo []RecordInfo
}

type RecordInfo struct {
	RecordIndexInfo
}

//RecordId	String	索引文件Id
//RecordUrl	String	索引文件地址
//DomainName	String	流所属加速域名
//AppName	String	流所属应用名称
//StreamName	String	直播流名称
//OssEndpoint	String	oss endpoit
//OssBucket	String	oss存储bucket名称
//OssObject	String	oss存储的录制文件名
//CreateTime	String	创建时间
//StartTime	String	开始时间，格式：2015-12-01T17:36:00Z
//EndTime	String	结束时间，格式：2015-12-01T17:36:00Z
//Duration	String	录制时长，单位：秒
//Height	String	视频高
//Width	String	视频宽
//CreateTime	String	创建时间

// 录制配置信息
type RecordIndexInfo struct {
	StreamBase
	OssInfo          //oss存储
	RecordId  string //	索引文件Id
	RecordUrl string //	索引文件地址
	//OssObject  string //	oss存储的录制文件名 包含在OssInfo
	Height     string //	视频高
	Width      string //	视频宽
	CreateTime string //    创建时间
	StartTime  string //	开始时间，格式：2015-12-01T17:36:00Z
	EndTime    string //	结束时间，格式：2015-12-01T17:36:00Z
	Duration   string //	录制时长
}

// ---------------------------------------------------------------------------------------------------------------------

type RecordContentInfoListResponse struct {
	aliyun.Response
	RecordContentInfoList RecordContentInfoList
}

// 录制内容列表
type RecordContentInfoList struct {
	RecordContentInfo []RecordContentInfo
}

type RecordContentInfo struct {
	OssInfo
	StartTime string `json:"StartTime" xml:"StartTime"` //	开始时间，格式：2015-12-01T17:36:00Z
	EndTime   string `json:"EndTime" xml:"EndTime"`     //	结束时间，格式：2015-12-01T17:36:00Z
	Duration  string `json:"Duration" xml:"Duration"`   //	录制时长
}

// ---------------------------------------------------------------------------------------------------------------------

// 截图配置(参数)
type SnapshotConfig struct {
	OssInfo
	TimeInterval       int    `json:"TimeInterval" xml:"TimeInterval"`             // 截图周期
	OverwriteOssObject string `json:"OverwriteOssObject" xml:"OverwriteOssObject"` //oss存储文件名
	SequenceOssObject  string `json:"SequenceOssObject" xml:"SequenceOssObject"`   //oss存储文件名
}

// 查询域名下的截图配置返回结构
type LiveSnapshotConfigResponse struct {
	aliyun.Response
	LiveSnapshotParam
	LiveStreamSnapshotConfigList struct {
		LiveStreamSnapshotConfig []LiveStreamSnapshotConfig `json:"LiveStreamSnapshotConfig" xml:"LiveStreamSnapshotConfig"`
	} `json:"LiveStreamSnapshotConfigList" xml:"LiveStreamSnapshotConfigList"` //	截图配置
	TotalPage int `json:"TotalPage" xml:"TotalPage"`                           //	总页数
	TotalNum  int `json:"TotalNum" xml:"TotalNum"`                             //	符合条件的总个数
}

// 查询域名下的截图配置参数
type LiveSnapshotParam struct {
	PageNum  int    `json:"PageNum" xml:"PageNum"`   //    分页的页码
	PageSize int    `json:"PageSize" xml:"PageSize"` //	每页大小
	Order    string `json:"Order" xml:"Order"`
}

type LiveStreamSnapshotConfig struct {
	LiveBase
	SnapshotConfig
	CreateTime string `json:"CreateTime" xml:"CreateTime"` //创建时间
}

// 查询截图信息
type StreamSnapshotInfoResponse struct {
	aliyun.Response
	LiveStreamSnapshotInfoList struct {
		StreamSnapshotInfo []StreamSnapshotInfo `json:"StreamSnapshotInfo" xml:"StreamSnapshotInfo"`
	} `json:"LiveStreamSnapshotInfoList" xml:"LiveStreamSnapshotInfoList"` //截图内容列表，没有则返回空数组
	NextStartTime string `json:"NextStartTime" xml:"NextStartTime"`        //
}

// 单个截图数据类型
type StreamSnapshotInfo struct {
	OssInfo
	CreateTime string `json:"CreateTime" xml:"CreateTime"` //截图产生时间，格式：2015-12-01T17:36:00Z
}

// ---------------------------------------------------------------------------------------------------------------------
// 转码配置信息返回结构体
type StreamTranscodeInfoResponse struct {
	aliyun.Response
	DomainTranscodeList struct {
		DomainTranscodeInfo []DomainTranscodeInfo `json:"DomainTranscodeInfo" xml:"DomainTranscodeInfo"`
	} `json:"DomainTranscodeList" xml:"DomainTranscodeList"` //转码配置信息
}

type DomainTranscodeInfo struct {
	TranscodeName     string `json:"TranscodeName" xml:"TranscodeName"`         //	播放域名
	TranscodeApp      string `json:"TranscodeApp" xml:"TranscodeApp"`           //	应用名称
	TranscodeId       string `json:"TranscodeId" xml:"TranscodeId"`             //	数据库ID
	TranscodeTemplate string `json:"TranscodeTemplate" xml:"TranscodeTemplate"` //	转码模版
	TranscodeSnapshot string `json:"TranscodeSnapshot" xml:"TranscodeSnapshot"` //	是否实施截图
	TranscodeRecord   string `json:"TranscodeRecord" xml:"TranscodeRecord"`     //	是否实施录制
}

// ---------------------------------------------------------------------------------------------------------------------

// 混流信息返回结构体
type MixStreamsInfoResponse struct {
	aliyun.Response
	MixStreamsInfoList MixStreamsInfoList `json:"MixStreamsInfoList" xml:"MixStreamsInfoList"`
}

type MixStreamsInfoList struct {
	MixStreamsInfo []MixStreamsInfo `json:"MixStreamsInfo" xml:"MixStreamsInfo"`
}

type MixStreamsInfo struct {
	StreamBase
}

// ---------------------------------------------------------------------------------------------------------------------

// 拉流
// https://help.aliyun.com/document_detail/57733.html?spm=5176.doc57735.6.658.58h5BX
type PullStreamConfigResponse struct {
	aliyun.Response
	LiveAppRecordList AppRecordList `json:"LiveAppRecordList" xml:"LiveAppRecordList"`
}

type AppRecordList struct {
	LiveAppRecord []AppRecord `json:"LiveAppRecord" xml:"LiveAppRecord"`
}

type AppRecord struct {
	StreamBase
	SourceUrl string `json:"SourceUrl" xml:"SourceUrl"` // 	拉流源站
	StartTime string `json:"StartTime" xml:"StartTime"`
	EndTime   string `json:"EndTime" xml:"EndTime"`
}

// ---------------------------------------------------------------------------------------------------------------------

// 状态通知
type NotifyUrlConfigResponse struct {
	aliyun.Response
	NotifyUrlConfig struct {
		DomainName string `json:"DomainName" xml:"DomainName"` //    加速域名
		NotifyUrl  string `json:"NotifyUrl" xml:"NotifyUrl"`   //    回调地址
	}
}

// ---------------------------------------------------------------------------------------------------------------------

type MixNotifyConfigResponse struct {
	aliyun.Response
	NotifyUrl string `json:"NotifyUrl" xml:"NotifyUrl"` //    当前域名下连麦回调通知URL
}
