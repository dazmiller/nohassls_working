package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "release/index.html"
	//this.Layout = "public/layout.html"

}
