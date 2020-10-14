package main

import "errors"


/**
优点
实现了解偶

缺点
违背开闭原则

适合
创建的对象比较少


**/
// 定义一个Cache接口，作为父类
type Cache interface {
	Set(key string, value string)
	Get(key string) string
}

// 实现具体的父类：RedisCache
type RedisCache struct {
	data map[string]string
}

func (r *RedisCache) Set(key string, value string)  {
	r.data[key] = value
}
func (r *RedisCache) Get(key string) string {
	return r.data[key]
}

// 实现具体的父类：MemCache
type MemCache struct {
	data map[string]string
}

func (m *MemCache) Set(key string, value string)  {
	m.data[key] = value
}
func (m *MemCache) Get(key string) string {
	return m.data[key]
}


// 实现简单工厂类
type CacheFactory struct {

}

type CacheType string

const (
	redis CacheType = "redis"
	memcache CacheType = "memchace"
)

func (c *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	switch cacheType {
	case redis:
		return &RedisCache{
			data: map[string]string{},
		}, nil
	case memcache:
		return &MemCache{
			data: map[string]string{},
		}, nil
	}
	return nil,errors.New("error cache type")
}