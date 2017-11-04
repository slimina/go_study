package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

//https://studygolang.com/articles/2492
//我们要加入对pprof包里的方法调用，程序才能将运行时候程序的堆内存分配状态记录到文件中，以便进一步的分析
//如果你的go程序只是一个应用程序，比如计算fabonacci数列，那么你就不能使用net/http/pprof包了，
//你就需要使用到runtime/pprof。具体做法就是用到pprof.StartCPUProfile和pprof.StopCPUProfile。

var cpuprofile = flag.String("cpuprofile", "cpuprofile.prof", "write cpu profile to file")

func main() {
	flag.Parse()

	f, err := os.Create(*cpuprofile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10000; i++ {
		fmt.Println("ddadad---->", i)
	}
}

//运行程序的时候加一个--cpuprofile参数，比如fabonacci --cpuprofile=fabonacci.prof
//这样程序运行的时候的cpu信息就会记录到XXX.prof中了。
