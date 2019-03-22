package service

import (
	"go-websocket/core"
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
)

//定义结构体
type WebSocketService struct {
	core.BaseWebSocket
	Config				*core.Config
}

// 开始监听WebSocket
func (t *WebSocketService) Listen(){
	//设置服务
	http.Handle("/", websocket.Handler(t.Handler))
	if t.Config.Test {
		http.Handle("/test/", http.FileServer(http.Dir("./")))
	}

	//添加事件
	t.registerEvents(t)
}

// 客户端事件
func (t *WebSocketService) registerEvents(ws *WebSocketService){
	ws.AddOnConnect(t.onConnect)
	ws.AddOnClose(t.onClose)
	ws.AddOnMessage(t.onMessage)
	ws.AddOnPublish(t.onPublish)
}

// 客户端连接
func (t *WebSocketService) onConnect(ws *websocket.Conn) {
	fmt.Print("[*] Client connect.\n")
}

// 客户端离线
func (t *WebSocketService) onClose(ws *websocket.Conn) {
	fmt.Print("[*] Client close.\n")
}

// 接受客户端信息
func (t *WebSocketService) onMessage(ws *websocket.Conn, data string) {
	fmt.Print("[*] Receive: " + data + "\n")
}

// 服务端发送信息
func (t *WebSocketService) onPublish(data string) {
	fmt.Print("[*] Publish: " + data + "\n")
}