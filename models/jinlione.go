package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)



type JinLiOne struct {
	Id         int
	DateUrl    string `orm:"size(32)"`
	Ba         string `orm:"size(64)"`
	Ip         string
	M1         string
	Mo         string
	Nm         string
	P          string
	Appv       string
	CategoryId string
	ClickType  string
	Eplatform  string
	From       string
	Ips        string
	Itemid     string
	ItemIdTwo  string
	OrderId    string
	Kw         string
	Uid        string
	Match      string
	Pid        string
	Position   string
	Name       string
	Qd         string
	Tm         string
	V          string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func init() {
	orm.RegisterModel(new(JinLiOne))
}

func InsertJinLiOne(dateUrl, ba, ip, m1, mo, nm, p, appv, categoryId, clickType, eplatform, from, ips, itemid, itemId, orderid, kw, uid, match, pid, position, name, qd, tm, v string) (int64, error) {
	o := orm.NewOrm()

	M := new(JinLiOne)
	M.DateUrl = dateUrl
	M.Ba = ba
	M.Ip = ip
	M.M1 = m1
	M.Mo = mo
	M.Nm = nm
	M.P = p
	M.Appv = appv
	M.CategoryId = categoryId
	M.ClickType = clickType
	M.Eplatform = eplatform
	M.From = from
	M.Ips = ips
	M.Itemid = itemid
	M.ItemIdTwo = itemId
	M.OrderId = orderid
	M.Kw = kw
	M.Uid = uid
	M.Match = match
	M.Pid = pid
	M.Position = position
	M.Name = name
	M.Qd = qd
	M.Tm = tm
	M.V = v
	M.CreatedAt = time.Now()
	M.UpdatedAt = time.Now()
	logs.Info("M1=", m1)
	return o.Insert(M)
}
