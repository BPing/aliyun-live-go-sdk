package example

import (
	"fmt"
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
	"time"
)

// LiveExample live例子
func LiveExample() {

	cert := client.NewCredentials(AccessKeyID, AccessKeySecret)
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
	cert := client.NewCredentials(AccessKeyID, AccessKeySecret)
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
	cert := client.NewCredentials(AccessKeyID, AccessKeySecret)
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
	cert := client.NewCredentials(AccessKeyID, AccessKeySecret)
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
	cert := client.NewCredentials(AccessKeyID, AccessKeySecret)
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

// StreamExample 流
func StreamExample() {
	cert := client.NewCredentials(AccessKeyID, AccessKeySecret)
	streamCert := live.NewStreamCredentials(PrivateKey, live.DefualtStreamTimeout)
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
}
