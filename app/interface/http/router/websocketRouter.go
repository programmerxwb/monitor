package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	socket "monitor/app/infrastructure/websocket"
	"net/http"
)

func WebsocketRouter(r *gin.Engine) {
	r.GET("/websocket", conn)
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func conn(context *gin.Context) {

	conn, err := upGrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println(err)
		http.NotFound(context.Writer, context.Request)
		return
	}
	socket.AddConn(conn)
	for {
		messageType, p, err := conn.ReadMessage()
		if messageType == -1 {
			socket.RemoveConn(conn)
			conn.Close()
			break
		}
		if err != nil {

		}
		fmt.Println(string(p))
	}
}
