package hook

import (
	"github.com/BPing/go-toolkit/http-client/core"
	"time"
	"fmt"
)

const (
	SlowReqRecord  = "SlowReqRecord"
	ReqRecord      = "ReqRecord"
	ErrorReqRecord = "ErrorReqRecord"

	// 默认慢请求时间
	defaultSlowReqLong = 5 * time.Second
)

type LogHook struct {
	// 超过SlowReqLong时间长度的请求，
	// 将记录为慢请求。
	// 如果为负数，代表不记录
	// 默认为2秒
	slowReqLong time.Duration

	// 函数参数
	// 记录信息；如日志记录
	record func(tag, msg string)
}

func (log *LogHook) BeforeRequest(req core.Request, client core.Client) error {
	return nil
}

func (log *LogHook) AfterRequest(cErr error, req core.Request, client core.Client) {
	if nil != cErr {
		if nil != log.record {
			log.record(ErrorReqRecord, fmt.Sprintf("query:: %s error:: %v ", req.String(), cErr))
		}
	} else {
		if nil != log.record {
			resp := req.Response()
			reqInfo := fmt.Sprintf(" http query:: %s status:%d \n response:%s \n ts:(%v) \n",
				req.String(),
				resp.StatusCode,
				resp.ToString(),
				req.ReqLongTime())
			if log.slowReqLong > 0 && req.ReqLongTime() >= log.slowReqLong {
				log.record(SlowReqRecord, reqInfo)
			}
			log.record(ReqRecord, reqInfo)
		}
	}
}

func NewLogHook(slowReqLong time.Duration, record func(tag, msg string)) *LogHook {
	if slowReqLong == 0 {
		slowReqLong = defaultSlowReqLong
	}
	return &LogHook{
		slowReqLong: slowReqLong,
		record:      record,
	}
}
