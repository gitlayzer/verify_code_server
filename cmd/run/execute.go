package run

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gitlayzer/callback_platform/api"
	"github.com/gitlayzer/callback_platform/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	// 初始化配置
	config.Init()
}

func Execute() {
	// 设置gin的运行模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化gin
	r := gin.Default()
	// 注册路由
	api.RegisterRouter(r)
	// 初始化http服务

	// 初始化http服务
	srv := &http.Server{
		Addr:    config.Listen,
		Handler: r,
	}

	// 启动http服务
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
