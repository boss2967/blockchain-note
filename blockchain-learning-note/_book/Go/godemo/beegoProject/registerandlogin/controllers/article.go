package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}
func (this *ArticleController)ShowArticleList(){
	this.TplName = "content.html"
}
func (this *ArticleController)AddArticle(){
	this.TplName = "add.html"
}
