package main

import "fmt"

/**
优点：
对开闭原则的完美支持
避免使用多重条件转移语句（if else）
提供了管理相关的算法族的办法

缺点：
客户端必须知道所有策略类，并自己决定使用哪个策略
策略模式将造成产生很多策略类
多个类区别仅在他们的行为或算法不同的场景

适合：
一个系统需要动态的在多个算法和行为中选择一种
**/

// 实现一个日志记录器（相当于context）
type LogManager struct {
	Logging
}

func NewLogManager(logging Logging) *LogManager {
	return &LogManager{logging}
}

// 抽象的日志
type Logging interface {
	Info()
	Error()
}

// 实现具体的日志：文件方式日志
type FileLogging struct {

}

func (f *FileLogging) Info() {
	fmt.Println("file logging info")
}

func (f *FileLogging) Error() {
	fmt.Println("file logging error")
}

// 实现具体的日志：数据库方式的日志
type DatabaseLogging struct {

}

func (d *DatabaseLogging) Info() {
	fmt.Println("database logging info")
}
func (d *DatabaseLogging) Error() {
	fmt.Println("database logging error")
}