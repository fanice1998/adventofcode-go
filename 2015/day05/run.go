package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dataFile, err := readData("data")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	count := 0
	for line := range strings.SplitSeq(dataFile, "\n") {
		ok := false
		ok = checkThreeVowels(line)
		if !ok {
			continue
		}

		ok = checkTwiceLetter(line)
		if !ok {
			continue
		}

		ok = checkContain(line)
		if !ok {
			continue
		}

		if ok {
			count++
			ok = false
			fmt.Println(line, checkThreeVowels(line), checkTwiceLetter(line), checkContain(line), count)
		}
	}

	fmt.Println("Part one")
	fmt.Println(count)

}

func checkThreeVowels(checkStr string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, vowel := range vowels {
		count += strings.Count(checkStr, vowel)
		if count >= 3 {
			return true
		}
	}
	return false
}

func checkTwiceLetter(checkStr string) bool {
	var tmp rune

	for _, v := range checkStr {
		switch tmp {
		case 0:
			tmp = v
		case v:
			return true
		default:
			tmp = v
		}
	}

	return false
}

func checkContain(checkStr string) bool {
	var ForbiddenWords = []string{"ab", "cd", "pq", "xy"}

	for _, v := range ForbiddenWords {
		if strings.Contains(checkStr, v) {
			return false
		}
	}

	return true
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
