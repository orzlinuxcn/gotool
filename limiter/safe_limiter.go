package limiter

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/orzlinuxcn/gotool/consts"
	redis2 "github.com/orzlinuxcn/gotool/utils/redis"
	"time"
)

// SafeLimiter 匀速放入token，至多放入burst个
type SafeLimiter struct {
	redisClient  *redis.Client
	soleID       string  // 对每个soleID试做一个单独的令牌桶，支持并发
	rate         float64 // 每秒放入几个token
	timePerToken int64   // 多长时间放入一个token，单位 ns
	burst        int64   // 桶里可以容纳几个token，应对激增流量
}

func NewSafeLimiter(r float64, b int64, soleID string) *SafeLimiter {
	if r <= 0 {
		r = consts.DefaultLimiterRate
	}
	if b <= 0 {
		b = consts.DefaultLimiterBurst
	}
	// redis 需存放tokens、lastUpdateTime
	return &SafeLimiter{
		rate:         r,
		burst:        b,
		soleID:       soleID,
		redisClient:  redis2.GetRedis(),
		timePerToken: int64(1 * 1000 * 1000 * 1000 / r),
	}
}

func (l *SafeLimiter) Wait(ctx context.Context) error {
	eval := l.redisClient.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		l.burst, time.Now().UnixNano(), l.timePerToken, "")
	tokens, _ := eval.Int64()
	duration := -tokens * l.timePerToken
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

func (l *SafeLimiter) Allow() bool {
	eval := l.redisClient.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		l.burst, time.Now().UnixNano(), l.timePerToken, "wait")
	tokens, _ := eval.Int64()
	return tokens > 0
}

func (l *SafeLimiter) Tokens() int64 {
	eval := l.redisClient.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		l.burst, time.Now().UnixNano(), l.timePerToken, "wait")
	tokens, _ := eval.Int64()
	return tokens
}
