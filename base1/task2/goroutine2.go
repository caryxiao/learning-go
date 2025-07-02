package main

import (
	"fmt"
	"sync"
	"time"
)

type Tasks2 struct {
	task []func() string
	wg   sync.WaitGroup
}

func (tasks *Tasks2) Add(task func() string) {
	tasks.task = append(tasks.task, task)
}

func (tasks *Tasks2) Run() {
	for _, task := range tasks.task {
		if task == nil {
			continue
		}
		tasks.wg.Add(1)
		go tasks.Work(task)
	}
	tasks.wg.Wait()
}

func (tasks *Tasks2) Work(task func() string) {
	start := time.Now()
	defer tasks.wg.Done()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Task took %s\n", elapsed)
	}()
	res := task()
	fmt.Println(res)
}

// sync.WaitGroup 版本
func main() {

	tasks := Tasks2{task: make([]func() string, 10)}
	tasks.Add(func() string {
		return "Hello World"
	})

	tasks.Add(func() string {
		v := 1 + 2
		return fmt.Sprintf("1 + 2 = %d", v)
	})

	tasks.Add(func() string {
		for i := 0; i < 10; i++ {
			fmt.Println("number: ", i)
		}
		return "Print Number"
	})

	tasks.Add(func() string {
		for i := 0; i < 10; i++ {
			fmt.Println("wait index:", i, " ms: ", 500)
			time.Sleep(500 * time.Millisecond)
		}

		return "sleep 500 ms * 10"
	})

	tasks.Run()
}
