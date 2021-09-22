## my-gin
记录下gin基础功能，通过它，你可以对gin有一个很好的了解

## 一、安装环境
1.首先需要安装Go（需要1.10+版本），然后可以使用下面的Go命令安装Gin
go get -u github.com/gin-gonic/gin

2、运行项目
go run main.go

## 二、进入main.go可以看到，写了一些常用的例子,如：

- [路由](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E8%B7%AF%E7%94%B1.html)

- [基础部分](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E5%9F%BA%E7%A1%80%E9%83%A8%E5%88%86.html)

- [配置](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E9%85%8D%E7%BD%AE.html)

- [模型绑定和验证](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E6%A8%A1%E5%9E%8B%E7%BB%91%E5%AE%9A%E5%92%8C%E9%AA%8C%E8%AF%81.html)

- [参数验证](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E5%8F%82%E6%95%B0%E9%AA%8C%E8%AF%81.html)

- [渲染](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E6%B8%B2%E6%9F%93.html)

- [中间件](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E4%B8%AD%E9%97%B4%E4%BB%B6.html)

- [会话控制](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/%E4%BC%9A%E8%AF%9D%E6%8E%A7%E5%88%B6.html)

- [token](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/token.html)

- [grpc](http://yzx-fjl.cn:88/gin%E6%A1%86%E6%9E%B6/grpc.html)

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

