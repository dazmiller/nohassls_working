package controllers

import (
	"github.com/astaxie/beego"
)

type BikeController struct {
	beego.Controller
}

func (this *BikeController) Get() {
	this.TplName = "public/bike.html"
	this.Layout = "public/layout.html"

}
