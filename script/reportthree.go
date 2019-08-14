package script

import (
	"strconv"
	_ "strconv"
	"strings"
)

var (
	url string
)
var (
	reportDataThree      map[string]int
	GoodsScanOne         map[string]string
	GoodsScanTwo         map[string]string
	GoodsDetail          map[string]string
	SearchResult         map[string]string
	SearchResultTwo      map[string]string
	HotUse               map[string]string
	KeyUse               map[string]string
	SearchResultClick    map[string]string
	SearchResultClickTwo map[string]string
	ShopClickUser        map[string]string
	DetailOrder          map[string]string
	DetailUser           map[string]string
	reportThree          = "reportthree.xls"
)

type ReportThree struct {
	AppQiDong         int
	GoodsScan         int
	GoodsClick        int
	GoodsDetail       int
	HotCode           int
	KeyCode           int
	SearchResult      int
	HotUse            int
	KeyUse            int
	SearchResultClick int
	ShopClick         int
	ShopClickUser     int
	DetailOrder       int
	DetailUser        int
}

// 执行脚本
func ReportThreeTask() {
	log.Info("开始执行脚本---ReportThreeTask")
	//获取数据
	formatNow = now.AddDate(0, 0, -1)
	dateUrl = formatNow.Format("20060102")
	url = logReadUrl + dateUrl
	log.Info("logUrl:", url)
	ReadCountLog(url)
	reportThreeCount()
	writeReportThreeXls()
}

func reportThreeCount() {
	reportDataThree = make(map[string]int)
	GoodsScanOne = make(map[string]string)
	GoodsScanTwo = make(map[string]string)
	GoodsDetail = make(map[string]string)
	SearchResult = make(map[string]string)
	SearchResultTwo = make(map[string]string)
	HotUse = make(map[string]string)
	KeyUse = make(map[string]string)
	SearchResultClick = make(map[string]string)
	SearchResultClickTwo = make(map[string]string)
	ShopClickUser = make(map[string]string)
	DetailOrder = make(map[string]string)
	DetailUser = make(map[string]string)
	for key, value := range logDataOne {
		//app
		if value.NM == "USER_START_APK_EVERYTIME" {
			reportDataThree["AppQiDong"] = reportDataThree["AppQiDong"] + 1 //当日启动打开APP次数
		}
		//商品浏览总坑位数
		if value.NM == "ZHUANQU_WHEN_SCAN_POINT" {
			itemidString := value.ITEMID
			chrstr := strings.Split(itemidString, ",")
			if len(chrstr) != 0 {
				for _, va := range chrstr {
					if len(va) != 0 {
						_, ok := GoodsScanOne[value.PID+"_"+va]
						if !ok {
							reportDataThree["GoodsScan"] = reportDataThree["GoodsScan"] + 1 //商品浏览总坑位数
						}
						GoodsScanOne[value.PID+"_"+va] = value.PID + "_" + va
					}
				}
			}
		}
		if value.NM == "FIRST_PAGE_WHEN_SCAN_POINT" {
			itemidStringTwo := value.ITEMID
			chrstrTwo := strings.Split(itemidStringTwo, ",")
			if len(chrstrTwo) != 0 {
				for _, val := range chrstrTwo {
					if len(val) != 0 {
						_, ok := GoodsScanOne[value.PID+"_"+val]
						if !ok {
							reportDataThree["GoodsScan"] = reportDataThree["GoodsScan"] + 1 //商品浏览总坑位数
						}
						GoodsScanOne[value.PID+"_"+val] = value.PID + "_" + val
					}
				}
			}
		}
		//除搜索粘贴板之外的商品点击数--点击访问商品详情页UV
		if value.NM == "EVENT_GOODS_DETAIL_PV" {
			if value.FROM == "category" || value.FROM == "zq" {
				log.Info(key)
				reportDataThree["GoodsClick"] = reportDataThree["GoodsClick"] + 1 //除搜索粘贴板之外的商品点击数
				_, ok := GoodsDetail[value.M1]
				if !ok {
					reportDataThree["GoodsDetail"] = reportDataThree["GoodsDetail"] + 1 //点击访问商品详情页UV
				}
				GoodsDetail[value.M1] = value.NM + "_" + value.M1
			}
		}

		//热词---热词使用用户数
		if value.NM == "EVENT_SEARCH_PV" {
			if value.FROM == "hot_search" || value.FROM == "3" || value.FROM == "category" || value.FROM == "6" {
				reportDataThree["HotCode"] = reportDataThree["HotCode"] + 1 //热词使用量
				_, ok := HotUse[value.M1]
				if !ok {
					reportDataThree["HotUse"] = reportDataThree["HotUse"] + 1 //热词使用用户数
				}
				HotUse[value.M1] = value.NM + "_" + value.M1
			}
		}
		//关键词---关键词使用用户数
		if value.NM == "EVENT_SEARCH_PV" {
			if value.FROM == "search_input" || value.FROM == "7" {
				reportDataThree["KeyCode"] = reportDataThree["KeyCode"] + 1 //关键词使用量
				_, ok := KeyUse[value.M1]
				if !ok {
					reportDataThree["KeyUse"] = reportDataThree["KeyUse"] + 1 //关键词使用用户数
				}
				KeyUse[value.M1] = value.NM + "_" + value.M1
			}
		}
		//搜索结果点击量---搜索结果点击用户数
		if value.NM == "EVENT_GOODS_DETAIL_PV" {
			if value.FROM == "searchResult" {
				_, ok := SearchResult[value.PID]
				if ok {
					reportDataThree["SearchResult"] = reportDataThree["SearchResult"] + 1 //搜索结果点击量
					_, okTwo := SearchResultClick[value.M1]
					if !okTwo {
						reportDataThree["SearchResultClick"] = reportDataThree["SearchResultClick"] + 1 //搜索结果点击用户数
					}
					SearchResultClick[value.M1] = value.NM + "_" + value.M1
				}
				SearchResult[value.PID] = value.NM + "_" + value.PID
			}
		}
		if value.NM == "EVENT_JUMP_GET_COUPON" {
			_, ok := SearchResultTwo[value.PID]
			if ok {
				reportDataThree["SearchResult"] = reportDataThree["SearchResult"] + 1 //搜索结果点击量
				_, okTwo := SearchResultClickTwo[value.M1]
				if !okTwo {
					reportDataThree["SearchResultClick"] = reportDataThree["SearchResultClick"] + 1 //搜索结果点击用户数
				}
				SearchResultClickTwo[value.M1] = value.NM + "_" + value.M1
			}
			SearchResultTwo[value.PID] = value.NM + "_" + value.PID
		}
		//电商点击量
		if value.NM == "EVENT_JUMP_GET_COUPON" {
			if value.CLICKTYPE == "1" || value.CLICKTYPE == "2" {
				reportDataThree["ShopClick"] = reportDataThree["ShopClick"] + 1 //电商点击量
				_, ok := ShopClickUser[value.M1]
				if !ok {
					reportDataThree["ShopClickUser"] = reportDataThree["ShopClickUser"] + 1 //电商点击用户数
				}
				ShopClickUser[value.M1] = value.NM + "_" + value.M1
				DetailOrder[value.PID] = value.PID + "_" + value.M1
			}
		}
		//详情页人均订单量
		if value.NM == "EVENT_USER_ORDER" {
			_, ok := DetailOrder[value.PID]
			if ok {
				reportDataThree["DetailOrder"] = reportDataThree["DetailOrder"] + 1
			}
		}
		//详情页用户数
		if value.NM == "EVENT_GOODS_DETAIL_PV" {
			_, ok := DetailUser[value.M1]
			if !ok {
				reportDataThree["DetailUser"] = reportDataThree["DetailUser"] + 1
			}
			DetailUser[value.M1] = value.PID + "_" + value.M1
		}
		_ = key
	}
}

func writeReportThreeXls() {

	countContentTitle := "日期" + "\t" + "APP启动量" + "\t" + "商品卡片有效展示量" + "\t" + "商品卡片点击量" + "\t" + "商品卡片点击用户数" + "\t" + "热词使用量" + "\t" + "关键词使用量" + "\t" + "搜索结果点击量" + "\t" + "热词使用用户数" + "\t" + "关键词使用用户数" + "\t" + "搜索结果点击用户数" + "\t" + "电商点击量" + "\t" + "电商点击用户数" + "\t" + "详情页人均订单量" + "\t" + "详情页用户数"
	WriteLog(reportThree, countContentTitle)

	countContent = dateUrl + "\t" + strconv.Itoa(reportDataThree["AppQiDong"]) + "\t" + strconv.Itoa(reportDataThree["GoodsScan"]) + "\t" + strconv.Itoa(reportDataThree["GoodsClick"]) + "\t" + strconv.Itoa(reportDataThree["GoodsDetail"]) + "\t" + strconv.Itoa(reportDataThree["HotCode"]) + "\t" + strconv.Itoa(reportDataThree["KeyCode"]) + "\t" + strconv.Itoa(reportDataThree["SearchResult"]) + "\t" + strconv.Itoa(reportDataThree["HotUse"]) + "\t" + strconv.Itoa(reportDataThree["KeyUse"]) + "\t" + strconv.Itoa(reportDataThree["SearchResultClick"]) + "\t" + strconv.Itoa(reportDataThree["ShopClick"]) + "\t" + strconv.Itoa(reportDataThree["ShopClickUser"]) + "\t" + strconv.Itoa(reportDataThree["DetailOrder"]) + "\t" + strconv.Itoa(reportDataThree["DetailUser"])
	WriteLog(reportThree, countContent)

}
