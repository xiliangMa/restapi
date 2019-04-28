package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// Cluster object api list
type OrderDetailController struct {
	beego.Controller
}

// @Title GetOrderDetail
// @Description get OrderDetails
// @Param token header string true "Auth token"
// @Param name query string false "OrderDetail name"
// @Param from query int 0 false "from"
// @Param limit query int 20 false "limit"
// @Success 200 {object} models.Result
// @router / [post]
func (this *OrderDetailController) OrderDetailList() {
	name := this.GetString("name")
	limit, _ := this.GetInt("limit")
	from, _ := this.GetInt("from")
	this.Data["json"] = models.GetOrderDetailList(name, from, limit)
	this.ServeJSON(false)
}

// @Title AddOrderDetail
// @Description Add OrderDetail
// @Param token header string true "Auth token"
// @Param OrderDetail body models.OrderDetail true "OrderDetail object"
// @Success 200 {object} models.Result
// @router /addOrderDetail [post]
func (this *OrderDetailController) AddOrderDetail() {
	var c models.OrderDetail
	json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	this.Data["json"] = models.AddOrderDetail(&c)
	this.ServeJSON(false)

}

// @Title DelOrderDetail
// @Description Delete OrderDetail
// @Param token header string true "Auth token"
// @Param id path int true "OrderDetail id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *OrderDetailController) DeleteOrderDetail() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteOrderDetail(id)
	this.ServeJSON(false)

}
