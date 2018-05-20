package live

import "strconv"

// 直播转点播
// -------------------------------------------------------------------------------

// AddLiveRecordVodConfig 增加直播录制转点播配置，录制内容保存到点播媒资库。
//
// @param bodTranscodeGroupId string  点播转码模板组ID
// @param cycleDuration  int 非必需字段
//                   周期录制时长。
//                      取值范围：[300，21600]
//                      单位：秒。
//                      不填则默认为3600秒。
//
// https://help.aliyun.com/document_detail/63968.html?spm=a2c4g.11186623.6.668.Deq7on
func (l *Live) AddLiveRecordVodConfig(bodTranscodeGroupId string, cycleDuration int, resp interface{}) (err error) {
	req := l.cloneRequest(AddLiveRecordVodConfigAction)
	req.SetArgs("VodTranscodeGroupId", bodTranscodeGroupId)
	if cycleDuration >= 300 && cycleDuration <= 21600 {
		req.SetArgs("CycleDuration", strconv.Itoa(cycleDuration))
	}
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveRecordVodConfig 删除直播录制转点播配置。
//
// https://help.aliyun.com/document_detail/63969.html?spm=a2c4g.11186623.6.669.mKTDcM
func (l *Live) DeleteLiveRecordVodConfig(streamName string, resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveRecordVodConfigAction)
	req.SetArgs("StreamName", streamName)
	err = l.rpc.Query(req, resp)
	return
}

type DescribeVodParam struct {
	//阿里云颁发给用户的访问服务所用的密钥ID。
	AccessKeyId string
	StreamName  string
	PageInfo
}

// DescribeLiveRecordVodConfigs 查询直转点配置列表。
//
// https://help.aliyun.com/document_detail/63970.html?spm=a2c4g.11186623.6.670.AmfuWu
func (l *Live) DescribeLiveRecordVodConfigs(vodParam DescribeVodParam, resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveRecordVodConfigsAction)

	if vodParam.AccessKeyId != "" {
		req.SetArgs("AccessKeyId", vodParam.AccessKeyId)
	}

	if vodParam.StreamName != "" {
		req.SetArgs("StreamName", vodParam.StreamName)
	}

	//分页的页码。
	//默认值：1
	if vodParam.PageNum <= 0 {
		vodParam.PageNum = 1
	}

	//每页大小。取值范围：
	//[5，100]
	//默认值：10
	if vodParam.PageSize < 5 || vodParam.PageSize > 100 {
		vodParam.PageSize = 10
	}

	req.SetArgs("PageNum", strconv.Itoa(vodParam.PageNum))
	req.SetArgs("PageSize", strconv.Itoa(vodParam.PageSize))

	err = l.rpc.Query(req, resp)
	return
}
