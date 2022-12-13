package basics

import "fmt"

func LearnIfs() {

	x := 5
	if x > 5 {
		fmt.Println("Bigger 5")
	} else if x > 0 {
		fmt.Println("Bigger zero")
	} else {
		fmt.Println("Else")
	}

	if b := 6; b > 5 {
		fmt.Println("Bigger 5")
	} else {
		fmt.Println("Else")
	}

}
