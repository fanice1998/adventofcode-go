package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	signStr, err := readData("data")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	signStr = strings.TrimSpace(signStr)

	enterSign := ""
	md5Str := ""

	for i := range 99999999 {
		md5Str = ""
		enterSign = ""

		enterSign = fmt.Sprintf("%s%d", signStr, i)
		hash := md5.Sum([]byte(enterSign))
		md5Str = hex.EncodeToString(hash[:])

		if md5Str[0:5] == "00000" {
			break
		}
	}

	fmt.Println("Part one:")
	fmt.Printf("Sign: %s\n", enterSign)
	fmt.Printf("md5Str: %s\n", md5Str)

	for i := range 99999999999 {
		md5Str = ""
		enterSign = ""

		enterSign = fmt.Sprintf("%s%d", signStr, i)
		hash := md5.Sum([]byte(enterSign))
		md5Str = hex.EncodeToString(hash[:])

		if md5Str[0:6] == "000000" {
			break
		}
	}

	fmt.Println("Part two:")
	fmt.Printf("Sign: %s\n", enterSign)
	fmt.Printf("md5Str: %s\n", md5Str)
}

func readData(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Open file err: %w", err)
	}

	return string(data), nil
}
