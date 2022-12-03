package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// input, _ := os.ReadFile("input.txt")
	input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	fmt.Println(data)
}
