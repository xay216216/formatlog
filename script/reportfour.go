package script

import (
	"formatlog/ebeasy"
	"github.com/astaxie/beego"
	"os"
	"strconv"
	_ "strconv"
	"strings"
)

var (
	reportDataFour    map[string]string
	reportDataFourTwo map[string]int
	YxllPv            map[string]string
	YxllPvM           map[string]string
	Pid               map[string]string
	PIdDjxqM          map[string]string
	PIdDstzM          map[string]string
	reportFour        = "reportfour.xls"
	goodsIdString     string
	dateUrlFour       string
)

// 执行脚本
func ReportFourTask() {
	log.Info("开始执行脚本---ReportFourTask")
	//获取数据
	for idx, args := range os.Args {
		log.Info("参数"+strconv.Itoa(idx)+":", args)
		if 1 == idx {
			goodsIdString = args //运营给短的id--对应日志里的itemid--对应goods表里的goods_id
		}
		if 2 == idx {
			dateUrlFour = args
		}
	}
	if len(dateUrlFour) == 0 {
		formatNow = now.AddDate(0, 0, -1)
		dateUrlFour = formatNow.Format("20060102")
	}
	url = logReadUrl + dateUrlFour
	log.Info("logUrl:", url)
	ReadCountLog(url)
	storePid()
	countContentTitle := "日期" + "\t" + "商品ID" + "\t" + "入口" + "\t" + "商品名称" + "\t" + "商品类别" + "\t" + "商品价格" + "\t" + "折扣" + "\t" + "券额度" + "\t" + "返利额度" + "\t" + "商品来源" + "\t" + "图片链接" + "\t" + "后台标签" + "\t" + "有效浏览pv" + "\t" + "有效浏览uv" + "\t" + "点击进详情pv" + "\t" + "点击进详情uv" + "\t" + "电商跳转pv" + "\t" + "电商跳转uv" + "\t" + "电商订单数"
	WriteLog(reportFour, countContentTitle)
	if len(goodsIdString) > 0 {
		goodsIdArray := strings.Split(goodsIdString, ",")
	Loop:
		for _, goodsId := range goodsIdArray {
			//根据mysql获取信息
			goodsId = "100000 266737"
			goodsInfo, err := ebeasy.GetInfoByItemIid(goodsId)
			//fmt.Printf("v1 type:%T\n", goodsInfo)
			//beego.Info("UserInfo:", goodsInfo)
			if err != nil {
				beego.Info("数据库操作错误:", err)
				os.Exit(2)
			}
			if len(goodsInfo) == 0 {
				continue Loop
			}
			for _, value := range goodsInfo {
				reportDataFour["Goods_id"] = strconv.Itoa(value.Goods_id)
				reportDataFour["Item_iid"] = value.Item_iid
				reportDataFour["Name"] = value.Name
				reportDataFour["Eplatform"] = strconv.Itoa(value.Eplatform)
				reportDataFour["Label_price"] = value.Label_price
				reportDataFour["Zk_price"] = value.Zk_price
				reportDataFour["Origin_coupon_price"] = value.Origin_coupon_price
				reportDataFour["Origin_commission_price"] = value.Origin_commission_price
				reportDataFour["Srcsite"] = value.Srcsite
				reportDataFour["List_img"] = value.List_img
				reportDataFour["Ht_biaoqian"] = "标签无"
			}
			//获取日志信息
			reportFourCount(goodsId)
			writeReportFourXls()
		}
	} else {
		os.Exit(2)
	}
}

func storePid() {
	reportDataFour = make(map[string]string)
	reportDataFourTwo = make(map[string]int)
	YxllPv = make(map[string]string)
	YxllPvM = make(map[string]string)
	Pid = make(map[string]string)
	PIdDjxqM = make(map[string]string)
	PIdDstzM = make(map[string]string)
	for _, value := range logDataOne {
		itemidString := value.ITEMID
		chrstr := strings.Split(itemidString, ",")
		if len(chrstr) != 0 {
			for _, va := range chrstr {
				if len(va) != 0 {
					YxllPv[value.PID+"_"+va] = value.PID + "_" + va
				}
			}
		}
	}
}

func reportFourCount(goodsId string) {

	for key, value := range logDataOne {
		//商品浏览数
		if value.NM == "ZHUANQU_WHEN_SCAN_POINT" || value.NM == "FIRST_PAGE_WHEN_SCAN_POINT" {
			dataKey := value.PID + "_" + goodsId
			_, ok := YxllPv[dataKey]
			if ok {
				reportDataFourTwo["YxllPv"] = reportDataFourTwo["YxllPv"] + 1
				YxllPvM[value.PID+"_"+goodsId+"_"+value.M1] = value.PID + "_" + goodsId + "_" + value.M1
				_, okt := YxllPvM[value.PID+"_"+goodsId+"_"+value.M1]
				if !okt {
					reportDataFourTwo["YxllUv"] = reportDataFourTwo["YxllUv"] + 1
				}
			}
		}
		if value.NM == "EVENT_HOME_PV" || value.NM == "EVENT_CATEGORY_LIST" || value.NM == "EVENT_ZHUANGQU_PV" {
			dataKey := value.PID + "_" + goodsId
			_, ok := YxllPv[dataKey]
			if ok {
				Pid[value.PID] = value.PID
			}
		}
		//点击详情
		if value.NM == "EVENT_GOODS_DETAIL_PV" {
			_, ok := Pid[value.PID]
			if ok {
				reportDataFourTwo["DjxqPv"] = reportDataFourTwo["DjxqPv"] + 1
				_, okTwo := PIdDjxqM[value.M1]
				if !okTwo {
					reportDataFourTwo["DjxqUv"] = reportDataFourTwo["DjxqUv"] + 1
				}
				PIdDjxqM[value.M1] = value.NM + "_" + value.M1
			}
		}

		//电商跳转
		if value.NM == "EVENT_JUMP_GET_COUPON" {
			_, ok := Pid[value.PID]
			if ok {
				reportDataFourTwo["DstzPv"] = reportDataFourTwo["DstzPv"] + 1
				_, okTwo := PIdDstzM[value.M1]
				if !okTwo {
					reportDataFourTwo["DstzUv"] = reportDataFourTwo["DstzUv"] + 1
				}
				PIdDstzM[value.M1] = value.NM + "_" + value.M1
			}
		}

		//电商订单数
		if value.NM == "EVENT_USER_ORDER" {
			_, ok := Pid[value.PID]
			if ok {
				reportDataFourTwo["DsOrderNum"] = reportDataFourTwo["DsOrderNum"] + 1
			}
		}
		_ = key
	}
}

func writeReportFourXls() {

	countContent = dateUrlFour + "\t" + reportDataFour["Item_iid"] + "\t" + "无" + "\t" + reportDataFour["Name"] + "\t" + reportDataFour["Eplatform"] + "\t" + reportDataFour["Label_price"] + "\t" +reportDataFour["Zk_price"] + "\t" + reportDataFour["Origin_coupon_price"] + "\t" + reportDataFour["Origin_commission_price"] + "\t" + reportDataFour["Srcsite"] + "\t" + reportDataFour["List_img"]  + "\t" + reportDataFour["Ht_biaoqian"] + "\t" + strconv.Itoa(reportDataFourTwo["YxllPv"]) + "\t" + strconv.Itoa(reportDataFourTwo["YxllUv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DjxqPv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DjxqUv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DstzPv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DstzUv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DsOrderNum"])
	WriteLog(reportFour, countContent)

}
