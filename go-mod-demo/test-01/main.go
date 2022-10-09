package main

/*
1.选择on模式：
a.GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
b.GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
c.GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持至于设置命令（默认mac/unix）：export env -w GO111MODULE=on  即可；
开启了之后，项目所需要的第三方包都会通过go module来下载管理，不会放在$GOPATH/src

2.设置完之后：开启代理，因为默认的代理没有在翻墙的情况下国内是访问不了的，
go env -w GOPROXY=https://goproxy.cn,direct
*/

import (
	"github.com/gin-gonic/gin" // go get -u github.com/gin-gonic/gin
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
