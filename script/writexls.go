package script

import (
	"bufio"
	_ "encoding/json"
	_ "fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	countFileName = "count.xls"
	countContent  string
)

// 执行脚本
func WriteXls() {
	log.Info("开始执行脚本---WriteXls")
	//获取数据
	formatNow = now.AddDate(0, 0, -1)
	dateUrl = formatNow.Format("20060102")
	url = logReadUrl + dateUrl
	log.Info("logUrl:", url)
	readLog(url)
	log.Info("脚本执行完毕---WriteXls")
	os.Exit(2)
}

//读取日志
func readLog(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	countContentTitle := "ba" + "\t" + "m1" + "\t" + "mo" + "\t" + "nm" + "\t" + "tm" + "\t" + "categoryId" + "\t" + "clickType" + "\t" + "eplatform" + "\t" + "from" + "\t" + "ip" + "\t" + "itemid" + "\t" + "itemId" + "\t" + "pid" + "\t" + "orderid" + "\t" + "position" + "\t" + "qd" + "\t" + "uid"
	WriteLog(countFileName, countContentTitle)
	for _, file := range files {
		filename := file.Name()
		if file.IsDir() == true {
			//fmt.Println("是目录")
			log.Warning("是目录")
		}
		//读取具体内容
		//log.Info("filename为:", filename)
		openUrlfile := dir + "/" + filename
		logFile, err := os.OpenFile(openUrlfile, os.O_RDWR, 0666)
		if err != nil {
			log.Info("Open file error!")
			return
		}
		defer logFile.Close()

		stat, err := logFile.Stat()
		if err != nil {
			panic(err)
		}

		var size = stat.Size()
		log.Info("file size=", size)

		buf := bufio.NewReader(logFile)
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			//log.Info(line)
			//content := new(datastruct.JsonOne)
			if len(line) != 0 {
				ba := gjson.Get(line, "ba").String()
				m1 := gjson.Get(line, "m1").String()
				mo := gjson.Get(line, "mo").String()
				nm := gjson.Get(line, "nm").String()
				tm := gjson.Get(line, "tm").String()
				//seg
				categoryId := gjson.Get(line, "seg.custom.categoryId").String()
				clickType := gjson.Get(line, "seg.custom.clickType").String()
				eplatform := gjson.Get(line, "seg.custom.eplatform").String()
				from := gjson.Get(line, "seg.custom.from").String()
				ip := gjson.Get(line, "seg.custom.ip").String()
				itemid := gjson.Get(line, "seg.custom.itemid").String()
				itemId := gjson.Get(line, "seg.custom.itemId").String()
				pid := gjson.Get(line, "seg.custom.pid").String()
				orderid := gjson.Get(line, "seg.custom.orderid").String()
				position := gjson.Get(line, "seg.custom.position").String()
				qd := gjson.Get(line, "seg.custom.qd").String()
				uid := gjson.Get(line, "seg.custom.uid").String()

				countContent = ba + "\t" + m1 + "\t" + mo + "\t" + nm + "\t" + tm + "\t" + categoryId + "\t" + clickType + "\t" + eplatform + "\t" + from + "\t" + ip + "\t" + itemid + "\t" + itemId + "\t" + pid + "\t" + orderid + "\t" + position + "\t" + qd + "\t" + uid
			}
			WriteLog(countFileName, countContent)
			//WriteLog(filename, line)
			if err != nil {
				if err == io.EOF {
					log.Info("File read ok!")
					break
				} else {
					log.Info("Read file error!", err)
					return
				}
			}
		}
	}
}


