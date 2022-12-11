package basics

import "fmt"

func LearnArrays() {

	var a [4]string
	a[2] = "a"
	var b [5]string
	b[0] = "b"
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Testing arrays of structs")
	var r [10]Rectangle
	fmt.Println(r)
	var c [4][4]int
	fmt.Println(c)
}
