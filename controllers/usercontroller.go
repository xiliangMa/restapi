package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// Users object api list
type UserController struct {
	beego.Controller
}

// @Title GetUserInfo
// @Description get User Info
// @Param token header string true "Auth token"
// @Success 200 {object} models.Result
// @router /info [get]
func (this *UserController) UserInfo() {
	token := this.Ctx.Input.Header("token")
	this.Data["json"] = models.GetUserInfo(token)
	this.ServeJSON(false)
}

// @Title GetUser
// @Description get Users
// @Param token header string true "Auth token"
// @Param mobile query string false "User mobile"
// @Param email query string false "User email"
// @Param from query int 0 false "from"
// @Param limit query int 20 false "limit"
// @Success 200 {object} models.Result
// @router / [post]
func (this *UserController) UserList() {
	mobile := this.GetString("mobile")
	email := this.GetString("email")
	limit, _ := this.GetInt("limit")
	from, _ := this.GetInt("from")
	this.Data["json"] = models.GetUserList(mobile, email, from, limit)
	this.ServeJSON(false)

}

// @Title AddUser
// @Description Add User
// @Param token header string true "Auth token"
// @Param User body models.User true "User object"
// @Success 200 {object} models.Result
// @router /addUser [post]
func (this *UserController) AddUser() {
	var h models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &h)
	this.Data["json"] = models.AddUser(&h)
	this.ServeJSON(false)

}

// @Title DelUser
// @Description Delete User
// @Param token header string true "Auth token"
// @Param id path int true "User id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *UserController) DeleteUser() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteUser(id)
	this.ServeJSON(false)

}
