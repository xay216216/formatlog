package models

import (
	"github.com/astaxie/beego"
)

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}

