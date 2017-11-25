package test6

import (
	. "../test6"
	"testing"
)

/*
压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似
1.压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
  func BenchmarkXXX(b *testing.B) { ... }
2.go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，
  语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数
3.在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行
4.文件名也必须以_test.go结尾
*/
func Benchmark_Division(b *testing.B) {
	//必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。
	//方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

// 测试并发效率
func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Division(4, 5)
		}
	})
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

//go test test6_bench_test.go -test.bench=".*"
