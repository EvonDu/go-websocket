package service

import (
	"fmt"
	"net/http"
)

//定义结构体
type ApiService struct {
	WebSocketService *WebSocketService
}

//WebSocket处理函数
func (t *ApiService) Publish(w http.ResponseWriter, r *http.Request) {
	//获取参数
	r.ParseForm()
	var message string
	if len(r.Form["message"]) > 0 {
		message = r.Form["message"][0]
	} else {
		message = ""
	}

	//广播消息
	t.WebSocketService.Publish(message)

	//接口返回
	fmt.Fprintln(w, "success publish message : "+message)
}