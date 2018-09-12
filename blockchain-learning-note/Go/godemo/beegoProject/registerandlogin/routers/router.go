package routers

import (
	"registerandlogin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/",&controllers.LoginController{},"get:ShowLogin")
	beego.Router("/register", &controllers.RegController{},"get:ShowReg;post:HandleReg")
	beego.Router("/login", &controllers.LoginController{},"get:ShowLogin;post:HandleLogin")


	beego.Router("/showArticle", &controllers.ArticleController{},"get:ShowArticleList")
	beego.Router("/addArticle", &controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/articleContent", &controllers.ArticleController{},"get:ShowContent")
	beego.Router("/deleteArticle", &controllers.ArticleController{},"get:HandleDeleteArticle")
	beego.Router("/updateArticle", &controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
}
