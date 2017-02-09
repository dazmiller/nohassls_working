package controllers

import (
	"github.com/astaxie/beego"
)

type FAQController struct {
	beego.Controller
}

func (this *FAQController) Get() {
	this.TplName = "public/tpd.html"
	this.Layout = "public/layout.html"

}
