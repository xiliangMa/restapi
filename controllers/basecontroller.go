package controllers

import "github.com/astaxie/beego"

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
