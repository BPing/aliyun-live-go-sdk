package core

// 钩子接口
type Hook interface {
	// 请求处理前执行
	// 如果返回错误
	// 将提前终止请求
	// 并将此错误返回
	BeforeRequest(req Request, client Client) error

	// 请求处理后执行
	// @params err 请求处理错误信息，如果不为nil，代表请求失败
	AfterRequest(cErr error, req Request, client Client)
}
