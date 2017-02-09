package controllers

import (
	"github.com/astaxie/beego"
)

type SocialController struct {
	beego.Controller
}

func (this *SocialController) Get() {
	this.TplName = "public/social.html"
	this.Layout = "public/layout.html"

}
