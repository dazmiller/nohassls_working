package controllers

import (
	"github.com/astaxie/beego"
)

type OfferController struct {
	beego.Controller
}

func (this *OfferController) Get() {
	this.TplName = "public/offers.html"
	this.Layout = "public/layout.html"

}
