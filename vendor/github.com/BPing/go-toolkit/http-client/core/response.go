package core

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"bytes"
)

type ResponseFormat string

const (
	JSONResponseFormat = ResponseFormat("JSON")
	XMLResponseFormat  = ResponseFormat("XML")
)

var (
	RawRespNilErr     = errors.New("raw http response is nil. ")
	RawRespBodyNilErr = errors.New("the body of raw http response is nil. ")
)

// 封装标准库中的Response
// 方便处理响应内容信息
type Response struct {
	*http.Response
	body []byte //缓存响应的Response的body字节内容
}

// 响应的Response的body字节内容保存到文件中去(文件请求)
func (resp *Response) ToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if resp.Response == nil || resp.Response.Body == nil {
		return nil
	}
	defer resp.Response.Body.Close()
	_, err = io.Copy(f, resp.Response.Body)
	return err
}

// 返回响应的Response的body字节内容
func (resp *Response) Bytes() ([]byte, error) {
	if resp.body != nil {
		return resp.body, nil
	}
	if resp.Response == nil {
		return nil, RawRespNilErr
	}
	if resp.Response.Body == nil {
		return nil, RawRespBodyNilErr
	}
	var err error
	defer resp.Response.Body.Close()
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.body, err = ioutil.ReadAll(reader)
	} else {
		resp.body, err = ioutil.ReadAll(resp.Body)
	}
	// 为了*http.Response能再次使用，重新复制回去
	resp.Response.Body = ioutil.NopCloser(bytes.NewReader(resp.body))
	return resp.body, err
}

// 将响应的Response的body字节内容以JSON格式转化
func (resp *Response) ToJSON(v interface{}) error {
	data, err := resp.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// 将响应的Response的body字节内容以XML格式转化
func (resp *Response) ToXML(v interface{}) error {
	data, err := resp.Bytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

// 将响应的Response的body字节内容以字符串格式
// 如果为空的话，有可能是转化失败
func (resp *Response) ToString() string {
	data, err := resp.Bytes()
	if err != nil {
		return ""
	}
	return string(data)
}

func (resp *Response) Close() error {
	//resp.body = nil
	return nil
}
