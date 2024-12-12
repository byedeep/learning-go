package main

import "fmt"

func main() {

	FirstName := "Walter"
	LastName := "White"
	age := 56
	Networth := 80.7

	fmt.Println(FirstName, LastName)
	fmt.Printf("His age is %d\n", age)
	fmt.Printf("%T \n", age)
	fmt.Printf("His total networth is %.1f Millon dollars\n", Networth)
	fmt.Printf("%T\n", Networth)

}
