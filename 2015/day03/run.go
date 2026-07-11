package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := readData("data")
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	roadMap := map[int]map[int]int{}
	x, y := 0, 0
	gift := 0
	for _, v := range data {
		switch string(v) {
		case "^":
			y++
			checkMap(roadMap, x, y)
		case "v":
			y--
			checkMap(roadMap, x, y)
		case ">":
			x++
			checkMap(roadMap, x, y)
		case "<":
			x--
			checkMap(roadMap, x, y)
		}
	}

	for x := range roadMap {
		for range roadMap[x] {
			gift++
		}
	}

	fmt.Println(gift)
}

func checkMap(roadMap map[int]map[int]int, x, y int) {
	if _, ok := roadMap[x]; !ok {
		roadMap[x] = map[int]int{}
	}

	roadMap[x][y]++
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
