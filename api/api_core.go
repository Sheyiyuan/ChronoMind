package api

import (
	"context"
	"github.com/Sheyiyuan/ChronoMind/logos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"

	// 导入Hertz框架
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// NewAPICore 创建并运行一个新的API核心实例
func NewAPICore(port int, isGlobal bool) {
	// 创建一个新的Hertz服务器实例
	address := "127.0.0.1:" + strconv.Itoa(port)
	if isGlobal {
		address = "0.0.0.0:" + strconv.Itoa(port)
	}
	h := server.Default(
		server.WithHostPorts(address),
	)

	// 设置日志级别为Fatal
	hlog.SetLevel(hlog.LevelFatal)

	// 使用日志中间件记录请求信息
	h.Use(LoggingMiddleware())

	// 定义一个路由处理函数
	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "Ciallo～(∠・ω< )⌒☆, ChronoMind!",
		})
	})

	h.Spin()
}

// LoggingMiddleware 是一个中间件函数，用于记录请求信息
func LoggingMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 获取客户端 IP
		clientIP := ctx.ClientIP()
		// 获取请求完整路径
		fullPath := ctx.Request.URI().PathOriginal()
		// 继续处理请求
		ctx.Next(c)
		// 获取请求处理状态码
		statusCode := ctx.Response.StatusCode()
		// 在请求处理后记录结束信息
		logos.Info("收到网络请求 | IP: %s | Path: %s | Status: %d ", clientIP, fullPath, statusCode)
	}
}
