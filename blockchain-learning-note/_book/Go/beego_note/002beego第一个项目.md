# post案例实现

1. 在views下增加对应的html文件，hello.html
	
	```
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <meta http-equiv="X-UA-Compatible" content="ie=edge">
	    <title>Document</title>
	</head>
	<body>
	<h1>hellobee</h1>
	<h1>{{.type}}</h1>
	<form method="POST" action="/index">
	    <input type="submit">
	</form>
	</body>
	</html>
	```

2. router.go 

	```
	package routers
	import (
		"hellobee/controllers"
		"github.com/astaxie/beego"
	)
	func init() {
	    beego.Router("/", &controllers.MainController{})
	    //默认使用Get或Post方法，可以自定义方法
	    beego.Router("/index",&controllers.IndexController{},"get:NewGet;post:Post")
	}
	```

3. index.go

	```
	package controllers
	import "github.com/astaxie/beego"
	type IndexController struct {
		beego.Controller
	}
	//默认的get
	func(this *IndexController)Get(){
		this.Data["type"]="默认的get"
		this.TplName = "hello.html"
	}
	//默认的post
	func(this *IndexController)Post(){
		this.Data["type"]="默认的post"
		this.TplName = "hello.html"
	}
	//自定义的get
	func(this *IndexController)NewGet(){
		this.Data["type"]="自定义的get"
		this.TplName = "hello.html"
	}
	```
	