package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type knot struct {
	x, y int
}

func (k *knot) move(direction string) {
	switch direction {
	case "U":
		k.y++
	case "D":
		k.y--
	case "L":
		k.x--
	case "R":
		k.x++
	}
}

func (k *knot) follow(leader knot) {
	dx := leader.x - k.x
	dy := leader.y - k.y
	if abs(dx)+abs(dy) > 2 { // diagonal
		k.x += dx / abs(dx)
		k.y += dy / abs(dy)
	} else if abs(dx) == 2 { // horizontal
		k.x += dx / abs(dx)
	} else if abs(dy) == 2 { // vertical
		k.y += dy / abs(dy)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	rope := [10]knot{}
	visited1 := map[knot]struct{}{}
	visited2 := map[knot]struct{}{}
	for l := range advent.GetLines() {
		f := strings.Fields(l)
		direction := f[0]
		distance, err := strconv.Atoi(f[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < distance; i++ {
			rope[0].move(direction)
			for i := 1; i < len(rope); i++ {
				rope[i].follow(rope[i-1])
			}
			visited1[rope[1]] = struct{}{}
			visited2[rope[9]] = struct{}{}
		}
	}

	fmt.Println(len(visited1))
	fmt.Println(len(visited2))
}
