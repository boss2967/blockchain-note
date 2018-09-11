package controllers

import (
	"github.com/astaxie/beego"
	//指定驱动 一定别忘了
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
	"beedemo/models"
	"strconv"
)

type OrmController struct {
	beego.Controller
}

func (this *OrmController) InsertOrm() {
	//	1.获取连接对象，rom
	o := orm.NewOrm()
	//	2.插入对象
	user := models.User{}
	user.Name = "张三"
	//	3.执行插入操作
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败", err)
		return
	}
	this.Ctx.WriteString("插入成功")
}

func (this *OrmController) QueryOrm() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		this.Ctx.WriteString("查询失败" + err.Error())
		return
	}
	this.Ctx.WriteString("查询成功" + user.Name)
}

func (this *OrmController) UpdateOrm() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	user.Name = "新名字"
	_, err = o.Update(&user)
	if err != nil {
		beego.Info("更新失败", err)
	}
	this.Ctx.WriteString("更新成功")
}

func (this *OrmController) DeleteOrm() {
	id := this.GetString(":id")
	o := orm.NewOrm()
	user := models.User{}
	user.Id, _ = strconv.Atoi(id)
	_, err := o.Delete(&user)
	if err != nil {
		this.Ctx.WriteString("删除失败" + err.Error())
	}
	this.Ctx.WriteString("删除成功")
}
