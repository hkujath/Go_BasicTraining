package advanced

import "fmt"

type Task struct {
	a, b   int
	result chan int
}

func LearnChannelOfChannels() {
	ch := make(chan Task)
	go CalcTask(ch)
	resCh := make(chan int)
	ch <- Task{
		a:      3,
		b:      4,
		result: resCh,
	}
	close(ch)
	fmt.Println(<-resCh)
}

func CalcTask(tasks chan Task) {
	for task := range tasks {
		r := task.a + task.b
		task.result <- r
	}
}
