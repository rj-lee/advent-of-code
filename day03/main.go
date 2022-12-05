package main

import (
	"fmt"
	"unicode"

	"in.thewardro.be/rjlee/advent"
)

func getPriority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 65 + 27
	}
	return int(r) - 96
}

func part1() {
	sum := 0
	for l := range advent.GetLines() {
		firstHalf, secondHalf := l[:len(l)/2], l[len(l)/2:]

		set := map[rune]bool{}
		for _, r := range firstHalf {
			if _, ok := set[r]; ok {
				continue
			}
			set[r] = true
		}

		for _, r := range secondHalf {
			if _, ok := set[r]; ok {
				sum += getPriority(r)
				break
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0
	group := []string{}
	for l := range advent.GetLines() {
		group = append(group, l)
		if len(group) == 3 {
			badges := map[rune]int{}
			for _, elf := range group {
				set := map[rune]bool{}
				for _, r := range elf {
					if _, ok := set[r]; ok {
						continue
					}
					if badges[r] == 2 {
						sum += getPriority(r)
						break
					}
					set[r] = true
					badges[r]++
				}
			}
			group = []string{}
		}
	}

	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
