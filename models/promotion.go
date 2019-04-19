package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type Promotion struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"not null comment('促销名称') VARCHAR(50)"`
	Content    string    `xorm:"comment('促销描述') TEXT"`
	BeginDate  time.Time `xorm:"comment('开始时间') DATETIME"`
	EndDate    time.Time `xorm:"comment('结束时间') DATETIME"`
	Status     int       `xorm:"comment('状态') INT(11)"`
	UpdateTime time.Time `xorm:"DATETIME"`
	CreateTime time.Time `xorm:"DATETIME"`
}

func init() {
	orm.RegisterModel(new(Promotion))
}

func GetPromotionList(name string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var PromotionList []*Promotion
	var ResultData Result
	_, err := o.QueryTable("promotion").Filter("name__icontains", name).Limit(limit, from).All(&PromotionList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetPromotionListErr
		logs.Error("GetPromotionList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	total, _ := o.QueryTable("promotion").Count()
	data := make(map[string]interface{})
	data["items"] = PromotionList
	data["total"] = total
	ResultData.Code = 200
	ResultData.Data = data
	return ResultData
}

func AddPromotion(Promotion *Promotion) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(Promotion)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddPromotionErr
		logs.Error("AddPromotion failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeletePromotion(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&Promotion{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeletePromotionErr
		logs.Error("DeletePromotion failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
