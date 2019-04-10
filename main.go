package main

import (
	_ "github.com/xiliangMa/restapi/routers"
	"github.com/xiliangMa/restapi/conf"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego"
)

func initDD() {
	datasource := fmt.Sprintf("%s%s%s%s%s%s%s%s", conf.MysqlUser, ":", conf.MysqlPass, "@tcp(", conf.MysqlUrls, ":3306)/", conf.MysqlDb, "?charset=utf8")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	orm.RegisterDataBase("default", "mysql", datasource)
}

func main() {
	initDD()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
