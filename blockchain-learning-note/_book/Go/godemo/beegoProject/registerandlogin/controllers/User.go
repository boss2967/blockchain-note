package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"registerandlogin/models"
)

type RegController struct {
	beego.Controller
}
func (this *RegController)ShowReg(){
	this.TplName="register.html"
}
func (this *RegController)HandleReg(){
	//1.获取前端数据
	//2，处理数据
	//3，注册成功，插入数据库
	//4，注册失败

	//key是input的name
	username:=this.GetString("username")
	password:=this.GetString("password")
	beego.Info(username,password)
	if username==""||password==""{
		beego.Info("用户名密码不能为空")
		this.TplName = "register.html"
		return
	}

	o:=orm.NewOrm()
	user:=models.User{}
	user.UserName = username
	err:=o.Read(&user,"user_name")
	if err == nil {
		beego.Info("用户已存在")
		this.TplName="register.html"
		return
	}
	user.UserName = username
	user.Passwd = password
	_,err=o.Insert(&user)
	if err != nil {
		beego.Info("插入失败",err)
		this.TplName="register.html"
		return
	}
	//this.TplName = "login.html"
	this.Redirect("/login",302)
}

type LoginController struct {
	beego.Controller
}
func (this *LoginController)ShowLogin(){
	this.TplName="login.html"
}
func (this *LoginController)HandleLogin(){
	username:=this.GetString("username")
	password:=this.GetString("password")
	if username==""||password==""{
		beego.Info("用户名密码不能为空")
		this.TplName = "register.html"
		return
	}
	o:=orm.NewOrm()
	user:=models.User{}
	user.UserName = username
	err:=o.Read(&user,"user_name")
	if err != nil {
		beego.Info("用户名错误")
		this.TplName = "login.html"
		return
	}
	if password!=user.Passwd{
		beego.Info("密码错误")
		this.TplName = "login.html"
		return
	}
	this.Ctx.WriteString("登录成功")
}