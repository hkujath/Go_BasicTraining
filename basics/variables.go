package basics

import (
	"fmt"
	"os"
)

func LearnVariables() {
	// Simple declaration
	var number int
	var txt string
	var checked bool
	var meinNumberPointer *int
	var list []string

	fmt.Printf("number %v (%T)\n", number, number)
	fmt.Printf("txt %s (%T)\n", txt, txt)
	fmt.Printf("checked %v (%T)\n", checked, checked)
	fmt.Printf("meinNumberPointer %b (%T)\n", meinNumberPointer, meinNumberPointer)
	fmt.Printf("list %+v (%T)\n", list, list)

	// Group declaration
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Printf("home %v (%T)\n", home, home)
	fmt.Printf("user %v (%T)\n", user, user)
	fmt.Printf("gopath %v (%T)\n", gopath, gopath)

	a, b := 12, 34

	fmt.Printf("a %v (%T), b %v (%T)\n", a, a, b, b)

	// Constants
	const c = 24
	const (
		maxWidth  = 100
		maxLength = 100
	)
	const d = 10 + 4

	fmt.Printf("Const variable c %v (%T)\n", c, c)
	fmt.Printf("maxWidth %v (%T), maxLength %v (%T)\n", maxWidth, maxWidth, maxLength, maxLength)
	fmt.Printf("Const variable d %v (%T)\n", d, d)

	const f = 2.0
	fmt.Println(add(f))
	//g := 2.0
	//fmt.Println(add(g))

	// IOTA

	const (
		Monday  = iota
		Tuesday = iota
		Wednesday
		Thursday
		Friday
	)

	const (
		_ = iota
		one
		_
		three
	)

}

func add(length int) int {
	return length + 4
}
