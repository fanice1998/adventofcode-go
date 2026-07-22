package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"strings"
)

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}

func splitToInt(originStr string) []int {
	splitStr := strings.Split(originStr, ",")
	x, _ := strconv.Atoi(splitStr[0])
	y, _ := strconv.Atoi(splitStr[1])

	return []int{x, y}
}

func main() {
	dataFile, err := readData("data")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	part1 := make([][]int, 1000)
	for i := range part1 {
		part1[i] = make([]int, 1000)
	}

	for line := range strings.SplitSeq(dataFile, "\n") {
		re := regexp.MustCompile(`(.+?)\s+(\d+,\d+)\s+.*?(\d+,\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			action := strings.TrimSpace(match[1])
			startCoord := match[2]
			endCoord := match[3]

			sc := splitToInt(startCoord)
			ec := splitToInt(endCoord)

			switch action {
			case "turn on":
				for x := sc[0]; x <= ec[0]; x++ {
					for y := sc[1]; y <= ec[1]; y++ {
						part1[x][y] = 1
					}
				}
			case "turn off":
				for x := sc[0]; x <= ec[0]; x++ {
					for y := sc[1]; y <= ec[1]; y++ {
						part1[x][y] = 0
					}
				}
			case "toggle":
				for x := sc[0]; x <= ec[0]; x++ {
					for y := sc[1]; y <= ec[1]; y++ {
						if part1[x][y] == 0 {
							part1[x][y] = 1
						} else {
							part1[x][y] = 0
						}
					}
				}
			}
		}
	}

	part2 := make([][]int, 1000)
	for i := range part2 {
		part2[i] = make([]int, 1000)
	}

	for line := range strings.SplitSeq(dataFile, "\n") {
		re := regexp.MustCompile(`(.+?)\s+(\d+,\d+)\s+.*?(\d+,\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			action := strings.TrimSpace(match[1])
			startCoord := match[2]
			endCoord := match[3]

			sc := splitToInt(startCoord)
			ec := splitToInt(endCoord)

			switch action {
			case "turn on":
				for x := sc[0]; x <= ec[0]; x++ {
					for y := sc[1]; y <= ec[1]; y++ {
						part2[x][y] += 1
					}
				}
			case "turn off":
				for x := sc[0]; x <= ec[0]; x++ {
					for y := sc[1]; y <= ec[1]; y++ {
						if part2[x][y] > 0 {
							part2[x][y] -= 1
						}
					}
				}
			case "toggle":
				for x := sc[0]; x <= ec[0]; x++ {
					for y := sc[1]; y <= ec[1]; y++ {
						part2[x][y] += 2
					}
				}
			}
		}
	}

	count := 0
	for i := range 1000 {
		for j := range 1000 {
			if part1[i][j] == 1 {
				count++
			}
		}
	}
	count2 := 0
	for i := range 1000 {
		for j := range 1000 {
			if part2[i][j] > 0 {
				count2 += part2[i][j]
			}
		}
	}

	fmt.Println("part1: ", count)
	fmt.Println("part2: ", count2)
}
