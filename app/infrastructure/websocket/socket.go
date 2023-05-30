package socket

import (
	"container/list"
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

func SendAll(message []byte) {
	for e := conns.Front(); e != nil; e = e.Next() {
		if value, ok := e.Value.(*websocket.Conn); ok {
			value.WriteMessage(websocket.TextMessage, message)
		} else {
			fmt.Println("无法将元素转换为整数类型")
		}
	}
}
