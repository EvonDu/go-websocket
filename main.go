package main

import (
	"fmt"
	"net/http"
	"go-websocket/service"
	"go-websocket/core"
	"strconv"
)


var port string

func main() {
	// 命令行参数
	config := core.Config{}
	config.Load()

	// 创建服务对象
	ws := service.WebSocketService{Config:&config}
	api := service.ApiService{Config:&config, WebSocketService:&ws}
	con := service.ConsoleService{Config:&config, WebSocketService:&ws}

	// 服务启动
	con.Run()
	ws.Listen()
	api.Listen()

	// 开始监听
	e := http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
	fmt.Println(e)
}