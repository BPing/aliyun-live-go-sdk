package live

type MixStreamParam struct {
	MainDomainName string //	主混流域名
	MainAppName    string //	主混流应用名称
	MainStreamName string //	主混流直播流名称
	MixDomainName  string //	副混流域名
	MixAppName     string //	副混流应用名称
	MixStreamName  string //	副混流直播流名称
	MixTemplate    string //	混流模版，支持 picture_in_picture 、side_by_side
	MixType        string //	混流类型， 支持 channel 和 stream （都是指的副流，主流必须是 channel）
}

// StartMixStreamsService 开始混流操作
// {@link https://help.aliyun.com/document_detail/44405.html?spm=5176.doc44406.6.698.Uh0xWX}
func (l *Live) StartMixStreamsService(mixStream MixStreamParam, resp interface{}) (err error) {
	req := l.cloneRequest(StartMixStreamsServiceAction)
	req.AppName = ""
	req.DomainName = ""
	req.SetArgs("MainDomainName", mixStream.MainDomainName)
	req.SetArgs("MainAppName", mixStream.MainAppName)
	req.SetArgs("MainStreamName", mixStream.MainStreamName)
	req.SetArgs("MixDomainName", mixStream.MixDomainName)
	req.SetArgs("MixAppName", mixStream.MixAppName)
	req.SetArgs("MixStreamName", mixStream.MixStreamName)
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
	req.SetArgs("MainDomainName", mixStream.MainDomainName)
	req.SetArgs("MainAppName", mixStream.MainAppName)
	req.SetArgs("MainStreamName", mixStream.MainStreamName)
	req.SetArgs("MixDomainName", mixStream.MixDomainName)
	req.SetArgs("MixAppName", mixStream.MixAppName)
	req.SetArgs("MixStreamName", mixStream.MixStreamName)
	err = l.rpc.Query(req, resp)
	return
}
