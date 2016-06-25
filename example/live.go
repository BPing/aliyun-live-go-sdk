package example

import (
	"aliyun-live-go-sdk/client"
	"aliyun-live-go-sdk/device/live"
	"time"
	"fmt"
)

const (
	AccessKeyId = ""
	AccessKeySecret = ""
	DomainName = ""
	AppName = "app-name"
	PrivateKey = ""
)

func LiveExample() {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(true)
	resp := make(map[string]interface{})
	liveM.StreamsPublishList(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
	fmt.Println(resp)
	resp1 := live.OnlineInfoResponse{}
	err := liveM.StreamOnlineUserNum("video-name", &resp1)
	fmt.Println(err, resp1)

	resp = make(map[string]interface{})
	err = liveM.StreamsBlockList(&resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.StreamsControlHistory(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.ForbidLiveStreamWithPublisher("video-name", nil, &resp)
	fmt.Println(err, resp)

	resp = make(map[string]interface{})
	err = liveM.ResumeLiveStreamWithPublisher("video-name", &resp)
	fmt.Println(err, resp)

	stream := liveM.GetStream("video-name")
	fmt.Println(stream.RtmpPublishUrl())
	fmt.Println(stream.String())

	StreamExample()
}

func StreamExample() {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	streamCert := live.NewStreamCredentials(PrivateKey, live.DefualtStreamTimeout)
	liveM := live.NewLive(cert, DomainName, AppName, streamCert).SetDebug(true)
	stream := liveM.GetStream("video-name")
	stream.RtmpPublishUrl()

	fmt.Println(stream.Online())
	fmt.Println(stream.String())
}