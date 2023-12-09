package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"in.thewardro.be/rjlee/advent"
)

func main() {
	elves := []int{0}
	for l := range advent.GetLines() {
		if l == "" {
			elves = append(elves, 0)
			continue
		}

		calories, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		elves[len(elves)-1] += calories
	}

	sort.Ints(elves)
	fmt.Println("Max calories:", elves[len(elves)-1])
	fmt.Println("Sum of top 3:", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3])
}
