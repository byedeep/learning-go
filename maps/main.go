package main

import (
	"fmt"
)

func main() {

	Ages := map[string]int{
		"Walter": 50,
		"Jesse":  24,
		"Saul":   48,
		"Hank":   40,
	}

	for key, value := range Ages {
		fmt.Println(key, value)
	}
	fmt.Println("")

	fmt.Printf("The age of Walter white is %d \n", Ages["Walter"])

	Ages["Mike"] = 50

	fmt.Printf("The age of Mike is %d\n", Ages["Mike"])
	if _, ok := Ages["Skyler"]; ok {
		fmt.Println("Value exists")
	} else {
		fmt.Println("This doesn't exist")
	}

	ShoppingList := make(map[string]map[string]int)

	VegetableList := map[string]int{
		"Broccoli":    2,
		"Cauliflower": 3,
		"Carrot":      2,
	}

	FruitList := map[string]int{
		"Mango": 2,
		"Apple": 1,
	}
	ShoppingList["Vegetables"] = VegetableList
	ShoppingList["Fruits"] = FruitList

	fmt.Println("")
	fmt.Println("Shopping list")
	fmt.Println("")

	for value := range ShoppingList {
		fmt.Println("Category", value)
		fmt.Println("Category details", ShoppingList[value])
	}
}
