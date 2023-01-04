package forms

type Former interface {
	Area() int
	Position() (int, int)
}

type Rectangle struct {
	Length int
	Width  int
	XPos   int
	YPos   int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func (r Rectangle) Position() (int, int) {
	return r.XPos, r.YPos
}

type Circle struct {
	Radius int
	XPos   int
	YPos   int
}

func (c Circle) Area() int {
	return c.Radius * c.Radius * 3
}

func (c Circle) Position() (int, int) {
	return c.XPos, c.YPos
}

func SameArea(a, b Former) bool {
	return a.Area() == b.Area()
}
