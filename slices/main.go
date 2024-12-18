package main

import "fmt"

func main() {

	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	Slice := arr[1:4]
	fmt.Println(Slice)
	Slice1 := []string{"I", "dream", "my", "painting", "and", "I", "paint", "my", "dream"}
	fmt.Println(Slice1)
	fmt.Printf("The length of the slice is %d\n", len(Slice1))

	for i, e := range Slice1 {
		fmt.Printf("The index is : %d and the element is : %s\n", i, e)
	}

	for _, i := range Slice {
		fmt.Printf("Element = %d\n", i)
	}

	Slice[0] = 6
	Slice[1] = 7
	Slice[2] = 8
	fmt.Println(Slice)
}
