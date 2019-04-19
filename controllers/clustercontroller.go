package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/xiliangMa/restapi/models"
)

// Cluster object api list
type ClusterController struct {
	beego.Controller
}

// @Title GetCluster
// @Description get Clusters
// @Param token header string true "Auth token"
// @Param name query string false "Cluster name"
// @Param ip query string false "Cluster ip"
// @Param from query int 0 false "from"
// @Param limit query int 20 false "limit"
// @Success 200 {object} models.Result
// @router / [post]
func (this *ClusterController) ClusterList() {
	name := this.GetString("name")
	limit, _ := this.GetInt("limit")
	from, _ := this.GetInt("from")
	this.Data["json"] = models.GetClusterList(name, from, limit)
	this.ServeJSON(false)
}

// @Title AddCluster
// @Description Add Cluster
// @Param token header string true "Auth token"
// @Param Cluster body models.Cluster true "Cluster object"
// @Success 200 {object} models.Result
// @router /addCluster [post]
func (this *ClusterController) AddCluster() {
	var c models.Cluster
	json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	this.Data["json"] = models.AddCluster(&c)
	this.ServeJSON(false)

}

// @Title DelCluster
// @Description Delete Cluster
// @Param token header string true "Auth token"
// @Param id path int true "Cluster id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *ClusterController) DeleteCluster() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteCluster(id)
	this.ServeJSON(false)

}
