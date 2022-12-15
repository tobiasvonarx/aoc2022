package main

import (
	"fmt"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	sensors := [][]int{}
	beacons := [][]int{}

	max_dist := 0
	min_x, max_x := 0, 0

	for _, line := range data {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		// fmt.Println(sx, sy, bx, by)
		dist := abs(sx-bx) + abs(sy-by)
		min_x, max_x = min(min_x, sx), max(max_x, sx)
		max_dist = max(max_dist, dist)
		sensors = append(sensors, []int{sx, sy, dist})
		beacons = append(beacons, []int{bx, by})
	}

	min_x, max_x = min_x-max_dist, max_x+max_dist

	res := 0
	for x := min_x; x < max_x; x++ {
		can := true
		already := false
		for _, sensor := range sensors {
			if abs(sensor[0]-x)+abs(sensor[1]-2_000_000) <= sensor[2] {
				can = false
			}
		}
		for _, beacon := range beacons {
			if beacon[0] == x && beacon[1] == 2_000_000 {
				already = true
			}
		}
		if !can && !already {
			res++
		}
	}

	fmt.Println(res)
	k := 4_000_000

	for _, sensor := range sensors {
		// candidates are dist+1 away from some beacon
		for xx := 0; xx <= sensor[2]+1; xx++ {
			yy := sensor[2] + 1 - xx
			for _, sgnx := range []int{1, -1} {
				for _, sgny := range []int{1, -1} {
					x, y := sensor[0]+xx*sgnx, sensor[1]+yy*sgny
					if x > k || x < 0 || y > k || y < 0 {
						continue
					}
					for _, sensor := range sensors {
						if abs(sensor[0]-x)+abs(sensor[1]-y) <= sensor[2] {
							goto skippy
						}
					}
					fmt.Println(x*k + y)
					goto done
				skippy:
				}
			}
		}
	}
done:
}
