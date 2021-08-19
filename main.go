package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"my-gin/routers"
	"my-gin/toolkit/config"
	"os"
	"time"
)

func main() {

	// 创建文件
	createFile("./logs")
	createFile("./upload")

	// 写日志
	witerLog()

	r := gin.Default()

	// 基础常用
	routers.BaseRouter(r)

	// 数据解析常用
	routers.DataRouter(r)

	// 渲染常用
	routers.RenderRouter(r)

	// 中间件常用
	routers.MiddlewareRouter(r)

	//// 会话控制常用 （测试其它时注释）
	//routers.ConversationRouter(r)

	// 参数验证常用
	routers.VerificationRouter(r)

	// token
	routers.TokenRouter(r)

	//grpc  测试时先开启grpc服务 运行 go run grpc/server.go
	routers.GrpcRouter(r)

	conf := config.Configure()
	// 默认8080端口
	if err := r.Run(conf.GetString("http.listen")); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

// 创建文件夹
func createFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

// 写日志
func witerLog() {
	gin.DisableConsoleColor()
	// Logging to a file.
	// 创建日志路径
	f, _ := os.Create("./logs/" + time.Now().Format("2006-01-02") + ".log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
