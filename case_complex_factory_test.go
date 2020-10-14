package main

import (
	"fmt"
	"testing"
)

func TestComplexFactory_Create(t *testing.T) {
	redisFactory := &RedisFactory{}
	redisCache := redisFactory.Create()
	redisCache.Set("redis", "hello redis")
	fmt.Println(redisCache.Get("redis"))

	memFactory := &MemFactory{}
	memCache := memFactory.Create()
	memCache.Set("memcache", "hello memcache")
	fmt.Println(memCache.Get("memcache"))
}