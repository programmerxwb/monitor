package file

import (
	"encoding/json"
	entity "monitor/app/domain/entity"
	"monitor/app/infrastructure/cache"
	socket "monitor/app/infrastructure/websocket"
	"os"
	"time"
)

type ServerRepository struct {
}

const dataName = "a.text"

func (m *ServerRepository) Inspect() {
	go doInspect()
}

func doInspect() {
	for {
		value := cache.GetCache("servers")
		var servers []*entity.Server
		if len(value) != 0 {
			err := json.Unmarshal([]byte(value), &servers)
			if err == nil {
				for _, server := range servers {
					server.DetectionByPort()
				}
			}
		}
		marshal, _ := json.Marshal(servers)
		socket.SendAll(marshal)
		time.Sleep(time.Second * 10)
	}
}

func (m *ServerRepository) LoadServerToCache() error {
	_, err := os.Stat(dataName)
	if os.IsNotExist(err) {
		os.Create(dataName)
	}
	content, err := os.ReadFile(dataName)
	if err != nil {
		return err
	}

	cache.AddCache("servers", string(content))
	return nil
}

func (m *ServerRepository) Insert(server *entity.Server) error {
	content, err := os.ReadFile(dataName)
	if err != nil {
		return err
	}
	var servers []*entity.Server
	if len(content) == 0 {
		servers = append(servers, server)
		if err != nil {
			return err
		}
	} else {
		err = json.Unmarshal(content, &servers)
		servers = append(servers, server)
	}
	marshal, err := json.Marshal(servers)
	os.WriteFile(dataName, marshal, os.ModePerm)
	m.LoadServerToCache()
	return nil
}

func (m *ServerRepository) DeleteServerByPort(port int) error {
	content, err := os.ReadFile(dataName)
	if err != nil {
		return err
	}
	var servers []*entity.Server
	if len(content) == 0 {
		return nil
	}
	err = json.Unmarshal(content, &servers)
	var deleteIndex int = 0
	for i, value := range servers {
		if value.Port == port {
			deleteIndex = i
			break
		}
	}
	servers = append(servers[:deleteIndex], servers[deleteIndex+1:]...)
	marshal, err := json.Marshal(servers)
	os.WriteFile(dataName, marshal, os.ModePerm)
	m.LoadServerToCache()
	return nil
}
