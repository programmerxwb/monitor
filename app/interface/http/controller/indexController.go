package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (t *IndexController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"messgae": "asd",
	})
}
