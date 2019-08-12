package models

import (
	"formatlog/conf"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type any = interface{}

type BaseModel struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"script_time"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JsonDataOne struct {
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

func (m *JsonDataOne) TableName() string {
	return TableName("jsondataone")
}

func (m *JsonDataOne) SaveJsonDataOne() (int64, error) {
	o := orm.NewOrm()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return o.Insert(m)
}

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
	orm.RegisterModel(new(JsonDataOne))
}

func InsertJsonOne(dateUrl, ba, ip, m1, mo, nm, p, appv, categoryId, clickType, eplatform, from, ips, itemid, itemId, orderid, kw, uid, match, pid, position, name, qd, tm, v string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	M := new(JsonDataOne)
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
