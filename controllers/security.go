package controllers

import (
	"github.com/astaxie/beego"
)

type SecurityController struct {
	beego.Controller
}

func (this *SecurityController) Get() {

	this.LayoutSections = make(map[string]string)
	this.TplName = "public/security.html"
	this.Layout = "public/layout.html"
	this.LayoutSections["PageScripts"] = "public/scripts_security.html"

}
