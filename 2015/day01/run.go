package main

import (
	"fmt"
	"os"
)

func main() {
	// read data file
	data, err := readData("data")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	up, down := 0, 0
	for _, v := range data {
		switch v {
		case '(':
			up++
		case ')':
			down++
		}
	}
	fmt.Println(up - down)
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
