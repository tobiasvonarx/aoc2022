package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

func main() {
	input, _ := aocutil.NewInputFromFile("session_id")
	data, _ := input.Strings(2022, 1)

	var calories []int
	tmpSum := 0

	for _, x := range data {
		if x == "" {
			calories = append(calories, tmpSum)
			tmpSum = 0
		}
		itemCalories, _ := strconv.Atoi(x)
		tmpSum += itemCalories
	}

	calories = append(calories, tmpSum)

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	// part 1
	fmt.Println(calories[0])

	// part 2
	fmt.Println(calories[0] + calories[1] + calories[2])
}
