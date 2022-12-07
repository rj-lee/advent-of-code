package main

import (
	"fmt"

	"in.thewardro.be/rjlee/advent"
)

func solve(s string, length int) {
	start := 0
	m := map[rune]int{}
	for i, r := range s {
		m[r] += 1
		if len(m) == length {
			start = i
			break
		}
		if i >= length-1 {
			first := rune(s[i-length+1])
			if m[first] > 1 {
				m[first] -= 1
			} else {
				delete(m, first)
			}
		}
	}

	fmt.Println(start + 1)
}

func main() {
	for l := range advent.GetLines() {
		solve(l, 4)
		solve(l, 14)
	}
}
