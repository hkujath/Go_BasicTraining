package advanced

import (
	"fmt"
	"testing"
)

var (
	s           string
	x           = "abcd"
	y           = "abcd"
	benchResult int
)

func BenchmarkSprintfVersion1(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = MySPrintFVersion1(x, y)
	}
	s = str
}

func BenchmarkSprintfVersion2(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = MySPrintFVersion2(x, y)
	}
	s = str
}

func BenchmarkFibonacci(b *testing.B) {
	inValues := []int{1, 10, 20}

	for _, v := range inValues {
		b.Run(fmt.Sprintf("fib(%d)", v), func(b *testing.B) {
			var f int
			for i := 0; i < b.N; i++ {
				f = fib(v)
			}
			benchResult = f
		})
	}

}
