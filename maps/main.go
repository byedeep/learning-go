package main

import "fmt"

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

	fmt.Printf("The age of Walter white is %d \n", Ages["Walter"])

	Ages["Mike"] = 50

	fmt.Printf("The age of Mike is %d\n", Ages["Mike"])
	if _, ok := Ages["Skyler"]; ok {
		fmt.Println("Value exists")
	} else {
		fmt.Println("This doesn't exist")
	}
}
