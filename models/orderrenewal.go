package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type OrderRenewal struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	RenewalNo      string    `xorm:"not null comment('续费单号') VARCHAR(15)"`
	BuyerId        int       `xorm:"not null comment('买家') INT(11)"`
	SellerId       int       `xorm:"not null comment('卖家') INT(11)"`
	OrderDetailId  int       `xorm:"comment('续费的订单Id') INT(11)"`
	OrderNo        string    `xorm:"not null comment('续费单号') VARCHAR(15)"`
	ProdId         int       `xorm:"not null comment('续费商品') INT(11)"`
	ProdType       string    `xorm:"not null comment('商品类型') VARCHAR(20)"`
	RentDays       int       `xorm:"not null comment('续费天数') INT(11)"`
	ProdAmount     float64   `xorm:"comment('数量') DOUBLE"`
	Poundage       float64   `xorm:"default 0 comment('续费单据的手续费') DOUBLE"`
	ProdStatus     int       `xorm:"comment('商品状态') INT(11)"`
	ProdName       string    `xorm:"comment('商品名称') VARCHAR(128)"`
	ProdPrice      float64   `xorm:"comment('价格') DOUBLE"`
	PaySuccessTime time.Time `xorm:"comment('续费成功时间') DATETIME"`
	OrderHash      string    `xorm:"comment('订单哈希') VARCHAR(255)"`
	CreateTime     time.Time `xorm:"DATETIME"`
	UpdateTime     time.Time `xorm:"DATETIME"`
}

func init() {
	orm.RegisterModel(new(OrderRenewal))
}

func GetOrderRenewalList(name string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var OrderRenewalList []*OrderRenewal
	var ResultData Result
	_, err := o.QueryTable("order_renewal").Limit(limit, from).All(&OrderRenewalList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetOrderRenewalListErr
		logs.Error("GetOrderRenewalList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	total, _ := o.QueryTable("order_renewal").Count()
	data := make(map[string]interface{})
	data["items"] = OrderRenewalList
	data["total"] = total
	ResultData.Code = utils.Success
	ResultData.Data = data
	return ResultData
}

func AddOrderRenewal(host *OrderRenewal) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(host)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddOrderRenewalErr
		logs.Error("AddOrderRenewal failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeleteOrderRenewal(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&OrderRenewal{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteOrderRenewalErr
		logs.Error("DeleteOrderRenewal failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
