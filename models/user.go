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

type ElementAdmin struct {
	Roles  []string `json:"roles"`
	Avatar string   `json:"avatar"`
	Name   string   `json:"name"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserInfo(token string) Result {
	claims, _ := utils.ParseToken(token, []byte("Hello World！This is jwt test demo!"))
	name := utils.GetNameFromClaims("jti", claims)
	var ResultData Result
	userInfo := ElementAdmin{
		Roles:  []string{"admin"},
		Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:   name,
	}
	ResultData.Code = utils.Success
	ResultData.Data = userInfo
	return ResultData
}

func GetUserList(mobile, email string, from, limit int) Result {
	o := orm.NewOrm()
	o.Using("default")
	var UserList []*User
	var ResultData Result
	_, err := o.QueryTable("user").Filter("mobile__icontains", mobile).Filter("email__icontains", email).Limit(limit, from).All(&UserList)
	if err != nil {
		ResultData.Message = err.Error()
		ResultData.Code = utils.GetUserListErr
		logs.Error("GetUserList failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}

	total, _ := o.QueryTable("user").Count()
	data := make(map[string]interface{})
	data["items"] = UserList
	data["total"] = total
	ResultData.Code = utils.Success
	ResultData.Data = data
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
		logs.Error("AddUser failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
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
		logs.Error("DeleteUser failed, code: %d, err: %s", ResultData.Code, ResultData.Message)
		return ResultData
	}
	ResultData.Code = utils.Success
	return ResultData
}
