package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {

	var s Shape
	s = Rectangle{length: 4, width: 3}
	val, ok := s.(Shape)
	if ok {
		fmt.Printf("Area of the rectangle: %.2f\n", val.Area())
		fmt.Printf("Perimeter of the rectagle: %.2f\n", val.Perimeter())
	} else {
		fmt.Println("Not a shape interface")
	}

}
