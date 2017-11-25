package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

//Go语言中提供了一个简易的log包，我们使用该包可以方便的实现日志记录的功能，
//这些日志都是基于fmt包的打印再结合panic之类的函数来进行一般的打印、抛出错误处理。

//logrus是用Go语言实现的一个日志系统，与标准库log完全兼容并且核心API很稳定,是Go语言目前最活跃的日志库
//还可以利用Hook机制实现日志功能扩展，例如Syslog hook，将输出的日志发送到指定的Syslog服务。
//go get -u github.com/sirupsen/logrus
/*
logrus和go lib里面一样有6个等级，可以直接调用
logrus.Debug("Useful debugging information.")
logrus.Info("Something noteworthy happened!")
logrus.Warn("You should probably take a look at this.")
logrus.Error("Something failed but I'm not quitting.")
logrus.Fatal("Bye.")   //log之后会调用os.Exit(1)
logrus.Panic("I'm bailing.")   //log之后会panic()
*/
//自定义日志处理
func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})
	//输出stdout而不是默认的stderr，可以为任何io.Writer，比如文件*os.File
	log.SetOutput(os.Stdout)
	// 只记录严重或以上警告
	log.SetLevel(log.InfoLevel)
}
func main() {
	log.WithFields(log.Fields{
		"username": "tom",
	}).Debug("hello world")

	// 通过日志语句重用字段
	// logrus.Entry返回自WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
	log.Info("===================")
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
