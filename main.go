package main

import (
	"github.com/gin-gonic/gin"
	"monitor/app/application/service"
	"monitor/app/interface/http/router"
)

func main() {
	service.LoadServer()
	service.Inspect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	router.InitRouter(r)
	r.Run()

}
