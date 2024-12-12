package main

import "fmt"

type Car struct {
	Brand            string
	Seats            int
	Colour           col
	EngineDislacment float64
}

type col struct {
	Interior string
	Exterior string
}

func main() {

	a1 := Car{"Honda", 2, col{"White", "Black"}, 3}

	fmt.Println("Details of the car 1")
	fmt.Printf("Brand name of the car %s\n", a1.Brand)
	fmt.Printf("Seats in the car %d\n", a1.Seats)
	fmt.Printf("Car's engine is %.1f L\n", a1.EngineDislacment)
	fmt.Printf("Colour of the car's interior %s\n", a1.Colour.Interior)
	fmt.Printf("Colour of the car's exterior %s\n", a1.Colour.Exterior)

}
