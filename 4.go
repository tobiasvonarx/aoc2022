package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	x, y := 0, 0
	for _, line := range data {
		tmp := strings.Split(line, ",")
		a, b := tmp[0], tmp[1]
		tmp = strings.Split(a, "-")
		aa, _ := strconv.Atoi(tmp[0])
		az, _ := strconv.Atoi(tmp[1])
		tmp = strings.Split(b, "-")
		ba, _ := strconv.Atoi(tmp[0])
		bz, _ := strconv.Atoi(tmp[1])

		if (aa >= ba && az <= bz) || (ba >= aa && bz <= az) {
			x += 1
		}
		if (aa >= ba && aa <= bz) || (ba >= aa && ba <= az) {
			y += 1
		}
	}

	// part 1
	fmt.Println(x)

	// part 2
	fmt.Println(y)
}
