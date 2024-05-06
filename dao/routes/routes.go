package routes

import (
	"GoofyLink/controller"
	"GoofyLink/logger"
	"GoofyLink/middleware"

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
	v1 := r.Group("/api/v1")
	// 登录
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	v1.Use(middleware.JWTAuthMiddleware()) // 应用验证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
	}
	//r.POST("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
	//	//如果登录的是 用户，则返回用户信息
	//	isLogin := true
	//	if isLogin {
	//		c.JSON(200, gin.H{
	//			"message": "pong",
	//		})
	//	} else {
	//		// 如果登录的不是用户，则返回 401
	//		c.JSON(200, gin.H{
	//			"message": "请登录",
	//		})
	//	}
	//})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	return r
}
