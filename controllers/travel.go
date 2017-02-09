package controllers

import (
	"github.com/astaxie/beego"
)

type TravelController struct {
	beego.Controller
}

func (this *TravelController) Get() {
	this.TplName = "public/travel.html"
	this.Layout = "public/layout.html"

}
