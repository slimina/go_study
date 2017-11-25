package main

import (
	"fmt"
	"github.com/cihub/seelog"
)

/*
seelog是用Go语言实现的一个日志系统，它提供了一些简单的函数来实现复杂的日志分配、
过滤和格式化。主要有如下特性：
1.XML的动态配置，可以不用重新编译程序而动态的加载配置信息
2.支持热更新，能够动态改变配置而不需要重启应用
3.支持多输出流，能够同时把日志输出到多种流中、例如文件流、网络流等
4.支持不同的日志输出
  命令行输出
  文件输出
  缓存输出
  支持log rotate
  SMTP邮件
*/
//基于seelog的自定义日志处理
var logger seelog.LoggerInterface

func loadAppConfig() {
	//根据配置文件初始化seelog的配置信息，这里我们把配置文件通过字符串读取设置好了，当然也可以通过读取XML文件。
	/*
		seelog: minlevel参数可选，如果被配置,高于或等于此级别的日志会被记录
		outputs: 输出信息的目的地，设置了filter, 如果这个错误级别是critical，那么将发送报警邮件。
		formats: 定义了各种日志的格式
	*/

	//通过recipient配置接收邮件的用户，如果有多个用户可以再添加一行
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="/data/logs/roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
            <file path="/data/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="xxx@qq.com" sendername="ShortUrl API" hostname="smtp.qq.com" hostport="25" username="xxx@qq.com" password="xxx">
                <recipient address="xxx@qq.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
	    <format id="critical" format="%File %FullPath %Func %Msg%n" />
	    <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
	`
	newLogger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(newLogger)
}

func init() {
	//初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
	DisableLog()
	loadAppConfig()
}

//设置当前的日志器为相应的日志处理
func UseLogger(newLogger seelog.LoggerInterface) {
	logger = newLogger
}

func DisableLog() {
	logger = seelog.Disabled
}
func main() {
	defer logger.Flush()
	logger.Warn("waring .....")
	logger.Critical("error .....")
}
