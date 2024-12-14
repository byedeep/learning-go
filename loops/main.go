package main

import "fmt"

func sheeps() {
	for i := 1; i <= 10; i++ {
		if i == 1 {
			fmt.Printf("%d sheep\n", i)
			continue
		}
		fmt.Printf("%d sheeps\n", i)
	}

	for i := 11; ; i++ {
		fmt.Printf("%d sheeps\n", i)
		if i >= 20 {
			break
		}
	}
	i := 21
	for i <= 30 {
		fmt.Printf("%d sheeps\n", i)
		i++
	}
}

func main() {

	sheeps()

}
