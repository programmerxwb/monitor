package repository

import "monitor/app/domain/entity"

type ServerRepository interface {
	// Insert 插入新的服务
	Insert(server *entity.Server) error

	// LoadServerToCache 加载服务到缓存
	LoadServerToCache() error

	// 检查服务是否启动并推送给前端
	Inspect()

	// 根据端口删除服务
	DeleteServerByPort(port int) error
}
