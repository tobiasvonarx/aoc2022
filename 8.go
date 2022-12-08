package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func highestTree(grid [][]int, i, j, di, dj int) int {
	// fmt.Println(i, j, di, dj)

	if di == -1 && i < 0 || di == 1 && i >= len(grid) ||
		dj == -1 && j < 0 || dj == 1 && j >= len(grid[0]) {
		return -1
	}

	return max(grid[i][j], highestTree(grid, i+di, j+dj, di, dj))
}

func viewingDistance(grid [][]int, i, j, di, dj int) int {
	i = max(0, min(i, len(grid)-1))
	j = max(0, min(j, len(grid[0])-1))

	tree := grid[i][j]
	n := 0
	i += di
	j += dj
	for i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		n += 1
		if grid[i][j] >= tree {
			break
		}
		i += di
		j += dj
	}
	return n
}

func isVisible(grid [][]int, i, j int) bool {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		x := highestTree(grid, i+dir[0], j+dir[1], dir[0], dir[1])
		if x < grid[i][j] {
			return true
		}
	}
	return false
}

func scenicScore(grid [][]int, i, j int) int {
	score := 1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		dist := viewingDistance(grid, i, j, dir[0], dir[1])
		// fmt.Println(dist)
		score *= dist
	}
	return score
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	grid := [][]int{}

	for _, line := range data {
		row := []int{}

		for _, cell := range line {
			digit, _ := strconv.Atoi(string(cell))
			row = append(row, digit)
		}
		grid = append(grid, row)
	}

	visibles := 0
	bestScore := 0
	for i, row := range grid {
		for j, _ := range row {
			if isVisible(grid, i, j) {
				// fmt.Printf("Tree at (%d, %d) is visible\n", i, j)
				visibles += 1
			}
			bestScore = max(bestScore, scenicScore(grid, i, j))
		}
	}

	fmt.Println(visibles)
	fmt.Println(bestScore)
}
