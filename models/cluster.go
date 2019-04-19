package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type Cluster struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	Rid           string    `xorm:"comment('集群在rancher服务上的Id') VARCHAR(255)"`
	RancherId     int       `xorm:"not null comment('rancher表主键') INT(11)"`
	OwnerId       int       `xorm:"not null comment('集群所有者') INT(11)"`
	Name          string    `xorm:"comment('集群名称') VARCHAR(128)"`
	NetworkType   string    `xorm:"default '' comment('网络类型：inner:内网；outer：外网') VARCHAR(10)"`
	Profit        float64   `xorm:"default 0 comment('收益,每次购买资源后 将订单金额累加到改字段') DOUBLE"`
	CpuKernel     float64   `xorm:"comment('核数') DOUBLE(11,2)"`
	CpuKernelUsed float64   `xorm:"not null default 0.00 comment('已用核数') DOUBLE(11,2)"`
	CpuKernelLock float64   `xorm:"not null default 0.00 comment('订单占用cpu核数(会自动释放15分)') DOUBLE(11,2)"`
	Mem           float64   `xorm:"comment('内存(G)') DOUBLE"`
	MemUsed       float64   `xorm:"not null default 0 comment('已用内存') DOUBLE"`
	MemLock       float64   `xorm:"not null default 0 comment('内存占用') DOUBLE"`
	Disk          float64   `xorm:"comment('硬盘(G)') DOUBLE"`
	DiskUsed      float64   `xorm:"not null default 0 comment('已用硬盘') DOUBLE"`
	DiskLock      float64   `xorm:"not null default 0 comment('硬盘占用') DOUBLE"`
	Capacity      string    `xorm:"comment('集群的容量') TEXT"`
	Limits        string    `xorm:"comment('集群的限额') TEXT"`
	Requested     string    `xorm:"comment('集群的需求') TEXT"`
	Network       float64   `xorm:"comment('宽带(M)') DOUBLE"`
	NetworkUsed   float64   `xorm:"not null default 0 comment('已用宽带') DOUBLE"`
	NetworkLock   float64   `xorm:"not null default 0 comment('宽带占用') DOUBLE"`
	Pods          int       `xorm:"default 0 comment('Pod总量') INT(11)"`
	PodsUsed      int       `xorm:"default 0 comment('已使用Pod数量') INT(11)"`
	State         string    `xorm:"default '0' comment('集群状态') VARCHAR(50)"`
	ClusterInfo   string    `xorm:"comment('同步主机的信息') TEXT"`
	Command       string    `xorm:"comment('集群中主机命令') VARCHAR(5000)"`
	TotalCompute  float64   `xorm:"default 0 comment('算力') DOUBLE"`
	UsedCompute   float64   `xorm:"default 0 DOUBLE"`
	RentPrice     float64   `xorm:"comment('出租价格') DOUBLE"`
	Deleted       int       `xorm:"not null default 0 comment('是否删除') INT(11)"`
	Online        int       `xorm:"default 0 comment('是否上架') TINYINT(1)"`
	BeginTime     time.Time `xorm:"comment('集群有效开始时间') DATETIME"`
	EndTime       time.Time `xorm:"comment('集群有效结束时间') DATETIME"`
	UpdateTime    time.Time `xorm:"DATETIME"`
	CreateTime    time.Time `xorm:"DATETIME"`
	SyncTime      time.Time `xorm:"comment('同步时间') DATETIME"`
}

func init() {
	orm.RegisterModel(new(Cluster))
}

func GetClusterList(name string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ClusterList []*Cluster
	var ResultData Result

	_, err := o.QueryTable("cluster").Filter("name__icontains", name).Limit(limit, from).All(&ClusterList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetClusterListErr
		logs.Error("GetClusterList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	data := make(map[string]interface{})
	total, _ := o.QueryTable("cluster").Count()
	data["items"] = ClusterList
	data["total"] = total
	ResultData.Code = utils.Success
	ResultData.Data = &data
	return ResultData
}

func AddCluster(cluster *Cluster) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(cluster)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddClusterErr
		logs.Error("AddCluster failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeleteCluster(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&Cluster{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteClusterErr
		logs.Error("DeleteCluster failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
