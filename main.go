package main

import (
	"log"
	"net/http"

	"github.com/arl/statsviz"
)

func main() {
	mux := http.NewServeMux()

	// 注册监控路由
	if err := statsviz.Register(mux); err != nil {
		log.Fatalf("statsviz.Register err: %s", err)
	}

	// 启动 http 服务器
	listenAddr := "localhost:8080"
	log.Printf("HTTP Server 监控页面: http://%s/debug/statsviz\n", listenAddr)

	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		log.Printf("http server exit: %s\n", err)
	}
}
