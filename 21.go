package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func eval(v1, v2 int, op string) int {
	switch op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	default:
		panic("cat")
	}
}

var m = map[string][]string{}

func rec(n string, humn int) int {
	expr := m[n]
	if n == "humn" && humn >= 0 {
		return humn
	}
	if len(expr) == 1 {
		v, _ := strconv.Atoi(expr[0])
		return v
	} else {
		expr1 := rec(expr[0], humn)
		expr2 := rec(expr[2], humn)

		return eval(expr1, expr2, expr[1])
	}

}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	for _, line := range data {
		s := strings.Split(line, ": ")
		m[s[0]] = strings.Split(s[1], " ")
	}

	res := rec("root", -1)
	fmt.Println(res)

	lhs := m["root"][0]
	rhs := m["root"][2]

	goal := rec(rhs, 0)

	if goal != rec(rhs, 1) {
		goal = rec(lhs, 0)
		if goal != rec(lhs, 1) {
			panic("cat")
		}
		lhs, rhs = rhs, lhs
	}
	// lhs humn dependent, rhs not
	// change humn for lhs to match rhs

	lo := math.MinInt64 >> 5
	hi := math.MaxInt64 >> 5
	// lo := 0
	// hi := 1 << 4
	var humn int
	for lo < hi {
		humn = lo>>1 + hi>>1
		diff := goal - rec(lhs, humn)
		if diff == 0 {
			break
		} else if diff > 0 {
			hi = humn
		} else {
			lo = humn
		}
	}
	fmt.Println(humn)
}
