package main

import (
	"fmt"
	"os"
	"strings"
)

type NiceString struct {
	data string
	nice bool
}

func main() {
	dataFile, err := readData("data")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	count := 0
	for line := range strings.SplitSeq(dataFile, "\n") {
		str := NiceString{
			data: line,
			nice: true,
		}

		str.checkThreeVowels()
		str.checkTwiceLetter()
		str.checkContain()

		if !str.nice {
			continue
		}

		if str.nice {
			count++
		}
	}

	fmt.Println("Part one")
	fmt.Println(count)

	count = 0
	for line := range strings.SplitSeq(dataFile, "\n") {
		str := NiceString{
			data: line,
			nice: true,
		}

		str.doubleWords()
		str.intervalWord()

		if !str.nice {
			continue
		}

		if str.nice {
			count++
		}
	}

	fmt.Println("Part two")
	fmt.Println(count)
}

func (s *NiceString) intervalWord() {
	if !s.nice {
		return
	}

	for idx := range len(s.data) - 2 {
		if s.data[idx] == s.data[idx+2] {
			s.nice = true
			return
		}
	}

	s.nice = false
}

func (s *NiceString) doubleWords() {
	if !s.nice {
		return
	}

	for idx := range len(s.data) - 3 {
		checkWords := s.data[idx : idx+2]
		remainWords := s.data[idx+2:]
		if strings.Contains(remainWords, checkWords) {
			s.nice = true
			return
		}
	}

	s.nice = false
}

func (s *NiceString) checkThreeVowels() {
	if !s.nice {
		return
	}
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, vowel := range vowels {
		count += strings.Count(s.data, vowel)
		if count >= 3 {
			s.nice = true
			return
		}
	}

	s.nice = false
}

func (s *NiceString) checkTwiceLetter() {
	var tmp rune
	if !s.nice {
		return
	}

	for _, v := range s.data {
		switch tmp {
		case 0:
			tmp = v
		case v:
			s.nice = true
			return
		default:
			tmp = v
		}
	}

	s.nice = false
}

func (s *NiceString) checkContain() {
	if !s.nice {
		return
	}

	var ForbiddenWords = []string{"ab", "cd", "pq", "xy"}

	for _, v := range ForbiddenWords {
		if strings.Contains(s.data, v) {
			s.nice = false
			return
		}
	}

	s.nice = true
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
