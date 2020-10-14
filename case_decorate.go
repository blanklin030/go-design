package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("tracing start 记录请求的url和方法: %s, %s", r.URL, r.Method)
		next.ServeHTTP(w, r)
		log.Printf("tracing end")
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("logging start记录请求的网络地址: %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("logging end")
	})
}

func timeRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("timeRecording start")
		startTime := time.Now()
		next.ServeHTTP(w, r)
		endTime := time.Since(startTime)
		log.Printf("timeRecording end 记录方法的执行时间: %s", endTime)
	})
}


// 1. 实现一个http server
// 2. 实现一个http handler：hello
// 3. 实现中间件的功能： 1）记录请求的url和方法 2）记录请求的网络地址 3）记录方法的执行时间
func hello(w http.ResponseWriter, r *http.Request)  {
	//log.Printf("记录请求的url和方法: %s, %s", r.URL, r.Method)
	//log.Printf("记录请求的网络地址: %s", r.RemoteAddr)
	//startTime := time.Now()
	fmt.Fprintf(w, "hello")
	//endTime := time.Since(startTime)
	//log.Printf("记录方法的执行时间: %s", endTime)
}
func main()  {
	http.Handle("/", tracing(logging(timeRecording(http.HandlerFunc(hello)))))
	http.ListenAndServe(":8980", nil)
}