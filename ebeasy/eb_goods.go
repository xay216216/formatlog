package ebeasy

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Goods_info struct {
	Goods_id                int
	Item_iid                string
	Name                    string
	Eplatform                int
	Label_price             string
	Zk_price                string
	Origin_coupon_price     string
	Origin_commission_price string
	Srcsite                 string
	List_img                string
}

func GetInfoByItemIid(goodsId string) (u []Goods_info, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var goods_info []Goods_info
	_, err = o.Raw("select goods_id,item_iid,name,eplatform,label_price, zk_price, origin_coupon_price, origin_commission_price, srcsite, list_img from eb_goods where eb_goods.goods_id=? ", goodsId).QueryRows(&goods_info)

	return goods_info, err
}
