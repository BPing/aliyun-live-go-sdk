package live

import (
	"errors"
	"fmt"
	"github.com/BPing/aliyun-live-go-sdk/util"
	"strings"
	"time"
)

// 截图 处理
// ---------------------------------------------------------------------------------------------------------------------
// AddLiveAppSnapshotConfigWithApp 添加截图配置
// {@link https://help.aliyun.com/document_detail/44718.html?spm=5176.doc44720.6.690.d4pYiq}
func (l *Live) AddLiveAppSnapshotConfigWithApp(appName string, config SnapshotConfig, resp interface{}) (err error) {
	req := l.cloneRequest(AddLiveAppSnapshotConfigAction)
	req.AppName = appName
	req.SetArgs("TimeInterval", fmt.Sprintf("%d", config.TimeInterval))
	req.SetArgs("OssEndpoint", config.OssEndpoint)
	req.SetArgs("OssBucket", config.OssBucket)
	req.SetArgs("OverwriteOssObject", config.OverwriteOssObject)
	req.SetArgs("SequenceOssObject", config.SequenceOssObject)
	err = l.rpc.Query(req, resp)
	return
}

// @see AddLiveAppSnapshotConfigWithApp
func (l *Live) AddLiveAppSnapshotConfig(config SnapshotConfig, resp interface{}) (err error) {
	err = l.AddLiveAppSnapshotConfigWithApp(l.liveReq.AppName, config, resp)
	return
}

// UpdateLiveAppSnapshotConfigWithApp 更新截图配置
// {@link https://help.aliyun.com/document_detail/44720.html?spm=5176.doc44722.6.700.8aPqov}
func (l *Live) UpdateLiveAppSnapshotConfigWithApp(appName string, config SnapshotConfig, resp interface{}) (err error) {
	req := l.cloneRequest(UpdateLiveAppSnapshotConfigAction)
	req.AppName = appName
	req.SetArgs("TimeInterval", fmt.Sprintf("%d", config.TimeInterval))
	req.SetArgs("OssEndpoint", config.OssEndpoint)
	req.SetArgs("OssBucket", config.OssBucket)
	req.SetArgs("OverwriteOssObject", config.OverwriteOssObject)
	req.SetArgs("SequenceOssObject", config.SequenceOssObject)
	err = l.rpc.Query(req, resp)
	return
}

// @see UpdateLiveAppSnapshotConfigWithApp
func (l *Live) UpdateLiveAppSnapshotConfig(config SnapshotConfig, resp interface{}) (err error) {
	err = l.UpdateLiveAppSnapshotConfigWithApp(l.liveReq.AppName, config, resp)
	return
}

// DeleteLiveAppSnapshotConfigWithApp 删除截图配置
// {@link https://help.aliyun.com/document_detail/44719.html?spm=5176.doc44718.6.692.pCd2yM}
func (l *Live) DeleteLiveAppSnapshotConfigWithApp(appName string, resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveAppSnapshotConfigAction)
	req.AppName = appName
	err = l.rpc.Query(req, resp)
	return
}

// @see DeleteLiveAppSnapshotConfigWithApp
func (l *Live) DeleteLiveAppSnapshotConfig(resp interface{}) (err error) {
	err = l.DeleteLiveAppSnapshotConfigWithApp(l.liveReq.AppName, resp)
	return
}

// LiveSnapshotConfigWithApp 查询域名截图配置
// pageNum	int	否	分页的页码，默认值：1
// pageSize	int	否	每页大小，默认值：10，范围：5~30
// order	string	否	排序，asc：升序，desc：降序，默认：asc
// {@link https://help.aliyun.com/document_detail/44721.html?spm=5176.doc44719.6.694.QyVLo8}
func (l *Live) LiveSnapshotConfigWithApp(appName string, param LiveSnapshotParam, resp interface{}) (err error) {
	param.Order = strings.ToLower(param.Order)
	if param.Order != "asc" && param.Order != "desc" {
		return errors.New("order:'asc' or 'desc'")
	}
	if param.PageNum <= 0 {
		param.PageNum = 1
	}
	if param.PageSize < 5 {
		param.PageSize = 10
	}
	req := l.cloneRequest(DescribeLiveSnapshotConfigAction)
	req.AppName = appName
	req.SetArgs("PageNum", fmt.Sprintf("%d", param.PageNum))
	req.SetArgs("PageSize", fmt.Sprintf("%d", param.PageSize))
	req.SetArgs("Order", param.Order)
	err = l.rpc.Query(req, resp)
	return
}

// @see LiveSnapshotConfigWithApp
func (l *Live) LiveSnapshotConfig(param LiveSnapshotParam, resp interface{}) (err error) {
	err = l.LiveSnapshotConfigWithApp(l.liveReq.AppName, param, resp)
	return
}

// LiveStreamSnapshotInfoWithApp 查询截图信息
// streamName	string	是	直播流名称
// startTime	time.Time	是	开始时间
// endTime	time.Time	是	结束时间
// limit	int	否	一次调用获取的数量，范围1~100，默认值：10
// {@link https://help.aliyun.com/document_detail/44722.html?spm=5176.doc44721.6.696.Xcp6VD}
func (l *Live) LiveStreamSnapshotInfoWithApp(appName string, streamName string, startTime, endTime time.Time, limit int, resp interface{}) (err error) {
	if limit < 1 || limit > 100 {
		limit = 10
	}
	req := l.cloneRequest(DescribeLiveStreamSnapshotInfoAction)
	req.AppName = appName
	req.SetArgs("StreamName", streamName)
	req.SetArgs("StartTime", util.GetISO8601TimeStamp(startTime))
	req.SetArgs("EndTime", util.GetISO8601TimeStamp(endTime))
	req.SetArgs("Limit", fmt.Sprintf("%d", limit))
	err = l.rpc.Query(req, resp)
	return
}

// @see LiveStreamSnapshotInfoWithApp
func (l *Live) LiveStreamSnapshotInfo(streamName string, startTime, endTime time.Time, limit int, resp interface{}) (err error) {
	err = l.LiveStreamSnapshotInfoWithApp(l.liveReq.AppName, streamName, startTime, endTime, limit, resp)
	return
}
