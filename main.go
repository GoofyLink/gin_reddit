package main

import (
	"GoofyLink/controller"
	"GoofyLink/dao/mysql"
	"GoofyLink/dao/redis"
	"GoofyLink/dao/routes"
	"GoofyLink/logger"
	"GoofyLink/pkg/snowflake"
	"GoofyLink/settings"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// go web 开发项目 模板
func main() {
	// 1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("settings.Init() is failed%v\n", err)
		return
	}
	// 2.初始化日志
	//fmt.Println("settings.Conf.Name", settings.Conf.Name)
	if err := logger.Init(settings.Conf.LogConfig, viper.GetString("app.mode")); err != nil {
		fmt.Printf("logger.Init() is failed%v\n", err)
		return
	}
	defer zap.L().Sync()
	fmt.Println()
	zap.L().Debug(" logger.Init() is success..")
	// 3.初始化mysql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql.Init() is failed%v\n", err)
		return
	}
	defer mysql.Close()
	// 4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis.Init() is failed %v\n", err)
		return
	}
	defer redis.Close()
	// 初始化雪花算法

	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("snowflake.Init is failed", err)
		return
	}

	// 初始化gin框架内置的校验器来翻译错误信息
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("controller.InitTrans is failed", err)
		return
	}
	// 5.注册路由
	r := routes.SingUp(viper.GetString("app.mode"))
	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	if err != nil {
		fmt.Printf("r.Run is failed %v\n", err)
	}
	//r := routes.SingUp()
	//serve := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
	//	Handler: r,
	//}

	// 6.启动服务
	//go func() {
	//	if err := serve.ListenAndServe(); err != nil {
	//		zap.L().Info("serve.ListenAndServe() is failed")
	//		return
	//	}
	//}()
	//// 7.优雅关机
	//quit := make(chan os.Signal, 1)
	//zap.L().Info("Shutdown server")
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := serve.Shutdown(ctx); err != nil {
	//	zap.L().Info("serve.Shutdown(ctx) is failed")
	//}
	//zap.L().Info("Server exiting")
}
