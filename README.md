# Go WebSocket
Golang搭建的WebSocket服务框架，可以直接当一个基础的WebSocket服务使用，也可以使用框架中的扩展服务完成客户端注册、私聊等高级功能

## 使用方法
* 运行方法：`go run main.go`
* 运行参数：

参数 | 名称    | 默认值 | 说明
-----|---------|--------|---------
-p   | port    | 8080   | 端口
-t   | test    | fasle  | 开启调试页面
-s   | swagger | false  | 开启接口文档

## 控制台操作

参数    | 例子                | 说明
--------|---------------------|-------------
help    | help                | 帮助信息
count   | count               | 当前连接数量
connect | connect             | 客户端列表
publish | publish HelloWord!  | 向所有WebSocket发布信息

## 使用例子
需要引入javascript：`/js/websocket_io.js`
```
ws = new websocket_io("ws://127.0.0.1:8080");
ws.onopen = function() {
    console.log("连接成功!");
    console.log("客户端编号："+ws.id);
};
ws.onclose = function() {
    console.log("连接断开!");
};
ws.onmessage = function(data) {
    console.log("收到信息："+ data);
};
ws.onevent["publish"] = function(data) {
    console.log("收到事件[publish]："+ data);
};
```
