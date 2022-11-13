package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

func main() {
	etcdReg := etcd.NewRegistry()

	ginRouter := gin.Default()
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("GET", "/hello", func(context *gin.Context) {
			context.JSON(200, "Hello World")
		})
	}
	server := web.NewService(
		web.Name("webserver"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(etcdReg),
		web.RegisterTTL(3*time.Second),
		web.RegisterInterval(2*time.Second),
		web.Version("1.0"),
	)

	server.Run()

}
