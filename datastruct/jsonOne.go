package datastruct

// mongo -----------------------------------
// 订单
type JsonOne struct {
	Ba         string `json:"ba"`
	Ip         string `json:"ip"`
	Ml         string `json:"m1"`
	Mo         string `json:"mo"`
	Nm         string `json:"nm"`
	P          string `json:"p"`
	Seg        string `json:"seg"`
	Custom     string `json:"seg.custom"`
	Appv       string `json:"seg.custom.appv"`
	CategoryId string `json:"seg.custom.categoryId"`
	Eplatform  string `json:"seg.custom.eplatform"`
	From       string `json:"seg.custom.from"`
	Ips        string `json:"seg.custom.ip"`
	ItemId     string `json:"seg.custom.itemId"`
	Itemid     string `json:"seg.custom.itemid"`
	Orderid    string `json:"seg.custom.orderid"`
	Kw    	   string `json:"seg.custom.kw"`
	Uid        string `json:"seg.custom.uid"`
	M1s        string `json:"seg.custom.m1"`
	Match      int 	  `json:"seg.custom.match"`
	Pid        string `json:"seg.custom.pid"`
	Position   string `json:"seg.custom.position"`
	Qd         string `json:"seg.custom.qd"`
	Tm         int    `json:"tm"`
	V          string `json:"v"`
}
