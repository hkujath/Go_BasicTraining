package advanced

import (
	"context"
	"fmt"
)

func LearnContext() {
	myFunc()
}

func myFunc() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := gen(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				// Function terminates, because done-channel was closed ("myFunc()" ended).
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
