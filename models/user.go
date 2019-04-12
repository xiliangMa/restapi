package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xiliangMa/restapi/utils"
	"time"
)

type User struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	AccountAddress string    `xorm:"not null default '1' comment('钱包地址') VARCHAR(255)"`
	LoginName      string    `xorm:"comment('登录名') VARCHAR(50)"`
	Password       string    `xorm:"comment('密码') VARCHAR(255)"`
	Salt           string    `xorm:"comment('密码加密随即字符串') VARCHAR(10)"`
	RoleId         int       `xorm:"not null default 0 comment('用户类型(1:前台注册的用户)') INT(11)"`
	NickName       string    `xorm:"default '' comment('界面显示昵称') VARCHAR(255)"`
	Mobile         string    `xorm:"comment('手机号') VARCHAR(50)"`
	Email          string    `xorm:"comment('邮箱') VARCHAR(50)"`
	RegisterType   string    `xorm:"comment('注册类型，手机还是邮箱') VARCHAR(255)"`
	CaptchaMode    string    `xorm:"comment('验证方式，手机或邮箱') VARCHAR(255)"`
	CreateTime     time.Time `xorm:"TIMESTAMP"`
	UpdateTime     time.Time `xorm:"TIMESTAMP"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserList(mobile, email string, page, number int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var UserList []*User
	var ResultData Result
	_, err := o.QueryTable("user").Filter("mobile__icontains", mobile).Filter("email__icontains", email).Limit(number, page).All(&UserList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetUserListErr
		logs.Error("GetUserList failed, code: %d, err: %s", ResultData.Code,  ResultData.Message )
		return ResultData
	}

	ResultData.Code = utils.Success
	ResultData.Data = UserList
	return ResultData
}

func AddUser(User *User) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	id, err := o.Insert(User)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.AddUserErr
		logs.Error("AddUser failed, code: %d, err: %s", ResultData.Code,  ResultData.Message )
		return ResultData
	}
	ResultData.Code = utils.Success
	ResultData.Data = id
	return ResultData
}

func DeleteUser(id int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var ResultData Result
	_, err := o.Delete(&User{Id: id})
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.DeleteUserErr
		logs.Error("DeleteUser failed, code: %d, err: %s", ResultData.Code,  ResultData.Message )
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
