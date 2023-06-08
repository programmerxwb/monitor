package main

import (
	"github.com/gin-gonic/gin"
	"monitor/app/application/service"
	"monitor/app/interface/http/router"
)

func main() {
	// 服务加载到内存中
	service.LoadServer()
	// 定时扫描服务的状态
	service.Inspect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	router.InitRouter(r)
	r.Run()

}
