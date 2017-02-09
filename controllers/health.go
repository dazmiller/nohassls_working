package controllers

import (
	"github.com/astaxie/beego"
)

type HealthController struct {
	beego.Controller
}

func (this *HealthController) Get() {
	this.TplName = "public/health.html"
	this.Layout = "public/layout.html"

}
