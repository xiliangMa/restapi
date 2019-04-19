package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type Host struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	Rid           string    `xorm:"comment('主机在rancher服务上的Id') VARCHAR(255)"`
	Name          string    `xorm:"comment('主机名称') VARCHAR(128)"`
	RancherId     int       `xorm:"INT(11)"`
	ClusterId     int       `xorm:"comment('所属集群主键') INT(11)"`
	OwnerId       int       `xorm:"not null comment('主机所有者') INT(11)"`
	HostIp        string    `xorm:"not null comment('主机IP地址') unique VARCHAR(63)"`
	Host          string    `xorm:"comment('预留,暂时未用') VARCHAR(63)"`
	CpuFrequency  float64   `xorm:"comment('频率(GHZ)') DOUBLE(11,1)"`
	CpuKernel     float64   `xorm:"comment('核数') DOUBLE(11)"`
	CpuKernelUsed float64   `xorm:"default 0 comment('已用核数') DOUBLE(11)"`
	CpuKernelLock float64   `xorm:"default 0 comment('订单占用cpu核数(会自动释放15分)') DOUBLE(11)"`
	Mem           float64   `xorm:"comment('内存(G)') DOUBLE"`
	MemUsed       float64   `xorm:"default 0 comment('已用内存') DOUBLE"`
	MemLock       float64   `xorm:"default 0 comment('内存占用') DOUBLE"`
	Disk          float64   `xorm:"comment('硬盘(G)') DOUBLE"`
	DiskUsed      float64   `xorm:"default 0 comment('已用硬盘') DOUBLE"`
	DiskLock      float64   `xorm:"default 0 comment('硬盘占用') DOUBLE"`
	Network       float64   `xorm:"comment('宽带(M)') DOUBLE"`
	NetworkUsed   float64   `xorm:"default 0 comment('已用宽带') DOUBLE"`
	NetworkLock   float64   `xorm:"default 0 comment('宽带占用') DOUBLE"`
	Pods          int       `xorm:"default 0 comment('Pod总量') INT(11)"`
	PodsUsed      int       `xorm:"default 0 comment('已使用Pod数量') INT(11)"`
	State         string    `xorm:"not null default '0' comment('主机状态') VARCHAR(50)"`
	HostInfo      string    `xorm:"comment('同步主机的信息') TEXT"`
	Command       string    `xorm:"comment('生成主机的命令') VARCHAR(5000)"`
	TotalCompute  float64   `xorm:"comment('算力') DOUBLE"`
	UsedCompute   float64   `xorm:"comment('算力') DOUBLE"`
	Deleted       int       `xorm:"default 0 comment('是否删除') INT(11)"`
	BeginTime     time.Time `xorm:"comment('主机有效开始时间') DATETIME"`
	Etcd          int       `xorm:"TINYINT(1)"`
	ControlPlane  int       `xorm:"TINYINT(1)"`
	Worker        int       `xorm:"TINYINT(1)"`
	EndTime       time.Time `xorm:"comment('主机有效结束时间') DATETIME"`
	UpdateTime    time.Time `xorm:"DATETIME"`
	CreateTime    time.Time `xorm:"DATETIME"`
	SyncTime      time.Time `xorm:"comment('同步时间') DATETIME"`
	CallbackUrl   string    `xorm:"comment('主机回调地址') VARCHAR(50)"`
	PublicAddress int       `xorm:"not null default 1 TINYINT(1)"`
}

func init() {
	orm.RegisterModel(new(Host))
}

func GetHostList(name, ip string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var HostList []*Host
	var ResultData Result
	_, err := o.QueryTable("host").Filter("name__icontains", name).Filter("host_ip__icontains", ip).Limit(limit, from).All(&HostList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetHostListErr
		logs.Error("GetHostList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	total, _ := o.QueryTable("host").Count()
	data := make(map[string]interface{})
	data["items"] = HostList
	data["total"] = total
	ResultData.Code = utils.Success
	ResultData.Data = data
	return ResultData
}

func AddHost(host *Host) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(host)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddHostErr
		logs.Error("AddHost failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeleteHost(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&Host{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteHostErr
		logs.Error("DeleteHost failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
