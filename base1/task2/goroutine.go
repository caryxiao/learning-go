package main

import (
	"fmt"
	"time"
)

type Tasks struct {
	task []func() string
}

func (tasks *Tasks) Add(task func() string) {
	tasks.task = append(tasks.task, task)
}

func (tasks *Tasks) Run() {
	for _, task := range tasks.task {
		if task == nil {
			continue
		}
		go tasks.Work(task)
	}
}

func (tasks *Tasks) Work(task func() string) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Task took %s\n", elapsed)
	}()
	res := task()
	fmt.Println(res)
}

// sleep版本
func main() {

	tasks := Tasks{task: make([]func() string, 10)}
	tasks.task[0] = func() string {
		return "Hello World"
	}

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
	time.Sleep(8 * time.Second)
}
