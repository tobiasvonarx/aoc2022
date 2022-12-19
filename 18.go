package main

import (
	"fmt"
	"os"
	"strings"
)

type Cube struct {
	x, y, z int
}

var memo = map[int]map[int]map[int]int{}
var vis map[int]map[int]map[int]bool

func reach_exterior(cs []Cube, x, y, z int) int {
	vis = map[int]map[int]map[int]bool{}
	if res, ok := memo[x][y][z]; ok {
		return res
	}

	viscnt := 0
	q := []*Cube{&Cube{x, y, z}}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		for _, cc := range cs {
			if c.x == cc.x && c.y == cc.y && c.z == cc.z {
				goto next
			}
		}

		if _, ok := vis[c.x][c.y][c.z]; ok {
			continue
		}

		if _, ok := vis[c.x]; !ok {
			vis[c.x] = map[int]map[int]bool{}
		}
		if _, ok := vis[c.x][c.y]; !ok {
			vis[c.x][c.y] = map[int]bool{}
		}
		vis[c.x][c.y][c.z] = true
		viscnt++

		// basically unbounded reachable cubes -> exterior
		if viscnt > 20000 {
			for xx, xxx := range vis {
				for yy, yyy := range xxx {
					for zz, _ := range yyy {
						if _, ok := memo[xx]; !ok {
							memo[xx] = map[int]map[int]int{}
						}
						if _, ok := memo[xx][yy]; !ok {
							memo[xx][yy] = map[int]int{}
						}
						memo[xx][yy][zz] = 1
					}
				}
			}

			return 1
		}

		q = append(q, &Cube{c.x + 1, c.y, c.z})
		q = append(q, &Cube{c.x - 1, c.y, c.z})
		q = append(q, &Cube{c.x, c.y + 1, c.z})
		q = append(q, &Cube{c.x, c.y - 1, c.z})
		q = append(q, &Cube{c.x, c.y, c.z + 1})
		q = append(q, &Cube{c.x, c.y, c.z - 1})
	next:
	}
	if _, ok := memo[x]; !ok {
		memo[x] = map[int]map[int]int{}
	}
	if _, ok := memo[x][y]; !ok {
		memo[x][y] = map[int]int{}
	}
	for xx, xxx := range vis {
		for yy, yyy := range xxx {
			for zz, _ := range yyy {
				if _, ok := memo[xx]; !ok {
					memo[xx] = map[int]map[int]int{}
				}
				if _, ok := memo[xx][yy]; !ok {
					memo[xx][yy] = map[int]int{}
				}

				memo[xx][yy][zz] = 0
			}
		}
	}
	return 0
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	cs := []Cube{}

	for _, line := range data {
		var a, b, c int
		fmt.Sscanf(line, "%d,%d,%d", &a, &b, &c)
		cs = append(cs, Cube{a, b, c})
	}

	sa := 0
	esa := 0
	for _, c := range cs {
		adjacent := 0
		for _, cc := range cs {
			if c.x == cc.x && c.y == cc.y && c.z == cc.z {
				continue
			}

			if (c.x == cc.x && c.y == cc.y && c.z == cc.z+1) ||
				(c.x == cc.x && c.y == cc.y && c.z == cc.z-1) ||
				(c.x == cc.x && c.y == cc.y+1 && c.z == cc.z) ||
				(c.x == cc.x && c.y == cc.y-1 && c.z == cc.z) ||
				(c.x == cc.x+1 && c.y == cc.y && c.z == cc.z) ||
				(c.x == cc.x-1 && c.y == cc.y && c.z == cc.z) {
				adjacent++
			}
		}
		sa += 6 - adjacent

		esa += reach_exterior(cs, c.x, c.y, c.z+1)
		esa += reach_exterior(cs, c.x, c.y, c.z-1)
		esa += reach_exterior(cs, c.x, c.y+1, c.z)
		esa += reach_exterior(cs, c.x, c.y-1, c.z)
		esa += reach_exterior(cs, c.x+1, c.y, c.z)
		esa += reach_exterior(cs, c.x-1, c.y, c.z)
	}
	fmt.Println(sa)
	fmt.Println(esa)

}
