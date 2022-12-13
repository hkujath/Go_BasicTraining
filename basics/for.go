package basics

import "fmt"

func LearnFor() {

	fmt.Println("Testing simple loop")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("Testing for loop as while")
	var i2 int
	for i2 <= 4 {
		fmt.Println(i2)
		i2++
	}

	fmt.Println("Testing loop without conditions but break")
	var i3 int
	for {
		if i3 > 4 {
			break
		}
		fmt.Println(i3)
		i3++
	}

	fmt.Println("Testing loop with continue")

	for i4 := 0; i4 <= 10; i4++ {
		if i4%2 == 0 {
			continue
		}
		fmt.Println(i4)
	}

}
