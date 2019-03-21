package main

import (
	"fmt"
	"net/http"
	"go-websocket/service"
)

func main() {
	//创建服务对象
	ws := service.WebSocketService{}
	api := service.ApiService{WebSocketService:&ws}
	con := service.ConsoleService{WebSocketService:&ws}

	//服务启动
	con.Run()
	ws.Listen()
	api.Listen()

	//开始监听
	e := http.ListenAndServe(":8080", nil)
	fmt.Println(e)
}