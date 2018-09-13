package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"github.com/astaxie/beego/orm"
	"registerandlogin/models"
	"time"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}
func (this *ArticleController)ShowArticleList(){
	o:=orm.NewOrm()
	qs:=o.QueryTable("article")
	var articles[]models.Article
	qs.All(&articles)
	beego.Info(articles)
	this.Data["article"]=articles
	this.TplName="index.html"
}

func (this *ArticleController)ShowAddArticle(){
	this.TplName = "add.html"
}


func (this *ArticleController)ShowContent(){
//	获取id
	id:=this.GetString("id")
//	查询数据
	o:=orm.NewOrm()
	article:=models.Article{}
	article.Id2,_= strconv.Atoi(id)
	err:=o.Read(&article)
	if err != nil {
		beego.Info("数据为空")
		return
	}
	article.Count+=1
	o.Update(&article)
//	传递给视图
	this.Data["article"] = article
	this.TplName = "content.html"
}
func (this *ArticleController)HandleAddArticle(){
	articleName:=this.GetString("articleName")
	artiContent:=this.GetString("content")
	beego.Info(articleName,artiContent)
	f,h,err:=this.GetFile("uploadname")
	var fileName string
	if err != nil {
		beego.Info("获取文件失败err:",err)
	}else {
		beego.Info(h.Filename)
		defer f.Close()
		//1.格式
		ext := path.Ext(h.Filename)
		if path.Ext(h.Filename) != ".jpg" && path.Ext(h.Filename) != ".png" && path.Ext(h.Filename) != ".jepg" {
			beego.Info("格式不正确")
			return
		}
		//2.大小
		if h.Size>1024*1024*10{
			beego.Info("文件过大")
			return
		}
		//3.不重名
		fileName = time.Now().Format("2006-01-02 15:04:05")+ext
		err = this.SaveToFile("uploadname", "./static/img/"+fileName+ext)
		if err != nil {
			beego.Info("上传失败", err)
			return
		}
		beego.Info("上传成功")
	}
	o:=orm.NewOrm()
	article:=models.Article{}
	article.Img="./static/img/"+fileName
	article.Title = articleName
	article.Content = artiContent
	_,err=o.Insert(&article)
	if err != nil {
		beego.Info("insert err:",err)
		return
	}
	beego.Info("插入成功")
	this.Redirect("showArticle",302)
}

func (this *ArticleController)HandleDeleteArticle(){
//	1,获取id
//	2，新建article，
//	3，delete
	id,err:=this.GetInt("id")
	if err != nil {
		beego.Info(err)
		return
	}
	o:=orm.NewOrm()
	article:= models.Article{Id2:id}
	_,err=o.Delete(&article)
	if err != nil {
		beego.Info("删除失败")
	}
	this.Redirect("/showArticle",302)
}

func (this *ArticleController)ShowUpdateArticle(){
	id,err:=this.GetInt("id")
	if err != nil {
		beego.Info("获取id失败",err)
		return
	}
	beego.Info("用户id为：",id)
	o:=orm.NewOrm()
	article:=models.Article{}
	article.Id2 = id
	o.Read(&article)
	this.Data["article"] = article
	this.TplName = "update.html"
}

func (this *ArticleController)HandleUpdateArticle(){
	title:=this.GetString("articleName")
	content:=this.GetString("content")
	id:= this.GetString("articleId")
	beego.Info("title",title)
	beego.Info("content",content)
	beego.Info("id",id)
	var fileName string
	f,h,err:=this.GetFile("uploadname")
	if err!=nil{
		beego.Info("获取文件失败err:",err)
	}else {
		defer f.Close()
		ext:=path.Ext(h.Filename)
		if ext!=".jpg"&&ext!=".png"&&ext!="jepg"{
			beego.Info("格式不正确")
			return
		}
		if h.Size>1024*1024*10{
			beego.Info("文件过大")
			return
		}
		fileName=time.Now().Format("2006-01-02 15:04:05")+ext
		err=this.SaveToFile("uploadname","./static/img/"+fileName)
		if err != nil {
			beego.Info("保存文件失败")
			return
		}
	}

	o := orm.NewOrm()
	article := models.Article{}
	article.Id2,_ = strconv.Atoi(id)
	article.Title = title
	article.Content = content
	article.Img = fileName
	_,err=o.Update(&article)
	if err != nil {
		beego.Info("err:",err)
		return
	}
	this.Redirect("showArticle",302)
}