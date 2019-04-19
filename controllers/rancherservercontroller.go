package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// RancherServers object api list
type RancherServerController struct {
	beego.Controller
}

// @Title GetRancherServer
// @Description get RancherServers
// @Param token header string true "Auth token"
// @Param region query string false "RancherServer name"
// @Param from query int 0 false "from"
// @Param limit query int 20 false "limit"
// @Success 200 {object} models.Result
// @router / [post]
func (this *RancherServerController) RancherServerList() {
	region := this.GetString("region")
	limit, _ := this.GetInt("limit")
	from, _ := this.GetInt("from")
	this.Data["json"] = models.GetRancherServerList(region, from, limit)
	this.ServeJSON(false)

}

// @Title AddRancherServer
// @Description Add RancherServer
// @Param token header string true "Auth token"
// @Param RancherServer body models.RancherServer true "RancherServer object"
// @Success 200 {object} models.Result
// @router /addRancherServer [post]
func (this *RancherServerController) AddRancherServer() {
	var h models.RancherServer
	json.Unmarshal(this.Ctx.Input.RequestBody, &h)
	this.Data["json"] = models.AddRancherServer(&h)
	this.ServeJSON(false)

}

// @Title DelRancherServer
// @Description Delete RancherServer
// @Param token header string true "Auth token"
// @Param id path int true "RancherServer id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *RancherServerController) DeleteRancherServer() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteRancherServer(id)
	this.ServeJSON(false)

}
