package main

import (
	"fmt"
	"os"
	"strings"
)

type Valve struct {
	flow_rate int
	leads_to  []string
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dp = make(map[string]map[int]map[int]map[bool]int)

func rec(status string, v, time int, elephant bool) int {
	if res, ok := dp[status][v][time][elephant]; ok {
		return res
	}

	if time == 0 {
		if elephant {
			return rec(status, m["AA"], 26, false)
		} else {
			return 0
		}
	}

	pressure := 0
	closed := status[v] == '0'
	flow := valves[v].flow_rate

	// open
	if closed && flow > 0 {
		status_ := status[:v] + "1" + status[v+1:]
		pressure = max(pressure, (time-1)*flow+rec(status_, v, time-1, elephant))
	}

	// move
	for _, vv := range valves[v].leads_to {
		pressure = max(pressure, rec(status, m[vv], time-1, elephant))
	}

	if _, ok := dp[status]; !ok {
		dp[status] = make(map[int]map[int]map[bool]int)
	}
	if _, ok := dp[status][v]; !ok {
		dp[status][v] = make(map[int]map[bool]int)
	}
	if _, ok := dp[status][v][time]; !ok {
		dp[status][v][time] = make(map[bool]int)
	}
	dp[status][v][time][elephant] = pressure

	return pressure
}

var valves []Valve
var m = make(map[string]int)

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	for i, line := range data {
		spl := strings.Split(line, "; ")

		var name string
		var flow_rate int

		fmt.Sscanf(spl[0], "Valve %s has flow rate=%d", &name, &flow_rate)

		leads_to := strings.Split(spl[1], " ")[4:]
		for j, v := range leads_to {
			leads_to[j] = strings.Trim(v, ",")
		}

		m[name] = i

		v := Valve{flow_rate, leads_to}
		valves = append(valves, v)
	}

	res := rec(strings.Repeat("0", len(valves)), m["AA"], 30, false)
	fmt.Println(res)
	res = rec(strings.Repeat("0", len(valves)), m["AA"], 26, true)
	fmt.Println(res)

}
