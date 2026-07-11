package main

import (
	"fmt"
	"os"
)

type roadMap struct {
	x    int
	y    int
	data map[int]map[int]int
}

func main() {
	data, err := readData("data")
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	santa := roadMap{
		x:    0,
		y:    0,
		data: map[int]map[int]int{},
	}
	gift := 0
	for _, v := range data {
		santa.runMap(v)
	}

	for x := range santa.data {
		for range santa.data[x] {
			gift++
		}
	}

	fmt.Println("Part one: ", gift)

	// -------------------
	santa = roadMap{
		x:    0,
		y:    0,
		data: map[int]map[int]int{},
	}
	santa_bot := roadMap{
		x:    0,
		y:    0,
		data: map[int]map[int]int{},
	}
	gift = 0

	option := true
	for _, v := range data {
		if option {
			santa.runMap(v)
		} else {
			santa_bot.runMap(v)
		}
		option = !option
	}
	for x := range santa.data {
		for y := range santa.data[x] {
			if _, ok := santa_bot.data[x][y]; !ok {
				gift++
			}
		}
	}
	for x := range santa_bot.data {
		for range santa_bot.data[x] {
			gift++
		}
	}

	fmt.Println("Part two: ", gift)
}

func (s *roadMap) runMap(v rune) {
	switch string(v) {
	case "^":
		s.y++
		s.checkMap()
	case "v":
		s.y--
		s.checkMap()
	case ">":
		s.x++
		s.checkMap()
	case "<":
		s.x--
		s.checkMap()
	}
}

func (s *roadMap) checkMap() {
	if _, ok := s.data[s.x]; !ok {
		s.data[s.x] = map[int]int{}
	}

	s.data[s.x][s.y]++
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
