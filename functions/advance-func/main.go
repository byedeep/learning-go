package main

import "fmt"

func concatter() func(string) string {
	doc := " "
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {

	Concatter := concatter()

	Concatter("hallo,")
	Concatter("how")
	Concatter("are")
	fmt.Print(Concatter("you?"))

}
