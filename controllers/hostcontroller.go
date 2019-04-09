package controllers

import (
	"restapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

// Hosts object api list
type HostController struct {
	beego.Controller
}

// @Title GetHost
// @Description get Hosts
// @Param name query string false "host name"
// @Param ip query string false "host ip"
// @Param page query int 1 false "page"
// @Param number query int 20 false "page"
// @Success 200 {object} models.Result
// @router / [post]
func (this *HostController) HostList() {
	name := this.GetString("name")
	ip := this.GetString("ip")
	number, _ := this.GetInt("number")
	page, _ := this.GetInt("page")
	this.Data["json"] = models.GetHostList(name, ip, page, number)
	this.ServeJSON(false)

}

// @Title AddHost
// @Description dd Host
// @Param Host body models.Host true "host object"
// @Success 200 {object} models.Result
// @router /addhost [post]
func (this *HostController) AddHost() {
	var h models.Host
	json.Unmarshal(this.Ctx.Input.RequestBody, &h)
	this.Data["json"] = models.AddHost(&h)
	this.ServeJSON(false)

}



// @Title AddHost
// @Description dd Host
// @Param id path int true "host id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *HostController) DeleteHost() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteHost(id)
	this.ServeJSON(false)

}