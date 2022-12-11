package basics

import "fmt"

func LearnObjectOrientation() {

	fmt.Println("Testing calcArea() function")
	r := Rectangle{
		height: 10,
		width:  5,
	}
	fmt.Println(calcAreaAsFunction(r))
	r.SetWidth(2)
	fmt.Println(r.CalcAreaAsMethod())
}

type Rectangle struct {
	height int
	width  int
}

func calcAreaAsFunction(r Rectangle) int {
	return r.width * r.height
}

func (r *Rectangle) CalcAreaAsMethod() int {
	return r.height * r.width
}

func (r *Rectangle) SetWidth(w int) {
	r.width = w
}
