# aliyun-live-go-sdk
    阿里云直播 golang SDK

# 快速开始

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

func main() {
    cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
    live := live.NewLive(cert, "<Yours' CDN>", "app-name").SetDebug(true)
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

# 贡献参与者
* cbping(452775680@qq.com)

# License
这个项目是采用 [Apache License, Version 2.0](https://github.com/denverdino/aliyungo/blob/master/LICENSE.txt)许可证授权原则。
 
# 参考文献
参考项目： https://github.com/denverdino/aliyungo
