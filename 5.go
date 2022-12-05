package main

import (
	"fmt"
	"os"
	"strings"
)

func print(ys [9][]rune) {
	fmt.Print("[")
	for _, l := range ys {
		fmt.Printf(" %s ", string(l))
	}
	fmt.Println("]")
}

func main() {
	// [J]             [F] [M]
	// [Z] [F]     [G] [Q] [F]
	// [G] [P]     [H] [Z] [S] [Q]
	// [V] [W] [Z] [P] [D] [G] [P]
	// [T] [D] [S] [Z] [N] [W] [B] [N]
	// [D] [M] [R] [J] [J] [P] [V] [P] [J]
	// [B] [R] [C] [T] [C] [V] [C] [B] [P]
	// [N] [S] [V] [R] [T] [N] [G] [Z] [W]
	// 1   2   3   4   5   6   7   8   9

	var xs [9][]rune

	xs[0] = []rune("JZGVTDBN")
	xs[1] = []rune("FPWDMRS")
	xs[2] = []rune("ZSRCV")
	xs[3] = []rune("GHPZJTR")
	xs[4] = []rune("FQZDNJCT")
	xs[5] = []rune("MFSGWPVN")
	xs[6] = []rune("QPBVCG")
	xs[7] = []rune("NPBZ")
	xs[8] = []rune("JPW")

	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	// print(xs)

	for _, l := range data {
		var count, from, to int
		fmt.Sscanf(l, "move %d from %d to %d", &count, &from, &to)
		from -= 1
		to -= 1
		tmp := make([]rune, count)
		copy(tmp, xs[from][:count])

		// uncomment these three lines for part 1
		// for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		// 	tmp[i], tmp[j] = tmp[j], tmp[i]
		// }

		xs[to] = append(tmp, xs[to]...)
		xs[from] = xs[from][count:]
		// print(xs)
	}

	var res []rune

	for _, l := range xs {
		res = append(res, l[0])
	}

	fmt.Println(string(res))
}
