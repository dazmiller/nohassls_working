package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/index.html"
	//this.Layout = "public/layout.html"
	//this.LayoutSections["PageScripts"] = "public/scripts_dash.html"
}

func (this *MainController) Signup() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/signup.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["signup"] = "release/signup.html"
	this.LayoutSections["Footer"] = "release/footer.html"

}

func (this *MainController) GetLandingPage() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/life.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/life.html"
	this.LayoutSections["Feature"] = "release/feature.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}
