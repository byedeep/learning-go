package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		fmt.Println(reader.Text())
	}

}
