var websocket_io = function(sdn,id){
    // 定义类
    var self = this;
    var ws = new WebSocket(sdn);

    // 定义属性
    this.id = id || Date.now().toString(36);
    this.onopen = null;
    this.onclose = null;
    this.onmessage = null;
    this.onevent = {};

    // 发送信息
    this.send = function(event, data){
        ws.send(JSON.stringify({"event":event, "data":data}));
    };

    // 连接成功
    ws.onopen = function() {
        // 发送登录信息
        ws.send(JSON.stringify({"event":"login", "data":self.id}));
        // 触发事件
        if(typeof self.onopen === 'function')
            self.onopen();
        //DEBUG
        //console.log("[*] WebSocket connected.");
    };
    // 关闭连接
    ws.onclose = function() {
        // 触发事件
        if(typeof self.onclose === 'function')
            self.onclose();
        // DEBUG
        console.log("[*] WebSocket close.");
    };
    // 接收信息
    ws.onmessage = function(event) {
        //解析报文
        var json = null;
        try{
            json = JSON.parse(event.data);
        }catch (e){
            json = null
        }
        //判断报文触发事件
        if(json && json.event && json.data){
            // 触发事件
            if(typeof self.onevent[json.event] === 'function')
                self.onevent[json.event](json.data);
        } else {
            // 触发事件
            if(typeof self.onmessage === 'function')
                self.onmessage(event.data);
        }
    };

    // 返回
    return this;
};