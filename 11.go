package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	res := a * b / gcd(a, b)
	for _, i := range integers {
		res = lcm(res, i)
	}

	return res
}

func solve(rounds, div int) {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n\n")

	worryLevels := [][]int{}
	divisors := []int{}

	for _, monkey := range data {
		lines := strings.Split(monkey, "\n")
		for i, line := range lines {
			lines[i] = strings.Trim(line, " ")
		}
		var divisor int
		fmt.Sscanf(lines[3], "Test: divisible by %d", &divisor)
		divisors = append(divisors, divisor)
		items := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		monkeyLevels := []int{}
		for _, x := range items {
			y, _ := strconv.Atoi(x)
			monkeyLevels = append(monkeyLevels, y)
		}
		worryLevels = append(worryLevels, monkeyLevels)
	}

	monkeyLcm := lcm(divisors[0], divisors[1], divisors[2:]...)
	inspectionCnt := make([]int, len(data))

	for i := 0; i < rounds; i++ {
		for _, monkey := range data {
			lines := strings.Split(monkey, "\n")
			for i, line := range lines {
				lines[i] = strings.Trim(line, " ")
			}
			var n int
			fmt.Sscanf(lines[0], "Monkey %d:", &n)
			rawOp := strings.Split(lines[2], "= ")[1]
			expr, _ := govaluate.NewEvaluableExpression(rawOp)
			ctxt := make(map[string]interface{}, 8)
			var divisor, trueBr, falseBr int
			fmt.Sscanf(lines[3], "Test: divisible by %d", &divisor)
			fmt.Sscanf(lines[4], "If true: throw to monkey %d", &trueBr)
			fmt.Sscanf(lines[5], "If false: throw to monkey %d", &falseBr)

			for _, worryLevel := range worryLevels[n] {
				inspectionCnt[n]++
				ctxt["old"] = worryLevel % monkeyLcm
				new, _ := expr.Evaluate(ctxt)
				if f64, ok := new.(float64); ok {
					worryLevel = int(f64)
				} else {
					panic("cat")
				}
				worryLevel /= div
				if worryLevel%divisor == 0 {
					worryLevels[trueBr] = append(worryLevels[trueBr], worryLevel)
				} else {
					worryLevels[falseBr] = append(worryLevels[falseBr], worryLevel)
				}
			}
			worryLevels[n] = []int{}
		}
	}
	// fmt.Println(worryLevels)

	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCnt)))

	monkeyBusiness := inspectionCnt[0] * inspectionCnt[1]

	fmt.Println(monkeyBusiness)
}

func main() {
	solve(20, 3)
	solve(10000, 1)
}
