package ebeasy

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Goods_info struct {
	Goods_id                string
	Item_iid                string
	Name                    string
	Eplatform               string
	Label_price             string
	Zk_price                string
	Origin_coupon_price     string
	Origin_commission_price string
	Pay_price               string
	Discount                string
	Srcsite                 string
	List_img                string
	Front_cate_ids          string
	categoryName            string
}

type Category_info struct {
	Cate_id string
	Name    string
}

var (
	categoryName map[int]string
)

func GetInfoByItemIid(itemIid string) (categoryName map[int]string, u []Goods_info, err error) {

	categoryName = make(map[int]string)

	o := orm.NewOrm()
	o.Using("default")
	var goods_info []Goods_info
	var category_info []Category_info
	_, err = o.Raw("select goods_id,item_iid,name,eplatform,label_price,zk_price,origin_coupon_price,origin_commission_price,pay_price,discount,srcsite,list_img,front_cate_ids from eb_goods where eb_goods.item_iid=? ", itemIid).QueryRows(&goods_info)
	if len(goods_info) > 0 {
		beego.Info("goods_info:", goods_info)
		for key, value := range goods_info {
			ids := value.Front_cate_ids
			chrstr := strings.Split(ids, ",")
			for _, va := range chrstr {
				if len(va) > 0 {
					fmt.Printf("va type is:%T\n", va)
					beego.Info("va:", va)
					_, err = o.Raw("select cate_id,name from eb_category where eb_category.cate_id = ? ", va).QueryRows(&category_info)
					if len(category_info) > 0 {
						for _, valuea := range category_info {
							//categoryName = valuea.Name + "," + categoryName
							categoryName[key] = valuea.Name + "," + categoryName[key]
							beego.Info("categoryName:", categoryName)
						}
					}
				}
			}
			categoryName[key] = strings.TrimRight(categoryName[key], ",")
		}
	}
	//categoryName = strings.TrimRight(categoryName, ",")
	return categoryName, goods_info, err
}
