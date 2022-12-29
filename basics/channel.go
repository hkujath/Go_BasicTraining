package basics

import (
	"fmt"
	"sync"
	"time"
)

func LearnChannel() {

	wg := &sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(2)
	go routine1(ch, wg)
	go routine2(ch, wg)
	wg.Wait()

	// Using anonymous function
	go func(ch chan string) {
		ch <- "Hello from anonymous function"
	}(ch)
	msg := <-ch
	fmt.Println("Received message from anonymous function: ", msg)

	//Channel as input and output
	in := make(chan int)
	out := make(chan int)
	go spaceShip(in, out)
	in <- 1
	in <- 2
	fmt.Println("Result of out channel: ", <-out)

	//Using buffered channel
	chBuffered := make(chan int, 2)
	chBuffered <- 2
	chBuffered <- 3
	fmt.Println(<-chBuffered)
	fmt.Println(<-chBuffered)
	chBuffered <- 4
	chBuffered <- 5
	for i := 0; i < cap(chBuffered); i++ {
		fmt.Println(<-chBuffered)
	}

	// Directional channel
	dirChan := make(chan int)
	go directionalChannel(dirChan)
	fmt.Println(<-dirChan)

	// Closing channel
	strChan := make(chan string)
	for i := 0; i < 5; i++ {
		go listening(strChan)
		time.Sleep(time.Second)
	}
	strChan <- "Single message"
	close(strChan)
	time.Sleep(time.Second)

	// Starting routines simultaneously
	running()

	// Using range with a channel
	intChan := make(chan int)
	go func(chan int) {
		for i := 0; i < 2; i++ {
			intChan <- i
		}
		close(intChan)
	}(intChan)
	for n := range intChan {
		fmt.Println(n)
	}

	// Using "select" with a channel
	timout := time.Second * 3
	result := make(chan int, 1)
	go func() {
		result <- 555
	}()
	select {
	case r := <-result:
		fmt.Println(r)
	case <-time.After(timout):
		fmt.Println("timeout")
		return
	}

	fmt.Println("End")
}

func running() {
	letsGo := make(chan struct{})
	fmt.Println("Ready, set, ...")
	go runner(letsGo)
	go runner(letsGo)
	go runner(letsGo)
	// After closing the channel, all routines will continue to work.
	close(letsGo)

}

func runner(letsGo chan struct{}) {
	<-letsGo
	// Do something
}

func listening(ch chan string) {
	fmt.Println("Waiting...")
	msg := <-ch
	fmt.Println("Message received: ", msg)
}

func Add(a, b int) int {
	return a + b
}

func spaceShip(in chan int, out chan int) {
	a := <-in
	b := <-in
	out <- Add(a, b)
}

func routine1(ch chan string, wg *sync.WaitGroup) {
	ch <- "Hello routine2"
	wg.Done()
}

func routine2(ch chan string, wg *sync.WaitGroup) {
	msg := <-ch
	fmt.Println("Message received in routing2: ", msg)
	wg.Done()
}

func directionalChannel(ch chan<- int) {
	ch <- 42
}
