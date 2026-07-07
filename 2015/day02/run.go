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

		// part 1
		// area += 2*l*w + 2*w*h + 2*l*h
		// m1, m2 := getMinSecond(val)
		// area += m1 * m2

		// part 2
		area += l * w * h
		m1, m2 := getMinSecond(val)
		area += 2 * (m1 + m2)

	}

	return area
}

func getMinSecond(val []int) (int, int) {
	l, w, h := val[0], val[1], val[2]
	switch {
	case l > w && l > h:
		return w, h
	case w > h && w > l:
		return h, l
	case h > l && h > w:
		return l, w
	case h == l && l > w:
		return l, w
	case h == l && l == w:
		return l, w
	case l == w && w > h:
		return l, h
	case h == w && w > l:
		return l, h
	}

	return 0, 0
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
