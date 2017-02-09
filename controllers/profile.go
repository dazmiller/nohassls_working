package controllers

import (
	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

func (this *ProfileController) Profile() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/profile.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/profile.html"
	this.LayoutSections["Feature"] = "release/feature.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}

func (this *ProfileController) ProfileClose() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/profile_close.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/profile_close.html"
	this.LayoutSections["Feature"] = "release/feature.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}

/*

import "github.com/astaxie/beego"

type ProfileController struct {
	beego.Controller
}


func (this *ProfileController) Post() {

	u := models.Profile{}
	if err := this.ParseForm(&u); err != nil {

		o := orm.NewOrm()

		//o.Insert(u)
		fmt.Println(o.Insert(u))
		this.LayoutSections = make(map[string]string)
		this.TplName = "public/profile.html"
		this.Layout = "public/layout.html"
		this.LayoutSections["PageScripts"] = "public/scripts_profile.html"
	}
	this.LayoutSections = make(map[string]string)
	this.TplName = "public/profile.html"
	this.Layout = "public/layout.html"
	this.LayoutSections["PageScripts"] = "public/scripts_profile.html"
}

//
//
//
func (this *ProfileController) Get() {
	flash := beego.ReadFromRequest(&this.Controller)

	if _, ok := flash.Data["notice"]; ok {
	}

	o := orm.NewOrm()
	o.Using("default")

	var profile []*models.Profile

	num, err := o.QueryTable("profiles").OrderBy("-id").All(&profile)

	if err != orm.ErrNoRows && num > 0 {
		this.TplName = "public/profile.html"
		this.Data["articles"] = profile
	} else {
		// No result
		this.TplName = "blog/error.tpl"
		this.Data["error"] = "No article in the database"
	}
	this.LayoutSections = make(map[string]string)
	this.Data["title"] = "My blog"
	this.Layout = "public/layout.html"
	this.LayoutSections["PageScripts"] = "public/scripts_profile.html"
}

func (this *ProfileController) GetOne() {
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

func (this *ProfileController) Add() {
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

func (this *ProfileController) Edit() {
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

func (this *ProfileController) Delete() {
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

*/
