## my-gin
记录下gin基础功能，通过它，你可以对gin有一个很好的了解

###一、安装环境
1.首先需要安装Go（需要1.10+版本），然后可以使用下面的Go命令安装Gin
go get -u github.com/gin-gonic/gin

2、运行项目
go run main.go

###二、进入main.go可以看到，写了一些常用的例子,如：

#####1、基础部分

#####2、数据解析

#####3、渲染

#####4、中间件

#####5、会话控制

#####6、参数验证

#####7、token

#####8、grpc
grpc 的安装
go get -u google.golang.org/grpc
protobuf 的安装
下载地址：https://github.com/google/protobuf/releases
选择protoc-xxx-win64.zip下载
安装 golang protobuf
go get -u github.com/golang/protobuf/proto // golang protobuf 库
go get -u github.com/golang/protobuf/protoc-gen-go //protoc --go_out 工具

列子运行 
E:\www\my\go-gin\proto\hello> protoc -I . --go_out=plugins=grpc:E:\www\my --proto_path . ./hello.proto

