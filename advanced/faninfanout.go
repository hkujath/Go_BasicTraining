package advanced

import (
	"fmt"
	"sync"
	"time"
)

func LearnFanInFanOut() {

	runFanOut()

	runFanIn()

}

func runFanOut() {
	var wg sync.WaitGroup
	useInt := func(id int, c <-chan int) {
		for i := range c {
			fmt.Println("ID:", id, "->", i)
		}
		fmt.Println("ID:", id, "ist fertig.")
		wg.Done()
	}

	in := make(chan int)
	wg.Add(3)
	go useInt(1, in)
	go useInt(2, in)
	go useInt(3, in)
	for i := 0; i < 5; i++ {
		in <- i
		time.Sleep(time.Millisecond)
	}
	close(in)
	wg.Wait()
}

func runFanIn() {

}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
