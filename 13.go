package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pre(s, t string) (string, string, int, int) {
	// fmt.Println("pre", s, t)
	sc := 0
	var i, j int
	for i = 1; i < len(s); i++ {
		if s[i] == '[' {
			sc++
		} else {
			break
		}
	}
	tc := 0
	for j = 1; j < len(t); j++ {
		if t[j] == '[' {
			tc++
		} else {
			break
		}
	}
	if sc > tc && t[j] >= '0' && t[j] <= '9' {
		// wrap [] in t
		jj := j
		for t[jj] >= '0' && t[jj] <= '9' {
			jj++
		}
		return s, t[:j] + strings.Repeat("[", sc-tc) + t[j:jj] + strings.Repeat("]", sc-tc) + t[jj:], sc, tc
	} else if tc > sc && s[i] >= '0' && s[i] <= '9' {
		// wrap [] in s
		ii := i
		for s[ii] >= '0' && s[ii] <= '9' {
			ii++
		}

		return s[:i] + strings.Repeat("[", tc-sc) + s[i:ii] + strings.Repeat("]", tc-sc) + s[ii:], t, sc, tc
	}
	return s, t, sc, tc
}

func pop(s string) (int, string) {
	// fmt.Println("popping from", s)
	k := 0
	if len(s) == 2 {
		return math.MinInt32, ""
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '[' {
			k++
		}
		c := s[i]
		if c == ']' {
			if s[i-1] != '[' {
				panic("cat")
			}
			if len(s) >= i+2 && s[i+1] == ',' {
				return math.MinInt32, s[:i-1] + s[i+2:]
			}
			return math.MinInt32, s[:i-1] + s[i+1:]

		} else if c >= '0' && c <= '9' {
			x := ""
			j := i
			for c >= '0' && c <= '9' {
				// consume whole number
				x += string(c)
				i++
				c = s[i]
			}
			n, _ := strconv.Atoi(x)
			if s[i] == ',' {
				return n, s[:j] + s[i+1:]
			}
			return n, s[:j] + s[i:]
		}
	}
	panic("dog")
}

func solve(s1 string, s2 string) bool {
	n1, n2 := 0, 0
	p1 := s1
	p2 := s2
	var sc, tc int

	for n1 == n2 {
		p1, p2, sc, tc = pre(p1, p2)
		// fmt.Println("post", p1, p2)
		n1, p1 = pop(p1)
		// fmt.Println("popped", n1, "left is", p1)
		n2, p2 = pop(p2)
		// fmt.Println("popped", n2, "left is", p2)
		// fmt.Println("l:", n1, "r:", n2)

		if n1 == math.MinInt32 && n2 == math.MinInt32 {
			// sc > tc means we have more [ in p1
			// so right side will run out of items
			// wrong order
			if sc > tc {
				break
			} else if tc > sc {
				break
			}

		}
	}

	if n1 == n2 && tc > sc || n1 < n2 {
		return true
		// fmt.Println(line, i+1)
	} else {
		return false
		// fmt.Println("incorrect")
	}
	// fmt.Println("\n\n")
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	res := 0

	thingies := []string{"[[2]]", "[[6]]"}

	for _, thing := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		if thing == "" {
			continue
		}
		thingies = append(thingies, thing)
	}

	for i, line := range data {
		pair := strings.Split(line, "\n")
		s1 := pair[0]
		s2 := pair[1]
		if solve(s1, s2) {
			res += i + 1
		}
	}

	fmt.Println(res)

	sort.Slice(thingies, func(i, j int) bool {
		return solve(thingies[i], thingies[j])
	})

	res = 1
	for i, thing := range thingies {
		if thing == "[[2]]" || thing == "[[6]]" {
			res *= i + 1
		}
	}

	fmt.Println(res)

}
