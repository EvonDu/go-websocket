package service

import (
	"net/http"
	"io/ioutil"
	"go-websocket/core"
	"encoding/json"
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
	http.HandleFunc("/publish", t.Publish)
	http.HandleFunc("/clients", t.Clients)
	http.HandleFunc("/events", t.Events)
}

// 按接口结构输出结果
func (t *ApiService) apiResponse(w http.ResponseWriter, r *http.Request, code int, message string, data interface{}){
	//设置HTTP头
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Content-type","application/json;charset='utf-8'")

	//设置返回格式
	result := make(map[string]interface{})
	result["code"] = code
	result["message"] = message
	result["data"] = data

	//设置返回内容
	response,_ := json.Marshal(result)
	w.Write(response)
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
	t.apiResponse(w, r, 0, "OK", len(t.WebSocketService.Connects))
}

/**
 * 发布信息
 * @OA\Post(
 *      path="/publish",
 *      tags={"WebSocket"},
 *      summary="发布信息",
 *      description="向所有WebSocket客户端发送信息",
 *      @OA\Parameter(name="message", required=true, in="query",description="信息内容", @OA\Schema(type="string", default="测试信息")),
 *      @OA\Response(response="default", description="返回结果"),
 * )
 */
func (t *ApiService) Publish(w http.ResponseWriter, r *http.Request) {
	//获取参数(BODY)
	body, _ := ioutil.ReadAll(r.Body)
	message := string(body)
	//广播消息
	t.WebSocketService.Publish(message)
	//接口返回
	t.apiResponse(w, r, 0, "OK", message)
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
	//接口返回
	t.apiResponse(w, r, 0, "OK", list)
}

/**
 * 发布事件
 * @OA\Post(
 *      path="/events",
 *      tags={"Extend"},
 *      summary="发布事件",
 *      description="向所有登录的WebSocket客户端发布事件信息",
 *      @OA\RequestBody(required=true, @OA\MediaType(
 *          mediaType="application/x-www-form-urlencoded", @OA\Schema(
 *              @OA\Property(description="事件名称", property="event", type="string", default="publish"),
 *              @OA\Property(description="目标编号", property="to", type="string"),
 *              @OA\Property(description="信息内容", property="data", type="string", default="sample data"),
 *          )
 *      )),
 *      @OA\Response(response="default", description="返回结果"),
 * )
 */
func (t *ApiService) Events(w http.ResponseWriter, r *http.Request) {
	//获取参数
	event := r.PostFormValue("event")
	data := r.PostFormValue("data")
	to := r.PostFormValue("to")
	// 整理格式
	result := make(map[string]string)
	result["__event"] = event
	result["data"] = data
	response,_ := json.Marshal(result)
	// 发布信息
	if to == "" {
		t.WebSocketService.Publish(string(response))
	} else {
		t.WebSocketService.Send(to, response)
	}
	//接口返回
	t.apiResponse(w, r, 0, "OK", result)
}