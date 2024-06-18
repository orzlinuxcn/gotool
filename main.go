package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orzlinuxcn/gotool/handler"
	"os"
)

func main() {
	r := gin.Default()
	logFile, _ := os.OpenFile("./app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	r.Use(gin.LoggerWithWriter(logFile))

	r.Any("/captcha/get", handler.QPSLimiter(10, 10, "generate_captcha"), handler.GetCaptchaCode)
	r.Any("/captcha/check", handler.CheckCaptchaCode)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
