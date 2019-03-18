package service

import (
	"fmt"
	"net/http"
)

/**
 * @OA\Tag(name="WebSocket",description="WebSocket")
 */
type ApiService struct {
	WebSocketService *WebSocketService
}

/**
 * 创建订单
 * @OA\Post(
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
	fmt.Fprintln(w, "success publish message : "+message)
}