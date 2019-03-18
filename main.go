package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"go-websocket/service"
)
func main() {
	//创建服务对象
	ws := service.WebSocketService{}
	api := service.ApiService{WebSocketService:&ws}

	//创建协程
	Console(&ws)

	//添加事件
	Events(&ws)

	//创建WebSocket服务
	http.Handle("/", websocket.Handler(ws.Handler))
	//创建文件服务器
	http.Handle("/test/", http.FileServer(http.Dir("./")))

	//添加Rest接口
	http.HandleFunc("/publish", api.Publish)

	//开始监听
	e := http.ListenAndServe(":8080", nil)
	fmt.Println(e)
}

//客户端事件
func Events(ws *service.WebSocketService){
	ws.AddOnConnect(func(ws *websocket.Conn) {
		fmt.Print("client connect.\n")
	})
}

//命令行协程
func Console(ws *service.WebSocketService){
	con := service.Console{WebSocketService:ws}
	var ch chan string
	go con.Run(ch)
}