package basics

import "fmt"

func LearnGoTo() {

	for row := 0; row <= 2; row++ {
		for column := 0; column <= 3; column++ {
			fmt.Printf("(%d,%d) ", row, column)
			if row+column == 4 {
				goto MyLabel
			}
		}
		fmt.Printf("\n")
	}
MyLabel:
	fmt.Println("\nMy label reached")
}
