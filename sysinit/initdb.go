package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/xiliangMa/restapi/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func initDB() {

	dbType := beego.AppConfig.String("db::db_type")

	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")

	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")

	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")

	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")

	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")

	//数据库端口
	//dbPort := beego.AppConfig.String(dbType + "::db_port")

	datasource := fmt.Sprintf("%s%s%s%s%s%s%s%s", dbUser, ":", dbPwd, "@tcp(", dbHost, ":3306)/", dbName, "?charset=utf8")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbAlias, "mysql", datasource)

	// local
	//orm.RegisterDataBase("default", "mysql", "root:abc123@tcp(127.0.0.1:3306)/uranus_local?charset=utf8")

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")

	// true: drop table 后再建表
	force := false

	//auto create db
	orm.RunSyncdb("default", force, isDev)
}