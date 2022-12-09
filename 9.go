package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var visited = 0

func visit(vis map[int]map[int]bool, x, y int) {
	if vis[x] == nil {
		vis[x] = make(map[int]bool)
	}

	if _, ok := vis[x][y]; !ok {
		visited++
		vis[x][y] = true
	}

}
func DBG(vis map[int]map[int]bool) {
	min_x, min_y := 0, 0
	max_x, max_y := 0, 0
	for k, v := range vis {
		for kk, _ := range v {
			if k < min_x {
				min_x = k
			}
			if k > max_x {
				max_x = k
			}
			if kk < min_y {
				min_y = kk
			}
			if kk > max_y {
				max_y = kk
			}
		}
	}

	for y := max_y; y >= min_y; y-- {
		for x := min_x; x <= max_x; x++ {
			if _, ok := vis[x][y]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func touching(x1, y1, x2, y2 int) bool {
	return int(math.Abs(float64(x1-x2))) <= 1 && int(math.Abs(float64(y1-y2))) <= 1
}

func solve(k int) {
	visited = 0
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	vis := make(map[int]map[int]bool)
	dirs := map[rune][]int{
		'U': {0, 1},
		'D': {0, -1},
		'L': {-1, 0},
		'R': {1, 0},
	}

	snake := [][]int{}
	for i := 0; i < k; i++ {
		snake = append(snake, []int{0, 0})
	}

	visit(vis, 0, 0)

	for _, line := range data {
		var dir rune
		var steps int

		fmt.Sscanf(line, "%c %d", &dir, &steps)

		// fmt.Println(line)

		dx, dy := dirs[dir][0], dirs[dir][1]

		for i := 0; i < steps; i++ {
			snake[0][0], snake[0][1] = snake[0][0]+dx, snake[0][1]+dy

			for j := 1; j < len(snake); j++ {
				if !touching(snake[j-1][0], snake[j-1][1], snake[j][0], snake[j][1]) {
					if snake[j-1][0] > snake[j][0] {
						snake[j][0]++
					} else if snake[j-1][0] < snake[j][0] {
						snake[j][0]--
					}
					if snake[j-1][1] > snake[j][1] {
						snake[j][1]++
					} else if snake[j-1][1] < snake[j][1] {
						snake[j][1]--
					}
				}
			}
			visit(vis, snake[len(snake)-1][0], snake[len(snake)-1][1])
		}
	}
	// DBG(vis)

	fmt.Println(visited)
}

func main() {
	solve(2)
	solve(10)
}
