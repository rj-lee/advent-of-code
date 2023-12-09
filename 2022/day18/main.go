package main

import (
	"fmt"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type cube [3]int

func (c *cube) sides() []cube {
	x, y, z := c[0], c[1], c[2]
	return []cube{
		{x + 1, y, z},
		{x - 1, y, z},
		{x, y + 1, z},
		{x, y - 1, z},
		{x, y, z + 1},
		{x, y, z - 1},
	}
}

func main() {
	cubes := map[cube]bool{}
	for l := range advent.GetLines() {
		f := strings.Split(l, ",")
		x, y, z := advent.ParseInt(f[0]), advent.ParseInt(f[1]), advent.ParseInt(f[2])
		c := cube{x, y, z}
		cubes[c] = true
	}
	area := 0
	for c := range cubes {
		for _, s := range c.sides() {
			if !cubes[s] {
				area++
			}
		}
	}
	fmt.Println(area)

	seen := map[cube]bool{}
	stack := []cube{{-1, -1, -1}} // BFS starting from outside
	for len(stack) > 0 {
		// Pop in Go is ugly
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, s := range c.sides() {
			if !cubes[s] && !seen[s] { // If we haven't seen it before
				notTooFar := true
				for _, n := range s {
					if n < -1 || n > 25 { // Pick a number that is big enough to cover the input
						notTooFar = false
						break
					}
				}
				if notTooFar {
					stack = append(stack, s) // Check its neighbors
				}
			}
		}
		seen[c] = true
	}
	// Count the number of neighbors of each cube that we have seen from the outside
	sum := 0
	for c := range cubes {
		for _, s := range c.sides() {
			if seen[s] {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
