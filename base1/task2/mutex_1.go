package main

import (
	"fmt"
	"sync"
)

func main() {
	//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//考察点 ： sync.Mutex 的使用、并发数据安全。
	var wg sync.WaitGroup
	country := struct {
		count int
		lock  sync.Mutex
	}{
		count: 0,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				country.lock.Lock()
				country.count++
				country.lock.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("count: ", country.count)
}
