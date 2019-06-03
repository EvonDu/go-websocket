package main

import (
	"fmt"
	"net/http"
	"go-websocket/service"
	"go-websocket/core"
	"strconv"
)

//主函数
func main() {
	// 命令行参数
	config := core.Config{}
	config.Load()

	// 注册Web
	registerWeb(&config)

	// 注册Service
	registerService(&config)

	// 开始监听
	e := http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
	fmt.Println(e)
}

//注册Web
func registerWeb(config *core.Config){
	//注册JS
	http.Handle("/js/", http.FileServer(http.Dir("./web/")))
	//注册文档
	if config.Swagger {
		http.Handle("/swagger/", http.FileServer(http.Dir("./web")))
	}
	//注册测试页
	if config.Test {
		http.Handle("/test/", http.FileServer(http.Dir("./web")))
	}
}

//注册Service
func registerService(config *core.Config){
	// 创建服务对象
	ws := service.WebSocketService{Config:config}
	api := service.ApiService{Config:config, WebSocketService:&ws}
	con := service.ConsoleService{Config:config, WebSocketService:&ws}

	// 服务启动
	con.Run()
	ws.Listen()
	api.Listen()
}