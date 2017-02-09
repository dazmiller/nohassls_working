package controllers

import (
	models "nohassls_material2/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
	flash := beego.ReadFromRequest(&this.Controller)

	if _, ok := flash.Data["notice"]; ok {
	}

	o := orm.NewOrm()
	o.Using("default")

	var articles []*models.Articles

	num, err := o.QueryTable("Articles").OrderBy("-id").All(&articles)

	if err != orm.ErrNoRows && num > 0 {
		this.TplName = "blog/listing.tpl"
		this.Data["articles"] = articles
	} else {
		// No result
		this.TplName = "blog/error.tpl"
		this.Data["error"] = "No article in the database"
	}

	this.Data["title"] = "My blog"
	this.Layout = "blog/layout_blog.tpl"
}

func (this *BlogController) GetOne() {
	o := orm.NewOrm()
	o.Using("default")

	// Get the ID page
	articlesId := this.Ctx.Input.Param(":id")

	var articles []*models.Articles

	err := o.QueryTable("articles").Filter("id", articlesId).One(&articles)

	if err == orm.ErrNoRows {
		// No result
		this.TplName = "blog/error.tpl"
		this.Data["title"] = "Error :("
		this.Data["error"] = "No available article"
	} else {
		this.TplName = "blog/content.tpl"
		for _, data := range articles {
			this.Data["title"] = data.Title
			this.Data["content"] = data.Content
		}

	}

	this.Layout = "blog/layout_blog.tpl"
}

func (this *BlogController) Add() {
	o := orm.NewOrm()
	o.Using("default")

	articles := models.Articles{}

	this.Data["Form"] = &articles

	if err := this.ParseForm(&articles); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		valid := validation.Validation{}
		valid.Required(articles.Title, "Title")
		valid.Required(articles.Content, "Content")
		isValid, _ := valid.Valid(articles)

		if this.Ctx.Input.Method() == "POST" {

			if !isValid {
				this.Data["errors"] = valid.ErrorsMap
				for _, err := range valid.Errors {
					beego.Error(err.Key, err.Message)
				}
			} else {
				_, err := o.Insert(&articles)
				flash := beego.NewFlash()

				if err == nil {
					flash.Notice("Article " + articles.Title + " added")
					flash.Store(&this.Controller)
					this.Redirect("/blog", 302)
				} else {
					beego.Debug("Couldn't insert new article. Reason: ", err)
				}
			}

		}

	}

	this.Layout = "blog/layout_blog.tpl"
	this.TplName = "blog/form.tpl"
	this.Data["title"] = "Add an article"
}

func (this *BlogController) Edit() {
	o := orm.NewOrm()
	o.Using("default")

	articlesId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	articles := models.Articles{}

	flash := beego.NewFlash()

	err := o.QueryTable("articles").Filter("id", articlesId).One(&articles)

	if err != orm.ErrNoRows {

		this.Data["Form"] = &articles

		if err := this.ParseForm(&articles); err != nil {
			beego.Error("Couldn't parse the form. Reason: ", err)
		} else {
			valid := validation.Validation{}
			valid.Required(articles.Title, "Title")
			valid.Required(articles.Content, "Content")
			isValid, _ := valid.Valid(articles)

			if this.Ctx.Input.Method() == "POST" {

				if !isValid {
					this.Data["errors"] = valid.ErrorsMap
					beego.Error("Form didn't validate.")
				} else {
					_, err := o.Update(&articles)

					if err == nil {
						flash.Notice("Article " + articles.Title + " updated")
						flash.Store(&this.Controller)

						this.Redirect("/blog", 302)
					} else {
						beego.Debug("Couldn't update new article. Reason: ", err)
					}
				}

			}

		}

		this.Data["title"] = "Edit this article"
		this.Layout = "blog/layout_blog.tpl"
		this.TplName = "blog/form.tpl"

	} else {
		flash.Notice("Article #%d doesn't exists", articlesId)
		flash.Store(&this.Controller)
		this.Redirect("/blog", 302)
	}

}

func (this *BlogController) Delete() {
	o := orm.NewOrm()
	o.Using("default")

	articlesId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	articles := models.Articles{}

	flash := beego.NewFlash()

	if exist := o.QueryTable(articles.TableName()).Filter("Id", articlesId).Exist(); exist {

		if num, err := o.Delete(&models.Articles{Id: articlesId}); err == nil {
			beego.Info("Record Deleted. ", num)
			flash.Notice("Article #%d deleted", articlesId)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}

	} else {
		flash.Notice("Article #%d doesn't exists", articlesId)
	}

	flash.Store(&this.Controller)

	this.Redirect("/blog", 302)
}
