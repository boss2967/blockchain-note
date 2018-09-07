# ?如果我阻止冒泡了，事件委托还能成立吗？
# 服务器，客户端联通
# ajax
>	ajax一个前后台配合的技术，它可以让javascript发送http请求，与后台通信，获取数据和信息。ajax技术的原理是实例化xmlhttp对象，使用此对象与后台通信。jquery将它封装成了一个函数$.ajax()，我们可以直接用这个函数来执行ajax请求。

ajax需要在服务器环境下运行。

## $.ajax使用方法

常用参数：

1. url 请求地址
2. type 请求方式，默认是'get'，常用的还有'post'
3. dataType 设置返回的数据格式，常用的是'json'格式，也可以设置为'text'
4. data 设置发送给服务器的数据
5. success 设置请求成功后的回调函数
6. error 设置请求失败后的回调函数
7. async 设置是否异步，默认值是'true'，表示异步

**以前的写法：**

```
$.ajax({
    url: '/change_data',
    type: 'get',
    dataType: 'json',
    data:{'code':300268}
    success:function(dat){
        alert(dat.name);
    },
    error:function(){
        alert('服务器超时，请重试！');
    }
});
```
**新的写法(推荐)：**

```
$.ajax({
    url: '/change_data',
    type: 'get',
    dataType: 'json',
    data:{'code':300268}
})
.done(function(dat) {
    alert(dat.name);
})
.fail(function() {
    alert('服务器超时，请重试！');
});
```

## $.ajax的简写方式

$.ajax按照请求方式可以简写成$.get或者$.post方式

```
$.get(URL,data,function(data){},dataType);
$.post(URL,data,function(data){},dataType);
```

## 注意

`<div class="login_btn fl">`

指的是使用了两个样式，用第一个选择就可以。

## 练习


1. 启动服务器环境`node server.js`
2. 编写index.html
3. 服务器访问localhost:8888

```
$(function () {
            $.ajax({
                url: 'js/data.json',
                type: 'get',
                dataType: 'json',
                success: function (data) {
                    console.log(data);
             
                },
                error: function () {
                    alert("失败")
                }
            })
        })
```