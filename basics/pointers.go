package basics

import "fmt"

func LearnPointers() {
	var a int
	var b *int
	fmt.Println("After init", a, b)
	a = 123
	b = &a
	fmt.Println("b got address from a", b, *b)
	*b = 100
	fmt.Println("a after new value to b", a)
}
