package basics

import "fmt"

func LearnSwitch() {

	a := 2

	switch a {
	case 0:
		fmt.Printf("a is %d (small)\n", a)
	case 1, 2, 3:
		fmt.Printf("a is %d\n", a)
	default:
		fmt.Println("everything else")
	}

	fmt.Println("Testing case with multiple values ")
	switch b := a - 1; b {
	case 0:
		fmt.Printf("b is %d (small)\n", b)
	case 1, 2, 3:
		fmt.Printf("b is %d\n", b)
	default:
		fmt.Println("everything else")
	}

	fmt.Println("Testing break and switch without statement")
	c := a - 2
	switch {
	case c == 0:
		fmt.Printf("c is %d (small)\n", c)
		break
		fmt.Println("Cant reach this part")
	default:
		fmt.Println("everything else")
	}

	fmt.Println("Testing fallthrough")
	d := 42
	switch {
	case d == 42:
		fmt.Printf("d is %d (small)\n", d)
		fallthrough
	case d == 0:
		fmt.Printf("d is %d (small) again\n", d)
		fallthrough
	default:
		fmt.Println("everything else")
	}

	fmt.Println("Testing type switch")

	var x interface{} = "holger"

	switch i := x.(type) {
	case int:
		fmt.Printf("i is a int with value %d", i)
	case float64:
		fmt.Printf("i is a float with value %f", i)
	case bool, string:
		fmt.Println("i is bool or string")
	default:
		fmt.Println("Dont know the type")

	}
}
