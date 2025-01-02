package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var Rows int
	fmt.Println("How many people's data you want to store?")
	fmt.Scan(&Rows)
	data := make([][]string, Rows)
	for i := range data {
		data[i] = make([]string, 4)
	}

	fmt.Println("Please enter the data to save")
	for i := 0; i < Rows; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("Enter data for the person %d, %d: ", i+1, j+1)
			fmt.Scan(&data[i][j])
		}
	}

	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the file")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, record := range data {
		line := ""
		for i, field := range record {
			line += field
			if i < len(record)-1 {
				line += ","
			}

		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing the file")
			return
		}
	}
	writer.Flush()

}
