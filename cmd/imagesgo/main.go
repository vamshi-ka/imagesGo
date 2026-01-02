package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {
	// Input to image path
	inputPath := flag.String("input", "", "path to image")
	flag.Parse()

	// 1 input validation
	// 1.1 is input given?
	if *inputPath == "" {
		fmt.Println(Red + "Error: Missing input file" + Reset)
		fmt.Println(Yellow + "Usage: go run main.go -input=photo.jpg" + Reset)
		os.Exit(1)
	}

	// 1.2 is it a valid path?
	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		fmt.Println(Red + "Error: Input file does not exist" + Reset)
	}

	fmt.Println(Green + "Valid file. Loading image" + Reset)

}
