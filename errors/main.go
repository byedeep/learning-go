package main

import "fmt"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("Error!")
	}
	return x / y, nil
}

func main() {
	div, err := divide(1, 0)
	if err != nil {
		fmt.Print("Error!")
	} else {
		fmt.Println(div)
	}
}
