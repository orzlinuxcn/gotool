package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/orzlinuxcn/gotool/utils/redis"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func GetCaptchaCode(c *gin.Context) {
	client := redis.GetRedis()
	randNum := rand.Int() % (1000 * 1000)
	randStr := strconv.FormatInt(int64(randNum), 10)
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
