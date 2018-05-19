package core

// 为了兼容1.7或者更早的版本
type Context interface {
	Done() <-chan struct{}
	Err() error
	// 因为这下面两个方法在此包中没有使用到，所以不定义
	//Deadline() (deadline time.Time, ok bool)
	//Value(key interface{}) interface{}
}

type emptyCtx int

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}
func (*emptyCtx) Err() error {
	return nil
}
var (
	background = new(emptyCtx)
)
func BackgroundContext() Context {
	return background
}
