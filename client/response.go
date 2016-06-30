package client

import "fmt"

type Response struct {
	RequestId string
}

//错误信息结构体
type ErrorResponse struct {
	RequestId string
	HostId    string
	Code      string
	Message   string
	StatusCode int //Status Code of HTTP Response
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Aliyun API Error: RequestId: %s Status Code: %d Code: %s Message: %s", e.RequestId, e.StatusCode, e.Code, e.Message)
}