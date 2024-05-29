package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orzlinuxcn/gotool/handler"
)

func main() {
	r := gin.Default()
	r.GET("/get_captcha_code", handler.GetCaptchaCode)
	r.GET("/check_captcha_code", handler.CheckCaptchaCode)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
