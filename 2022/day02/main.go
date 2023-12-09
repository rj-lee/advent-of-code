package main

import (
	"fmt"

	"in.thewardro.be/rjlee/advent"
)

func part1() {
	// Map the elf's move to the player's move to the outcome of the game
	// For example, elf: A (Rock), player: X (Rock), outcome: 3 (Draw)
	outcomes := map[rune]map[rune]int{
		'A': {'X': 3, 'Y': 6, 'Z': 0},
		'B': {'X': 0, 'Y': 3, 'Z': 6},
		'C': {'X': 6, 'Y': 0, 'Z': 3},
	}

	choices := map[rune]int{
		'X': 1, // Rock
		'Y': 2, // Paper
		'Z': 3, // Scissors
	}

	score := 0
	for l := range advent.GetLines() {
		elf := rune(l[0])
		player := rune(l[2])
		score += choices[player] + outcomes[elf][player]
	}

	fmt.Println(score)
}

func part2() {
	// Map the elf's move to the outcome of the game to the player's move's score
	// For example, elf: A (Rock), outcome: Y (Draw), player: 1 (Rock)
	choices := map[rune]map[rune]int{
		'A': {'X': 3, 'Y': 1, 'Z': 2},
		'B': {'X': 1, 'Y': 2, 'Z': 3},
		'C': {'X': 2, 'Y': 3, 'Z': 1},
	}

	outcomes := map[rune]int{
		'X': 0, // Lose
		'Y': 3, // Draw
		'Z': 6, // Win
	}

	score := 0
	for l := range advent.GetLines() {
		elf := rune(l[0])
		outcome := rune(l[2])
		score += outcomes[outcome] + choices[elf][outcome]
	}

	fmt.Println(score)
}

func main() {
	part1()
	part2()
}
