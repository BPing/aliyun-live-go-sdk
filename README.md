# aliyun-live-go-sdk
    阿里云直播 golang SDK

# Quick Start

```go
    package main
    
    import (
    "aliyun-live-go-sdk/client"
    "aliyun-live-go-sdk/device/live"
    "aliyun-live-go-sdk/util"
    "time"
    "fmt"
    )
    
    const (
    AccessKeyId = "<Yours' Id>"
    AccessKeySecret = "<...>"
    
    )
    
    func main(){
    cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
    live := live.NewLive(cert, "alilive.strongwind.cn", "app-name").SetDebug(true)
    resp := make(map[string]interface{})
    live.StreamsPublishList(util.NewISO6801Time(time.Now().Add(-time.Hour * 12).UTC()), util.NewISO6801Time(time.Now().UTC()), &resp)
    fmt.Println(resp)
    }
```
 
# References
The GO API design of Live refer the implementation from https://github.com/denverdino/aliyungo
