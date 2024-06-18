package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/orzlinuxcn/gotool/limiter"
	"net/http"
)

func QPSLimiter(rate float64, burst int64, soleID string) gin.HandlerFunc {
	safeLimiter := limiter.NewSafeLimiter(rate, burst, soleID)
	return func(c *gin.Context) {
		if safeLimiter.Allow() {
			safeLimiter.Wait(c)
			c.Next()
		} else {
			c.String(http.StatusInternalServerError, "限流")
			c.Abort()
		}
	}
}
