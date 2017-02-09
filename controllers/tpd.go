package controllers

import (
	"github.com/astaxie/beego"
)

type TPDController struct {
	beego.Controller
}

func (this *TPDController) Get() {
	this.TplName = "public/tpd.html"
	this.Layout = "public/layout.html"

}
