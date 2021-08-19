package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "my-gin/proto/hello"
	"net/http"
)

// 数据解析常用
func GrpcRouter(e *gin.Engine) {

	// 使用路由组
	r := e.Group("/grpc")

	// http://127.0.0.1:8080/grpc/client?name=world
	r.GET("/client", client)
}


const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8081"
)

func client(c *gin.Context) {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()

	// 初始化客户端
	client := pb.NewHelloClient(conn)

	name := c.DefaultQuery("name", "yzx-fjl.cn")

	// 调用方法
	req := &pb.HelloRequest{Name: name}
	res, err := client.SayHello(c, req)

	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "message": res})
}