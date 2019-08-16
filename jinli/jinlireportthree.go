package jinli

import (
	"strconv"
	_ "strconv"
)

var (
	url string
)
var (
	reportDataThree map[string]int
	flagNum         map[string]string
	MpLh            map[string]string
	Tcxs            map[string]string
	Jhdj            map[string]string
	Tcywgb          map[string]string
	Xctxdj          map[string]string
	Fbdj            map[string]string
	Xtcjh           map[string]string
	Xtcgb           map[string]string

	reportThree = "reportthree.xls"
	Dspt        string
)

type ReportThree struct {
	MplhPv   int
	MplhUv   int
	TcxsPv   int
	TcxsUv   int
	JhdjPv   int
	JhdjUv   int
	Jhl      float32
	TcywgbPv int
	TcywgbUv int
	Tcywgbl  float32
	XctxdjPv int
	XctxdjUv int
	Gbl      float32
	FbdjPv   int
	FbdjUv   int
	Rjfbdj   float32
	XtcjhPv  int
	XtcjhUv  int
	Fbjhl    float32
	XtcgbPv  int
	XtcgbUv  int
	Xcgbl    float32
}

// 执行脚本
func ReportThreeTask() {
	log.Info("开始执行脚本---jinliReportThreeTask")
	//获取数据
	formatNow = now.AddDate(0, 0, -8)
	dateUrl = formatNow.Format("20060102")
	url = logReadUrl + dateUrl
	log.Info("logUrl:", url)
	ReadCountLog(url)
	for i := 0; i < 3; i++ {
		reportThreeCount(i)
		writeReportThreeXls()
	}

}

func reportThreeCount(typess int) {
	reportDataThree = make(map[string]int)
	flagNum = make(map[string]string)
	MpLh = make(map[string]string)
	Tcxs = make(map[string]string)
	Jhdj = make(map[string]string)
	Tcywgb = make(map[string]string)
	Xctxdj = make(map[string]string)
	Fbdj = make(map[string]string)
	Xtcjh = make(map[string]string)
	Xtcgb = make(map[string]string)
	if typess == 0 {
		flagNum["1"] = "1"
		flagNum["2"] = "2"
		Dspt = "全部"
	} else if typess == 1 {
		flagNum["1"] = "1"
		Dspt = "淘宝"
	} else {
		flagNum["2"] = "2"
		Dspt = "拼多多"
	}
	for key, value := range logDataOne {
		_, okss := flagNum[value.EPLATFORM]
		if okss {
			//灭屏拉活
			if value.NM == "GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN" {
				reportDataThree["MplhPv"] = reportDataThree["MplhPv"] + 1
				_, ok := MpLh[value.M1]
				if !ok {
					reportDataThree["MplhUv"] = reportDataThree["MplhUv"] + 1
				}
				MpLh[value.M1] = value.NM + "_" + value.M1
			}
			//弹窗显示
			if value.NM == "GIONEE_TAOBAO_HELPER_DIALOG_SHOW" {
				reportDataThree["TcxsPv"] = reportDataThree["TcxsPv"] + 1
				_, ok := Tcxs[value.M1]
				if !ok {
					reportDataThree["TcxsUv"] = reportDataThree["TcxsUv"] + 1
				}
				Tcxs[value.M1] = value.NM + "_" + value.M1
			}
			//激活点击
			if value.NM == "GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK" {
				reportDataThree["JhdjPv"] = reportDataThree["JhdjPv"] + 1
				_, ok := Jhdj[value.M1]
				if !ok {
					reportDataThree["JhdjUv"] = reportDataThree["JhdjUv"] + 1
				}
				Jhdj[value.M1] = value.NM + "_" + value.M1
			}
			//弹窗意外关闭
			if value.NM == "GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE" {
				reportDataThree["TcywgbPv"] = reportDataThree["TcywgbPv"] + 1
				_, ok := Tcywgb[value.M1]
				if !ok {
					reportDataThree["TcywgbUv"] = reportDataThree["TcywgbUv"] + 1
				}
				Tcywgb[value.M1] = value.NM + "_" + value.M1
			}
			//下次提醒点击
			if value.NM == "GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE" {
				reportDataThree["XctxdjPv"] = reportDataThree["XctxdjPv"] + 1
				_, ok := Xctxdj[value.M1]
				if !ok {
					reportDataThree["XctxdjUv"] = reportDataThree["XctxdjUv"] + 1
				}
				Xctxdj[value.M1] = value.NM + "_" + value.M1
			}
			//浮标点击
			if value.NM == "GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK" {
				reportDataThree["FbdjPv"] = reportDataThree["FbdjPv"] + 1
				_, ok := Fbdj[value.M1]
				if !ok {
					reportDataThree["FbdjUv"] = reportDataThree["FbdjUv"] + 1
				}
				Fbdj[value.M1] = value.NM + "_" + value.M1
			}
			//小弹窗激活
			if value.NM == "GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK" {
				reportDataThree["XtcjhPv"] = reportDataThree["XtcjhPv"] + 1
				_, ok := Xtcjh[value.M1]
				if !ok {
					reportDataThree["XtcjhUv"] = reportDataThree["XtcjhUv"] + 1
				}
				Xtcjh[value.M1] = value.NM + "_" + value.M1
			}
			//小弹窗关闭
			if value.NM == "GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE" {
				reportDataThree["XtcgbPv"] = reportDataThree["XtcgbPv"] + 1
				_, ok := Xtcgb[value.M1]
				if !ok {
					reportDataThree["XtcgbUv"] = reportDataThree["XtcgbUv"] + 1
				}
				Xtcgb[value.M1] = value.NM + "_" + value.M1
			}

		}

		_ = key
	}
}

func writeReportThreeXls() {

	countContentTitle := "日期" + "\t" + "电商平台" + "\t" + "灭屏拉活电商PV" + "\t" + "灭屏拉活UV" + "\t" + "弹窗显示PV" + "\t" + "弹窗显示UV" + "\t" + "激活点击PV" + "\t" + "激活点击UV" + "\t" + "激活率" + "\t" + "弹窗意外关闭PV" + "\t" + "弹窗意外关闭UV" + "\t" + "意外关闭率" + "\t" + "下次提醒点击PV" + "\t" + "下次提醒点击UV" + "\t" + "关闭率" + "\t" + "浮标点击PV" + "\t" + "浮标点击UV" + "\t" + "人均浮标点击" + "\t" + "小弹窗激活PV" + "\t" + "小弹窗激活UV" + "\t" + "浮标激活率" + "\t" + "小弹窗关闭PV" + "\t" + "小弹窗关闭UV" + "\t" + "小窗关闭率"
	WriteLog(reportThree, countContentTitle)

	countContent = dateUrl + "\t" + Dspt + "\t" + strconv.Itoa(reportDataThree["MplhPv"]) + "\t" + strconv.Itoa(reportDataThree["MplhUv"]) + "\t" + strconv.Itoa(reportDataThree["TcxsPv"]) + "\t" + strconv.Itoa(reportDataThree["TcxsUv"]) + "\t" + strconv.Itoa(reportDataThree["JhdjPv"]) + "\t" + strconv.Itoa(reportDataThree["JhdjUv"]) + "\t" + "激活率" + "\t" + strconv.Itoa(reportDataThree["TcywgbPv"]) + "\t" + strconv.Itoa(reportDataThree["TcywgbUv"]) + "\t" + "意外关闭率" + "\t" + strconv.Itoa(reportDataThree["XctxdjPv"]) + "\t" + strconv.Itoa(reportDataThree["XctxdjUv"]) + "\t" + "关闭率" + "\t" + strconv.Itoa(reportDataThree["FbdjPv"]) + "\t" + strconv.Itoa(reportDataThree["FbdjUv"]) + "\t" + "人均浮标点击" + "\t" + strconv.Itoa(reportDataThree["XtcjhPv"]) + "\t" + strconv.Itoa(reportDataThree["XtcjhUv"]) + "\t" + "浮标激活率" + "\t" + strconv.Itoa(reportDataThree["XtcgbPv"]) + "\t" + strconv.Itoa(reportDataThree["XtcgbUv"]) + "\t" + "小窗关闭率"
	WriteLog(reportThree, countContent)

}
