package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"

}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Post() {
	c.Data["type"] = "post对应的方法"
	c.TplName = "test.html"
}


func (c *IndexController) ShowGet() {
	c.Data["type"] = "get对应的方法"
	c.TplName = "test.html"
	id:=c.GetString(":id")
	beego.Info(id)
}