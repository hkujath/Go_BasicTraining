package basics

import "fmt"

func LearnSlices() {

	var a []int
	printSlice("a", a)
	a = append(a, 1)
	printSlice("a", a)
	a = append(a, 2)
	printSlice("a", a)
	a = append(a, 3)
	printSlice("a", a)

	b := []int{4, 5, 6}
	printSlice("b", b)
	a = append(a, b...)
	printSlice("a", a)

	fmt.Println("Testing slices with strings")
	s := []string{"null", "one", "two"}
	for i, v := range s {
		fmt.Printf("%02d: %s\n", i, v)
	}
	fmt.Println(s[1])
	fmt.Println(s[2])

	fmt.Println("Testing slices with length")
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(d[0:3])
	fmt.Println(d[:3])
	fmt.Println(d[3:5])
	fmt.Println(d[3:])
	fmt.Println(d[:])

	fmt.Println("Testing slices with append")
	f := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	g := a[1:5]
	printSlice("g := f[1:5]", f, g)
	g = append(g, 11)
	printSlice("after append", f, g)
	h := f[1:5:5]
	printSlice("h := f[1:5:5]", h)
	h = append(h, 14)
	printSlice("after h append", f, h)
	h[1] = 100
	printSlice("after changed entry at pos 1", f, h)

	fmt.Println("Testing slices with copy and make")
	f1 := make([]int, 2, 4)
	num1 := copy(f1, f[1:5])
	fmt.Println("Copied elements:", num1)
	printSlice("copy(f1, f[1:5])", f, f1)
	f2 := make([]int, 4)
	num2 := copy(f2, f[1:5])
	fmt.Println("Copied elements:", num2)
	printSlice("copy(f2, f[1:5])", f, f2)
}

func printSlice(name string, slice ...[]int) {
	fmt.Println(name, ":")
	for _, s := range slice {
		fmt.Printf("%p - len: %d cap: %d - %#v\n", s, len(s), cap(s), s)
	}

}
