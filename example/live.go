package example

import (
	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
	"time"
	"fmt"
)

func LiveExample() {

	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(false)
	resp := make(map[string]interface{})
	err := liveM.StreamsPublishList(time.Now().Add(-time.Hour * 24 * 10), time.Now(), &resp)
	fmt.Println(resp)

	resp1 := live.OnlineInfoResponse{}
	err = liveM.StreamOnlineUserNum("test-video-name", &resp1)
	fmt.Println(err, resp1)

	resp2 := live.StreamListResponse{}
	err = liveM.StreamsBlockList(&resp2)
	fmt.Println(err, resp2)

	resp = make(map[string]interface{})
	err = liveM.StreamsControlHistory(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
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
		OssBucket:OssBucket,
		OssEndpoint:OssEndpoint,
		OssObject:OssObject,
		OssObjectPrefix:OssObjectPrefix,
	}

	resp = make(map[string]interface{})
	err = liveM.AddLiveAppRecordConfig(oss, &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.CreateLiveStreamRecordIndexFiles("test-video-name", oss, time.Now().Add(-time.Hour * 24 * 10), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveRecordConfig(&resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveStreamRecordContent("test-video-name", time.Now().Add(-time.Hour * 24 * 10), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveStreamRecordIndexFiles("test-video-name", time.Now().Add(-time.Hour * 24 * 10), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.DescribeLiveStreamsFrameRateAndBitRateData("test-video-name", &resp)
	fmt.Println(err, resp)

	StreamExample()
}

func StreamExample() {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
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
		OssBucket:OssBucket,
		OssEndpoint:OssEndpoint,
		OssObject:OssObject,
		OssObjectPrefix:OssObjectPrefix,
	}

	//录制
	fmt.Println(stream1.CreateRecordIndexFiles(oss, time.Now().Add(-time.Hour * 24 * 10), time.Now()))

	fmt.Println(stream1.RecordContent(time.Now().Add(-time.Hour * 24 * 20), time.Now()))

	fmt.Println(stream1.RecordIndexFiles(time.Now().Add(-time.Hour * 24 * 20), time.Now()))

	fmt.Println(stream1.FrameRateAndBitRateData())
}