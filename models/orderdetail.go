package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type OrderDetail struct {
	Id         int       `xorm:"not null pk autoincr comment('订单号') INT(11)"`
	OrderNo    string    `xorm:"not null VARCHAR(15)"`
	ProdId     int       `xorm:"not null comment('商品编号') INT(11)"`
	ProdType   string    `xorm:"not null comment('商品类型') VARCHAR(20)"`
	RentDays   int       `xorm:"not null default 0 comment('租用天数') INT(11)"`
	ProdAmount float64   `xorm:"default 0 comment('金额') DOUBLE"`
	ProdName   string    `xorm:"comment('商品名称') VARCHAR(128)"`
	ProdPrice  float64   `xorm:"comment('价格') DOUBLE"`
	ProdInfo   string    `xorm:"TEXT"`
	Rid        string    `xorm:"comment('商品在rancher服务上的Id.') VARCHAR(128)"`
	DeployTime time.Time `xorm:"DATETIME"`
	BeginTime  time.Time `xorm:"comment('租用开始时间') DATETIME"`
	EndTime    time.Time `xorm:"comment('租用到期时间') DATETIME"`
}


func init() {
	orm.RegisterModel(new(OrderDetail))
}

func GetOrderDetailList(name string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var OrderDetailList []*OrderDetail
	var ResultData Result
	_, err := o.QueryTable("order_detail").Limit(limit, from).All(&OrderDetailList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetOrderDetailListErr
		logs.Error("GetOrderDetailList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	total, _ := o.QueryTable("order_detail").Count()
	data := make(map[string]interface{})
	data["items"] = OrderDetailList
	data["total"] = total
	ResultData.Code = utils.Success
	ResultData.Data = data
	return ResultData
}

func AddOrderDetail(host *OrderDetail) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(host)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddOrderDetailErr
		logs.Error("AddOrderDetail failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeleteOrderDetail(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&OrderDetail{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteOrderDetailErr
		logs.Error("DeleteOrderDetail failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
