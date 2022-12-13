package basics

import "fmt"

func LearnMaps() {

	m := make(map[int]string)
	m[1] = "Eintrag 1"
	fmt.Println(m)

	m1 := map[int]string{
		1: "Eintrag 1",
		2: "Eintrag 2",
	}
	fmt.Println(m1)
}
