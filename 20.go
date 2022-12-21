package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(key, iters int) {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	n := len(data)
	xs := make([]int, n)
	mod := n - 1

	for i, line := range data {
		x, _ := strconv.Atoi(line)
		xs[i] = x * key
	}

	r := ring.New(n)
	index := map[int]*ring.Ring{}
	var zero *ring.Ring

	for i, x := range xs {
		if x == 0 {
			zero = r
		}
		r.Value = x
		index[i] = r
		r = r.Next()
	}

	for j := 0; j < iters; j++ {
		for i := 0; i < n; i++ {
			// remove current element
			r = index[i].Prev()
			el := r.Unlink(1)

			// move
			r.Move(el.Value.(int) % mod).Link(el)
		}
	}

	a := zero.Move(1000).Value.(int)
	b := zero.Move(2000).Value.(int)
	c := zero.Move(3000).Value.(int)

	fmt.Println(a + b + c)
}

func main() {
	solve(1, 1)
	solve(811589153, 10)
}
