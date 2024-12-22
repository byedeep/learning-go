package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("file created")

	Data := bufio.NewWriter(file)
	_, err = Data.WriteString("Hello world\n")
	if err != nil {
		fmt.Println("Error!", err)
		return
	}
	Data.Flush()

	fmt.Println("Data written!")
}
