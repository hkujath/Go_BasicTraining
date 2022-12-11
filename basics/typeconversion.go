package basics

import "fmt"

type (
	meter      int
	centimeter int
)

func LearnTypeConversion() {
	fmt.Println("Converted to centimeter:", meterToCentimeter(4))

	type (
		apple  int
		cherry int
	)

	a := apple(10)
	b := cherry(5)

	anzahl := int(a) + int(b)
	fmt.Printf("Number fruits: %d\n", anzahl)

}

func meterToCentimeter(m meter) centimeter {
	return centimeter(m * 100)
}
