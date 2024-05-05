package routes

import (
	"GoofyLink/controller"
	"GoofyLink/logger"

	"github.com/gin-gonic/gin"
)

// 注册路由
func SingUp(mode string) *gin.Engine {
	// 如果是开发模式，则打印日志
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 使用中间件 注册
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 登录
	r.POST("/login", controller.LoginHandler)
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	return r
}
