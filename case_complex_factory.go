package main


/**
优点
实现了解偶

缺点
违背开闭原则

适合
创建的对象比较少


**/
// 定义一个Cache接口，作为父类
type Cache1 interface {
	Set(key string, value string)
	Get(key string) string
}

// 实现具体的父类：RedisCache
type RedisCache1 struct {
	data map[string]string
}

func (r *RedisCache1) Set(key string, value string)  {
	r.data[key] = value
}
func (r *RedisCache1) Get(key string) string {
	return r.data[key]
}

// 实现具体的父类：MemCache
type MemCache1 struct {
	data map[string]string
}

func (m *MemCache1) Set(key string, value string)  {
	m.data[key] = value
}
func (m *MemCache1) Get(key string) string {
	return m.data[key]
}
// 实现抽象的工厂类
type Factory1 interface {
	Create() Cache1
}

// 实现具体工厂类
type RedisFactory struct {

}
func (r *RedisFactory) Create() Cache1{
	return &RedisCache1{data: map[string]string{}}
}
// 实现具体工厂类
type MemFactory struct {

}
func (m *MemFactory) Create() Cache1{
	return &MemCache1{data: map[string]string{}}
}