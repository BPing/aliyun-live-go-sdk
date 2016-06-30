package cdn

import (
	"github.com/BPing/aliyun-live-go-sdk/client"
	"errors"
)

// 服务操作接口
// -------------------------------------------------------------------------------


// OpenCdnService 开通CDN服务
// @param  internetChargeType 开通服务的计费类型(默认按流量) 按流量(PayByTraffic)、按带宽峰值(PayByBandwidth)。
//                            常量  PayByTrafficType(PayByTraffic)和PayByBandwidthType(PayByBandwidth)
// @link https://help.aliyun.com/document_detail/27157.html?spm=0.0.0.0.t6wFRF
func (c *CDN)OpenCdnService(internetChargeType CdnPayType, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = OpenCdnServiceAction
	if (PayNullType == internetChargeType) {
		internetChargeType = PayByTrafficType
	}
	req.SetArgs("InternetChargeType", string(internetChargeType))
	err = c.rpc.Query(req, resp)
	return
}

// ScanCdnService 查询CDN服务状态。包括：当前计费类型，服务开通时间，下次生效的计费类型，当前业务状态等。
//
// @link https://help.aliyun.com/document_detail/27158.html?spm=0.0.0.0.2RRuSQ
func (c *CDN)ScanCdnService(resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DescribeCdnServiceAction
	err = c.rpc.Query(req, resp)
	return
}

//  ModifyCdnServicePayType 变更CDN服务的计费类型.
//  @param  internetChargeType 开通服务的计费类型(默认按流量) 按流量(PayByTraffic)、按带宽峰值(PayByBandwidth)。
//                            常量  PayByTrafficType(PayByTraffic)和PayByBandwidthType(PayByBandwidth)
// @link https://help.aliyun.com/document_detail/27159.html?spm=0.0.0.0.AuPm7B
func (c *CDN)ModifyCdnServicePayType(internetChargeType CdnPayType, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = ModifyCdnServiceAction
	if (PayNullType == internetChargeType) {
		err = errors.New("internetChargeType should not be empty")
		return
	}
	req.SetArgs("InternetChargeType", string(internetChargeType))
	err = c.rpc.Query(req, resp)
	return
}

