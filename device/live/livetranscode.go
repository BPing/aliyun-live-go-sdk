package live

import (
	"errors"
	"strings"
)

// 转码 处理
// ---------------------------------------------------------------------------------------------------------------------

// AddLiveStreamTranscodeWithApp 添加转码配置
// template	string	是	转码模版
// record	string	是	yes or no，是否需要录制
// snapshot	string	是	yes or no，是否需要截图
// {@link https://help.aliyun.com/document_detail/44041.html?spm=5176.doc44719.6.691.U0gtfy}
func (l *Live) AddLiveStreamTranscode(template, record, snapshot string, resp interface{}) (err error) {
	record = strings.ToLower(record)
	snapshot = strings.ToLower(snapshot)
	if record != "yes" && record != "no" {
		return errors.New("record:'yes' or 'no'")
	}
	if snapshot != "yes" && snapshot != "no" {
		return errors.New("snapshot:'yes' or 'no'")
	}
	req := l.cloneRequest(AddLiveStreamTranscodeAction)
	req.SetArgs("Domain", req.DomainName)
	req.SetArgs("App", req.AppName)
	req.SetArgs("Template", template)
	req.SetArgs("Record", record)
	req.SetArgs("Snapshot", snapshot)
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveStreamTranscodeWithApp 删除转码配置
// template	string	是	转码模版
// {@link https://help.aliyun.com/document_detail/44042.html?spm=5176.doc44041.6.693.Xu9vuV}
func (l *Live) DeleteLiveStreamTranscode(template string, resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveStreamTranscodeAction)
	req.SetArgs("Domain", req.DomainName)
	req.SetArgs("App", req.AppName)
	req.SetArgs("Template", template)
	err = l.rpc.Query(req, resp)
	return
}

// LiveStreamTranscodeInfo 查询转码配置信息.
// {@link https://help.aliyun.com/document_detail/44043.html?spm=5176.doc44042.6.697.7sPQts}
func (l *Live) LiveStreamTranscodeInfo(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamTranscodeInfoAction)
	req.AppName = ""
	// domainTranscodeName	String	是	您的加速域名
	req.SetArgs("DomainTranscodeName", req.DomainName)
	err = l.rpc.Query(req, resp)
	return
}
