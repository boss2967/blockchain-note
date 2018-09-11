package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

//1.放结构体
//2.初始化操作

type User struct {
	//不写id会报错
	Id int
	Name string
}

func init()  {
	//1.注册数据库 必须有一个默认的default
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	//2，注册表
	orm.RegisterModel(new(User))
	//3.生成表
	err:=orm.RunSyncdb("default",false,true)
	if err != nil {
		beego.Info("注册表失败")
	}
}
