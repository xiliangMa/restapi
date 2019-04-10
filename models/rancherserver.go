package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type RancherServer struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	Url         string    `xorm:"not null VARCHAR(128)"`
	AccessKey   string    `xorm:"not null VARCHAR(128)"`
	SecretKey   string    `xorm:"not null VARCHAR(128)"`
	Region      string    `xorm:"not null VARCHAR(128)"`
	CreateTime  time.Time `xorm:"TIMESTAMP"`
	UpdateTime  time.Time `xorm:"TIMESTAMP"`
	RegionEnUs  string    `xorm:"VARCHAR(128)"`
	NetworkType string    `xorm:"default '' comment('网络类型：inner:内网；outer：外网') VARCHAR(10)"`
}

func init() {
	orm.RegisterModel(new(RancherServer))
}

func GetRancherServerList(region string, page, number int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var RancherServerList []*RancherServer
	var ResultData Result
	_, err := o.QueryTable("rancher_server").Filter("region__icontains", region).Limit(number, page).All(&RancherServerList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetRancherServerListErr
		return ResultData
	}

	ResultData.Code = 200
	ResultData.Data = RancherServerList
	return ResultData
}

func AddRancherServer(rancherserver *RancherServer) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(rancherserver)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddRancherServerErr
		return ResultData
	}
	ResultData.Code = 200
	ResultData.Data = id
	return ResultData
}

func DeleteRancherServer(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&RancherServer{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteRancherServerErr
		return ResultData
	}
	ResultData.Code = 200
	return ResultData

}
