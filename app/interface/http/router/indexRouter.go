package router

import (
	"github.com/gin-gonic/gin"
	"monitor/app/interface/http/controller"
)

func IndexRouter(r *gin.Engine) {
	r.GET("/", (&controller.IndexController{}).Index)
}
