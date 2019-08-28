package script

import (
	"fmt"
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
	itemIidString     string
	dateUrlFour       string
)

// 执行脚本
func ReportFourTask() {
	eplatformArray := map[string]string{
		"1": "淘宝",
		"2": "天猫",
		"3": "京东",
		"4": "拼多多",
	}
	log.Info("开始执行脚本---ReportFourTask")
	//获取数据
	for idx, args := range os.Args {
		log.Info("参数"+strconv.Itoa(idx)+":", args)
		if 1 == idx {
			itemIidString = args //运营给短的id--对应日志里的itemid--对应goods表里的goods_id
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
	countContentTitle := "日期" + "\t" + "itermID" + "\t" + "商品ID" + "\t" + "入口" + "\t" + "商品名称" + "\t" + "商品类别" + "\t" + "商品价格" + "\t" + "折扣" + "\t" + "券额度" + "\t" + "返利额度" + "\t" + "商品来源" + "\t" + "图片链接" + "\t" + "后台标签" + "\t" + "有效浏览pv" + "\t" + "有效浏览uv" + "\t" + "点击进详情pv" + "\t" + "点击进详情uv" + "\t" + "电商跳转pv" + "\t" + "电商跳转uv" + "\t" + "电商订单数"
	WriteLog(reportFour, countContentTitle)
	if len(itemIidString) > 0 {
		itemIidArray := strings.Split(itemIidString, ",")
	Loop:
		for _, itemIid := range itemIidArray {
			//根据mysql获取信息
			categoryName, goodsInfo, err := ebeasy.GetInfoByItemIid(itemIid)
			beego.Info("itemIid:", itemIid)
			if err != nil {
				beego.Info("数据库操作错误:", err)
				os.Exit(2)
			}
			if len(goodsInfo) == 0 {
				continue Loop
			}
			for key, value := range goodsInfo {
				reportDataFour = make(map[string]string)
				reportDataFour["Goods_id"] = value.Goods_id
				reportDataFour["Item_iid"] = value.Item_iid
				reportDataFour["Name"] = value.Name
				reportDataFour["Eplatform"] = eplatformArray[value.Eplatform]
				reportDataFour["Label_price"] = value.Label_price
				reportDataFour["Zk_price"] = value.Zk_price
				reportDataFour["Back_cate_ids"] = categoryName[key]
				//v1, _ := strconv.ParseFloat(value.Label_price, 64)
				//Discount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value.Discount), 64)   strconv.ParseFloat(fmt.Sprintf("%.2f", diffPrice), 64)
				reportDataFour["Discount"] = value.Discount                       //折扣
				reportDataFour["Origin_coupon_price"] = value.Origin_coupon_price //券额度
				Pay_price, _ := strconv.ParseFloat(value.Pay_price, 64)
				Label_price, _ := strconv.ParseFloat(value.Label_price, 64)
				Discount, _ := strconv.ParseFloat(value.Discount, 64)
				diffPrice := Pay_price - Label_price*Discount
				reportDataFour["Fan_li_price"] = fmt.Sprintf("%.2f", diffPrice)
				reportDataFour["Srcsite"] = value.Srcsite
				reportDataFour["List_img"] = value.List_img
				reportDataFour["Ht_biaoqian"] = "标签无"

				//获取日志信息
				reportFourCount(value.Goods_id)
				//beego.Info("goodsId:", value.Goods_id)
				writeReportFourXls()
			}
		}
	} else {
		beego.Info("please input id:")
		os.Exit(2)
	}
	log.Info("脚本执行完毕---ReportFourTask")
	os.Exit(2)
}

func storePid() {
	YxllPv = make(map[string]string)
	Pid = make(map[string]string)
	for _, value := range logDataOne {
		if value.NM == "ZHUANQU_WHEN_SCAN_POINT" || value.NM == "FIRST_PAGE_WHEN_SCAN_POINT" {
			itemidString := value.ITEMID
			chrstr := strings.Split(itemidString, ",")
			if len(chrstr) != 0 {
				for _, va := range chrstr {
					if len(va) != 0 {
						YxllPv[value.PID+"_"+va] = value.PID + "_" + va
					}
					//log.Info("YxllPv::",value.PID + "_" + va)
				}
			}
		}
		if value.NM == "EVENT_HOME_PV" || value.NM == "EVENT_CATEGORY_LIST" || value.NM == "EVENT_ZHUANGQU_PV" {
			Pid[value.PID] = value.PID
		}
	}
}

func reportFourCount(goodsId string) {

	reportDataFourTwo = make(map[string]int) //map清空
	YxllPvM = make(map[string]string)
	PIdDjxqM = make(map[string]string)
	PIdDstzM = make(map[string]string)
	for key, value := range logDataOne {
		dataKey := value.PID + "_" + goodsId
		//商品浏览数
		if value.NM == "ZHUANQU_WHEN_SCAN_POINT" || value.NM == "FIRST_PAGE_WHEN_SCAN_POINT" {
			_, ok := YxllPv[dataKey]
			if ok {
				itemidString1 := value.ITEMID
				chrstr1 := strings.Split(itemidString1, ",")
				if len(chrstr1) != 0 {
					for _, va := range chrstr1 {
						if va == goodsId {
							reportDataFourTwo["YxllPv"] = reportDataFourTwo["YxllPv"] + 1
							_, okt := YxllPvM[value.PID+"_"+goodsId+"_"+value.M1]
							if !okt {
								reportDataFourTwo["YxllUv"] = reportDataFourTwo["YxllUv"] + 1
							}
							YxllPvM[value.PID+"_"+goodsId+"_"+value.M1] = value.PID + "_" + goodsId + "_" + value.M1
						}
					}
				}
			}
		}
		/*if value.NM == "EVENT_HOME_PV" || value.NM == "EVENT_CATEGORY_LIST" || value.NM == "EVENT_ZHUANGQU_PV" {
			_, ok := YxllPv[dataKey]
			if ok {
				itemidString2 := value.ITEMID
				chrstr2 := strings.Split(itemidString2, ",")
				log.Info("itemidString::",itemidString2)
				if len(chrstr2) != 0 {
					for _, va := range chrstr2 {
						log.Info("goodsId::",goodsId)
						log.Info("va::",va)
						if va == goodsId {
							log.Info("1111::",va)
							Pid[value.PID] = value.PID
							log.Info("Pid::",value.PID + "_" + va+ "_" + goodsId)
						}
					}
				}
			}
		}*/
		//点击详情
		if value.NM == "EVENT_GOODS_DETAIL_PV" {
			_, ok := Pid[value.PID]
			if ok {
				itemidString3 := value.ITEMID
				chrstr3 := strings.Split(itemidString3, ",")
				if len(chrstr3) != 0 {
					for _, va := range chrstr3 {
						if va == goodsId {
							reportDataFourTwo["DjxqPv"] = reportDataFourTwo["DjxqPv"] + 1
							_, okTwo := PIdDjxqM[value.M1]
							if !okTwo {
								reportDataFourTwo["DjxqUv"] = reportDataFourTwo["DjxqUv"] + 1
							}
							PIdDjxqM[value.M1] = value.NM + "_" + value.M1
						}
					}
				}
			}
		}

		//电商跳转
		if value.NM == "EVENT_JUMP_GET_COUPON" {
			_, ok := Pid[value.PID]
			if ok {
				itemidString3 := value.ITEMID
				chrstr3 := strings.Split(itemidString3, ",")
				if len(chrstr3) != 0 {
					for _, va := range chrstr3 {
						if va == goodsId {
							reportDataFourTwo["DstzPv"] = reportDataFourTwo["DstzPv"] + 1
							_, okTwo := PIdDstzM[value.M1]
							if !okTwo {
								reportDataFourTwo["DstzUv"] = reportDataFourTwo["DstzUv"] + 1
							}
							PIdDstzM[value.M1] = value.NM + "_" + value.M1
						}
					}
				}
			}
		}

		//电商订单数
		if value.NM == "EVENT_USER_ORDER" {
			_, ok := Pid[value.PID]
			if ok {
				itemidString3 := value.ITEMID
				chrstr3 := strings.Split(itemidString3, ",")
				if len(chrstr3) != 0 {
					for _, va := range chrstr3 {
						if va == goodsId {
							reportDataFourTwo["DsOrderNum"] = reportDataFourTwo["DsOrderNum"] + 1
						}
					}
				}
			}
		}
		_ = key
	}
}

func writeReportFourXls() {

	countContent = dateUrlFour + "\t" + reportDataFour["Goods_id"] + "\t" + reportDataFour["Item_iid"] + ",\t " + "无" + "\t" + reportDataFour["Name"] + "\t" + reportDataFour["Back_cate_ids"] + "\t" + reportDataFour["Zk_price"] + "\t " + reportDataFour["Discount"] + "\t " + reportDataFour["Origin_coupon_price"] + "\t " + reportDataFour["Fan_li_price"] + "\t" + reportDataFour["Eplatform"] + "\t" + reportDataFour["List_img"] + "\t" + reportDataFour["Ht_biaoqian"] + "\t" + strconv.Itoa(reportDataFourTwo["YxllPv"]) + "\t" + strconv.Itoa(reportDataFourTwo["YxllUv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DjxqPv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DjxqUv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DstzPv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DstzUv"]) + "\t" + strconv.Itoa(reportDataFourTwo["DsOrderNum"])
	WriteLog(reportFour, countContent)
}
