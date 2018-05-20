package live

// 状态通知
// -------------------------------------------------------------------------------

// SetStreamsNotifyUrlConfig 设置回调链接
// @param notifyUrl	String	是	设置直播流信息推送到的URL地址，必须以http://开头；
//                          详情请查看：@link https://help.aliyun.com/document_detail/54787.html?spm=5176.doc51836.6.673.KhDGLP
// @link https://help.aliyun.com/document_detail/35415.html?spm=5176.doc51835.6.670.Q6iOaA
func (l *Live) SetStreamsNotifyUrlConfig(notifyUrl string, resp interface{}) (err error) {
	req := l.cloneRequest(SetLiveStreamsNotifyUrlConfigAction)
	req.AppName = ""
	req.SetArgs("NotifyUrl", notifyUrl)
	err = l.rpc.Query(req, resp)
	return
}

// StreamsNotifyUrlConfig 查询推流回调配置
// @link https://help.aliyun.com/document_detail/51835.html?spm=5176.doc35415.6.671.ndNzEi
func (l *Live) StreamsNotifyUrlConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DescribeLiveStreamsNotifyUrlConfigAction)
	req.AppName = ""
	err = l.rpc.Query(req, resp)
	return
}

// DeleteLiveStreamsNotifyUrlConfig 删除推流回调配置
// @link https://help.aliyun.com/document_detail/51835.html?spm=5176.doc35415.6.671.ndNzEi
func (l *Live) DeleteLiveStreamsNotifyUrlConfig(resp interface{}) (err error) {
	req := l.cloneRequest(DeleteLiveStreamsNotifyUrlConfigAction)
	req.AppName = ""
	err = l.rpc.Query(req, resp)
	return
}
