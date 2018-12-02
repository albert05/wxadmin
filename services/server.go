package services

import (
	"github.com/astaxie/beego"
	"wxadmin/util/mysql"
	"fmt"
)

func init() {
	// init mysql
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
	mysql.Init(dsn)
}
