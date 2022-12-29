package basics

import (
	"fmt"
	"sync"
)

func LearnRoutines() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// Use an own function definition
	go greeter(wg, "Holger")

	// Use an anonymous function (closure)
	go func(name string) {
		fmt.Printf("Hi %s\n", name)
		wg.Done()
	}("Klaus")
	wg.Wait()
}

func greeter(wg *sync.WaitGroup, name string) {
	fmt.Printf("Hi %s\n", name)
	wg.Done()
}
