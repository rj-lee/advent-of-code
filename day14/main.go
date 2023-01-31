package main

import (
	"fmt"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type coord struct {
	x, y int
}

func parseCoords(s string) coord {
	cs := strings.Split(s, ",")
	x, y := advent.ParseInt(cs[0]), advent.ParseInt(cs[1])
	return coord{x, y}
}

type cave map[coord]bool

func (c *cave) fall(sand *coord) {
	if down := (coord{sand.x, sand.y + 1}); !(*c)[down] {
		(*c)[down] = true
		(*c)[*sand] = false
		*sand = down
	} else if downLeft := (coord{sand.x - 1, sand.y + 1}); !(*c)[downLeft] {
		(*c)[downLeft] = true
		(*c)[*sand] = false
		*sand = downLeft
	} else if downRight := (coord{sand.x + 1, sand.y + 1}); !(*c)[downRight] {
		(*c)[downRight] = true
		(*c)[*sand] = false
		*sand = downRight
	} else {
		*sand = coord{-1, -1}
	}
}

func main() {
	cave1, cave2 := cave{}, cave{}
	maxy := 0
	for l := range advent.GetLines() {
		coords := strings.Split(l, " -> ")
		for i := 0; i < len(coords)-1; i++ {
			start := parseCoords(coords[i])
			end := parseCoords(coords[i+1])
			dx := end.x - start.x
			dy := end.y - start.y
			for x, y := start.x, start.y; x != end.x+sign(dx) || y != end.y+sign(dy); x, y = x+sign(dx), y+sign(dy) {
				cave1[coord{x, y}] = true
				cave2[coord{x, y}] = true
				if y > maxy {
					maxy = y
				}
			}
		}
	}

	part1(cave1, maxy)
	part2(cave2, maxy)
}

func part1(cave cave, maxy int) {
	sand := coord{500, 0}
	count := 0
	for sand.y < maxy {
		cave.fall(&sand)
		if sand == (coord{-1, -1}) {
			sand = coord{500, 0}
			cave[sand] = true
			count++
		}
	}
	fmt.Println(count)
}

func part2(cave cave, maxy int) {
	floor := maxy + 2
	minx := 500 - floor - 3
	maxx := 500 + floor + 3
	for x := minx; x <= maxx; x++ {
		cave[coord{x, floor}] = true
	}

	sand := coord{500, 0}
	count := 1
	for true {
		start := sand
		cave.fall(&sand)
		if sand == (coord{-1, -1}) {
			if start == (coord{500, 0}) {
				break
			}
			sand = coord{500, 0}
			cave[sand] = true
			count++
		}
	}
	fmt.Println(count)
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}
