package main

import (
	"fmt"
	"regexp"

	"in.thewardro.be/rjlee/advent"
)

type comparable interface{ string }
type set[T comparable] map[T]struct{}

func (s *set[T]) add(e T) {
	(*s)[e] = struct{}{}
}

func (s *set[T]) remove(e T) {
	delete(*s, e)
}

func (s *set[T]) copy() set[T] {
	c := set[T]{}
	for e := range *s {
		c.add(e)
	}
	return c
}

var valves = set[string]{}
var rates = map[string]int{}
var distances = map[string]map[string]int{}

func main() {
	for l := range advent.GetLines() {
		f := regexp.MustCompile(`[A-Z]{2}|\d+`).FindAllString(l, -1)
		valve := f[0]
		valves.add(valve)
		rate := advent.ParseInt(f[1])
		if rate > 0 {
			rates[valve] = rate
		}
		tunnels := f[2:]
		for _, t := range tunnels {
			if distances[valve] == nil {
				distances[valve] = map[string]int{}
			}
			distances[valve][t] = 1
		}
	}

	// Floyd-Warshall
	for _, p := range product(valves, valves, valves) {
		k, i, j := p[0], p[1], p[2]

		// Default value
		if distances[i][j] == 0 {
			distances[i][j] = 1000
		}
		if distances[i][k] == 0 {
			distances[i][k] = 1000
		}
		if distances[k][j] == 0 {
			distances[k][j] = 1000
		}

		distances[i][j] = min(distances[i][j], distances[i][k]+distances[k][j])
	}

	toVisit := mapToSet(rates)
	fmt.Println(solve(30, "AA", toVisit, false))
	fmt.Println(solve(26, "AA", toVisit, true))
}

var cache = map[string]int{}

func solve(minutes int, current string, toVisit set[string], elephant bool) int {
	key := fmt.Sprintf("%d:%s:%v:%v", minutes, current, toVisit, elephant)
	if cache[key] > 0 {
		return cache[key]
	}

	if minutes == 0 {
		return 0
	}

	max := 0
	for v := range toVisit {
		if distances[current][v] < minutes {
			m := minutes - distances[current][v] - 1
			vs := toVisit.copy()
			vs.remove(v)
			p := rates[v]*m + solve(m, v, vs, elephant)
			if p > max {
				max = p
			}
		}
	}
	if elephant {
		p := solve(26, "AA", toVisit, false)
		if p > max {
			max = p
		}
	}

	cache[key] = max
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mapToSet[T comparable](m map[T]int) set[T] {
	s := set[T]{}
	for k := range m {
		s.add(k)
	}
	return s
}

func product[T comparable](sets ...map[T]struct{}) [][3]T {
	var res [][3]T
	for k := range sets[0] {
		for i := range sets[1] {
			for j := range sets[2] {
				res = append(res, [3]T{k, i, j})
			}
		}
	}
	return res
}
