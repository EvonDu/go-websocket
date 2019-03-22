package service

import (
	"fmt"
	"net/http"
	"strconv"
	"go-websocket/core"
)

/**
 * @OA\Tag(name="WebSocket",description="WebSocket")
 */
type ApiService struct {
	Config				*core.Config
	WebSocketService	*WebSocketService
}

// 开始监听API
func (t *ApiService) Listen(){
	//注册文档
	if t.Config.Swagger {
		http.Handle("/swagger/", http.FileServer(http.Dir("./")))
	}

	//注册接口
	http.HandleFunc("/publish", t.Publish)
	http.HandleFunc("/count", t.Count)
}

/**
 * 广播信息
 * @OA\Get(
 *      path="/publish",
 *      tags={"WebSocket"},
 *      summary="广播信息",
 *      description="向所有WebSocket客户端发送信息",
 *      @OA\Parameter(name="message", required=true, in="query",description="信息内容", @OA\Schema(type="string", default="测试信息")),
 *      @OA\Response(response="default", description="返回结果"),
 * )
 */
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
	fmt.Fprintln(w, "Success publish message : "+message)
}

/**
 * 连接数量
 * @OA\Get(
 *      path="/count",
 *      tags={"WebSocket"},
 *      summary="连接数量",
 *      description="获取当前WebSocket客户端的连接数量",
 *      @OA\Response(response="default", description="返回结果"),
 * )
 */
func (t *ApiService) Count(w http.ResponseWriter, r *http.Request) {
	//接口返回
	fmt.Fprintln(w, "Client connect count : "+strconv.Itoa(len(t.WebSocketService.Connects)))
}