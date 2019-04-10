package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// Promotion object api list
type PromotionController struct {
	beego.Controller
}

// @Title GetPromotion
// @Description get Promotions
// @Param name query string false "Promotion name"
// @Param page query int 1 false "page"
// @Param number query int 20 false "page"
// @Success 200 {object} models.Result
// @router / [post]
func (this *PromotionController) PromotionList() {
	name := this.GetString("name")
	number, _ := this.GetInt("number")
	page, _ := this.GetInt("page")
	this.Data["json"] = models.GetPromotionList(name, page, number)
	this.ServeJSON(false)
}

// @Title AddPromotion
// @Description Add Promotion
// @Param Promotion body models.Promotion true "Promotion object"
// @Success 200 {object} models.Result
// @router /addPromotion [post]
func (this *PromotionController) AddPromotion() {
	var c models.Promotion
	json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	this.Data["json"] = models.AddPromotion(&c)
	this.ServeJSON(false)

}

// @Title AddPromotion
// @Description dd Promotion
// @Param id path int true "Promotion id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *PromotionController) DeletePromotion() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeletePromotion(id)
	this.ServeJSON(false)
}
