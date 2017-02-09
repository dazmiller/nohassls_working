package controllers

import (
	"github.com/astaxie/beego"
)

type LifeControllerPage struct {
	beego.Controller
}

func (this *LifeControllerPage) LifeQuotes() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/life_quotes.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/life_quotes.html"
	this.LayoutSections["Feature"] = "release/feature.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}

func (this *LifeControllerPage) LifeExplore() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/life_quotes.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/life_explore.html"
	this.LayoutSections["Feature"] = "release/feature.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}
