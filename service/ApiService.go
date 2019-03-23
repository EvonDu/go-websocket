package service

import (
	"fmt"
	"net/http"
	"strconv"
	"go-websocket/core"
	"io/ioutil"
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
	http.HandleFunc("/count", t.Count)
	http.HandleFunc("/clients", t.Clients)
	http.HandleFunc("/publish", t.Publish)
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

/**
 * 连接列表
 * @OA\Get(
 *      path="/clients",
 *      tags={"WebSocket"},
 *      summary="连接列表",
 *      description="获取所有注册过的客户端列表",
 *      @OA\Response(response="default", description="返回结果"),
 * )
 */
func (t *ApiService) Clients(w http.ResponseWriter, r *http.Request) {
	//获取列表
	var list []map[string]interface{}
	for i:=0;i<len(t.WebSocketService.Clients);i++ {
		item := make(map[string]interface{})
		item["Id"] = t.WebSocketService.Clients[i].Id
		item["Time"] = t.WebSocketService.Clients[i].Time.Format("2006-01-02 15:04:05")
		list = append(list, item)
	}
	//输出
	fmt.Fprintln(w, "Client id list : ")
	for i:=0;i<len(list);i++ {
		fmt.Fprintln(w, list[i])
	}
}

/**
 * 广播信息
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
	//获取参数(GET)
	/*r.ParseForm()
	var message string
	if len(r.Form["message"]) > 0 {
		message = r.Form["message"][0]
	} else {
		message = ""
	}*/

	//获取参数(BODY)
	body, _ := ioutil.ReadAll(r.Body)
	message := string(body)

	//广播消息
	t.WebSocketService.Publish(message)

	//接口返回
	fmt.Fprintln(w, "Success publish message : "+message)
}