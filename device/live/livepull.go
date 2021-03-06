package live

import (
	"github.com/BPing/aliyun-live-go-sdk/util"
	"time"
)

// 直播拉流
// -------------------------------------------------------------------------------

// AddLivePullStreamInfoConfig 添加拉流信息
//
// https://help.aliyun.com/document_detail/57734.html?spm=5176.doc57733.6.656.YS8uOK
func (l *Live) AddLivePullStreamInfoConfig(streamName, sourceUrl string, startTime, endTime time.Time, resp interface{}) (err error) {
	req := l.cloneRequest(AddLivePullStreamInfoConfigAction)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("SourceUrl", sourceUrl)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLivePullStreamInfoConfig 删除拉流信息
//
// https://help.aliyun.com/document_detail/57735.html?spm=5176.doc57734.6.657.wRW6P7
func (l *Live) DeleteLivePullStreamInfoConfig(streamName string, resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLivePullStreamInfoConfigAction)
	req.SetArgs("StreamName", streamName)
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLivePullStreamConfig 查询域名下拉流配置信息
//
// https://help.aliyun.com/document_detail/57733.html?spm=5176.doc57735.6.658.8TrTGR
func (l *Live) DescribeLivePullStreamConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLivePullStreamConfigAction)
	err = l.rpc.Query(req, resp)
	return
}
