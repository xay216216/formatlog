package script

import (
	"strconv"
)

type BaseDataReportOne struct {
	EnterNum  int
	EnterPeop int
	ClickNum  int
	ClikPeop  int
	BrowNum   int
	BrowPeop  int
}

type ReportOne struct {
	BcEnterNum  int
	BcEnterPeop int
	BcClickNum  int
	BcClikPeop  int
	BcBrowNum   int
	BcBrowPeop  int
	PpEnterNum  int
	PpEnterPeop int
	PpClickNum  int
	PpClikPeop  int
	PpBrowNum   int
	PpBrowPeop  int
	DeEnterNum  int
	DeEnterPeop int
	DeClickNum  int
	DeClikPeop  int
	DeBrowNum   int
	DeBrowPeop  int
	LhEnterNum  int
	LhEnterPeop int
	LhClickNum  int
	LhClikPeop  int
	LhBrowNum   int
	LhBrowPeop  int
}

var (
	reportDataOne  map[string]int
	setData        map[string]string
	setBcData      map[string]string
	setBcBrowData  map[string]string
	setPpData      map[string]string
	setPpBrowData  map[string]string
	setDeData      map[string]string
	setDeBrowData  map[string]string
	setLhData      map[string]string
	setLhClickData map[string]string
	reportOne      = "reportone.xls"
)

// 执行脚本
func ReportOneTask() {
	log.Info("开始执行脚本---CountLog")
	//获取数据
	formatNow     = now.AddDate(0, 0, -1)
	dateUrl       = formatNow.Format("20060102")
	url := logReadUrl + dateUrl
	log.Info("logUrl:", url)
	ReadCountLog(url)
	reportOneCount()
	writeReportOneXls()
}

func reportOneCount() {
	reportDataOne = make(map[string]int)
	setData = make(map[string]string)
	setBcData = make(map[string]string)
	setBcBrowData = make(map[string]string)
	setPpData = make(map[string]string)
	setPpBrowData = make(map[string]string)
	setDeData = make(map[string]string)
	setDeBrowData = make(map[string]string)
	setLhData = make(map[string]string)
	setLhClickData = make(map[string]string)
	for _, value := range logDataOne {
		//log.Info("key:"+key, value)
		//报表一
		if value.NM == "HOME_JINGANG_PV" {
			_, ok := setData[value.M1]
			if !ok {
				reportDataOne["BcEnterPeop"] = reportDataOne["BcEnterPeop"] + 1
			}
			setData[value.M1] = value.NM + "_" + value.M1
			reportDataOne["BcEnterNum"] = reportDataOne["BcEnterNum"] + 1
		}

		if value.NM == "JINGANG_CLICK_EVERYTIME" {
			if value.NAME == "淘宝白菜" {
				_, ok := setBcData[value.M1]
				if !ok {
					reportDataOne["BcClikPeop"] = reportDataOne["BcClikPeop"] + 1
				}
				setBcData[value.M1] = value.NM + "_" + value.M1
				reportDataOne["BcClickNum"] = reportDataOne["BcClickNum"] + 1
			}
			if value.NAME == "品牌特价" {
				_, ok := setPpData[value.M1]
				if !ok {
					reportDataOne["PpClikPeop"] = reportDataOne["PpClikPeop"] + 1
				}
				setPpData[value.M1] = value.NM + "_" + value.M1
				reportDataOne["PpClickNum"] = reportDataOne["PpClickNum"] + 1
			}
			if value.NAME == "大额神券" {
				_, ok := setDeData[value.M1]
				if !ok {
					reportDataOne["DeClikPeop"] = reportDataOne["DeClikPeop"] + 1
				}
				setDeData[value.M1] = value.NM + "_" + value.M1
				reportDataOne["DeClickNum"] = reportDataOne["DeClickNum"] + 1
			}
		}

		if value.NAME == "淘宝白菜" {
			_, ok := setBcBrowData[value.M1]
			if !ok {
				reportDataOne["BcBrowPeop"] = reportDataOne["BcBrowPeop"] + 1
			}
			setBcBrowData[value.M1] = value.NM + "_" + value.M1
			reportDataOne["BcBrowNum"] = reportDataOne["BcBrowNum"] + 1
		}

		if value.NAME == "品牌特价" {
			_, ok := setPpBrowData[value.M1]
			if !ok {
				reportDataOne["PpBrowPeop"] = reportDataOne["PpBrowPeop"] + 1
			}
			setPpBrowData[value.M1] = value.NM + "_" + value.M1
			reportDataOne["PpBrowNum"] = reportDataOne["PpBrowNum"] + 1
		}

		if value.NAME == "大额神券" {
			_, ok := setDeBrowData[value.M1]
			if !ok {
				reportDataOne["DeBrowPeop"] = reportDataOne["DeBrowPeop"] + 1
			}
			setDeBrowData[value.M1] = value.NM + "_" + value.M1
			reportDataOne["DeBrowNum"] = reportDataOne["DeBrowNum"] + 1
		}

		if value.NM == "FLEXIBLE_ACTIVI_VISIBILE" {
			_, ok := setLhData[value.M1]
			if !ok {
				reportDataOne["LhEnterPeop"] = reportDataOne["LhEnterPeop"] + 1
			}
			setLhData[value.M1] = value.NM + "_" + value.M1
			reportDataOne["LhEnterNum"] = reportDataOne["LhEnterNum"] + 1
		}

		if value.NM == "FLEXIBLEITEM_CLICK" {
			_, ok := setLhClickData[value.M1]
			if !ok {
				reportDataOne["LhClikPeop"] = reportDataOne["LhClikPeop"] + 1
			}
			setLhClickData[value.M1] = value.NM + "_" + value.M1
			reportDataOne["LhClickNum"] = reportDataOne["LhClickNum"] + 1
		}
	}
}

func writeReportOneXls() {

	countContentTitle := "日期" + "\t" + "模块名称" + "\t" + "专区入口展示次数" + "\t" + "专区入口展示人数" + "\t" + "专区入口点击次数" + "\t" + "专区入口点击人数" + "\t" + "专区访问次数" + "\t" + "专区访问人数"
	WriteLog(reportOne, countContentTitle)

	countContent = dateUrl + "\t" + "淘宝白菜" + "\t" + strconv.Itoa(reportDataOne["BcEnterNum"]) + "\t" + strconv.Itoa(reportDataOne["BcEnterPeop"]) + "\t" + strconv.Itoa(reportDataOne["BcClickNum"]) + "\t" + strconv.Itoa(reportDataOne["BcClikPeop"]) + "\t" + strconv.Itoa(reportDataOne["BcBrowNum"]) + "\t" + strconv.Itoa(reportDataOne["BcBrowPeop"]) + "\r\n" + dateUrl + "\t" + "品牌特价" + "\t" + strconv.Itoa(reportDataOne["BcEnterNum"]) + "\t" + strconv.Itoa(reportDataOne["BcEnterPeop"]) + "\t" + strconv.Itoa(reportDataOne["PpClickNum"]) + "\t" + strconv.Itoa(reportDataOne["PpClikPeop"]) + "\t" + strconv.Itoa(reportDataOne["PpBrowNum"]) + "\t" + strconv.Itoa(reportDataOne["PpBrowPeop"]) + "\r\n" + dateUrl + "\t" + "大额神券" + "\t" + strconv.Itoa(reportDataOne["BcEnterNum"]) + "\t" + strconv.Itoa(reportDataOne["BcEnterPeop"]) + "\t" + strconv.Itoa(reportDataOne["DeClickNum"]) + "\t" + strconv.Itoa(reportDataOne["DeClikPeop"]) + "\t" + strconv.Itoa(reportDataOne["DeBrowNum"]) + "\t" + strconv.Itoa(reportDataOne["DeBrowPeop"]) + "\r\n" + dateUrl + "\t" + "灵活运营位" + "\t" + strconv.Itoa(reportDataOne["LhEnterNum"]) + "\t" + strconv.Itoa(reportDataOne["LhEnterPeop"]) + "\t" + strconv.Itoa(reportDataOne["LhClickNum"]) + "\t" + strconv.Itoa(reportDataOne["LhClikPeop"]) + "\t" + strconv.Itoa(reportDataOne["LhClickNum"]) + "\t" + strconv.Itoa(reportDataOne["LhClikPeop"])

	WriteLog(reportOne, countContent)
}
