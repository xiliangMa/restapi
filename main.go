package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/xiliangMa/restapi/routers"
)

func initDD() {

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
}

func main() {
	initDD()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
