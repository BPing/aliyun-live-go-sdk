package live

import (
	"time"
	"github.com/BPing/aliyun-live-go-sdk/util"
)

// 资源监控
// -------------------------------------------------------------------------------

// StreamOnlineUserNum 获取在线人数
// @appname 应用名 为空时，忽略此参数
// @link https://help.aliyun.com/document_detail/27195.html?spm=0.0.0.0.n6eAJJ
func (l *Live) StreamOnlineUserNum(streamName string, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamOnlineUserNumAction)
	if "" != streamName {
		req.SetArgs("StreamName", streamName)
	}
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveDomainBpsData 查询直播域名的网络带宽监控数据
//          单位：bit/second。
//         1、不指定StartTime和EndTime时，默认读取过去24小时的数据，同时支持按指定的起止时间查询，两者需要同时指定。
//         2、支持批量域名查询，多个域名用逗号（半角）分隔。
//         3、最多可获取最近90天的数据。
// @link https://help.aliyun.com/document_detail/67406.html?spm=a2c4g.11186623.6.731.1Nki94
func (l *Live) DescribeLiveDomainBpsData(startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveDomainBpsDataAction)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveDomainRecordData 查询直播域名录制时长数据
//
// @link https://help.aliyun.com/document_detail/68943.html?spm=a2c4g.11186623.6.732.idI3mC
func (l *Live) DescribeLiveDomainRecordData(startTime, endTime time.Time, recordType RecordType, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveDomainRecordDataAction)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	if recordType != NilRecordType {
		req.SetArgs("RecordType", string(recordType))
	}
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveDomainSnapshotData 查询直播域名截图张数数据
//
// @link https://help.aliyun.com/document_detail/68944.html?spm=a2c4g.11186623.6.733.NuYuod
func (l *Live) DescribeLiveDomainSnapshotData(startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveDomainSnapshotDataAction)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveDomainTrafficData 查询直播域名网络流量监控数据，单位：byte。
//
// @link https://help.aliyun.com/document_detail/67409.html?spm=a2c4g.11186623.6.734.kgaNJn
func (l *Live) DescribeLiveDomainTrafficData(startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveDomainTrafficDataAction)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveDomainTranscodeData 查询直播域名转码时长数据，
//
// @link https://help.aliyun.com/document_detail/68942.html?spm=a2c4g.11186623.6.735.v9TmHH
func (l *Live) DescribeLiveDomainTranscodeData(startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveDomainTranscodeDataAction)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveStreamHistoryUserNum 查询直播流历史在线人数，
//
// @link https://help.aliyun.com/document_detail/61267.html?spm=a2c4g.11186623.6.736.znw3Xi
func (l *Live) DescribeLiveStreamHistoryUserNum(streamName string, startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamHistoryUserNumAction)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}
