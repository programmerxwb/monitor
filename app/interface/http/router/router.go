package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	IndexRouter(r)
	ServerRouter(r)
	WebsocketRouter(r)
}
