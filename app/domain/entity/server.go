package entity

import (
	"fmt"
	"net"
)

type Server struct {
	Port   int    `json:"port"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func NewServer(name string, port int) *Server {
	return &Server{
		Name: name,
		Port: port,
	}
}

/*
*
根据端口检测
*/
func (res *Server) DetectionByPort() {
	address := fmt.Sprintf("127.0.0.1:%d", res.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		res.Status = "close"
		return
	}
	conn.Close()
	res.Status = "open"
}
