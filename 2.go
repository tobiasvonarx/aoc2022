package main

import (
	"fmt"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	input, _ := aocutil.NewInputFromFile("session_id")
	data, _ := input.Strings(2022, 2)
	score := 0
	score2 := 0

	for _, x := range data {
		tmp := strings.Split(x, " ")
		them, me := tmp[0], tmp[1]
		// fmt.Printf("%s %s\n", them, me)

		switch them {
		case "A":
			switch me {
			case "X":
				score += 3
			case "Y":
				score += 6
			}
		case "B":
			switch me {
			case "Y":
				score += 3
			case "Z":
				score += 6
			}
		case "C":
			switch me {
			case "X":
				score += 6
			case "Z":
				score += 3
			}
		}

		switch me {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		switch me {
		case "X":
			switch them {
			case "A":
				score2 += 3
			case "B":
				score2 += 1
			case "C":
				score2 += 2
			}
		case "Y":
			switch them {
			case "A":
				score2 += 1
			case "B":
				score2 += 2
			case "C":
				score2 += 3
			}
			score2 += 3
		case "Z":
			switch them {
			case "A":
				score2 += 2
			case "B":
				score2 += 3
			case "C":
				score2 += 1
			}
			score2 += 6
		}
	}

	// part 1
	fmt.Println(score)

	// part 2
	fmt.Println(score2)
}
