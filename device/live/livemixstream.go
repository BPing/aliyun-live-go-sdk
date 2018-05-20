package live

// 混流
// -------------------------------------------------------------------------------

type MixStreamParam struct {
	Main StreamBase //主混流
	Mix  StreamBase //副混流

	//MainDomainName string //	主混流域名
	//MainAppName    string //	主混流应用名称
	//MainStreamName string //	主混流直播流名称
	//MixDomainName  string //	副混流域名
	//MixAppName     string //	副混流应用名称
	//MixStreamName  string //	副混流直播流名称
	MixTemplate string //	混流模版，支持 picture_in_picture 、side_by_side
	MixType     string //	混流类型， 支持 channel 和 stream （都是指的副流，主流必须是 channel）
}

// StartMixStreamsService 开始混流操作
// {@link https://help.aliyun.com/document_detail/44405.html?spm=5176.doc44406.6.698.Uh0xWX}
func (l *Live) StartMixStreamsService(mixStream MixStreamParam, resp interface{}) (err error) {
	req := l.cloneRequest(StartMixStreamsServiceAction)
	req.AppName = ""
	req.DomainName = ""
	req.SetArgs("MainDomainName", mixStream.Main.DomainName)
	req.SetArgs("MainAppName", mixStream.Main.AppName)
	req.SetArgs("MainStreamName", mixStream.Main.StreamName)
	req.SetArgs("MixDomainName", mixStream.Mix.DomainName)
	req.SetArgs("MixAppName", mixStream.Mix.AppName)
	req.SetArgs("MixStreamName", mixStream.Mix.StreamName)
	req.SetArgs("MixTemplate", mixStream.MixTemplate)
	req.SetArgs("MixType", mixStream.MixType)
	err = l.rpc.Query(req, resp)
	return
}

// StopMixStreamsService 结束混流操作
// {@link https://help.aliyun.com/document_detail/44406.html?spm=5176.doc44405.6.699.lDciGj}
func (l *Live) StopMixStreamsService(mixStream MixStreamParam, resp interface{}) (err error) {
	req := l.cloneRequest(StopMixStreamsServiceAction)
	req.AppName = ""
	req.DomainName = ""
	req.SetArgs("MainDomainName", mixStream.Main.DomainName)
	req.SetArgs("MainAppName", mixStream.Main.AppName)
	req.SetArgs("MainStreamName", mixStream.Main.StreamName)
	req.SetArgs("MixDomainName", mixStream.Mix.DomainName)
	req.SetArgs("MixAppName", mixStream.Mix.AppName)
	req.SetArgs("MixStreamName", mixStream.Mix.StreamName)
	err = l.rpc.Query(req, resp)
	return
}

// 连麦
// -------------------------------------------------------------------------------

// AddLiveMixConfig 添加连麦配置
// @param  template  你所需要配置的连麦转码模板，取值: mhd 或者 msd
// {@link https://help.aliyun.com/document_detail/52718.html?spm=5176.doc52726.6.679.SzBcKQ}
func (l *Live) AddLiveMixConfig(template string, resp interface{}) (err error) {
	req := l.cloneRequest(AddLiveMixConfigAction)
	req.SetArgs("Template", template)
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveMixConfig 在指定的域名下查询所有的连麦配置。
// {@link https://help.aliyun.com/document_detail/52722.html?spm=5176.doc52718.6.680.1EYw6L}
func (l *Live) DescribeLiveMixConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveMixConfigAction)
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveMixConfig 删除连麦配置
// {@link https://help.aliyun.com/document_detail/52720.html?spm=5176.doc52722.6.681.F9sNc1}
func (l *Live) DeleteLiveMixConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveMixConfigAction)
	err = l.rpc.Query(req, resp)
	return
}

// StartMultipleStreamMixService 开启多人连麦服务
// {@link https://help.aliyun.com/document_detail/51313.html?spm=5176.doc52720.6.682.0eSfxw}
func (l *Live) StartMultipleStreamMixService(streamName, mixTemplate string, resp interface{}) (err error) {
	req := l.cloneRequest(StartMultipleStreamMixServiceAction)
	req.SetArgs("StreamName", streamName)
	req.SetArgs("MixTemplate", mixTemplate)
	err = l.rpc.Query(req, resp)
	return
}

// StopMultipleStreamMixService 停止多人连麦服务
// {@link https://help.aliyun.com/document_detail/51314.html?spm=5176.doc51313.6.683.mdMFhc}
func (l *Live) StopMultipleStreamMixService(streamName string, resp interface{}) (err error) {
	req := l.cloneRequest(StopMultipleStreamMixServiceAction)
	req.SetArgs("StreamName", streamName)
	err = l.rpc.Query(req, resp)
	return
}

// AddMultipleStreamMixService 往主流添加一路流
// {@link https://help.aliyun.com/document_detail/51315.html?spm=5176.doc51314.6.684.e6WbAe}
func (l *Live) AddMultipleStreamMixService(mixStream MixStreamParam, resp interface{}) (err error) {
	req := l.cloneRequest(AddMultipleStreamMixServiceAction)
	if mixStream.Main.DomainName != "" {
		req.DomainName = mixStream.Main.DomainName
	}
	if mixStream.Main.AppName != "" {
		req.AppName = mixStream.Main.AppName
	}
	req.SetArgs("StreamName", mixStream.Main.StreamName)
	req.SetArgs("MixDomainName", mixStream.Mix.DomainName)
	req.SetArgs("MixAppName", mixStream.Mix.AppName)
	req.SetArgs("MixStreamName", mixStream.Mix.StreamName)
	err = l.rpc.Query(req, resp)
	return
}

// RemoveMultipleStreamMixService 从主流移除一路流
// {@link https://help.aliyun.com/document_detail/51316.html?spm=5176.doc51315.6.685.9SN11d}
func (l *Live) RemoveMultipleStreamMixService(mixStream MixStreamParam, resp interface{}) (err error) {
	req := l.cloneRequest(RemoveMultipleStreamMixServiceAction)
	if mixStream.Main.DomainName != "" {
		req.DomainName = mixStream.Main.DomainName
	}
	if mixStream.Main.AppName != "" {
		req.AppName = mixStream.Main.AppName
	}
	req.SetArgs("StreamName", mixStream.Main.StreamName)
	req.SetArgs("MixDomainName", mixStream.Mix.DomainName)
	req.SetArgs("MixAppName", mixStream.Mix.AppName)
	req.SetArgs("MixStreamName", mixStream.Mix.StreamName)
	err = l.rpc.Query(req, resp)
	return
}

// AddLiveMixNotifyConfig 添加连麦回调配置
// {@link https://help.aliyun.com/document_detail/52719.html?spm=5176.doc51316.6.686.imrZss}
func (l *Live) AddLiveMixNotifyConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.cloneRequest(AddLiveMixNotifyConfigAction)
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.rpc.Query(req, resp)
	return
}

// DescribeLiveMixNotifyConfig 查询连麦回调配置
// {@link https://help.aliyun.com/document_detail/52723.html?spm=5176.doc52725.6.687.d1vPFC}
func (l *Live) DescribeLiveMixNotifyConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveMixNotifyConfigAction)
	err = l.rpc.Query(req, resp)
	return
}

// UpdateLiveMixNotifyConfig 更新连麦回调配置
// {@link tps://help.aliyun.com/document_detail/52725.html?spm=5176.doc52726.6.688.HTKozg}
func (l *Live) UpdateLiveMixNotifyConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.cloneRequest(UpdateLiveMixNotifyConfigAction)
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveMixNotifyConfig 删除连麦回调配置
// {@link https://help.aliyun.com/document_detail/52721.html?spm=5176.doc52723.6.689.2MWUeX}
func (l *Live) DeleteLiveMixNotifyConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveMixNotifyConfigAction)
	err = l.rpc.Query(req, resp)
	return
}
