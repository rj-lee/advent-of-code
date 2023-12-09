package main

import (
	"fmt"
	"math"

	"in.thewardro.be/rjlee/advent"
)

type square struct {
	row      int
	col      int
	distance int
}

const (
	S = int('a')
	E = int('z') + 1
)

func findShortestPath(multiStart bool) {
	grid := [][]int{}
	visited := [][]bool{}
	queue := []square{}
	for l := range advent.GetLines() {
		grid = append(grid, []int{})
		visited = append(visited, []bool{})
		i := len(grid) - 1
		for j, c := range l {
			grid[i] = append(grid[i], int(c))
			visited[i] = append(visited[i], false)
			if c == 'S' || (multiStart && c == 'a') {
				grid[i][j] = S
				queue = append(queue, square{i, j, 0})
			}
			if c == 'E' {
				grid[i][j] = E
			}
		}
	}

	min := math.MaxInt
	for len(queue) > 0 {
		source := queue[0]
		queue = queue[1:]

		if grid[source.row][source.col] == E {
			if source.distance < min {
				min = source.distance
			}
			break
		}

		adjacent := [4][2]int{
			{source.row - 1, source.col},
			{source.row + 1, source.col},
			{source.row, source.col - 1},
			{source.row, source.col + 1},
		}

		for _, a := range adjacent {
			row, col := a[0], a[1]
			inGrid := row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
			if inGrid && grid[row][col] <= grid[source.row][source.col]+1 && !visited[row][col] {
				queue = append(queue, square{row, col, source.distance + 1})
				visited[row][col] = true
			}
		}
	}
	fmt.Println(min)
}

func main() {
	findShortestPath(false)
	findShortestPath(true)
}
