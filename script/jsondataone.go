package script

import (
	"bufio"
	_ "encoding/json"
	_ "fmt"
	"formatlog/models"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	formatNow = now.AddDate(0, 0, -1)
	dateUrl = formatNow.Format("20060102")
	url = logReadUrl + dateUrl

}

// 执行脚本
func InsertJsonDataOne() {
	log.Info("开始执行脚本---InsertJsonDataOne")
	//获取数据
	log.Info("logUrl:", url)
	InsertMysql(url)

}

//insert mysql
func InsertMysql(dir string) {
	logDataOne = make(map[string]JsonFormatOne)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
Loop:
	for _, file := range files {
		filename := file.Name()
		if file.IsDir() == true {
			log.Warning("是目录")
			continue Loop
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
		_ = size
		//log.Info("file size=", size)

		buf := bufio.NewReader(logFile)
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			//log.Info(line)
			if len(line) != 0 {
				ba := gjson.Get(line, "ba").String()
				ip := gjson.Get(line, "ip").String()
				m1 := gjson.Get(line, "m1").String()
				mo := gjson.Get(line, "mo").String()
				nm := gjson.Get(line, "nm").String()
				p := gjson.Get(line, "p").String()
				tm := gjson.Get(line, "tm").String()
				v := gjson.Get(line, "v").String()
				//seg
				appv := gjson.Get(line, "seg.custom.appv").String()
				categoryId := gjson.Get(line, "seg.custom.categoryId").String()
				clickType := gjson.Get(line, "seg.custom.clickType").String()
				eplatform := gjson.Get(line, "seg.custom.eplatform").String()
				from := gjson.Get(line, "seg.custom.from").String()
				ips := gjson.Get(line, "seg.custom.ip").String()
				itemid := gjson.Get(line, "seg.custom.itemid").String()
				itemId := gjson.Get(line, "seg.custom.itemId").String()
				pid := gjson.Get(line, "seg.custom.pid").String()
				kw := gjson.Get(line, "seg.custom.kw").String()
				name := gjson.Get(line, "seg.custom.name").String()
				match := gjson.Get(line, "seg.custom.match").String()
				orderid := gjson.Get(line, "seg.custom.orderid").String()
				position := gjson.Get(line, "seg.custom.position").String()
				qd := gjson.Get(line, "seg.custom.qd").String()
				uid := gjson.Get(line, "seg.custom.uid").String()

				//写入数据库
				models.InsertJsonOne(dateUrl,ba, ip, m1, mo, nm, p, appv, categoryId, clickType, eplatform, from, ips, itemid, itemId, orderid, kw, uid, match, pid, position, name, qd, tm, v)
			}
			//WriteLog(countFileName, countContent)
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
	log.Info("Read file over!")
}
