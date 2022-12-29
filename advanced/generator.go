package advanced

func LearnGenerator() {

}

func gen2(done chan struct{}) chan int {
	ch := make(chan int)

	go func() {
		i := 0

		for {
			select {
			case ch <- i:
				i++
			case <-done:
				close(ch)
				return
			}
		}
	}()
	return ch
}
