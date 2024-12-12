package main

import "fmt"

func calc(x, y int) (add int, sub int, mul int) {
	return x + y, x - y, x * y
}

func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("can't devide by zero")
	}
	return x / y
}

func main() {

	var add, sub, mul = calc(3, 4)
	var div = divide(3, 4)

	fmt.Printf("The addition is %d\n", add)
	fmt.Printf("The subtraction is %d\n", sub)
	fmt.Printf("The multiplication is %d\n", mul)
	fmt.Printf("The division is %.2f", div)
}
