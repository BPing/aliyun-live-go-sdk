package example

import (
	"fmt"
	"github.com/BPing/aliyun-live-go-sdk/aliyun"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
	"time"
)

// LiveExample live例子
func LiveExample() {

	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)
	resp := make(map[string]interface{})
	err := liveM.StreamsPublishList(time.Now().Add(-time.Hour*24*10), time.Now(), &resp)
	fmt.Println(resp)

	resp1 := live.OnlineInfoResponse{}
	err = liveM.StreamOnlineUserNum("test-video-name", &resp1)
	fmt.Println(err, resp1)

	resp2 := live.StreamListResponse{}
	err = liveM.StreamsBlockList(&resp2)
	fmt.Println(err, resp2)

	resp = make(map[string]interface{})
	err = liveM.StreamsControlHistory(time.Now().Add(-time.Hour*12), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.ForbidLiveStreamWithPublisher("test-video-name", nil, &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.ResumeLiveStreamWithPublisher("test-video-name", &resp)
	fmt.Println(err, resp)

	stream := liveM.GetStream("test-video-name")
	fmt.Println(stream.RtmpPublishUrl())
	fmt.Println(stream.String())

	// 录制

	oss := live.OssInfo{
		OssBucket:       OssBucket,
		OssEndpoint:     OssEndpoint,
		OssObject:       OssObject,
		OssObjectPrefix: OssObjectPrefix,
	}

	resp = make(map[string]interface{})
	err = liveM.AddLiveAppRecordConfig(oss, &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.CreateLiveStreamRecordIndexFiles("test-video-name", oss, time.Now().Add(-time.Hour*24*10), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveRecordConfig(&resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveStreamRecordContent("test-video-name", time.Now().Add(-time.Hour*24*10), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveStreamRecordIndexFiles("test-video-name", time.Now().Add(-time.Hour*24*10), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveStreamsFrameRateAndBitRateData("test-video-name", &resp)
	fmt.Println(err, resp)

	StreamExample()
}

// LiveSnapshotExample 截图例子
func LiveSnapshotExample() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	oss := live.OssInfo{
		OssBucket:       OssBucket,
		OssEndpoint:     OssEndpoint,
		OssObject:       OssObject,
		OssObjectPrefix: OssObjectPrefix,
	}
	config := live.SnapshotConfig{
		OssInfo:            oss,
		TimeInterval:       5,
		OverwriteOssObject: "{AppName}/{StreamName}.jpg",
	}

	fmt.Println("添加截图配置：")
	resp := make(map[string]interface{})
	err := liveM.AddLiveAppSnapshotConfig(config, &resp)
	fmt.Println(err, resp)

	config.SequenceOssObject = "{AppName}/{StreamName}.jpg"

	fmt.Println("更新截图配置：")
	resp = make(map[string]interface{})
	err = liveM.UpdateLiveAppSnapshotConfig(config, &resp)
	fmt.Println(err, resp)

	fmt.Println("查询域名截图配置：")
	param := live.LiveSnapshotParam{
		PageNum:  1,
		PageSize: 10,
		Order:    "asc",
	}
	resp = make(map[string]interface{})
	err = liveM.LiveSnapshotConfig(param, &resp)
	fmt.Println(err, resp)

	fmt.Println("查询域名截图配置(2):")
	respStruct := &live.LiveSnapshotConfigResponse{}
	err = liveM.LiveSnapshotConfig(param, respStruct)
	fmt.Println(err, respStruct)

	fmt.Println("查询截图信息")
	resp = make(map[string]interface{})
	err = liveM.LiveStreamSnapshotInfo("test-video-name1", time.Now().Add(-time.Hour*24*20), time.Now(), 10, &resp)
	fmt.Println(err, resp)

	fmt.Println("删除截图配置：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveAppSnapshotConfig(&resp)
	fmt.Println(err, resp)
}

// LiveTranscodeExample 转码
func LiveTranscodeExample() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("添加转码配置：")
	resp := make(map[string]interface{})
	err := liveM.AddLiveStreamTranscode("a", "no", "no", &resp)
	fmt.Println(err, resp)

	fmt.Println("查询转码配置信息：")
	resp = make(map[string]interface{})
	err = liveM.LiveStreamTranscodeInfo(&resp)
	fmt.Println(err, resp)

	fmt.Println("查询转码配置信息（2）：")
	respStruct := &live.StreamTranscodeInfoResponse{}
	err = liveM.LiveStreamTranscodeInfo(respStruct)
	fmt.Println(err, respStruct)

	fmt.Println("删除转码配置：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveStreamTranscode("a", &resp)
	fmt.Println(err, resp)
}

// 拉流
func LivePullStreamInfo() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("添加拉流信息：")
	resp := make(map[string]interface{})
	err := liveM.AddLivePullStreamInfoConfig("test-video-name", "http://", time.Now().Add(-time.Hour*24*20), time.Now(), &resp)
	fmt.Println(err, resp)

	fmt.Println("查看：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLivePullStreamConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("删除拉流信息：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLivePullStreamInfoConfig("test-video-name", &resp)
	fmt.Println(err, resp)

	fmt.Println("查看：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLivePullStreamConfig(&resp)
	fmt.Println(err, resp)
}

// 状态通知
func NotifyUrlConfig() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("设置回调链接：")
	resp := make(map[string]interface{})
	err := liveM.SetStreamsNotifyUrlConfig("http://1.1.1.1:8888", &resp)
	fmt.Println(err, resp)

	fmt.Println("查看：")
	resp = make(map[string]interface{})
	err = liveM.StreamsNotifyUrlConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("删除推流回调配置：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveStreamsNotifyUrlConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("查看：")
	resp = make(map[string]interface{})
	err = liveM.StreamsNotifyUrlConfig(&resp)
	fmt.Println(err, resp)
}

//
// 直播转点播
func RecordVodExample() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("增加直播录制转点播配置：")
	resp := make(map[string]interface{})
	err := liveM.AddLiveRecordVodConfig("A", 300, &resp)
	fmt.Println(err, resp)

	fmt.Println("查询直转点配置列表：")
	params := live.DescribeVodParam{
		StreamName: "test-video-name",
	}
	vodResp := live.RecordVodConfigsResponse{}
	err = liveM.DescribeLiveRecordVodConfigs(params, &vodResp)
	fmt.Println(err, vodResp)

	fmt.Println("删除直播录制转点播配置：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveRecordVodConfig("test-video-name", &resp)
	fmt.Println(err, resp)
}

// 直播审核
func VerifyExample() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("添加审核配置：")
	params := live.AddSnapshotDetectPornParam{
		OssBucket:   OssBucket,
		OssEndpoint: OssEndpoint,
		OssObject:   OssObject,
	}
	resp := make(map[string]interface{})
	err := liveM.AddLiveSnapshotDetectPornConfig(params, &resp)
	fmt.Println(err, resp)

	fmt.Println("查询审核配置：")
	sdppcResp := live.SnapshotDetectPornConfigResponse{}
	err = liveM.DescribeLiveSnapshotDetectPornConfig(live.SnapshotDetectPornParam{
		Order: live.DescOrderType,
	}, &sdppcResp)
	fmt.Println(err, sdppcResp)

	fmt.Println("更新审核回调：")
	uParams := live.AddSnapshotDetectPornParam{
		OssBucket:   OssBucket,
		OssEndpoint: OssEndpoint,
		OssObject:   OssObject,
		SceneN:      live.TerrorismSceneN,
	}
	resp = make(map[string]interface{})
	err = liveM.UpdateLiveSnapshotDetectPornConfig(uParams, &resp)
	fmt.Println(err, resp)

	fmt.Println("删除审核回调：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveSnapshotDetectPornConfig(live.LiveBase{}, &resp)
	fmt.Println(err, resp)

	fmt.Println("添加回调通知：")
	resp = make(map[string]interface{})
	err = liveM.AddLiveDetectNotifyConfig("http://www.yourdomain.cn/examplecallback.action", &resp)
	fmt.Println(err, resp)

	fmt.Println("查看回调通知：")
	resp = make(map[string]interface{})
	dncResp := live.DetectNotifyConfigResponse{}
	err = liveM.DescribeLiveDetectNotifyConfig(&dncResp)
	fmt.Println(err, dncResp)

	fmt.Println("更新回调通知：")
	resp = make(map[string]interface{})
	err = liveM.UpdateLiveDetectNotifyConfig("http://www.yourdomain.cn/examplecallback", &resp)
	fmt.Println(err, resp)

	fmt.Println("删除审核回调：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveDetectNotifyConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("查看回调通知：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLiveDetectNotifyConfig(&dncResp)
	fmt.Println(err, dncResp)
}

// 资源监控
func MonitorExample() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("查询直播域名的网络带宽监控数据：")
	dbdEesp := live.DomainBpsDataResponse{}
	err := liveM.DescribeLiveDomainBpsData(time.Now().Add(-time.Hour*24*10), time.Now().Add(time.Hour), &dbdEesp)
	fmt.Println(err, dbdEesp)

	fmt.Println("查询直播域名录制时长数据：")
	rdEesp := live.RecordDataInfoResponse{}
	err = liveM.DescribeLiveDomainRecordData(time.Now().Add(-time.Hour*24), time.Now().Add(-time.Hour*8), live.NilRecordType, &rdEesp)
	fmt.Println(err, rdEesp)

	fmt.Println("查询直播域名截图张数数据：")
	sdEesp := live.SnapshotDataInfoResponse{}
	err = liveM.DescribeLiveDomainSnapshotData(time.Now().Add(-time.Hour*24), time.Now().Add(-time.Hour*8), &sdEesp)
	fmt.Println(err, sdEesp)

	fmt.Println("查询直播域名网络流量监控数据：")
	dtdEesp := live.DomainTrafficDataResponse{}
	err = liveM.DescribeLiveDomainTrafficData(time.Now().Add(-time.Hour*24), time.Now().Add(-time.Hour), &dtdEesp)
	fmt.Println(err, dtdEesp)

	fmt.Println("查询直播域名转码时长数据：")
	tdIEesp := live.TranscodeDataInfoResponse{}
	err = liveM.DescribeLiveDomainTranscodeData(time.Now().Add(-time.Hour*24), time.Now().Add(-time.Hour*8), &tdIEesp)
	fmt.Println(err, tdIEesp)

	fmt.Println("查询直播流历史在线人数：")
	sunResp := live.StreamUserNumInfoResponse{}
	err = liveM.DescribeLiveStreamHistoryUserNum("test-video-name", time.Now().Add(-time.Hour*24),  time.Now().Add(-time.Hour*8), &sunResp)
	fmt.Println(err, sunResp)
}

// 状态通知
func MixStream() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)

	fmt.Println("添加连麦配置：")
	resp := make(map[string]interface{})
	err := liveM.AddLiveMixConfig("mhd ", &resp)
	fmt.Println(err, resp)

	fmt.Println("查询连麦配置：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLiveMixConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("删除连麦配置：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveMixConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("查询连麦配置：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLiveMixConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("开启多人连麦服务：")
	resp = make(map[string]interface{})
	err = liveM.StartMultipleStreamMixService("test-video-name", "pip4a", &resp)
	fmt.Println(err, resp)

	fmt.Println("停止多人连麦服务：")
	resp = make(map[string]interface{})
	err = liveM.StopMultipleStreamMixService("test-video-name", &resp)
	fmt.Println(err, resp)

	config := live.MixStreamParam{
		Mix: live.StreamBase{
			LiveBase: live.LiveBase{
				DomainName: DomainName,
				AppName:    AppName,
			},
			StreamName: "test-video-name-mix",
		},
		Main: live.StreamBase{
			LiveBase: live.LiveBase{
				DomainName: DomainName,
				AppName:    AppName,
			},
			StreamName: "test-video-name",
		}}
	fmt.Println("往主流添加一路流：")
	resp = make(map[string]interface{})
	err = liveM.AddMultipleStreamMixService(config, &resp)
	fmt.Println(err, resp)

	fmt.Println("从主流移除一路流：")
	resp = make(map[string]interface{})
	err = liveM.RemoveMultipleStreamMixService(config, &resp)
	fmt.Println(err, resp)

	fmt.Println("添加连麦回调配置：")
	resp = make(map[string]interface{})
	err = liveM.AddLiveMixNotifyConfig("http://1.1.1.1:8888", &resp)
	fmt.Println(err, resp)

	fmt.Println("查询连麦回调配置：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLiveMixNotifyConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("更新连麦回调配置：")
	resp = make(map[string]interface{})
	err = liveM.UpdateLiveMixNotifyConfig("http://1.1.1.1:8889", &resp)
	fmt.Println(err, resp)

	fmt.Println("查询连麦回调配置：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLiveMixNotifyConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("添加连麦回调配置：")
	resp = make(map[string]interface{})
	err = liveM.DeleteLiveMixNotifyConfig(&resp)
	fmt.Println(err, resp)

	fmt.Println("查询连麦回调配置：")
	resp = make(map[string]interface{})
	err = liveM.DescribeLiveMixNotifyConfig(&resp)
	fmt.Println(err, resp)
}

// StreamExample 流
func StreamExample() {
	cert := aliyun.NewCredentials(AccessKeyID, AccessKeySecret)
	streamCert := live.NewStreamCredentials(PrivateKey, live.DefaultStreamTimeout)
	liveM := live.NewLive(cert, DomainName, AppName, streamCert).SetDebug(false)
	stream := liveM.GetStream("test-video-name")
	stream.RtmpPublishUrl()

	fmt.Println("Online", stream.Online())
	fmt.Println("Blocked", stream.Blocked())
	fmt.Println(stream.String())

	stream1 := liveM.GetStream("test-video-name1")
	fmt.Println("Blocked", stream1.Blocked())

	oss := live.OssInfo{
		OssBucket:       OssBucket,
		OssEndpoint:     OssEndpoint,
		OssObject:       OssObject,
		OssObjectPrefix: OssObjectPrefix,
	}

	//录制
	fmt.Println(stream1.CreateRecordIndexFiles(oss, time.Now().Add(-time.Hour*24*10), time.Now()))

	fmt.Println(stream1.RecordContent(time.Now().Add(-time.Hour*24*20), time.Now()))

	fmt.Println(stream1.RecordIndexFiles(time.Now().Add(-time.Hour*24*20), time.Now()))

	fmt.Println(stream1.FrameRateAndBitRateData())

	//截图
	fmt.Println("查询截图信息：")
	fmt.Println(stream1.SnapshotInfo(time.Now().Add(-time.Hour*24*20), time.Now(), 10))

	fmt.Println(stream1.StartMultipleStreamMixService("pip4a"))

	fmt.Println(stream1.StopMultipleStreamMixService())

	fmt.Println(stream1.AddMultipleStream(
		live.StreamBase{
			LiveBase: live.LiveBase{
				DomainName: DomainName,
				AppName:    AppName,
			},
			StreamName: "test-video-name-mix",
		},
	))

	fmt.Println(stream1.RemoveMultipleStream(
		live.StreamBase{
			LiveBase: live.LiveBase{
				DomainName: DomainName,
				AppName:    AppName,
			},
			StreamName: "test-video-name-mix",
		},
	))
}
