package cache

import (
	"fmt"
	"github.com/coocood/freecache"
	"sync"
)

var cache *freecache.Cache

var mutex sync.Mutex

func init() {
	cache = freecache.NewCache(100 * 1024 * 1024)
}

func AddCache(key string, value string) {
	mutex.Lock()
	err := cache.Set([]byte(key), []byte(value), 0)
	if err != nil {
		fmt.Println("插入缓存失败")
	}
	mutex.Unlock()
}

func GetCache(key string) string {
	mutex.Lock()
	value, err := cache.Get([]byte(key))
	if err != nil {
		fmt.Println("插入缓存失败")
	}
	mutex.Unlock()

	return string(value)
}

func InitCache(key string) string {
	mutex.Lock()
	value, err := cache.Get([]byte(key))
	if err != nil {
		fmt.Println("插入缓存失败")
	}
	mutex.Unlock()

	return string(value)
}
