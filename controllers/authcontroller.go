package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xiliangMa/restapi/models"
	"github.com/xiliangMa/restapi/utils"
)

type AuthController struct {
	beego.Controller
}

// @Title Signin
// @Description get auth
// @Param name query string true "user name"
// @Param pwd query string true "user pwd"
// @Success 200 {object} models.Result
// @router /signin [post]
func (this *AuthController) Signin() {
	name := this.GetString("name")
	pwd := this.GetString("pwd")
	var ResultData models.Result
	result, code := utils.GenToken(name, pwd)
	ResultData.Code = code
	if code != utils.Success {
		ResultData.Message = result
	} else {
		ResultData.Data = result
	}
	this.Data["json"] = ResultData
	this.ServeJSON(false)
}

// @Title authorize
// @Description authorize
// @Param token header string true "Auth token"
// @Success 200 {object} models.Result
// @router /authorize [post]
func (this *AuthController) Authorize() {
	token := this.GetString("token")

	var ResultData models.Result
	result, code := utils.CheckToken(token)
	ResultData.Code = code
	if code != utils.Success {
		ResultData.Message = result
		logs.Error("Authorize failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
	} else {
		ResultData.Data = result
	}
	this.Data["json"] = ResultData
	this.ServeJSON(false)
}
