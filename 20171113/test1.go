package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/dpatel06/logrus_gomail"
	"os"
	"time"
)

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})
	//输出stdout而不是默认的stderr，可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	// 只记录严重或以上警告
	log.SetLevel(log.InfoLevel)
}
func main() {
	logger := log.New()
	//parameter"APPLICATION_NAME", "HOST", PORT, "FROM", "TO"
	//首先开启smtp服务，最后两个参数是smtp的用户名和密码
	hook, err := logrus_gomail.NewGoMailAuthHook("app name", "smtp.qq.com", 25, "xxx@qq.com",
		"xxx@qq.com", "xxx@qq.com", "1111111111")
	if err == nil {
		logger.Hooks.Add(hook)
	} else {
		logger.Fatal(err)
	}

	var filename = "123.txt"
	contextLogger := logger.WithFields(log.Fields{
		"file":    filename,
		"content": "GG",
	})
	//设置时间戳和message
	contextLogger.Time = time.Now()
	contextLogger.Message = "这是一个hook发来的邮件"
	//只能发送Error,Fatal,Panic级别的log
	contextLogger.Level = log.FatalLevel

	//使用Fire发送,包含时间戳，message
	hook.Fire(contextLogger)
}
