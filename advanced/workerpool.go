package advanced

import (
	"fmt"
	"runtime"
	"time"
)

type Job func(chan result)

type result struct {
	foo string
	sum int
	err error
}

func LearnWorkerpool() {
	jobs := make(chan Job)
	results := make(chan result)

	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()

	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		startWorker(jobs, results)
	}

	job1 := func(resCh chan result) {
		r := result{"10 + 5", 10 + 5, nil}
		resCh <- r
	}

	counter := 0
	job2 := func(resCh chan result) {
		counter++
		r := result{"Counter: ", counter, nil}
		resCh <- r
	}

	fmt.Println("Stop wit Ctrl+C")

	for {
		jobs <- job1
		jobs <- job2
		time.Sleep(time.Second)

	}
}

func startWorker(jobs chan Job, results chan result) {
	fmt.Println("Worker started.")
	go func() {
		for j := range jobs {
			j(results)
		}
		fmt.Println("Worker finished.")
	}()
}
