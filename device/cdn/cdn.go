//Copyright cbping
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License

// cdn CDN服务Api
// @author cbping
package cdn

import (
	"github.com/BPing/aliyun-live-go-sdk/client"
	"errors"
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"fmt"
)

const (
//action

	OpenCdnServiceAction = "OpenCdnService"
	DescribeCdnServiceAction = "DescribeCdnService"
	ModifyCdnServiceAction = "ModifyCdnService"

	AddCdnDomainAction = "AddCdnDomain"
	DeleteCdnDomainAction = "DeleteCdnDomain"
	DescribeUserDomainsAction = "DescribeUserDomains"
	DescribeCdnDomainDetailAction = "DescribeCdnDomainDetail"
	ModifyCdnDomainAction = "ModifyCdnDomain"
	StartCdnDomainAction = "StartCdnDomain"
	StopCdnDomainAction = "StopCdnDomain"

	RefreshObjectCachesAction = "RefreshObjectCaches"
	PushObjectCacheAction = "PushObjectCache"
	DescribeRefreshTasksAction = "DescribeRefreshTasks"
	DescribeRefreshQuotaAction = "DescribeRefreshQuota"

	DescribeDomainConfigsAction = "DescribeDomainConfigs"
	SetOptimizeConfigAction = "SetOptimizeConfig"
	SetPageCompressConfigAction = "SetPageCompressConfig"
	SetIgnoreQueryStringConfigAction = "SetIgnoreQueryStringConfig"
	SetRangeConfigAction = "SetRangeConfig"
	SetVideoSeekConfigAction = "SetVideoSeekConfig"
	SetSourceHostConfigAction = "SetSourceHostConfig"
	SetErrorPageConfigAction = "SetErrorPageConfig"
	SetForceRedirectConfigAction = "SetForceRedirectConfig"
	SetReferConfigAction = "SetReferConfig"
	SetFileCacheExpiredConfigAction = "SetFileCacheExpiredConfig"
	SetPathCacheExpiredConfigAction = "SetPathCacheExpiredConfig"
	ModifyFileCacheExpiredConfigAction = "ModifyFileCacheExpiredConfig"
	ModifyPathCacheExpiredConfigAction = "ModifyPathCacheExpiredConfig"
	DeleteCacheExpiredConfigAction = "DeleteCacheExpiredConfig"
	SetReqAuthConfigAction = "SetReqAuthConfig"
	SetHttpHeaderConfigAction = "SetHttpHeaderConfig"
	ModifyHttpHeaderConfigAction = "ModifyHttpHeaderConfig"
	DeleteHttpHeaderConfigAction = "DeleteHttpHeaderConfig"
	SetCcConfigAction = "SetCcConfig"
	SetWafConfigAction = "SetWafConfig"
)

//
// CDN接口控制器
// 记住，很多操作需先开通CDN服务才可执行。
type CDN struct {
	rpc    *client.Client
	cdnReq *client.CDNRequest

	debug  bool
}

// 新建"CDN接口控制器"
// @param cert  请求凭证
func NewCDN(cert *client.Credentials) *CDN {
	return &CDN{
		rpc:        client.NewClient(cert),
		cdnReq:    client.NewCDNRequest(""),
		debug:      false,
	}
}

func (c *CDN)SetDebug(debug bool) *CDN {
	c.debug = debug
	c.rpc.SetDebug(debug)
	return c
}


// 域名操作接口
// -------------------------------------------------------------------------------


// AddCdnDomain 添加加速域名
//
// @link https://help.aliyun.com/document_detail/27161.html?spm=0.0.0.0.ShLybr
func (c *CDN)AddCdnDomain(domainInfo DomainInfo, resp interface{}) (err error) {

	if (NullCdnType == domainInfo.CdnType || global.EmptyString == domainInfo.DomainName) {
		return errors.New("DomainName or CdnType should not be empty")
	}

	if (0 != domainInfo.SourcePort&&443 != domainInfo.SourcePort&&80 != domainInfo.SourcePort) {
		return errors.New("SourcePort  should  be 443 or 80 ")
	}

	//if (domainInfo.CdnType == HttpsDeliveryStreamCdnType&&(global.EmptyString == domainInfo.ServerCertificate)&&(global.EmptyString == domainInfo.PrivateKey)) {
	//	//如果是HttpsDelivery，需要上传的安全证书和私钥。
	//}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = AddCdnDomainAction
	req.SetArgs("DomainName", domainInfo.DomainName)
	req.SetArgs("CdnType", string(domainInfo.CdnType))

	if (global.EmptyString != domainInfo.SourceType) {
		req.SetArgs("SourceType", string(domainInfo.SourceType))
	}

	if (0 != domainInfo.SourcePort) {
		req.SetArgs("SourcePort", fmt.Sprintf("%d", domainInfo.SourcePort))
	}

	if (global.EmptyString != domainInfo.Sources) {
		req.SetArgs("Sources", domainInfo.Sources)
	}

	if (global.EmptyString != domainInfo.ServerCertificate) {
		req.SetArgs("ServerCertificate", domainInfo.ServerCertificate)
	}

	if (global.EmptyString != domainInfo.PrivateKey) {
		req.SetArgs("PrivateKey", domainInfo.PrivateKey)
	}

	err = c.rpc.Query(req, resp)
	return
}

//  DeleteCdnDomain 删除已添加的加速域名
//
//  @link https://help.aliyun.com/document_detail/27167.html?spm=0.0.0.0.SyHloH
func (c *CDN)DeleteCdnDomain(domainName string, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DeleteCdnDomainAction
	req.SetArgs("DomainName", domainName)

	err = c.rpc.Query(req, resp)
	return
}

// ReadUserDomains 查询用户名下所有的域名与状态。 支持域名模糊匹配过滤和域名状态过滤
//                 所有参数都是可选的。
// @param pageSize  分页大小，默认20，最大50，取值：1~50之前的任意整数。小于等于零则采用默认值
// @param pageNumber 取得第几页，取值范围为：1~100000。小于等于零则采用默认值
// @param domainName 域名模糊匹配过滤。为空时，忽略此参数
// @link https://help.aliyun.com/document_detail/27162.html?spm=0.0.0.0.COpoXo
func (c *CDN)ReadUserDomains(domainName string, pageSize, pageNumber int64, domainStatus DomainStatus, domainSearchType DomainSearchType, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DescribeUserDomainsAction
	if (global.EmptyString != domainName) {
		req.SetArgs("DomainName", domainName)
	}
	if (global.EmptyString != domainStatus) {
		req.SetArgs("DomainStatus", string(domainStatus))
	}

	if (pageSize >= 1) {
		req.SetArgs("PageSize", fmt.Sprintf("%d", pageSize))
	}

	if (pageNumber >= 1) {
		req.SetArgs("PageNumber", fmt.Sprintf("%d", pageNumber))
	}

	if (global.EmptyString != domainSearchType) {
		req.SetArgs("DomainSearchType", string(domainSearchType))
	}

	err = c.rpc.Query(req, resp)
	return
}


// CdnDomainDetail 获取指定加速域名配置的基本信息
//
// @link https://help.aliyun.com/document_detail/27162.html?spm=0.0.0.0.COpoXo
func (c *CDN)CdnDomainDetail(domainName string, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DescribeCdnDomainDetailAction
	req.SetArgs("DomainName", domainName)

	err = c.rpc.Query(req, resp)
	return
}

// ModifyCdnDomain 修改加速域名，目前支持修改源站
//
// @link https://help.aliyun.com/document_detail/27164.html?spm=0.0.0.0.rOMSJ4
func (c *CDN)ModifyCdnDomain(domainInfo DomainInfo, resp interface{}) (err error) {
	if (global.EmptyString == domainInfo.SourceType || global.EmptyString == domainInfo.DomainName) {
		return errors.New("DomainName or SourceType should not be empty")
	}

	if (0 != domainInfo.SourcePort&&443 != domainInfo.SourcePort&&80 != domainInfo.SourcePort) {
		return errors.New("SourcePort  should  be 443 or 80 ; if it is zero, 80 defualt  ")
	}

	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = ModifyCdnDomainAction
	req.SetArgs("DomainName", domainInfo.DomainName)
	req.SetArgs("SourceType", string(domainInfo.SourceType))

	if (0 != domainInfo.SourcePort) {
		req.SetArgs("SourcePort", fmt.Sprintf("%d", domainInfo.SourcePort))
	}

	if (global.EmptyString != domainInfo.Sources) {
		req.SetArgs("Sources", domainInfo.Sources)
	}

	err = c.rpc.Query(req, resp)
	return
}

//  StartCdnDomain 启用状态为“停用”的加速域名，将DomainStatus变更为online
//
//  @link https://help.aliyun.com/document_detail/27165.html?spm=0.0.0.0.8cQhXd
func (c *CDN)StartCdnDomain(domainName string, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = StartCdnDomainAction
	req.SetArgs("DomainName", domainName)
	err = c.rpc.Query(req, resp)
	return
}

//  StopCdnDomain 停用某个加速域名，将DomainStatus变更为offline
//
//  @link https://help.aliyun.com/document_detail/27166.html?spm=0.0.0.0.JcQVpK
func (c *CDN)StopCdnDomain(domainName string, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = StopCdnDomainAction
	req.SetArgs("DomainName", domainName)
	err = c.rpc.Query(req, resp)
	return
}


