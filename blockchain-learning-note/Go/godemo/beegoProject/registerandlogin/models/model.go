package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego"
)

type User struct {
	Id int
	UserName string
	Passwd string
}

type Article struct {
	Id2 int `orm:"pk;auto"`//手动设置为主键，自增长
	Title string `orm:size(30)`//不为空
	Content string `orm:"size(500)"`
	Img string `orm:"size(500);null"`
	//Type string
	Time time.Time `orm:type(datetime);`
	Count int
}

func init()  {
	beego.Info("---------------init-------------------")
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/newweb?charset=utf8")
	orm.RegisterModel(new(User),new(Article))
	orm.RunSyncdb("default",false,true)
}
