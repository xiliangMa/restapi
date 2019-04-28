package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// Cluster object api list
type OrderMasterController struct {
	beego.Controller
}

// @Title GetOrderMaster
// @Description get OrderMasters
// @Param token header string true "Auth token"
// @Param name query string false "OrderMaster name"
// @Param from query int 0 false "from"
// @Param limit query int 20 false "limit"
// @Success 200 {object} models.Result
// @router / [post]
func (this *OrderMasterController) OrderMasterList() {
	name := this.GetString("name")
	limit, _ := this.GetInt("limit")
	from, _ := this.GetInt("from")
	this.Data["json"] = models.GetOrderMasterList(name, from, limit)
	this.ServeJSON(false)
}

// @Title AddOrderMaster
// @Description Add OrderMaster
// @Param token header string true "Auth token"
// @Param OrderMaster body models.OrderMaster true "OrderMaster object"
// @Success 200 {object} models.Result
// @router /addOrderMaster [post]
func (this *OrderMasterController) AddOrderMaster() {
	var c models.OrderMaster
	json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	this.Data["json"] = models.AddOrderMaster(&c)
	this.ServeJSON(false)

}

// @Title DelOrderMaster
// @Description Delete OrderMaster
// @Param token header string true "Auth token"
// @Param id path int true "OrderMaster id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *OrderMasterController) DeleteOrderMaster() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteOrderMaster(id)
	this.ServeJSON(false)

}
