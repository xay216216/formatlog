package ebeasy

import (
	"formatlog/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func init() {
	myConfig := new(conf.Config)
	myConfig.InitConfig("./conf/app.conf")
	dbhost := myConfig.Read("mysqlbaice", "dbhost")
	dbport := myConfig.Read("mysqlbaice", "dbport")
	dbuser := myConfig.Read("mysqlbaice", "dbuser")
	dbpassword := myConfig.Read("mysqlbaice", "dbpassword")
	dbname := myConfig.Read("mysqlbaice", "dbname")
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

