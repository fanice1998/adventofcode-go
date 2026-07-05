package main

import (
	"fmt"
	"os"
)

func main() {
	advanced()
}

func advanced() {
	// read data file
	data, err := readData("data")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	up, down := 0, 0
	ans := 0
	for i, v := range data {
		switch v {
		case '(':
			up++
		case ')':
			down++
		}
		if (up - down) == -1 {
			ans = i
			break
		}
	}
	fmt.Println(ans + 1)

}

func simple() {
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

	os.Exit(0)
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
