package controller

import (
	"github.com/gin-gonic/gin"
	"monitor/app/application/service"
	"monitor/app/interface/http/result"
	"net/http"
)

type ServerController struct {
}

type ServerVo struct {
	Port int    `json:"port"`
	Name string `json:"name"`
}

func (t *ServerController) DetectionPort(context *gin.Context) {
	var param ServerVo
	if err := context.BindJSON(&param); err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}
	service.InspectServer(param.Name, param.Port)
	context.JSON(http.StatusOK, result.OK)
}

func (t *ServerController) GetAllServer(context *gin.Context) {
	servers := service.GetAllServer()
	context.JSON(http.StatusOK, servers)
}

func (t *ServerController) DeleteServer(context *gin.Context) {
	var param ServerVo
	if err := context.BindJSON(&param); err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}
	service.DeleteServerByPort(param.Port)

	context.JSON(http.StatusOK, nil)
}
