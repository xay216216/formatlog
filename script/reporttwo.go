package script

import (
	"os"
	"strconv"
	_ "strconv"
	"strings"
)

type ReportTwoBaseData struct {
	PID             string
	GoodsName      string
	GoodsId        string
	GoodsCategory  string
	POSITION       string
	Price          string
	QuanDue        string
	MoneyRate      int
	OnlinePosition string
	OrderMoney     string
}
type ReportTwo struct {
	GoodsBrowPv    int
	GoodsBrowUv    int
	GoodsDetailPv  int
	GoodsDetailUv  int
	GoodsQuanPv    int
	GoodsQuanUv    int
	GoodsOrderPv   int
	GoodsOrderUv   int
}

var (
	reportDataTwo  map[string]ReportTwoBaseData
	GoodsId        map[string]string
	GoodsCategory  map[string]string
	GoodsBrow      map[string]string
	GoodsBrowPv    map[string]int
	GoodsBrowUv    map[string]int
	GoodsDetailTwo map[string]string
	GoodsDetailPv  map[string]int
	GoodsDetailUv  map[string]int
	GoodsQuan      map[string]string
	GoodsQuanPv    map[string]int
	GoodsQuanUv    map[string]int
	GoodsOrder     map[string]string
	GoodsOrderPv   map[string]int
	GoodsOrderUv   map[string]int
	OrderMoneyUv   map[string]int
	reportTwo      = "reporttwo.xls"
)

// 执行脚本
func ReportTwoTask() {
	log.Info("开始执行脚本---ReportTwoTask")
	//获取数据
	formatNow = now.AddDate(0, 0, -1)
	dateUrl = formatNow.Format("20060102")
	url := logReadUrl + dateUrl
	log.Info("logUrl:", url)
	ReadCountLog(url)
	reportTwoCount()
	writeReportTwoXls()
	log.Info("脚本执行完毕---ReportTwoTask")
	os.Exit(2)
}

func reportTwoCount() {
	reportDataTwo = make(map[string]ReportTwoBaseData)
	GoodsId = make(map[string]string)
	GoodsCategory = make(map[string]string)
	GoodsBrow = make(map[string]string)
	GoodsBrowPv = make(map[string]int)
	GoodsBrowUv = make(map[string]int)
	GoodsDetailTwo = make(map[string]string)
	GoodsDetailPv = make(map[string]int)
	GoodsDetailUv = make(map[string]int)
	GoodsQuan = make(map[string]string)
	GoodsQuanPv = make(map[string]int)
	GoodsQuanUv = make(map[string]int)
	GoodsOrder = make(map[string]string)
	GoodsOrderPv = make(map[string]int)
	GoodsOrderUv = make(map[string]int)
	OrderMoneyUv = make(map[string]int)
	for _, value := range logDataOne {
		//app
		if value.NM == "ZHUANQU_WHEN_SCAN_POINT" {
			itemidString := value.ITEMID
			chrstr := strings.Split(itemidString, ",")
			if len(chrstr) != 0 {
				for _, va := range chrstr {
					if len(va) != 0 {
						reportDataTwo[value.PID+"_"+value.POSITION] = ReportTwoBaseData{value.PID,"商品名称",va,"商品类目",value.POSITION,"0.00","0.00",0,"上线位置","付款金额"}
						GoodsBrowPv[value.PID+"_"+value.POSITION] = GoodsBrowPv[value.PID+"_"+value.POSITION] + 1
						_, ok := GoodsBrow[value.PID+"_"+value.POSITION]
						if !ok {
							GoodsBrowUv[value.PID+"_"+value.POSITION] = GoodsBrowUv[value.PID+"_"+value.POSITION] + 1
						}
						GoodsBrow[value.PID+"_"+value.POSITION] = value.PID + "_" + value.POSITION
					}
				}
			}
		}
		//app
		if value.NM == "EVENT_GOODS_DETAIL_PV" {
			GoodsDetailPv[value.PID+"_"+value.POSITION] = GoodsDetailPv[value.PID+"_"+value.POSITION] + 1
			_, ok := GoodsDetailTwo[value.PID+"_"+value.POSITION]
			if !ok {
				GoodsDetailUv[value.PID+"_"+value.POSITION] = GoodsDetailUv[value.PID+"_"+value.POSITION] + 1
			}
			GoodsDetailTwo[value.PID+"_"+value.POSITION] = value.PID + "_" + value.POSITION
		}
		//app
		if value.NM == "EVENT_JUMP_GET_COUPON" {
			GoodsQuanPv[value.PID+"_"+value.POSITION] = GoodsQuanPv[value.PID+"_"+value.POSITION] + 1
			_, ok := GoodsQuan[value.PID+"_"+value.POSITION]
			if !ok {
				GoodsQuanUv[value.PID+"_"+value.POSITION] = GoodsQuanUv[value.PID+"_"+value.POSITION] + 1
			}
			GoodsQuan[value.PID+"_"+value.POSITION] = value.PID + "_" + value.POSITION
		}
		//app
		if value.NM == "EVENT_USER_ORDER" {
			GoodsOrderPv[value.PID+"_"+value.POSITION] = GoodsOrderPv[value.PID+"_"+value.POSITION] + 1
			_, ok := GoodsOrder[value.PID+"_"+value.POSITION]
			if !ok {
				GoodsOrderUv[value.PID+"_"+value.POSITION] = GoodsOrderUv[value.PID+"_"+value.POSITION] + 1
			}
			GoodsOrder[value.PID+"_"+value.POSITION] = value.PID + "_" + value.POSITION
		}

	}
}

func writeReportTwoXls() {

	countContentTitle := "日期" + "\t" + "模块名称" + "\t" + "专区有效浏览商品名称" + "\t" + "专区有效浏览商品id" + "\t" + "专区有效浏览商品类目" + "\t" + "价格" + "\t" + "券额度" + "\t" + "佣金比例" + "\t" + "上线位置" + "\t" + "专区商品有效浏览PV" + "\t" + "专区商品有效浏览UV" + "\t" + "专区有效浏览商品进详情pv" + "\t" + "专区有效浏览商品进详情uv" + "\t" + "专区商品领券pv" + "\t" + "专区商品领券uv" + "\t" + "专区商品下单pv" + "\t" + "专区商品下单uv" + "\t" + "付款金额"
	WriteLog(reportTwo, countContentTitle)

	var MkName string
	for _, value := range reportDataTwo {
		switch value.POSITION {
		case "0":
			MkName = "淘宝白菜"
		case "1":
			MkName = "品牌特价"
		case "2":
			MkName = "大额神券"
		default:
			MkName = "淘宝白菜"
		}

		countContent = dateUrl + "\t" + MkName + "\t" + value.GoodsName + "\t" + value.GoodsId + "\t" + value.GoodsCategory + "\t" + value.Price + "\t" + value.QuanDue + "\t" + strconv.Itoa(value.MoneyRate) + "\t" + value.OnlinePosition + "\t" + strconv.Itoa(GoodsBrowPv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsBrowUv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsDetailPv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsDetailUv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsQuanPv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsQuanUv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsOrderPv[value.PID+"_"+value.POSITION]) + "\t" + strconv.Itoa(GoodsOrderUv[value.PID+"_"+value.POSITION]) + "\t" + value.OrderMoney
		WriteLog(reportTwo, countContent)
	}

}
