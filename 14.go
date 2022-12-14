package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func print(m *map[int]map[int]bool) {
	for y := 0; y < 10; y++ {
		for x := 494; x < 504; x++ {
			if _, ok := (*m)[x]; ok {
				if _, ok2 := (*m)[x][y]; ok2 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func move(m *map[int]map[int]bool, x, y int, part bool) bool {

	if part {
		for k, _ := range (*m)[x] {
			if k > y {
				goto notdone
			}
		}
		return true
	}

	if y+1 >= floor {
		if _, ok := (*m)[x]; !ok {
			(*m)[x] = map[int]bool{}
		}
		(*m)[x][y] = true
		return true
	}

notdone:

	if _, ok := (*m)[x][y+1]; ok {
		if _, ok2 := (*m)[x-1][y+1]; ok2 {
			if _, ok3 := (*m)[x+1][y+1]; ok3 {
				// stuck
				if _, ok4 := (*m)[x]; !ok4 {
					(*m)[x] = map[int]bool{}
				}
				(*m)[x][y] = true
				return false
			} else {
				// bottom right
				return move(m, x+1, y+1, part)
			}
		} else {
			// bottom left
			return move(m, x-1, y+1, part)
		}
	} else {
		// move down
		return move(m, x, y+1, part)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var floor int

func solve(part bool) {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	m := map[int]map[int]bool{}

	highestY := 0

	for _, line := range data {
		p := strings.Split(line, " -> ")

		for i := 0; i < len(p)-1; i++ {
			// path from p[i] to p[i+1]
			// fmt.Println(p[i], p[i+1])
			tmp := strings.Split(p[i], ",")
			x1, _ := strconv.Atoi(tmp[0])
			y1, _ := strconv.Atoi(tmp[1])
			tmp = strings.Split(p[i+1], ",")
			x2, _ := strconv.Atoi(tmp[0])
			y2, _ := strconv.Atoi(tmp[1])

			if y1+y2-min(y1, y2) > highestY {
				if y1 > y2 {
					highestY = y1
				} else {
					highestY = y2
				}
			}

			for x := min(x1, x2); x <= x1+x2-min(x1, x2); x++ {
				if _, ok := m[x]; !ok {
					m[x] = map[int]bool{}
				}
				for y := min(y1, y2); y <= y1+y2-min(y1, y2); y++ {
					m[x][y] = true
				}
			}
		}
	}
	floor = highestY + 2

	if part {
		n := 0
		for !move(&m, 500, 0, true) {
			n++
		}
		fmt.Println(n)
	} else {
		n := 0
		ok := false
		for !ok {
			move(&m, 500, 0, false)
			n++
			_, ok = m[500][0]
		}
		fmt.Println(n)
	}

}

func main() {
	solve(true)
	solve(false)
}
