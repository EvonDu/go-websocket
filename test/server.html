<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>测试服务端</title>
    <script src="assets/layui/layui.js"></script>
    <script src="assets/jquery/jquery-3.2.1.min.js"></script>
    <link rel="stylesheet" href="assets/layui/css/layui.css"/>
    <link rel="stylesheet" href="assets/layui/css/modules/code.css"/>
    <link rel="stylesheet" href="css/main.css"/>
</head>
<body>
    <!-- 发送元素 -->
    <div class="input-group">
        <input name="input" lay-verify="title" autocomplete="off" placeholder="请输入信息" class="layui-input" type="text">
        <button class="layui-btn" onclick="sendMessage()">发送信息</button>
    </div>
    <!-- 信息板面 -->
    <pre class="layui-code">信息面板：</pre>
</body>
</html>

<!-- layui -->
<script>
    layui.use('code', function(){
        //初始化面板
        layui.code({skin: 'notepad'});
    });
</script>
<!-- ws -->
<script>
    //发送信息
    function sendMessage(){
        //获取元素
        var input = document.getElementsByName('input')[0];

        //发送信息
        $.ajax({
            url:"../api/publish.php",
            data:{"msg":input.value},
            success:function(res){
                printMessage("接口返回:"+JSON.stringify(res));
                input.value = "";
            }
        });
    }

    //输出信息
    function printMessage(msg){
        //获取元素
        var ol = document.getElementsByClassName('layui-code-ol')[0];

        //新增元素
        var li = document.createElement('li');
        li.innerText = msg;
        ol.appendChild(li);
    }
</script>