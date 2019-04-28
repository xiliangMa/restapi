package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// Cluster object api list
type OrderRenewalController struct {
	beego.Controller
}

// @Title GetOrderRenewal
// @Description get OrderRenewals
// @Param token header string true "Auth token"
// @Param name query string false "OrderRenewal name"
// @Param from query int 0 false "from"
// @Param limit query int 20 false "limit"
// @Success 200 {object} models.Result
// @router / [post]
func (this *OrderRenewalController) OrderRenewalList() {
	name := this.GetString("name")
	limit, _ := this.GetInt("limit")
	from, _ := this.GetInt("from")
	this.Data["json"] = models.GetOrderRenewalList(name, from, limit)
	this.ServeJSON(false)
}

// @Title AddOrderRenewal
// @Description Add OrderRenewal
// @Param token header string true "Auth token"
// @Param OrderRenewal body models.OrderRenewal true "OrderRenewal object"
// @Success 200 {object} models.Result
// @router /addOrderRenewal [post]
func (this *OrderRenewalController) AddOrderRenewal() {
	var c models.OrderRenewal
	json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	this.Data["json"] = models.AddOrderRenewal(&c)
	this.ServeJSON(false)

}

// @Title DelOrderRenewal
// @Description Delete OrderRenewal
// @Param token header string true "Auth token"
// @Param id path int true "OrderRenewal id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *OrderRenewalController) DeleteOrderRenewal() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteOrderRenewal(id)
	this.ServeJSON(false)

}
