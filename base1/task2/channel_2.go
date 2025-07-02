package main

import (
	"fmt"
	"sync"
)

func main() {
	//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	//考察点 ：通道的缓冲机制。
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			fmt.Println("send: ", i)
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("received: ", value)
		}
	}()

	wg.Wait()
}
