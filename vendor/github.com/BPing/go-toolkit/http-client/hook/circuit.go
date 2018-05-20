package hook

import (
	"fmt"
	"errors"
	"github.com/BPing/go-toolkit/http-client/core"
	"sync"
	"time"
)

// 断路器状态
type State int

// 状态常量
const (
	StateClosed   State = iota
	StateHalfOpen
	StateOpen
)

var (
	ErrTooManyRequests = errors.New("too many requests")
	ErrOpenState       = errors.New("circuit breaker is open")
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateHalfOpen:
		return "half-open"
	case StateOpen:
		return "open"
	default:
		return fmt.Sprintf("unknown state: %d", s)
	}
}

// 记录请求的数量以及失败和成功数量
type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

func (c *Counts) onRequest() {
	c.Requests++
}

func (c *Counts) onSuccess() {
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

func (c *Counts) onFailure() {
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

func (c *Counts) clear() {
	c.Requests = 0
	c.TotalSuccesses = 0
	c.TotalFailures = 0
	c.ConsecutiveSuccesses = 0
	c.ConsecutiveFailures = 0
}

// 断路器钩子
type CircuitHook struct {
	// cd 断路器map，保存不同服务器的断路器实例。
	cb map[string]*CircuitBreaker

	// setting 断路器配置，请查看结构体CircuitSettings。
	//         注意：Name 配置不起作用，因为钩子内部将会改变它。
	settings CircuitSettings

	//自定义失败。可以根据返回内容，自定义归类为失败请求.
	//@notice 不建议使用
	//@params cErr 原始错误信息
	//@params req  请求结构体。
	//@return 返回的错误替换原来的错误
	handleCErr func(cErr error, req core.Request) error
}

func (ch *CircuitHook) SetHandleCErr(handleFunc func(cErr error, req core.Request) error) {
	ch.handleCErr = handleFunc
}

const CircuitHookKey = "CircuitHook"

func (ch *CircuitHook) getGeneration(req core.Request) uint64 {
	generation, ok := req.HookData(CircuitHookKey)
	if !ok {
		return 0
	}
	return generation.(uint64)
}

func (ch *CircuitHook) setGeneration(req core.Request, generation uint64) (ok bool) {
	ok = req.SetHookData(CircuitHookKey, generation)
	return
}

func (ch *CircuitHook) getCircuitBreaker(req core.Request, ifCreate bool) (*CircuitBreaker, error) {
	cb, ok := ch.cb[req.ServerName()]
	if ok {
		return cb, nil
	}
	setting := ch.settings.Clone().(*CircuitSettings)
	setting.Name = req.ServerName()
	cb = NewCircuitBreaker(*setting)
	ch.cb[req.ServerName()] = cb
	return cb, nil
}

func (ch *CircuitHook) BeforeRequest(req core.Request, client core.Client) error {
	cb, _ := ch.getCircuitBreaker(req, true)
	generation, err := cb.beforeRequest()
	if nil != err {
		return err
	}
	ch.setGeneration(req, generation)
	return nil
}

func (ch *CircuitHook) AfterRequest(cErr error, req core.Request, client core.Client) {
	cb, err := ch.getCircuitBreaker(req, false)
	if nil != err {
		return
	}
	generation := ch.getGeneration(req)

	if nil != ch.handleCErr {
		// 对请求错误重新定义
		cErr = ch.handleCErr(cErr, req)
	}
	cb.afterRequest(generation, cErr == nil)
}

func NewCircuitHook(settings CircuitSettings) *CircuitHook {
	return &CircuitHook{
		settings:   settings,
		cb:         make(map[string]*CircuitBreaker),
		handleCErr: nil,
	}
}

// Name  名字，请务必保障名字的唯一性
//
// MaxRequests Half-Open状态下允许通过的最大请求数
//
// Interval 重置时间间隔（Closed状态下有效）。如果为零，永远不重置。
//
// Timeout  超时时间（Open状态下有效）。
//          超时之后，状态将转变为Half-Open状态。
//          如果为零，默认为60秒
//
// ReadyToTrip   测试是否应该从Closed状态转变为Open状态。
//               true 表示可以转变，否则不可以。
//               如果不配置，则采用默认的。默认失败次数达到5次则进入Open状
//
// OnStateChange 状态变化将调用此方法。
//
type CircuitSettings struct {
	Name          string
	MaxRequests   uint32
	Interval      time.Duration
	Timeout       time.Duration
	ReadyToTrip   func(counts Counts) bool
	OnStateChange func(name string, from State, to State)
}

func (set *CircuitSettings) Clone() interface{} {
	new_obj := *set
	return &new_obj
}

const defaultTimeout = time.Minute

func defaultReadyToTrip(counts Counts) bool {
	return counts.ConsecutiveFailures > 5
}

// 断路器
//   断路器的状态改变都是延迟懒惰的
type CircuitBreaker struct {
	name        string
	maxRequests uint32
	interval    time.Duration
	timeout     time.Duration

	readyToTrip   func(counts Counts) bool
	onStateChange func(name string, from State, to State)

	mutex  sync.Mutex
	state  State
	counts Counts
	expiry time.Time

	// 每一次状态改变加一。
	// 相当于状态改变标识。
	// 避免前一个的数据污染现在的
	generation uint64
}

func NewCircuitBreaker(st CircuitSettings) *CircuitBreaker {
	cb := new(CircuitBreaker)

	cb.name = st.Name
	cb.interval = st.Interval
	cb.onStateChange = st.OnStateChange

	if st.MaxRequests == 0 {
		cb.maxRequests = 1
	} else {
		cb.maxRequests = st.MaxRequests
	}
	if st.Timeout == 0 {
		cb.timeout = defaultTimeout
	} else {
		cb.timeout = st.Timeout
	}
	if st.ReadyToTrip == nil {
		cb.readyToTrip = defaultReadyToTrip
	} else {
		cb.readyToTrip = st.ReadyToTrip
	}
	cb.reset(time.Now())
	return cb
}

func (cb *CircuitBreaker) Name() string {
	return cb.name
}

// 返回断路器状态
func (cb *CircuitBreaker) State() State {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, _ := cb.currentState(now)
	return state
}

func (cb *CircuitBreaker) beforeRequest() (uint64, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, generation := cb.currentState(now)

	if state == StateOpen {
		return generation, ErrOpenState
	} else if state == StateHalfOpen && cb.counts.Requests >= cb.maxRequests {
		return generation, ErrTooManyRequests
	}

	cb.counts.onRequest()
	return generation, nil
}

func (cb *CircuitBreaker) afterRequest(before uint64, success bool) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, generation := cb.currentState(now)
	if generation != before {
		return
	}

	if success {
		cb.onSuccess(state, now)
	} else {
		cb.onFailure(state, now)
	}
}

func (cb *CircuitBreaker) currentState(now time.Time) (State, uint64) {
	switch cb.state {
	case StateClosed:
		if !cb.expiry.IsZero() && cb.expiry.Before(now) {
			cb.reset(now)
		}
	case StateOpen:
		if cb.expiry.Before(now) {
			cb.setState(StateHalfOpen, now)
		}
	}
	return cb.state, cb.generation
}

func (cb *CircuitBreaker) onSuccess(state State, now time.Time) {
	switch state {
	case StateClosed:
		cb.counts.onSuccess()
	case StateHalfOpen:
		cb.counts.onSuccess()
		if cb.counts.ConsecutiveSuccesses >= cb.maxRequests {
			cb.setState(StateClosed, now)
		}
	}
}

func (cb *CircuitBreaker) onFailure(state State, now time.Time) {
	switch state {
	case StateClosed:
		cb.counts.onFailure()
		if cb.readyToTrip(cb.counts) {
			cb.setState(StateOpen, now)
		}
	case StateHalfOpen:
		cb.setState(StateOpen, now)
	}
}

func (cb *CircuitBreaker) setState(state State, now time.Time) {
	if cb.state == state {
		return
	}

	prev := cb.state
	cb.state = state

	cb.reset(now)

	if cb.onStateChange != nil {
		cb.onStateChange(cb.name, prev, state)
	}
}

// 重置. 状态改变之后重置数据
//     1、清理计量器;
//     2、重置过期时间
//     3、状态改变标识加一
func (cb *CircuitBreaker) reset(now time.Time) {
	cb.generation++
	cb.counts.clear()
	var zero time.Time
	switch cb.state {
	case StateClosed:
		if cb.interval == 0 {
			cb.expiry = zero
		} else {
			cb.expiry = now.Add(cb.interval)
		}
	case StateOpen:
		cb.expiry = now.Add(cb.timeout)
	default: // StateHalfOpen
		cb.expiry = zero
	}
}
