package router

import (
	"github.com/gin-gonic/gin"
	"monitor/app/interface/http/controller"
)

func ServerRouter(r *gin.Engine) {
	r.POST("/server", (&controller.ServerController{}).DetectionPort)
	r.GET("/server", (&controller.ServerController{}).GetAllServer)
	r.DELETE("/server", (&controller.ServerController{}).DeleteServer)
}
