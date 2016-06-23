package example

import (
	"aliyun-live-go-sdk/client"
	"aliyun-live-go-sdk/device/live"
	"aliyun-live-go-sdk/util"
	"time"
	"fmt"
)

const (
	AccessKeyId = "Pvu6bYlDPNkpAZJs"
	AccessKeySecret = "ENHXIhHhFZcD9cxAOn3sY72xQ9XueL"
)

func LiveExample(){
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	live := live.NewLive(cert, "alilive.strongwind.cn", "app-name").SetDebug(true)
	resp := make(map[string]interface{})
	live.StreamsPublishList(util.NewISO6801Time(time.Now().Add(-time.Hour * 12).UTC()), util.NewISO6801Time(time.Now().UTC()), &resp)
	fmt.Println(resp)
	resp1 := make(map[string]interface{})
	live.StreamsBlockList(resp1)
	fmt.Println(&resp1)
}
