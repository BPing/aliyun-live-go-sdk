//Copyright cbping
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License

// client 包 核心代码包，请求相关代码，
// 如：client request response
// 其中也包括整个sdk的一下基本信息
// @author cbping
package client

import (
	"net/http"
	"errors"
	"time"
	"encoding/json"
	"io/ioutil"
	"net"
	"strings"
	"encoding/xml"
	"log"
)

//
//  客户端
//  处理http请求
type Client struct {
	*Credentials
	//ConnectTimeout小于或等于零时，
	// 采用默认&http.Client{}
	httpClient     *http.Client
	//版本号
	version        string
	//
	debug          bool

	ConnectTimeout time.Duration
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// 初始化
func (c *Client) Init() *Client {
	if (c.ConnectTimeout <= 0) {
		c.httpClient = &http.Client{}
	}else {
		c.httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   c.ConnectTimeout,
				}).Dial,
			},
		}
	}
	return c
}

type Unmarshal func(data []byte, v interface{}) error

// 默认json解析
func (c *Client)responseUnmarshal(req Request) (u Unmarshal) {
	if (req.ResponseFormat() == XMLResponseFormat) {
		u = xml.Unmarshal
	}else {
		u = json.Unmarshal
	}
	return
}

// 处理请求
func (c *Client)Query(req Request, resp interface{}) error {
	if (nil == c.httpClient) {
		return clientError(errors.New("httpClient is nil"))
	}

	if (nil == req) {
		return clientError(errors.New("Request is nil"))
	}

	req.Sign(c.Credentials)
	httpReq, err := req.HttpRequestInstance()
	if (nil != err) {
		return clientError(err)
	}

	//必要头部信息设置
	httpReq.Header.Set("X-SDK-Client", `AliyunLiveGoSDK/` + Version)
	if (req.ResponseFormat() == XMLResponseFormat) {
		httpReq.Header.Set("Content-Type", `application/` + strings.ToLower(XMLResponseFormat))
	}else {
		httpReq.Header.Set("Content-Type", `application/` + strings.ToLower(JSONResponseFormat))
	}

	t0 := time.Now()
	httpResp, err := c.httpClient.Do(httpReq)
	t1 := time.Now()
	if (nil != err) {
		return clientError(err)
	}

	if c.debug {
		log.Printf("http query %s %d (%v) ", req.String(), httpResp.StatusCode, t1.Sub(t0))
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return clientError(err)
	}

	if c.debug {
		log.Printf("body of response:%s", string(body))
	}

	respUnmarshal := c.responseUnmarshal(req)
	//失败响应处理
	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = respUnmarshal(body, &errorResponse)
		errorResponse.StatusCode = httpResp.StatusCode
		return &errorResponse
	}

	err = respUnmarshal(body, resp)
	if err != nil {
		return clientError(err)
	}

	if c.debug {
		log.Printf("AliyunLiveGoClient.> decoded response into %#v", resp)
	}

	//if (req.ResponseFormat() == XMLResponseFormat) {
	//	//Xml
	//
	//	//失败响应处理
	//	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
	//		errorResponse := ErrorResponse{}
	//		err = xml.NewDecoder(httpResp.Body).Decode(&errorResponse)
	//		xml.Unmarshal(body, &errorResponse)
	//		errorResponse.StatusCode = httpResp.StatusCode
	//		return errorResponse
	//	}
	//	err = xml.NewDecoder(httpResp.Body).Decode(resp)
	//	if err != nil {
	//		return clientError(err)
	//	}
	//}else {
	//	//Json
	//
	//	//失败响应处理
	//	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
	//		errorResponse := ErrorResponse{}
	//		err = json.NewDecoder(httpResp.Body).Decode(&errorResponse)
	//		errorResponse.StatusCode = httpResp.StatusCode
	//		return errorResponse
	//	}
	//
	//	err = json.NewDecoder(httpResp.Body).Decode(&errorResponse)
	//	if err != nil {
	//		return clientError(err)
	//	}
	//}

	return nil

}

func NewClientTimeout(cert *Credentials, connectTimeout time.Duration) (c *Client) {
	c = (&Client{
		Credentials:cert,
		ConnectTimeout:connectTimeout,
		version:Version,
		debug:false,
	}).Init()
	return
}

func NewClient(cert *Credentials) (c *Client) {
	return NewClientTimeout(cert, time.Duration(0))
}

func clientError(err error) error {
	if (nil == err) {
		return nil
	}
	return errors.New("AliyunLiveGoClientFailure:" + err.Error())
}