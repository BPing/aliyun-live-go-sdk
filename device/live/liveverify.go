package live

import (
	"strconv"
)

// 直播审核
// -------------------------------------------------------------------------------

type SnapshotDetectPornParam struct {
	LiveBase // 不设置 默认实例的配置
	Order OrderType
	PageInfo
}

// DescribeLiveSnapshotDetectPornConfig 查询审核配置。
//
// @link https://help.aliyun.com/document_detail/56044.html?spm=a2c4g.11186623.6.693.eqB4c9
func (l *Live) DescribeLiveSnapshotDetectPornConfig(snapshotParam SnapshotDetectPornParam, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveSnapshotDetectPornConfigAction)
	if snapshotParam.Order != NilOrderType {
		req.SetArgs("Order", string(snapshotParam.Order))
	}
	if snapshotParam.AppName != "" {
		req.AppName = snapshotParam.AppName
	}
	if snapshotParam.DomainName != "" {
		req.DomainName = snapshotParam.DomainName
	}
	//分页的页码。
	//默认值：1
	if snapshotParam.PageNum <= 0 {
		snapshotParam.PageNum = 1
	}
	//每页大小。取值范围：
	//[5，30]
	//默认值：10
	if snapshotParam.PageSize < 5 || snapshotParam.PageSize > 30 {
		snapshotParam.PageSize = 10
	}
	req.SetArgs("PageNum", strconv.Itoa(snapshotParam.PageNum))
	req.SetArgs("PageSize", strconv.Itoa(snapshotParam.PageSize))
	err = l.rpc.Query(req, resp)
	return
}

//
type AddSnapshotDetectPornParam struct {
	LiveBase           // 不设置 默认实例的配置
	OssEndpoint string //	oss endpoint，如：oss-cn-hangzhou.aliyuncs.com，详细请参照oss相关文档
	OssBucket   string //	oss存储bucket名称
	OssObject   string //	保存涉黄涉政等违规图片的对象模板， 如不明确给出，默认为{AppName}/{StreamName}/{Date}/{Hour}/{Minute}_{Second}.jpg。

	Interval int // 采样间隔。单位：秒。
	SceneN   SceneN
}

// AddLiveSnapshotDetectPornConfig 添加审核配置，
//
// @link https://help.aliyun.com/document_detail/56040.html?spm=a2c4g.11186623.6.691.xZ4nz7
func (l *Live) AddLiveSnapshotDetectPornConfig(snapshotParam AddSnapshotDetectPornParam, resp interface{}) (err error) {
	req := l.cloneRequest(AddLiveSnapshotDetectPornConfigAction)
	l.parseSnapshotDetectPornParam(req, snapshotParam)
	err = l.rpc.Query(req, resp)
	return
}

// 解析参数
func (l *Live) parseSnapshotDetectPornParam(req *Request, snapshotParam AddSnapshotDetectPornParam) {
	if snapshotParam.Interval >= 5 && snapshotParam.Interval <= 3600 {
		req.SetArgs("Interval", strconv.Itoa(snapshotParam.Interval))
	}
	if snapshotParam.AppName != "" {
		req.AppName = snapshotParam.AppName
	}
	if snapshotParam.DomainName != "" {
		req.DomainName = snapshotParam.DomainName
	}
	if snapshotParam.SceneN != NilSceneN {
		req.SetArgs("Scene.N", string(snapshotParam.SceneN))
	}
	req.SetArgs("OssEndpoint", snapshotParam.OssEndpoint)
	req.SetArgs("OssBucket", snapshotParam.OssBucket)
	if snapshotParam.OssObject != "" {
		req.SetArgs("OssObject", snapshotParam.OssObject)
	}
}

// UpdateLiveSnapshotDetectPornConfig 更新审核回调
//
// @link https://help.aliyun.com/document_detail/56046.html?spm=a2c4g.11186623.6.695.NcXR3v
func (l *Live) UpdateLiveSnapshotDetectPornConfig(snapshotParam AddSnapshotDetectPornParam, resp interface{}) (err error) {
	req := l.cloneRequest(UpdateLiveSnapshotDetectPornConfigAction)
	l.parseSnapshotDetectPornParam(req, snapshotParam)
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveSnapshotDetectPornConfig 删除审核回调
//
// @link https://help.aliyun.com/document_detail/56042.html?spm=a2c4g.11186623.6.697.NwRDuP
func (l *Live) DeleteLiveSnapshotDetectPornConfig(snapshotParam LiveBase, resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveSnapshotDetectPornConfigAction)
	if snapshotParam.AppName != "" {
		req.AppName = snapshotParam.AppName
	}
	if snapshotParam.DomainName != "" {
		req.DomainName = snapshotParam.DomainName
	}
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveDetectNotifyConfig 查询审核回调
//
// @link https://help.aliyun.com/document_detail/56043.html?spm=a2c4g.11186623.6.694.bqbBJW
func (l *Live) DescribeLiveDetectNotifyConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveDetectNotifyConfigAction)
	err = l.rpc.Query(req, resp)
	return
}

// AddLiveDetectNotifyConfig 添加回调通知 URL
// @param notifyUrl string 如：http://www.yourdomain.cn/examplecallback.action
// @link https://help.aliyun.com/document_detail/56039.html?spm=a2c4g.11186623.6.692.XGmSDW
func (l *Live) AddLiveDetectNotifyConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.cloneRequest(AddLiveDetectNotifyConfigAction)
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.rpc.Query(req, resp)
	return
}

// UpdateLiveDetectNotifyConfig 更新回调通知 URL
// @param notifyUrl string 如：http://www.yourdomain.cn/examplecallback.action
// @link https://help.aliyun.com/document_detail/56045.html?spm=a2c4g.11186623.6.696.rhGRzO
func (l *Live) UpdateLiveDetectNotifyConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.cloneRequest(UpdateLiveDetectNotifyConfigAction)
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveDetectNotifyConfig 删除审核回调
//
// @link https://help.aliyun.com/document_detail/56045.html?spm=a2c4g.11186623.6.696.rhGRzO
func (l *Live) DeleteLiveDetectNotifyConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveDetectNotifyConfigAction)
	err = l.rpc.Query(req, resp)
	return
}
