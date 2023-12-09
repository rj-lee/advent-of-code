package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

func main() {
	fs := map[string]int{}
	parents := []string{}
	current := ""
	for l := range advent.GetLines() {
		switch rune(l[0]) {
		case '$':
			cmd := strings.Split(l, " ")
			if cmd[1] == "cd" {
				if cmd[2] == ".." {
					i := len(parents) - 1
					current = parents[i]
					parents = parents[:i]
				} else {
					if current != "" {
						parents = append(parents, current)
					}
					current = strings.Join(parents, "") + cmd[2]
					fs[current] = 0
				}
			}
		case 'd':
		default:
			size, err := strconv.Atoi(strings.Split(l, " ")[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fs[current] += size
			for _, p := range parents {
				fs[p] += size
			}
		}
	}

	sum := 0
	available := 70_000_000 - fs["/"]
	needed := 30_000_000 - available
	min_needed := math.MaxInt
	for _, size := range fs {
		if size <= 100_000 {
			sum += size
		}
		if size >= needed && size < min_needed {
			min_needed = size
		}
	}

	fmt.Println(sum)
	fmt.Println(min_needed)
}
