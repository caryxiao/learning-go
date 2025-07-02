package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	value int32
}

func (c *counter) Incr() {
	atomic.AddInt32(&c.value, 1)
}

func (c *counter) Load() int32 {
	return atomic.LoadInt32(&c.value)
}

func main() {
	//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//考察点 ：原子操作、并发数据安全。
	var wg sync.WaitGroup
	counter := &counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println("count: ", counter.Load())

}
