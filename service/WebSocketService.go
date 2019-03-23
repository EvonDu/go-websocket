package service

import (
	"encoding/json"
	"go-websocket/core"
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
)

//定义客户端
type Client struct {
	Id			string
	Connect		*websocket.Conn
}

//定义结构体
type WebSocketService struct {
	core.BaseWebSocket
	Clients				[]*Client
	Config				*core.Config
}

// 开始监听WebSocket
func (t *WebSocketService) Listen(){
	//设置服务
	http.Handle("/", websocket.Handler(t.Handler))
	http.Handle("/js/", http.FileServer(http.Dir("./")))
	if t.Config.Test {
		http.Handle("/test/", http.FileServer(http.Dir("./")))
	}

	//添加事件
	t.registerEvents(t)
}

// 原始事件注册
func (t *WebSocketService) registerEvents(ws *WebSocketService){
	ws.AddOnMessage(t.onMessage)
}

// 接受客户端信息（按照报文扩展事件）
func (t *WebSocketService) onMessage(ws *websocket.Conn, data string) {
	//DEBUG
	fmt.Print("[*] Receive: " + data + "\n")

	//定义扩展事件方法
	events := make(map[string]func(ws *websocket.Conn, data interface{}, to interface{}))
	events["login"] = t.extendOnLogin
	events["publish"] = t.extendOnPublish

	//解析JSON请求信息
	var request map[string]interface{}
	json.Unmarshal([]byte(data), &request)

	//事件处理
	if request["__event"] != nil && request["data"] != nil {
		event := request["__event"].(string)
		data := request["data"]
		to := request["to"]
		if events[event] != nil {
			events[event](ws, data, to)
		}
	}
}

// 扩展事件：用户登录
func (t *WebSocketService) extendOnLogin(ws *websocket.Conn, data interface{}, to interface{}) {
	// 注册客户端
	id := data.(string)
	client := Client{Id:id, Connect:ws}
	t.Clients = append(t.Clients, &client)
}

// 扩展事件：发布信息
func (t *WebSocketService) extendOnPublish(ws *websocket.Conn, data interface{}, to interface{}) {
	// 整理格式
	result := make(map[string]string)
	result["__event"] = "publish"
	result["data"] = data.(string)
	response,_ := json.Marshal(result)
	// 发布信息
	t.BaseWebSocket.Publish(string(response))
}