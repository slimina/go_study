package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//对于更复杂的状态，可以使用互斥体来安全地访问多个goroutine中的数据。
//对于每个读取，我们选择一个键来访问，
//Lock()互斥体以确保对状态的独占访问，读取所选键的值，Unlock()互斥体，并增加readOps计数。
func main() {
	var state = make(map[int]int)

	var readOps uint64 = 0
	var writeOps uint64 = 0

	var mutex = &sync.Mutex{} //lock

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)

			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(atomic.LoadUint64(&readOps), atomic.LoadUint64(&writeOps))
	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
}
