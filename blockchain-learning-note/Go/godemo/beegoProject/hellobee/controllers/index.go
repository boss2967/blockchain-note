package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func(this *IndexController)Get(){
	this.Data["type"]="默认的get"
	this.TplName = "hello.html"
}

func(this *IndexController)Post(){
	this.Data["type"]="默认的post"
	this.TplName = "hello.html"
}

func(this *IndexController)NewGet(){
	this.Data["type"]="自定义的get"
	this.TplName = "hello.html"
}