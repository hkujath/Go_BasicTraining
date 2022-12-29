package advanced

func LearnPipeline() {

}

func sq(in <-chan int) <-chan int {
	// Function can block, when nobody writes into in channel
	out := make(chan int)
	go func() {
		for n := range in {
			// function blocks, if nobody is reading the out channel
			out <- n * n
		}
		close(out)
	}()
	return out
}
