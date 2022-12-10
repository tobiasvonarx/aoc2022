package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cursor = 0

func pixel(x int) {
	// fmt.Printf("%d call of pixel, x has value %d\n", cursor, x)

	if x-1 <= cursor%40 && cursor%40 <= x+1 {
		// lit
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	cursor++

	if cursor%40 == 0 {
		fmt.Println()
	}

}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	vals := []int{1}

	x := 1

	for _, line := range data {
		ins := strings.Split(line, " ")
		if ins[0] == "noop" {
			// 1 cycle
			pixel(x)
			vals = append(vals, x)
		} else if ins[0] == "addx" {
			// 2 cycles
			amt, _ := strconv.Atoi(ins[1])
			pixel(x)
			vals = append(vals, x)
			pixel(x)
			x += amt
			vals = append(vals, x)
		} else {
			panic("cat")
		}
	}
	fmt.Println()

	sum := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		sum += i * vals[i-1]
	}
	fmt.Println(sum)
}
