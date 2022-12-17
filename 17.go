package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Shape struct {
	points []Point
}

type diff struct {
	i, topY int
}

func moveRight(s *Shape) {
	for _, p := range s.points {
		for _, ss := range b.shapes {
			for _, pp := range ss.points {
				if p.x+1 == pp.x && p.y == pp.y {
					// fmt.Println("blocked right")
					return
				}
			}
		}
		if p.x+1 >= 7 {
			// fmt.Println("out of bounds right")
			return
		}
	}
	for i := range s.points {
		s.points[i].x++
	}
}
func moveLeft(s *Shape) {
	for _, p := range s.points {
		for _, ss := range b.shapes {
			for _, pp := range ss.points {
				if p.x-1 == pp.x && p.y == pp.y {
					// fmt.Println("blocked left")
					return
				}
			}
		}
		if p.x-1 < 0 {
			// fmt.Println("out of bounds left")
			return
		}
	}
	for i := range s.points {
		s.points[i].x--
	}
}
func moveDown(s *Shape) bool {
	for _, p := range s.points {
		for _, ss := range b.shapes {
			for _, pp := range ss.points {
				if p.x == pp.x && p.y-1 == pp.y {
					// fmt.Println("resting position")
					return false
				}
			}
		}
		if p.y-1 < 0 {
			// fmt.Println("reached floor")
			return false
		}
	}
	for i := range s.points {
		s.points[i].y--
	}
	return true
}

func newShape(i int) Shape {
	points := [][]Point{
		[]Point{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		[]Point{
			{1, 0},
			{0, 1},
			{1, 1},
			{2, 1},
			{1, 2},
		},
		[]Point{
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 1},
			{2, 2},
		},
		[]Point{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},
		[]Point{
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		},
	}

	maxY := 0
	for _, s := range b.shapes {
		for _, p := range s.points {
			if p.y+1 > maxY {
				maxY = p.y + 1
			}
		}
	}
	for j := range points[i%5] {
		points[i%5][j].x += 2
		points[i%5][j].y += 3 + maxY
	}

	return Shape{points[i%5]}
}

func print(ss []Shape) {
	maxY := 7
	for _, s := range b.shapes {
		for _, p := range s.points {
			if p.y+1 > maxY {
				maxY = p.y + 1
			}
		}
	}
	for y := maxY; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			for _, s := range ss {
				for _, p := range s.points {
					if p.x == x && p.y == y {
						fmt.Print("#")
						goto next
					}
				}
			}
			fmt.Print(".")
		next:
		}
		fmt.Println()
	}
	fmt.Println()
}

type Board struct {
	shapes []Shape
}

var b = Board{}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Trim(string(input), "\n")

	// init shape pos: minx = 2, miny = 3
	falling := false
	i, j := 0, 0
	var s Shape
	m := make(map[int]map[int]diff)
	maxY := 0
	k := 1_000_000_000_000
	found := -1
	// rem := 0
	lastPart := false
	skippedY := 0
	topY := 0

	for i <= k {
		c := data[j]
		j++
		j %= len(data)
		if !falling {
			s = newShape(i)
			falling = true
			i++
		}
		switch c {
		case '<':
			// fmt.Println("moving left")
			moveLeft(&s)
		case '>':
			// fmt.Println("moving right")
			moveRight(&s)
		}
		if !moveDown(&s) {
			falling = false
			b.shapes = append(b.shapes, s)
			// print(b.shapes)
		}
		// print([]Shape{s})
		topY = 0
		for _, s := range b.shapes {
			for _, p := range s.points {
				if p.y+1 > topY {
					topY = p.y + 1
				}
			}
		}

		if i <= 2022 {
			maxY = topY
			// fmt.Println(maxY)
		} else if i > 2022 && !lastPart {
			// rock type and move pos same => cycle potentially
			// check how many times we need to get to the same pos again
			if found != j {
				if _, ok := m[i%5][j]; ok {
					if found == -1 {
						found = j
					}
					// fmt.Println("idx", i)
					// panic("loop")
				} else if _, ok := m[i%5]; !ok {
					m[i%5] = make(map[int]diff)
				}
				if found == -1 {
					m[i%5][j] = diff{i, topY}
				}
			} else if prev, ok := m[i%5][j]; ok {
				// saw loop for 2nd time
				// fmt.Println("aaa", i)
				di := i - prev.i
				// fmt.Println("prev.i", prev.i, "prev.topY", prev.topY)
				dy := topY - prev.topY
				// fmt.Println("di", di, "dy", dy)
				// fmt.Println("topY", topY)

				cyclesLeft := int(math.Floor(float64((k - i)) / float64(di)))
				skippedY = cyclesLeft * dy
				skippedI := cyclesLeft * di
				i += skippedI
				lastPart = true

				// fmt.Println("cyclesLeft", cyclesLeft)
				// fmt.Println("skippedY", skippedY)
				// fmt.Println("skippedI", skippedI)

				continue
			}
		}
	}

	fmt.Println(maxY)
	fmt.Println(skippedY + topY)
}
