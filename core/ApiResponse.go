package core

import (
	"net/http"
	"encoding/json"
)

//定义结构体
type ApiResponse struct {}

//按接口结构输出结果
func (t *ApiResponse) Send(w http.ResponseWriter, r *http.Request, code int, message string, data interface{}){
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