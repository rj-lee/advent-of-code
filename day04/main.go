package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type elfRange = [2]int

func stringToRange(s string) elfRange {
	parts := strings.Split(s, "-")
	low, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	high, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return [2]int{low, high}
}

func contains(r1 elfRange, r2 elfRange) bool {
	return r1[0] <= r2[0] && r1[1] >= r2[1]
}

func overlaps(r1 elfRange, r2 elfRange) bool {
	return r1[1] >= r2[0] && r1[0] <= r2[1]
}

func main() {
	totalContains := 0
	totalOverlaps := 0
	for l := range advent.GetLines() {
		elves := strings.Split(l, ",")
		elf1, elf2 := stringToRange(elves[0]), stringToRange(elves[1])
		if contains(elf1, elf2) || contains(elf2, elf1) {
			totalContains++
		}
		if overlaps(elf1, elf2) || overlaps(elf2, elf1) {
			totalOverlaps++
		}
	}

	fmt.Println(totalContains)
	fmt.Println(totalOverlaps)
}
