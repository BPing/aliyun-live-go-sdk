package live

const (
	//action
	DescribeLiveStreamsPublishListAction    = "DescribeLiveStreamsPublishList"
	DescribeLiveStreamsOnlineListAction     = "DescribeLiveStreamsOnlineList"
	DescribeLiveStreamsBlockListAction      = "DescribeLiveStreamsBlockList"
	DescribeLiveStreamsControlHistoryAction = "DescribeLiveStreamsControlHistory"
	ForbidLiveStreamAction                  = "ForbidLiveStream"
	ResumeLiveStreamAction                  = "ResumeLiveStream"

	// 录制处理
	AddLiveAppRecordConfigAction                     = "AddLiveAppRecordConfig"
	CreateLiveStreamRecordIndexFilesAction           = "CreateLiveStreamRecordIndexFiles"
	DeleteLiveAppRecordConfigAction                  = "DeleteLiveAppRecordConfig"
	DescribeLiveAppRecordConfigAction                = "DescribeLiveAppRecordConfig"
	DescribeLiveRecordConfigAction                   = "DescribeLiveRecordConfig"
	DescribeLiveStreamRecordContentAction            = "DescribeLiveStreamRecordContent"
	DescribeLiveStreamRecordIndexFileAction          = "DescribeLiveStreamRecordIndexFile"
	DescribeLiveStreamRecordIndexFilesAction         = "DescribeLiveStreamRecordIndexFiles"
	DescribeLiveStreamsFrameRateAndBitRateDataAction = "DescribeLiveStreamsFrameRateAndBitRateData"

	// 截图处理
	AddLiveAppSnapshotConfigAction       = "AddLiveAppSnapshotConfig"
	UpdateLiveAppSnapshotConfigAction    = "UpdateLiveAppSnapshotConfig"
	DeleteLiveAppSnapshotConfigAction    = "DeleteLiveAppSnapshotConfig"
	DescribeLiveSnapshotConfigAction     = "DescribeLiveSnapshotConfig"
	DescribeLiveStreamSnapshotInfoAction = "DescribeLiveStreamSnapshotInfo"

	// 转码处理
	AddLiveStreamTranscodeAction          = "AddLiveStreamTranscode"
	DeleteLiveStreamTranscodeAction       = "DeleteLiveStreamTranscode"
	DescribeLiveStreamTranscodeInfoAction = "DescribeLiveStreamTranscodeInfo"

	// 混流处理
	StartMixStreamsServiceAction = "StartMixStreamsService"
	StopMixStreamsServiceAction  = "StopMixStreamsService"

	// 直播连麦
	AddLiveMixConfigAction               = "AddLiveMixConfig"
	DescribeLiveMixConfigAction          = "DescribeLiveMixConfig"
	DeleteLiveMixConfigAction            = "DeleteLiveMixConfig"
	StartMultipleStreamMixServiceAction  = "StartMultipleStreamMixService"
	StopMultipleStreamMixServiceAction   = "StopMultipleStreamMixService"
	AddMultipleStreamMixServiceAction    = "AddMultipleStreamMixService"
	RemoveMultipleStreamMixServiceAction = "RemoveMultipleStreamMixService"
	AddLiveMixNotifyConfigAction         = "AddLiveMixNotifyConfig"
	DescribeLiveMixNotifyConfigAction    = "DescribeLiveMixNotifyConfig"
	UpdateLiveMixNotifyConfigAction      = "UpdateLiveMixNotifyConfig"
	DeleteLiveMixNotifyConfigAction      = "DeleteLiveMixNotifyConfig"

	// 直播拉流
	AddLivePullStreamInfoConfigAction    = "AddLivePullStreamInfoConfig"
	DeleteLivePullStreamInfoConfigAction = "DeleteLivePullStreamInfoConfig"
	DescribeLivePullStreamConfigAction   = "DescribeLivePullStreamConfig"

	// 状态通知
	SetLiveStreamsNotifyUrlConfigAction      = "SetLiveStreamsNotifyUrlConfig"
	DescribeLiveStreamsNotifyUrlConfigAction = "DescribeLiveStreamsNotifyUrlConfig"
	DeleteLiveStreamsNotifyUrlConfigAction   = "DeleteLiveStreamsNotifyUrlConfig"

	// 直播转点播
	AddLiveRecordVodConfigAction       = "AddLiveRecordVodConfig"
	DeleteLiveRecordVodConfigAction    = "DeleteLiveRecordVodConfig"
	DescribeLiveRecordVodConfigsAction = "DescribeLiveRecordVodConfigs"

	// 资源监控
	DescribeLiveDomainBpsDataAction        = "DescribeLiveDomainBpsData"
	DescribeLiveDomainRecordDataAction     = "DescribeLiveDomainRecordData"
	DescribeLiveDomainSnapshotDataAction   = "DescribeLiveDomainSnapshotData"
	DescribeLiveDomainTrafficDataAction    = "DescribeLiveDomainTrafficData"
	DescribeLiveDomainTranscodeDataAction  = "DescribeLiveDomainTranscodeData"
	DescribeLiveStreamHistoryUserNumAction = "DescribeLiveStreamHistoryUserNum"
	DescribeLiveStreamOnlineUserNumAction  = "DescribeLiveStreamOnlineUserNum"

	// 直播审核
	AddLiveSnapshotDetectPornConfigAction      = "AddLiveSnapshotDetectPornConfig"
	AddLiveDetectNotifyConfigAction            = "AddLiveDetectNotifyConfig"
	DescribeLiveSnapshotDetectPornConfigAction = "DescribeLiveSnapshotDetectPornConfig"
	DescribeLiveDetectNotifyConfigAction       = "DescribeLiveDetectNotifyConfig"
	UpdateLiveSnapshotDetectPornConfigAction   = "UpdateLiveSnapshotDetectPornConfig"
	UpdateLiveDetectNotifyConfigAction         = "UpdateLiveDetectNotifyConfig"
	DeleteLiveSnapshotDetectPornConfigAction   = "DeleteLiveSnapshotDetectPornConfig"
	DeleteLiveDetectNotifyConfigAction         = "DeleteLiveDetectNotifyConfig"

	//直播中心服务器域名
	DefaultVideoCenter = "video-center.alivecdn.com"

	APILiveVersion = "2016-11-01" //2016-11-01
	APILiveHost    = "https://live.aliyuncs.com"
)

type RecordType string

const (
	NilRecordType RecordType = "" // 所有类型
	TSRecordType  RecordType = "TS"
	MP4RecordType RecordType = "MP4"
	FLVRecordType RecordType = "FLV"
)

type SceneN string

const (
	NilSceneN       SceneN = "" // 所有类型
	PornSceneN      SceneN = "porn"
	TerrorismSceneN SceneN = "terrorism"
	AdSceneN        SceneN = "ad"
	LiveSceneN      SceneN = "live"
)

type OrderType string

const (
	NilOrderType  OrderType = ""
	AscOrderType  OrderType = "asc"
	DescOrderType OrderType = "desc"
)
