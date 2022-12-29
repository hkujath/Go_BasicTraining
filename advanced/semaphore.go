package advanced

import (
	"fmt"
	"runtime"
)

func LearnSemaphore() {

	ch := make(chan Task)
	go CalcTaskWithSemaphore(ch)
	resCh := make(chan int)
	ch <- Task{
		a:      3,
		b:      4,
		result: resCh,
	}
	close(ch)
	fmt.Println(<-resCh)
}

func CalcTaskWithSemaphore(tasks chan Task) {
	limit := runtime.NumCPU()
	sem := make(chan bool, limit)

	for task := range tasks {
		sem <- true
		go func() {
			r := task.a + task.b
			task.result <- r
			<-sem
		}()

	}
}
