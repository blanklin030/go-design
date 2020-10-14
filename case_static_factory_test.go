package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	redisCache, err := cacheFactory.Create(redis)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	redisCache.Set("redis", "hello redis")
	fmt.Println(redisCache.Get("redis"))

	memCache, err := cacheFactory.Create(memcache)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	memCache.Set("memcache", "hello memcache")
	fmt.Println(memCache.Get("memcache"))
}