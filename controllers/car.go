package controllers

import (
	"github.com/astaxie/beego"
)

type CarController struct {
	beego.Controller
}

func (this *CarController) Get() {
	this.TplName = "public/car.html"
	this.Layout = "public/layout.html"

}
