package models

import (
	"formatlog/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func init() {
	myConfig := new(conf.Config)
	myConfig.InitConfig("./conf/app.conf")
	dbhost := myConfig.Read("mysql", "dbhost")
	dbport := myConfig.Read("mysql", "dbport")
	dbuser := myConfig.Read("mysql", "dbuser")
	dbpassword := myConfig.Read("mysql", "dbpassword")
	dbname := myConfig.Read("mysql", "dbname")
	logs.Info("dbname=", dbname)
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", dsn)
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}

