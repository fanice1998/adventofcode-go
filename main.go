package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: go run main.go (YYYY) (DD)")
		return
	}

	year := args[1]
	day, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("WTF!!! please enter day number")
		return
	}
	strDay := ""
	if day < 10 {
		strDay = fmt.Sprintf("day0%d", day)
	} else {
		strDay = fmt.Sprintf("day%d", day)
	}

	// check year directory
	if !dirExists(year) {
		fmt.Printf("Not find year %s\n", args[1])
		return
	}

	// check day directory
	targetFile := filepath.Join(year, strDay)
	if !dirExists(targetFile) {
		fmt.Printf("Not find target dir, %s\n", targetFile)
	}

	// run file
	runGoFile(targetFile)
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func runGoFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file not exists")
	}

	cmd := exec.Command("go", "run", "run.go")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Executing: %s/run.go\n", path)
	return cmd.Run()
}
