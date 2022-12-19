package main

import (
	"fmt"
	"os"
	"strings"
)

func max(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func min(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

var blueprint, oreOre, clayOre, obsOre, obsClay, geoOre, geoObs int

func key(time, ore_, clay_, obs_, geo_, oreR_, clayR_, obsR_, geoR_ int) string {
	return fmt.Sprintf("%d %d %d %d %d %d %d %d %d", time, ore_, clay_, obs_, geo_, oreR_, clayR_, obsR_, geoR_)
}

func revkey(key string) (time, ore_, clay_, obs_, geo_, oreR_, clayR_, obsR_, geoR_ int) {
	fmt.Sscanf(key, "%d %d %d %d %d %d %d %d %d", &time, &ore_, &clay_, &obs_, &geo_, &oreR_, &clayR_, &obsR_, &geoR_)
	return
}

func explore(time, ore_, clay_, obs_, geo_, oreR_, clayR_, obsR_, geoR_ int) int {
	q := []string{key(time, ore_, clay_, obs_, geo_, oreR_, clayR_, obsR_, geoR_)}
	vis := map[string]bool{}
	m := 0

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		var ore, clay, obs, geo, oreR, clayR, obsR, geoR int
		time, ore, clay, obs, geo, oreR, clayR, obsR, geoR = revkey(p)

		m = max(m, geo)
		if time == 0 {
			continue
		}

		oreR = min(oreR, max(oreOre, clayOre, obsOre, geoOre))
		clayR = min(clayR, obsClay)
		obsR = min(obsR, geoObs)

		ore = min(ore, time*max(oreOre, clayOre, obsOre, geoOre)-oreR*(time-1))
		clay = min(clay, time*obsClay-clayR*(time-1))
		obs = min(obs, time*geoObs-obsR*(time-1))

		if _, ok := vis[key(time, ore, clay, obs, geo, oreR, clayR, obsR, geoR)]; ok {
			continue
		}
		vis[key(time, ore, clay, obs, geo, oreR, clayR, obsR, geoR)] = true

		// buy nothing
		q = append(q, key(time-1, ore+oreR, clay+clayR, obs+obsR, geo+geoR, oreR, clayR, obsR, geoR))

		// buy robot
		if ore >= oreOre {
			q = append(q, key(time-1, ore+oreR-oreOre, clay+clayR, obs+obsR, geo+geoR, oreR+1, clayR, obsR, geoR))
		}
		if ore >= clayOre {
			q = append(q, key(time-1, ore+oreR-clayOre, clay+clayR, obs+obsR, geo+geoR, oreR, clayR+1, obsR, geoR))
		}
		if ore >= obsOre && clay >= obsClay {
			q = append(q, key(time-1, ore+oreR-obsOre, clay+clayR-obsClay, obs+obsR, geo+geoR, oreR, clayR, obsR+1, geoR))
		}
		if ore >= geoOre && obs >= geoObs {
			q = append(q, key(time-1, ore+oreR-geoOre, clay+clayR, obs+obsR-geoObs, geo+geoR, oreR, clayR, obsR, geoR+1))
		}

	}

	return m
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	mimi := 0
	mimimi := 1
	for _, line := range data {
		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &blueprint, &oreOre, &clayOre, &obsOre, &obsClay, &geoOre, &geoObs)

		r := explore(24, 0, 0, 0, 0, 1, 0, 0, 0)
		mimi += r * blueprint
		if blueprint <= 3 {
			rr := explore(32, 0, 0, 0, 0, 1, 0, 0, 0)
			mimimi *= rr
		}
	}

	fmt.Println(mimi)
	fmt.Println(mimimi)
}
