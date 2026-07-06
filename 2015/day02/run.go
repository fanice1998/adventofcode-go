package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := simple()
	fmt.Println(ans)
}

func simple() int {
	data, err := readData("data")
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	nums := convertIntSlice(data)
	area := 0
	for _, val := range nums {
		l, w, h := val[0], val[1], val[2]
		area += 2*l*w + 2*w*h + 2*l*h
		switch {
		case l > w && l > h:
			area += w * h
		case w > h && w > l:
			area += h * l
		case h > l && h > w:
			area += l * w
		case h == l && l > w:
			area += l * w
		case h == l && l == w:
			area += l * w
		case l == w && w > h:
			area += l * h
		case h == w && w > l:
			area += l * h
		}
	}

	return area
}

func convertIntSlice(data string) [][]int {
	lines := strings.Split(data, "\n")

	numArr := make([][]int, len(lines))
	for ind, val := range lines {
		numArr[ind] = make([]int, 3)
		nums := strings.Split(val, "x")
		for i, num := range nums {
			v, err := strconv.Atoi(num)
			if err != nil {
				v = 0
			}

			numArr[ind][i] = v
		}
	}

	return numArr
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
