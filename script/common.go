package script

import (
	"bufio"
	_ "encoding/json"
	_ "fmt"
	"formatlog/conf"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type JsonFormatOne struct {
	BA         string
	IP         string
	M1         string
	MO         string
	NM         string
	P          string
	APPV       string
	CATEGORYID string
	CLICKTYPE  string
	EPLATFORM  string
	FROM       string
	IPS        string
	ITEMID     string
	ITEMIDTWO  string
	ORDERID    string
	KW         string
	UID        string
	MATCH      string
	PID        string
	POSITION   string
	NAME       string
	QD         string
	TM         string
	V          string
}

var (
	logDataOne      map[string]JsonFormatOne
	log             = logs.NewLogger()
	LineFeed        = "\r\n"
	now             = time.Now()
	formatNow       = now.AddDate(0, 0, -1)
	dateUrl         = formatNow.Format("20060102")

	logReadUrl      string
	logWriteUrl     string
)

func init() {
	//hdfsLogWriteUrl := "logs/" + dateUrl + "/formatlog.log"
	//log := logs.NewLogger()
	jsonConfig := `{
        "filename" : "logs/formatlog.log",
        "maxlines" : 0,     
		"level"	:	7,         
        "maxsize"  : 0,     
		"daily"  : true,      
		"maxdays "  : 10      
    }`
	//log.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	//log.SetLogger("file", `{"filename":"logs/formatlog.log","level":7,"maxlines":1000,"maxsize":10240,"daily":true,"maxdays":10}`)
	log.SetLogger(logs.AdapterConsole, jsonConfig)
	log.SetLogger(logs.AdapterFile, jsonConfig)
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	//log.Flush() // 将日志从缓冲区读出，写入到文件
	//log.Close()

	myConfig := new(conf.Config)
	myConfig.InitConfig("./conf/app.conf")
	logReadUrl = myConfig.Read("log", "logReadPath")
	logWriteUrl = myConfig.Read("log", "logWritePath")

}

//WriteLog return error
func WriteLog(fileName, msg string) error {

	path := logWriteUrl + dateUrl + "/"
	if !IsExist(path) {
		CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, LineFeed+msg)

	defer f.Close()
	//f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	return err
}

//CreateDir  文件夹创建
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Info("errurl:", err)
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

//读取日志
func ReadCountLog(dir string) {
	logDataOne = make(map[string]JsonFormatOne)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var fileNum = 0
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
		var lineNum = 0
		fileNum++
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			lineNum++
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

				//写入map
				//mapString := m1 + "_" + tm + "_" + nm
				mapString := m1 + "_" + tm + "_" + nm + "_" + from + "_" + strconv.Itoa(lineNum) + "_" + strconv.Itoa(fileNum)
				//log.Info(mapString)
				logDataOne[mapString] = JsonFormatOne{ba, ip, m1, mo, nm, p, appv, categoryId, clickType, eplatform, from, ips, itemid, itemId, orderid, kw, uid, match, pid, position, name, qd, tm, v}

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
