package models

import (
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

func GetPromotionList(name, ip string, page, number int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var PromotionList []*Promotion
	var ResultData Result
	_, err := o.QueryTable("Promotion").Filter("name__icontains", name).Filter("Promotion_ip__icontains", ip).Limit(number, page).All(&PromotionList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetPromotionListErr
		return ResultData
	}

	ResultData.Code = 200
	ResultData.Data = PromotionList
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
		return ResultData
	}
	ResultData.Code = 200
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
		return ResultData
	}
	ResultData.Code = 200
	return ResultData
}