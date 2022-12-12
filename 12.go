package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type tile struct {
	i    int
	j    int
	v    rune
	vis  bool
	dist int
}

func bfs(g []*tile, start, end *tile) int {
	for _, v := range g {
		(*v).vis = false
		(*v).dist = math.MaxInt
	}

	q := []*tile{start}
	start.vis = true
	start.dist = 0

	for len(q) > 0 {
		u := *q[0]
		q = q[1:]
		for _, v := range g {

			if !((*v).i <= u.i+1 && (*v).i >= u.i-1 && (*v).j <= u.j+1 && (*v).j >= u.j-1 &&
				((*v).i == u.i || (*v).j == u.j)) {
				continue
			}

			if !(*v).vis && (*v).v <= u.v+1 {
				// fmt.Println("visiting", v, "from", u)
				(*v).vis = true
				(*v).dist = u.dist + 1
				q = append(q, v)
			}
		}
	}

	return end.dist
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	g := []*tile{}
	var start, end *tile

	for i, line := range data {
		for j, c := range line {
			t := tile{i, j, c, false, math.MaxInt}
			if c == 'S' {
				start = &t
				start.v = 'a'
			} else if c == 'E' {
				end = &t
				end.v = 'z'
			}
			g = append(g, &t)
		}
	}

	d := bfs(g, start, end)

	best := math.MaxInt

	for _, t := range g {
		if (*t).v == 'a' {
			dd := bfs(g, t, end)
			if dd < best {
				best = dd
			}
		}
	}

	fmt.Println(d)
	fmt.Println(best)

}
