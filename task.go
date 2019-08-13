package main

import (
	"formatlog/conf"
	"formatlog/script"
)

func main() {
	/*beego.SetLogger("file", `{"filename":"logs/test.log"}`)
		log := logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
		// 设置配置文件
		jsonConfig := `{
	        "filename" : "test.log", // 文件名
	        "maxlines" : 1000,       // 最大行
	        "maxsize"  : 10240       // 最大Size
	    }`
		log.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
		log.SetLevel(logs.LevelDebug)     // 设置日志写入缓冲区的等级
		log.EnableFuncCallDepth(true)     // 输出log时能显示输出文件名和行号（非必须）

		log.Emergency("Emergency")
		log.Alert("Alert")
		log.Critical("Critical")
		log.Error("Error")
		log.Warning("Warning")
		log.Notice("Notice")
		log.Informational("Informational")
		log.Debug("Debug")

		log.Flush() // 将日志从缓冲区读出，写入到文件
		log.Close()*/
	dotask()

}

func dotask() {
	myConfig := new(conf.Config)
	myConfig.InitConfig("./conf/app.conf")

	//script.WriteXls()
	//script.GetUidM1()
	//script.ReportOnetask()
	//script.ReportTwoTask()
	//script.ReportThreeTask()
	script.InsertJsonDataOne()

	/*taskTime := myConfig.Read("task", "taskTime")
	setLogTask := toolbox.NewTask("setLog", taskTime, newsBot)
	//添加定时任务
	toolbox.AddTask("setLog", setLogTask)
	//启动定时任务
	toolbox.StartTask()
	defer toolbox.StopTask()
	select {}*/

}

func newsBot() error {
	//fmt.Println(11)
	script.WriteXls()
	script.ReportOneTask()
	script.ReportTwoTask()
	script.ReportThreeTask()
	script.GetUidM1()
	script.InsertJsonDataOne()

	return nil
}
