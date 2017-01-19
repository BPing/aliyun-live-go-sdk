package live

import (
	"errors"
	"github.com/BPing/aliyun-live-go-sdk/util"
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"time"
)


// 录制视频
// -------------------------------------------------------------------------------

// AddLiveAppRecordConfig 配置APP录制，输出内容保存到OSS中
//
// https://help.aliyun.com/document_detail/35231.html?spm=5176.doc27193.6.221.xU2Kqb
func (l *Live) AddLiveAppRecordConfigWithApp(appName string, ossInfo OssInfo, resp interface{}) (err error) {
	if global.EmptyString == appName || ossInfo.OssEndpoint == global.EmptyString || ossInfo.OssBucket == global.EmptyString || ossInfo.OssObjectPrefix == global.EmptyString {
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
	if global.EmptyString == appName || global.EmptyString == streamName {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(CreateLiveStreamRecordIndexFilesAction)

	req.AppName = appName
	req.SetArgs("StreamName", streamName)

	if ossInfo.OssEndpoint != global.EmptyString {
		req.SetArgs("OssEndpoint", ossInfo.OssEndpoint)
	}
	if ossInfo.OssBucket != global.EmptyString {
		req.SetArgs("OssBucket", ossInfo.OssBucket)
	}
	if ossInfo.OssObject != global.EmptyString {
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
	if global.EmptyString == appName || global.EmptyString == streamName {
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
	if global.EmptyString == appName || global.EmptyString == streamName {
		return errors.New(" appName|streamName should not to be empty")
	}
	req := l.cloneRequest(DescribeLiveStreamRecordIndexFileAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	if recordId != global.EmptyString {
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
	if global.EmptyString == appName || global.EmptyString == streamName {
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
	if global.EmptyString == appName || global.EmptyString == streamName {
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
