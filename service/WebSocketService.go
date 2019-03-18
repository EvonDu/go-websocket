package service

import (
	"golang.org/x/net/websocket"
)

//定义结构体
type WebSocketService struct {
	Connects []*websocket.Conn
	EventOnConnect []func(ws *websocket.Conn)
	EventOnMessage []func(ws *websocket.Conn, data string)
	EventOnPublish []func(data string)
	EventOnClose []func(ws *websocket.Conn)
}

//WebSocket处理函数
func (t *WebSocketService) Handler(ws *websocket.Conn) {
	//执行事件客户端连接
	t.onConnect(ws)
	//轮询
	request := make([]byte, 512)
	for {
		//读取客户端内容
		read, _ := ws.Read(request)
		//检查WebSocket是否断开
		if read == 0 {
			t.onClose(ws)
			break
		}
		//接受客户端信息
		data := string(request[:read])
		t.onMessage(ws, data)
	}
}

//连接客户端
func (t *WebSocketService) onConnect(ws *websocket.Conn) {
	//执行事件客户端连接
	t.Connects = append(t.Connects, ws)
	for i:=0;i<len(t.EventOnConnect);i++ {
		t.EventOnConnect[i](ws)
	}
}

//断开客户端
func (t *WebSocketService) onClose(ws *websocket.Conn) {
	//移除连接
	for i:=0;i<len(t.Connects);i++{
		if t.Connects[i] == ws {
			t.Connects = append(t.Connects[:i], t.Connects[i+1:]...)
		}
	}
	//触发事件
	for i:=0;i<len(t.EventOnClose);i++ {
		t.EventOnClose[i](ws)
	}
}

//接受客户端信息
func (t *WebSocketService) onMessage(ws *websocket.Conn, data string) {
	//触发事件
	for i:=0;i<len(t.EventOnMessage);i++ {
		t.EventOnMessage[i](ws, data)
	}
}

//添加客户端连接事件
func (t *WebSocketService) AddOnConnect(event func(ws *websocket.Conn)){
	t.EventOnConnect = append(t.EventOnConnect,event)
}

//添加客户端断开事件
func (t *WebSocketService) AddOnClose(event func(ws *websocket.Conn)){
	t.EventOnClose = append(t.EventOnClose,event)
}

//添加接收信息事件
func (t *WebSocketService) AddOnMessage(event func(ws *websocket.Conn, data string)){
	t.EventOnMessage = append(t.EventOnMessage,event)
}

//添加发送信息事件
func (t *WebSocketService) AddOnPublish(event func(data string)){
	t.EventOnPublish = append(t.EventOnPublish,event)
}

//广播信息
func (t *WebSocketService) Publish(message string){
	//广播发送
	for i:=0;i<len(t.Connects);i++ {
		t.Connects[i].Write([]byte(message))
	}
	//触发事件
	for i:=0;i<len(t.EventOnPublish);i++ {
		t.EventOnPublish[i](message)
	}
}