package micro

import (
	"errors"
	"log"
	"sync"
	"time"
)

type Breakpoint struct {
	BreakCount   int
	LastFailedAt time.Time
}

var (
	ErrCircuitBreakerMessage = errors.New("unavailable service call, holding circuit breaker")
)

type Circuit struct {
	sync.Mutex
	breaks map[string]*Breakpoint
	//熔断触发次数
	touchOffTimes int
	//熔断保持时间
	blockHoldingDuration time.Duration
}

var once sync.Once
var circuit *Circuit

func init() {
	if circuit == nil {
		circuit = &Circuit{
			Mutex:                sync.Mutex{},
			breaks:               map[string]*Breakpoint{},
			touchOffTimes:        5,
			blockHoldingDuration: time.Minute * 15,
		}
	}
}

// InitializeCircuit 初始化熔断器
// touchOffTimes 熔断触发次数
// blockHoldingDuration 熔断保持时间
func InitializeCircuit(touchOffTimes int, blockHoldingDuration time.Duration) {
	once.Do(func() {
		circuit = &Circuit{
			Mutex:                sync.Mutex{},
			breaks:               map[string]*Breakpoint{},
			touchOffTimes:        touchOffTimes,
			blockHoldingDuration: blockHoldingDuration,
		}
		log.Println(circuit.touchOffTimes)
	})
}

func (m *Circuit) Failed(url string) {
	if m == nil {
		return
	}
	m.Lock()
	defer m.Unlock()

	bp, ok := m.breaks[url]

	if !ok {
		bp = &Breakpoint{}
		m.breaks[url] = bp
	}

	bp.LastFailedAt = time.Now()
	bp.BreakCount += 1
}

// IsHolding
// 检查是否断路保持
func (m *Circuit) IsHolding(url string) bool {
	if m == nil {
		return false
	}
	m.Lock()
	defer m.Unlock()
	bp, ok := m.breaks[url]
	if !ok {
		return false
	}
	return bp.BreakCount >= m.touchOffTimes && time.Now().Sub(bp.LastFailedAt) <= m.blockHoldingDuration
}
