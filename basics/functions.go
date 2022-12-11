package basics

import (
	"errors"
	"fmt"
	"io"
)

func LearnFunctions() {
	greeting := greet("Holger")
	fmt.Println(greeting)

	fmt.Println("Testing swap function")
	a, b := Swap(12, 2)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Testing multiAdd() function")
	fmt.Println(multiAdd(1, 2))
	fmt.Println(multiAdd())
	fmt.Println(multiAdd(1, 2, 3, 4, 5))

	fmt.Println("Testing myFilter() function")

	f := func(input string) bool {
		if len(input) < 3 {
			return false
		}
		return true
	}

	stringsToTest := []string{"a", "abcd", "abc", "ab"}
	fmt.Println(myFilter(stringsToTest, f))

	fmt.Println("Testing defer function")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")

	defer func() {
		fmt.Println(4)
		fmt.Println(5)
	}()
	fmt.Println(6)
}

// Simple function
func greet(name string) string {
	return fmt.Sprintf("Hi %s!", name)
}

func add3(a, b, c int) int {
	return a + b + c
}

// Swap Function with two return values
func Swap(a, b int) (int, int) {
	return b, a
}

// ReadSomething Function with throwing error
func ReadSomething(r io.Reader) (int, error) {

	n := 100

	if r == nil {
		return 0, errors.New("Error while reading")
	}

	return n, nil
}

// multiAdd function to test variable length of input parameters
// a variadic parameter
// summe the result of summing up all input data
func multiAdd(a ...int) (summe int) {
	summe = 0

	for _, v := range a {
		summe += v
	}
	return summe
}

// Type definition of a function
type filterFunc func(string) bool

// myFilter is used to get a function as input parameter, to filter a string.
func myFilter(s []string, filter filterFunc) []string {
	var out []string
	for _, str := range s {
		if filter(str) {
			out = append(out, str)
		}
	}
	return out
}
