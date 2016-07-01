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
	liveM.StreamsPublishList(time.Now().Add(-time.Hour * 24 * 10), time.Now(), &resp)
	fmt.Println(resp)
	resp1 := live.OnlineInfoResponse{}
	err := liveM.StreamOnlineUserNum("test-video-name", &resp1)
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
}