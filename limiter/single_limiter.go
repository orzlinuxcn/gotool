package limiter

import (
	"context"
	"github.com/orzlinuxcn/gotool/consts"
	"sync"
	"time"
)

// SingleLimiter 匀速放入token，至多放入burst个
type SingleLimiter struct {
	rate           float64    // 每秒放入几个token
	timePerToken   int64      // 多长时间放入一个token，单位 ns
	burst          int64      // 桶里可以容纳几个token，应对激增流量
	tokens         int64      // 当前桶里的token量
	mutex          sync.Mutex // lock
	lastUpdateTime int64      // 最近更新token时间，惰性增加bucket数量
}

func NewSingleLimiter(r float64, b int64) *SingleLimiter {
	if r <= 0 {
		r = consts.DefaultLimiterRate
	}
	if b <= 0 {
		b = consts.DefaultLimiterBurst
	}

	return &SingleLimiter{
		rate:           r,
		burst:          b,
		timePerToken:   int64(1 * 1000 * 1000 * 1000 / r),
		mutex:          sync.Mutex{},
		lastUpdateTime: time.Now().UnixNano(),
		tokens:         0,
	}
}

func (l *SingleLimiter) Wait(ctx context.Context) error {
	l.putToken()
	l.mutex.Lock()
	l.tokens = l.tokens - 1
	// 有token，放行
	if l.tokens >= 0 {
		return nil
	}
	duration := -l.tokens * l.timePerToken
	l.mutex.Unlock()
	timer := time.NewTimer(time.Duration(duration))
	select {
	// 等待多长时间放行
	case <-timer.C:
		return nil
	case <-ctx.Done():
		timer.Stop()
		return ctx.Err()
	}
}

func (l *SingleLimiter) Allow() bool {
	l.putToken()
	l.mutex.Lock()
	defer l.mutex.Unlock()
	// 桶里有余量，直接放行
	if l.burst > 0 {
		l.burst = l.burst - 1
		return true
	}
	return false
}

func (l *SingleLimiter) Tokens() int64 {
	return l.tokens
}

func (l *SingleLimiter) putToken() int64 {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	// 惰性增加桶里token
	curTime := time.Now().UnixNano()
	periodTimePass := curTime - l.lastUpdateTime
	addToken := periodTimePass / l.timePerToken
	oldTokenNum := l.tokens
	l.tokens = l.tokens + addToken
	// 只有增加token时才更新pass time
	if addToken > 0 {
		l.lastUpdateTime = curTime
	}
	if l.tokens > l.burst {
		l.tokens = l.burst
	}
	return l.tokens - oldTokenNum
}
