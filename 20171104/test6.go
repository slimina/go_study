package main

import (
	"fmt"
	"time"
)

func main() {
	out := fmt.Println

	now := time.Now()
	out(now)
	out(now.Location())

	then := time.Date(2017, 11, 4, 15, 32, 34, 123456000, time.UTC) //北京CST 比UTC 早8小时
	out(then)

	out(then.Year())
	out(then.Month())
	out(then.Day())
	out(then.Hour())
	out(then.Minute())
	out(then.Second())
	out(then.Nanosecond())
	out(then.Location())

	out(then.Weekday())

	out(then.Before(now)) //then 是否在 now之前
	out(then.After(now))  //then 是否在 now之后
	out(then.Equal(now))
	fmt.Println("--------------------")
	diff := now.Sub(then) // now - then 相差时间
	out(diff)
	out(diff.Hours())
	out(diff.Minutes())
	out(diff.Seconds())
	out(diff.Nanoseconds())

	out(then.Add(diff))
	out(then.Add(-diff))

	fmt.Println("----------------------")
	secs := now.Unix()
	nanos := now.UnixNano()
	out(secs)            //自Unix纪元以来的秒数
	out(nanos)           //自Unix纪元以来的纳秒数
	out(nanos / 1000000) //自Unix纪元以来的毫秒数
	//将整数秒或纳秒从历元转换为相应的时间
	out(time.Unix(secs, 0))
	out(time.Unix(0, nanos))

	fmt.Println("---------------")
	//通过基于模式的布局进行时间格式化和解析
	//时间解析使用Format相同的布局值。
	out(now.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2017-11-04T16:21:00+08:00")
	out(t1)
	//格式化，解析，使用如下时间值作为模板，不然无效：Mon Jan 2 15:04:05 -0700 MST 2006
	out(t1.Format("3:04PM"))
	out(t1.Format("2006-01-02 15:04:05"))
	out(t1.Format("2006/01/02 3:04:05"))
	out(time.Parse("2006-01-02 15:04:05", "2017-11-04 16:21:00"))
}
