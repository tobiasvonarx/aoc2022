package main

import (
	"fmt"
	"os"
	"strings"
)

func dup(a string) bool {
	for i, c := range a {
		for _, cc := range a[i+1:] {
			if c == cc {
				return true
			}
		}
	}
	return false
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")[0]

	// k := 4 for part 1
	k := 14

	for i, _ := range data[k-1:] {
		if !dup(data[i : k+i]) {
			fmt.Println(i + k)
			break
		}
	}
}
