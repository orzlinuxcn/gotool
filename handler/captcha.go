package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orzlinuxcn/gotool/utils/redis"
	"math/rand"
	"net/http"
	"time"
)

func GetCaptchaCode(c *gin.Context) {
	client := redis.GetRedis()
	randNum := rand.Int() % (1000 * 1000)
	randStr := fmt.Sprintf("%06d", randNum)
	cmd := client.Set(randStr, randStr, time.Minute*10)
	if cmd.Err() != nil {
		c.String(http.StatusInternalServerError, "err")
		return
	}
	c.String(http.StatusOK, randStr)
	return
}

func CheckCaptchaCode(c *gin.Context) {
	client := redis.GetRedis()
	code := c.Query("code")
	val := client.Get(code).Val()
	if len(val) > 0 {
		c.String(http.StatusOK, "success")
		return
	}
	c.String(http.StatusOK, "fail")
	return
}
