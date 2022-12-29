package advanced

import (
	"fmt"
	"runtime"
	"sync"
)

func LearnRaceConditions() {

	// Counter number is unpredictable because every sub routine wants to increase the counter -> data race
	// Because of this we can prevent this race by using a mutex
	// Another way would be to use a atomic function e.g. atomic.AddInt64(&counter,1)
	mutex := &sync.Mutex{}
	counter := 0
	//var counter2 int64 = 0

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			c := counter
			runtime.Gosched()
			counter = c + 1
			//atomic.AddInt64(&counter2, 1)
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	//fmt.Println(counter2)

	// Call "go run -race main.go" from command line so detect race conditions
}
