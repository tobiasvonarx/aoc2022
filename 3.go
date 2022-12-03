package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	priorities := 0
	for _, rucksack := range data {
		middle := len(rucksack) / 2
		// fmt.Println(rucksack)

		hash := make(map[rune]bool)
		for _, item := range rucksack[:middle] {
			hash[item] = true
		}
		for _, item := range rucksack[middle:] {
			if hash[item] {
				var priority int
				if item >= 'a' {
					priority = int(item-'a') + 1
				} else {
					priority = int(item-'A') + 27
				}
				priorities += priority
				break
			}
		}
	}

	// part 1
	fmt.Println(priorities)

	priorities = 0
	for i := 0; i < len(data); i += 3 {
		r1 := data[i]
		r2 := data[i+1]
		r3 := data[i+2]

		hash := make(map[rune]bool)
		hash2 := make(map[rune]bool)
		for _, item := range r1 {
			hash[item] = true
		}
		for _, item := range r2 {
			if hash[item] {
				hash2[item] = true
			}
		}
		for _, item := range r3 {
			if hash2[item] {
				var priority int
				if item >= 'a' {
					priority = int(item-'a') + 1
				} else {
					priority = int(item-'A') + 27
				}
				priorities += priority
				// fmt.Printf("%c\n", item)
				break
			}
		}
	}

	// part 2
	fmt.Println(priorities)

}
