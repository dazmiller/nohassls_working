package controllers

import (
	"github.com/astaxie/beego"
)

type RequestsController struct {
	beego.Controller
}

func (this *RequestsController) Get() {
	this.TplName = "public/profile.html"
	this.Layout = "public/layout.html"

}
