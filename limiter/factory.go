package limiter

import (
	"context"
	"github.com/orzlinuxcn/gotool/consts"
)

type ILimiter interface {
	Wait(ctx context.Context) error
	Allow() bool
	Tokens() int64
}

func NewLimiter(limiterType string, r float64, b int64) ILimiter {
	switch limiterType {
	case consts.LimiterTypeSingle:
		return NewSingleLimiter(r, b)
	default:
		return NewSafeLimiter(r, b, consts.RedisLimiterDefaultName)
	}
}
