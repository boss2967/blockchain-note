package routers

import (
	"beedemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/insert/?:id",&controllers.OrmController{},"get:InsertOrm")
	beego.Router("/query",&controllers.OrmController{},"get:QueryOrm")
	beego.Router("/update",&controllers.OrmController{},"get:UpdateOrm")
	beego.Router("/delete/?:id",&controllers.OrmController{},"get:DeleteOrm")

	beego.Router("/mysql/create",&controllers.MysqlController{},"get:CreateMysql")
	beego.Router("/mysql/insert",&controllers.MysqlController{},"get:InsertMysql")
	beego.Router("/mysql/query",&controllers.MysqlController{},"get:QueryMysql")
	beego.Router("/mysql/delete/?:id",&controllers.MysqlController{},"get:DeleteMysql")
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/index/?:id", &controllers.IndexController{},"get:ShowGet;post:Post")//指定路由
    beego.Router("/index", &controllers.IndexController{},"get:ShowGet;post:Post")//指定路由
    //mysql路由
    //beego.Router("/orm",&controllers.OrmController{},"get:CreateOrm")


}
