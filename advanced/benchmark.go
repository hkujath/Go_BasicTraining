package advanced

import "fmt"

func LearnBenchmark() {
	// benchmark is executed with "go test -bench ."
}

func MySPrintFVersion1(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func MySPrintFVersion2(a, b string) string {
	return a + b
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
