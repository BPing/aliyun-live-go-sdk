# aliyun-live-go-sdk [![Build Status](https://travis-ci.org/BPing/aliyun-live-go-sdk.svg?branch=master)](https://travis-ci.org/BPing/aliyun-live-go-sdk) [![Coverage Status](https://coveralls.io/repos/github/BPing/aliyun-live-go-sdk/badge.svg?branch=master)](https://coveralls.io/github/BPing/aliyun-live-go-sdk?branch=master)
    阿里云直播 golang SDK

# 快速开始

```go
package main

import (
    "github.com/BPing/aliyun-live-go-sdk/client"
    "github.com/BPing/aliyun-live-go-sdk/device/live"
    "github.com/BPing/aliyun-live-go-sdk/util"
    "time"
    "fmt"
)

const (
    AccessKeyId = "<Yours' Id>"
    AccessKeySecret = "<...>"
)

func main() {
    cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
    live := live.NewLive(cert, "<Yours' CDN>", "app-name",nil).SetDebug(true)
    resp := make(map[string]interface{})
    live.StreamsPublishList(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(resp)
}
```
 
# 构建安装

go get:

```sh
go get github.com/BPing/aliyun-live-go-sdk
```

# 文档
* Sdk:https://godoc.org/github.com/BPing/aliyun-live-go-sdk [![GoDoc](https://godoc.org/github.com/BPing/aliyun-live-go-sdk?status.svg)](https://godoc.org/github.com/BPing/aliyun-live-go-sdk)
* CDN:https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/cdn   [![GoDoc](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/cdn?status.svg)](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/cdn)
* Live:https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/live [![GoDoc](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/live?status.svg)](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/live)

# Example
* [example](https://github.com/BPing/aliyun-live-go-sdk/tree/master/example)

## 直播(Live)

        方法名以"WithApp"结尾代表可以更改请求中  "应用名字（AppName）"，
        否则按默认的(初始化时设置的AppName)。
        如果为空，代表忽略参数AppName

```go
    cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
    liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(true)
```

* 获取流列表
```go
    resp := make(map[string]interface{})
    liveM.StreamsPublishList(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(resp)
    // @appname 应用名 为空时，忽略此参数
    resp := make(map[string]interface{})
    liveM.StreamsPublishListWithApp(AppName,time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(resp)
```

* 获取黑名单
```go
    resp = make(map[string]interface{})
    err = liveM.StreamsBlockList(&resp)
    fmt.Println(err, resp)
```

* 获取流的在线人数
```go
    resp1 := live.OnlineInfoResponse{}
    err := liveM.StreamOnlineUserNum("video-name", &resp1)
    fmt.Println(err, resp1)  
   // @appname 应用名 为空时，忽略此参数
    resp1 := live.OnlineInfoResponse{}
    err := liveM.StreamOnlineUserNumWithApp(AppName,"video-name", &resp1)
    fmt.Println(err, resp1)  
```

* 获取控制历史
```go
    resp = make(map[string]interface{})
    err = liveM.StreamsControlHistory(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    //err = liveM.StreamsControlHistoryWithApp(AppName,time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(err, resp)
```

* 禁止
```go
    resp = make(map[string]interface{})
    err = liveM.ForbidLiveStreamWithPublisher("video-name", nil, &resp)
    fmt.Println(err, resp)
```

* 恢复
```go
    resp = make(map[string]interface{})
    err = liveM.ResumeLiveStreamWithPublisher("video-name", &resp)
    fmt.Println(err, resp)
```

### 录制（请看文档）
### 截图（2017-01-18）

* 添加截图配置
```go
	oss := live.OssInfo{
		OssBucket:       OssBucket,
		OssEndpoint:     OssEndpoint,
		OssObject:       OssObject,
		OssObjectPrefix: OssObjectPrefix,
	}
	config:=live.SnapshotConfig{
		OssInfo:oss,
		TimeInterval       : 5,
		OverwriteOssObject  : "{AppName}/{StreamName}.jpg",
	}
  	resp := make(map[string]interface{})
  	err:=liveM.AddLiveAppSnapshotConfig(config,&resp)
  	fmt.Println(err, resp)
```

* 更新截图配置
```go
	config.SequenceOssObject="{AppName}/{StreamName}.jpg"
	resp = make(map[string]interface{})
	err=liveM.UpdateLiveAppSnapshotConfig(config,&resp)
	fmt.Println(err, resp)
```

* 查询域名截图配置
```go
	param:=live.LiveSnapshotParam{
    		PageNum:1,
    		PageSize:10,
    		Order:"asc",
    	}
    	resp = make(map[string]interface{})
    	err=liveM.LiveSnapshotConfig(param,&resp)
    	fmt.Println(err, resp)
```

* 查询截图信息
```go
	resp = make(map[string]interface{})
	err=liveM.LiveStreamSnapshotInfo("test-video-name1",time.Now().Add(-time.Hour*24*20), time.Now(),10,&resp)
	fmt.Println(err, resp)
```

* 删除截图配置
```go
	resp = make(map[string]interface{})
	err=liveM.DeleteLiveAppSnapshotConfig(&resp)
	fmt.Println(err, resp)
```


### 转码（2017-01-19）
* 添加转码配置
```go
	resp := make(map[string]interface{})
	err:=liveM.AddLiveStreamTranscode("a","no","no",&resp)
	fmt.Println(err, resp)
```

* 查询转码配置信息
```go
	resp = make(map[string]interface{})
	err=liveM.LiveStreamTranscodeInfo(&resp)
	fmt.Println(err, resp))
```

* 删除转码配置
```go
	resp = make(map[string]interface{})
	err=liveM.DeleteLiveStreamTranscode("a",&resp)
	fmt.Println(err, resp)
```

### 混流（2017-01-19）

* 开始混流操作
```go
	err=liveM.StartMixStreamsService(...)
```

* 结束混流操作
```go
	err=liveM.StopMixStreamsService(...)
```

## 流(Stream)
```go
  //如果 streamCert 为空的话，则代表不开启直播流鉴权
   cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
   streamCert := live.NewStreamCredentials(PrivateKey, live.DefualtStreamTimeout)
   liveM := live.NewLive(cert, DomainName, AppName, streamCert)
  // GetStream 获取直播流
  // @describe 每一次都生成新的流实例，不检查流名的唯一性，并且同一个名字会生成不同的实例的，
  //          所以，使用时候，请自行确保流名的唯一性
   stream := liveM.GetStream("video-name")
```

* 获取RTMP推流地址
```go
    // RTMP 推流地址
    // 如果开启了直播流鉴权，签名失效后，会重新生成新的有效的推流地址
    stream.RtmpPublishUrl()
```

* RTMP 直播播放地址
```go
    url:=stream.RtmpLiveUrls()
```

* HLS 直播播放地址
```go
    url:=stream.HlsLiveUrls()
```

* FLV 直播播放地址
```go
    url:=stream.HttpFlvLiveUrls()
```

* 获取在线人数
```go
    num:=stream.OnlineUserNum()
```

* 是否在线
```go
    isOnline:=stream.Online()
```

* 是否被禁止
```go
    isBlocked:=stream.Blocked()
```

* 获取直播流的帧率和码率
```go
    // type FrameRateAndBitRateInfos struct {
    //	FrameRateAndBitRateInfo []FrameRateAndBitRateInfo
    //}
    //
    //// 各直播流的帧率/码率信息
    //type FrameRateAndBitRateInfo struct {
    //	StreamUrl      string // 直播流的URL
    //	VideoFrameRate int    // 直播流的视频帧率
    //	AudioFrameRate int    // 直播流的音频帧率
    //	BitRate        int    // 直播流的码率
    //}  
    frameRateAndBitRateInfo,err:=stream.FrameRateAndBitRateData()
```

* 获取截图信息
```go
// 查询截图信息
    // type StreamSnapshotInfoResponse struct {
    //    client.Response
    //	  LiveStreamSnapshotInfoList struct {
    //		StreamSnapshotInfo []StreamSnapshotInfo `json:"StreamSnapshotInfo" xml:"StreamSnapshotInfo"`
    //	 } `json:"LiveStreamSnapshotInfoList" xml:"LiveStreamSnapshotInfoList"` //截图内容列表，没有则返回空数组
    //    NextStartTime              string                //
    // }
    streamSnapshotInfo,err:=stream.SnapshotInfo(time.Now().Add(-time.Hour * 12), time.Now(), 10)
```



# 贡献参与者
* cbping(452775680@qq.com)

# License
 采用 [Apache License, Version 2.0](https://github.com/denverdino/aliyungo/blob/master/LICENSE.txt)许可证授权原则。
 
# 参考文献
参考项目： https://github.com/denverdino/aliyungo
