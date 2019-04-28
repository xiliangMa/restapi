package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type OrderMaster struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	OrderNo        string    `xorm:"not null comment('订单号') VARCHAR(15)"`
	BuyerId        int       `xorm:"comment('购买人') INT(11)"`
	BuyerAccount   string    `xorm:"VARCHAR(128)"`
	BuyerName      string    `xorm:"VARCHAR(100)"`
	SellerId       int       `xorm:"default 0 comment('卖家用户') INT(11)"`
	SellerAccount  string    `xorm:"VARCHAR(128)"`
	SellerName     string    `xorm:"VARCHAR(100)"`
	ProdType       string    `xorm:"not null comment('订单类型 host/app') VARCHAR(50)"`
	OrderAmount    float64   `xorm:"not null default 0 comment('订单金额') DOUBLE"`
	Poundage       float64   `xorm:"comment('手续费') DOUBLE"`
	OrderStatus    int       `xorm:"not null default 0 comment('状态') INT(11)"`
	OrderHash      string    `xorm:"comment('订单哈希') VARCHAR(128)"`
	PaySuccessTime time.Time `xorm:"comment('支付成功时间') DATETIME"`
	OriginalOrder  string    `xorm:"VARCHAR(15)"`
	CreateTime     time.Time `xorm:"not null DATETIME"`
	UpdateTime     time.Time `xorm:"not null DATETIME"`
}


func init() {
	orm.RegisterModel(new(OrderMaster))
}

func GetOrderMasterList(name string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var OrderMasterList []*OrderMaster
	var ResultData Result
	_, err := o.QueryTable("order_master").Limit(limit, from).All(&OrderMasterList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetOrderMasterListErr
		logs.Error("GetOrderMasterList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	total, _ := o.QueryTable("order_master").Count()
	data := make(map[string]interface{})
	data["items"] = OrderMasterList
	data["total"] = total
	ResultData.Code = utils.Success
	ResultData.Data = data
	return ResultData
}

func AddOrderMaster(host *OrderMaster) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(host)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddOrderMasterErr
		logs.Error("AddOrderMaster failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeleteOrderMaster(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&OrderMaster{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteOrderMasterErr
		logs.Error("DeleteOrderMaster failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
