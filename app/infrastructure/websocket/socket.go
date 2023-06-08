package socket

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

var conns = list.New()

func AddConn(conn *websocket.Conn) {
	conns.PushBack(conn)
}

func RemoveConn(conn *websocket.Conn) {
	for e := conns.Front(); e != nil; e = e.Next() {
		if e.Value == conn {
			conns.Remove(e)
		}
	}
}

func SendAll(dataType string, data any) {
	body := &SocketBody{
		Type: dataType,
		Data: data,
	}

	message, _ := json.Marshal(body)
	for e := conns.Front(); e != nil; e = e.Next() {
		if value, ok := e.Value.(*websocket.Conn); ok {
			value.WriteMessage(websocket.TextMessage, message)
		} else {
			fmt.Println("无法将元素转换为整数类型")
		}
	}
}

type SocketBody struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
