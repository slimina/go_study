package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	fmt.Println("worker ", id, " start")
	for i := range jobs {
		fmt.Println("worker ", id, " start job ", i)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " finished job ", i)
		results <- i * 2
	}
	fmt.Println("worker", id, " end")
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	fmt.Println("初始化worker完成")
	//接收任务
	for i := 1; i < 6; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 5; i > 0; i-- { //不能使用range 会一直堵塞
		fmt.Println("result => ", <-results)
	}

}
