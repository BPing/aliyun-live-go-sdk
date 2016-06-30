package cdn

import (
	"strings"
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"errors"
	"fmt"
)


// 配置操作接口
// -------------------------------------------------------------------------------

//  DomainConfigs 获取指定加速域名的配置
//  @configList 为空时代表查询所有
//  @link https://help.aliyun.com/document_detail/27169.html?spm=0.0.0.0.fewJpA
func (c *CDN)DomainConfigs(domainName string, resp interface{}, configList...ConfigName) (err error) {
	config := []string{}
	for _, val := range configList {
		config = append(config, string(val))
	}
	err = c.domainConfigs(domainName, resp, config ...)
	return
}

func (c *CDN)domainConfigs(domainName string, resp interface{}, configList...string) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DescribeDomainConfigsAction
	req.SetArgs("DomainName", domainName)
	if (len(configList) > 0) {
		req.SetArgs("ConfigList", strings.Join(configList, ","))
	}
	err = c.rpc.Query(req, resp)
	return
}

//
// @enable
func (c *CDN)enabledConfig(action, domainName string, enable bool, resp interface{}) (err error) {
	if (global.EmptyString == domainName) {
		return errors.New("domainName should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = action
	req.SetArgs("DomainName", domainName)
	enableConfig := "Off"
	if (enable) {
		enableConfig = "On"
	}
	req.SetArgs("Enable", enableConfig)
	err = c.rpc.Query(req, resp)
	return
}

// SetOptimizeConfig 设置页面优化功能，开启后可以删除 html ，<br>
//                    内嵌 javascript 和 css 中的注释以及重复的空白符；这样可以有效地去除页面的冗余内容，减小文件体积，提高加速分发效率<br>
// @link https://help.aliyun.com/document_detail/27170.html?spm=0.0.0.0.w21KyD
func (c *CDN)SetOptimizeConfig(domainName string, enable bool, resp interface{}) (err error) {
	err = c.enabledConfig(SetOptimizeConfigAction, domainName, enable, resp)
	return
}

// SetPageCompressConfig 设置智能压缩功能
//
// @link https://help.aliyun.com/document_detail/27170.html?spm=0.0.0.0.w21KyD
func (c *CDN)SetPageCompressConfig(domainName string, enable bool, resp interface{}) (err error) {
	err = c.enabledConfig(SetPageCompressConfigAction, domainName, enable, resp)
	return
}

//  SetIgnoreQueryStringConfig 设置过滤参数功能
//
//  @link https://help.aliyun.com/document_detail/27172.html?spm=0.0.0.0.1cmlUC
func (c *CDN)SetIgnoreQueryStringConfig(domainName string, enable bool, resp interface{}) (err error) {
	err = c.enabledConfig(SetIgnoreQueryStringConfigAction, domainName, enable, resp)
	return
}

//  SetRangeConfig 设置range回源功能
//
//  @link https://help.aliyun.com/document_detail/27173.html?spm=0.0.0.0.dKGEgt
func (c *CDN)SetRangeConfig(domainName string, enable bool, resp interface{}) (err error) {
	err = c.enabledConfig(SetRangeConfigAction, domainName, enable, resp)
	return
}

//  SetVideoSeekConfig 设置拖拽播放功能
//
//  @link https://help.aliyun.com/document_detail/27174.html?spm=0.0.0.0.r7OBIj
func (c *CDN)SetVideoSeekConfig(domainName string, enable bool, resp interface{}) (err error) {
	err = c.enabledConfig(SetVideoSeekConfigAction, domainName, enable, resp)
	return
}

//  SetWafConfig 设置加速域名的Waf防护功能
//
//  @link https://help.aliyun.com/document_detail/27189.html?spm=5176.doc27188.6.192.POkElI
func (c *CDN)SetWafConfig(domainName string, enable bool, resp interface{}) (err error) {
	err = c.enabledConfig(SetWafConfigAction, domainName, enable, resp)
	return
}

//  SetCcConfig 设置加速域名的CC防护功能、IP黑白名单设置
//
//  @link https://help.aliyun.com/document_detail/27188.html?spm=5176.doc27189.6.191.F7funw
func (c *CDN)SetCcConfig(domainName, allowIps, blockIps string, enable bool, resp interface{}) (err error) {
	if (global.EmptyString == domainName) {
		return errors.New("domainName should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetCcConfigAction
	req.SetArgs("DomainName", domainName)
	enableConfig := "Off"
	if (enable) {
		enableConfig = "On"
	}
	req.SetArgs("Enable", enableConfig)
	if (allowIps != global.EmptyString) {
		req.SetArgs("AllowIps", allowIps)
	}
	if (blockIps != global.EmptyString) {
		req.SetArgs("BlockIps", blockIps)
	}
	err = c.rpc.Query(req, resp)
	return
}

//  SetSourceHostConfig 设置回源host功能
//
//  @link https://help.aliyun.com/document_detail/27175.html?spm=0.0.0.0.RLyJt0
func (c *CDN)SetSourceHostConfig(domainName, backSrcDomain string, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == backSrcDomain) {
		return errors.New("domainName or backSrcDomain should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetSourceHostConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("BackSrcDomain", backSrcDomain)
	err = c.rpc.Query(req, resp)
	return
}

//  SetErrorPageConfig 设置加速域名自定义404错误页面跳转
//
//  @link https://help.aliyun.com/document_detail/27176.html?spm=0.0.0.0.2WOcZu
func (c *CDN)SetErrorPageConfig(domainName, customPageUrl string, pageType PageType, resp interface{}) (err error) {
	if (global.EmptyString == domainName) {
		return errors.New("domainName should not be empty")
	}

	if (pageType == OtherPageType&&global.EmptyString == customPageUrl) {
		return errors.New("'customPageUrl' should not be empty when the value of 'pageType' is 'other'")
	}

	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetErrorPageConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("PageType", string(pageType))
	if (global.EmptyString != customPageUrl) {
		req.SetArgs("CustomPageUrl", customPageUrl)
	}
	err = c.rpc.Query(req, resp)
	return
}

//  SetForceRedirectConfig 设置强制访问跳转方式, 目前支持强制Http或Https跳转.
//
//  @link https://help.aliyun.com/document_detail/27177.html?spm=0.0.0.0.kKSTQo
func (c *CDN)SetForceRedirectConfig(domainName string, redirectType RedirectType, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == redirectType) {
		return errors.New("domainName or redirectType should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetForceRedirectConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("RedirectType", string(redirectType))
	err = c.rpc.Query(req, resp)
	return
}

//  SetReferConfig 设置加速域名的Refer防盗链功能 <br>
//  @param allowEmpty bool 是否允许空refer访问 <br>
//  @link https://help.aliyun.com/document_detail/27178.html?spm=0.0.0.0.11Tr6z
func (c *CDN)SetReferConfig(domainName, referList string, referType ReferType, allowEmpty bool, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == referType) {
		return errors.New("domainName or referType should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetReferConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("ReferType", string(referType))
	enableConfig := "off"
	if (allowEmpty) {
		enableConfig = "on"
	}
	req.SetArgs("AllowEmpty", enableConfig)
	if (global.EmptyString != referList) {
		req.SetArgs("ReferList", referList)
	}
	err = c.rpc.Query(req, resp)
	return
}

//  SetFileCacheExpiredConfig 设置文件过期配置 <br>
//  Qparam ttl 	缓存时间设置，单位为秒 <br>
//  @link https://help.aliyun.com/document_detail/27179.html?spm=0.0.0.0.PrOMF2
func (c *CDN)SetFileCacheExpiredConfig(domainName, cacheContent string, ttl, weight int64, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == cacheContent || ttl < 0) {
		return errors.New("domainName or cacheContent should not be empty and ttl should be bigger than zero")
	}
	if (weight < 1 || weight > 99) {
		return errors.New("the number of weight should between 1 to 99")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetFileCacheExpiredConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("CacheContent", cacheContent)
	req.SetArgs("TTL", fmt.Sprintf("%d", ttl))
	req.SetArgs("Weight", fmt.Sprintf("%d", weight))
	err = c.rpc.Query(req, resp)
	return
}

//  SetPathCacheExpiredConfig 修改目录过期配置 <br>
//  Qparam ttl 	缓存时间设置，单位为秒 <br>
//  @link https://help.aliyun.com/document_detail/27179.html?spm=0.0.0.0.PrOMF2
func (c *CDN)SetPathCacheExpiredConfig(domainName, cacheContent string, ttl, weight int64, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == cacheContent || ttl < 0) {
		return errors.New("domainName or cacheContent should not be empty and ttl should be bigger than zero")
	}
	if (weight < 1 || weight > 99) {
		return errors.New("the number of weight should between 1 to 99")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetPathCacheExpiredConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("CacheContent", cacheContent)
	req.SetArgs("TTL", fmt.Sprintf("%d", ttl))
	req.SetArgs("Weight", fmt.Sprintf("%d", weight))
	err = c.rpc.Query(req, resp)
	return
}

//  ModifyFileCacheExpiredConfig 修改文件过期配置 <br>
//  Qparam ttl 	缓存时间设置，单位为秒 <br>
//  @link https://help.aliyun.com/document_detail/27179.html?spm=0.0.0.0.PrOMF2
func (c *CDN)ModifyFileCacheExpiredConfig(domainName, configID, cacheContent string, ttl, weight int64, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == cacheContent || global.EmptyString == configID || ttl < 0) {
		return errors.New("'domainName' 、 'cacheContent' and 'configID' should not be empty and ttl should be bigger than zero")
	}
	if (weight < 1 || weight > 99) {
		return errors.New("the number of weight should between 1 to 99")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = ModifyFileCacheExpiredConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("ConfigID", configID)
	req.SetArgs("CacheContent", cacheContent)
	req.SetArgs("TTL", fmt.Sprintf("%d", ttl))
	req.SetArgs("Weight", fmt.Sprintf("%d", weight))
	err = c.rpc.Query(req, resp)
	return
}

//  ModifyPathCacheExpiredConfig 修改目录过期配置
//  Qparam ttl 	缓存时间设置，单位为秒
//  @link https://help.aliyun.com/document_detail/27179.html?spm=0.0.0.0.PrOMF2
func (c *CDN)ModifyPathCacheExpiredConfig(domainName, configID, cacheContent string, ttl, weight int64, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == cacheContent || global.EmptyString == configID || ttl < 0) {
		return errors.New("'domainName' 、 'cacheContent' and 'configID' should not be empty and ttl should be bigger than zero")
	}
	if (weight < 1 || weight > 99) {
		return errors.New("the number of weight should between 1 to 99")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = ModifyPathCacheExpiredConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("ConfigID", configID)
	req.SetArgs("CacheContent", cacheContent)
	req.SetArgs("TTL", fmt.Sprintf("%d", ttl))
	req.SetArgs("Weight", fmt.Sprintf("%d", weight))
	err = c.rpc.Query(req, resp)
	return
}

//  DeleteCacheExpiredConfig 删除自定义缓存策略.
//
//  @link https://help.aliyun.com/document_detail/27183.html?spm=5176.doc27182.6.186.DVvk01
func (c *CDN)DeleteCacheExpiredConfig(domainName, configID string, cacheType CacheType, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == cacheType || global.EmptyString == configID ) {
		return errors.New("'domainName' 、 'cacheType' and 'configID' should not be empty ")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DeleteCacheExpiredConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("ConfigID", configID)
	req.SetArgs("CacheType", string(cacheType))
	err = c.rpc.Query(req, resp)
	return
}

//  SetReqAuthConfig 设置加速域名的访问鉴权配置.
//  @timeout 鉴权缓存时间，单位为秒；如果不大于零时，则忽略此参数，也就是说采用默认值
//  @link https://help.aliyun.com/document_detail/27184.html?spm=5176.doc27183.6.187.r0X7ox1
func (c *CDN)SetReqAuthConfig(domainName, key1, key2 string, authType AuthType, timeout int64, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == authType ) {
		return errors.New("'domainName' or 'authType' should not be empty ")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetReqAuthConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("AuthType", string(authType))
	if (global.EmptyString != key1) {
		req.SetArgs("Key1", key1)
	}
	if (global.EmptyString != key2) {
		req.SetArgs("Key2", key2)
	}
	if (timeout > 0) {
		req.SetArgs("Timeout", fmt.Sprintf("%d", timeout))
	}
	err = c.rpc.Query(req, resp)
	return
}

//  SetHttpHeaderConfig 设置自定义http头
//
//  @link https://help.aliyun.com/document_detail/27185.html?spm=5176.doc27184.6.188.o8SCqc
func (c *CDN)SetHttpHeaderConfig(domainName, headerKey, headerValue string, resp interface{}) (err error) {
	if (global.EmptyString == domainName  ) {
		return errors.New("'domainName'should not be empty ")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = SetHttpHeaderConfigAction
	req.SetArgs("DomainName", domainName)
	if (global.EmptyString != headerKey) {
		req.SetArgs("HeaderKey", headerKey)
	}

	req.SetArgs("HeaderValue", headerValue)

	err = c.rpc.Query(req, resp)
	return
}

//  ModifyHttpHeaderConfig 修改自定义http头
//
//  @link https://help.aliyun.com/document_detail/27186.html?spm=5176.doc27185.6.189.LlynbW
func (c *CDN)ModifyHttpHeaderConfig(domainName, configID, headerKey, headerValue string, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == configID  ) {
		return errors.New("'domainName'or 'configID' should not be empty ")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = ModifyHttpHeaderConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("ConfigID", configID)
	if (global.EmptyString != headerKey) {
		req.SetArgs("HeaderKey", headerKey)
	}

	req.SetArgs("HeaderValue", headerValue)

	err = c.rpc.Query(req, resp)
	return
}

//  DeleteHttpHeaderConfig 删除加速域名的Refer防盗链配置
//
//  @link https://help.aliyun.com/document_detail/27187.html?spm=5176.doc27186.6.190.EvwqYw
func (c *CDN)DeleteHttpHeaderConfig(domainName, configID string, resp interface{}) (err error) {
	if (global.EmptyString == domainName || global.EmptyString == configID  ) {
		return errors.New("'domainName'or 'configID' should not be empty ")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DeleteHttpHeaderConfigAction
	req.SetArgs("DomainName", domainName)
	req.SetArgs("ConfigID", configID)
	err = c.rpc.Query(req, resp)
	return
}