package service

import (
	"encoding/json"
	"monitor/app/domain/entity"
	repository "monitor/app/domain/repository"
	"monitor/app/infrastructure/cache"
)
import file "monitor/app/infrastructure/persistence/file"

var server repository.ServerRepository

func InspectServer(name string, port int) bool {
	newServer := entity.NewServer(name, port)
	server.Insert(newServer)
	return true
}

func LoadServer() error {
	return server.LoadServerToCache()
}

func Inspect() {
	server.Inspect()
}

func GetAllServer() []*entity.Server {
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
	return servers
}

func DeleteServerByPort(port int) {
	server.DeleteServerByPort(port)
}

func init() {
	server = &file.ServerRepository{}
}
