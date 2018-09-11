package routers

import (
	"hellobee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index",&controllers.IndexController{},"get:NewGet;post:Post")
}
