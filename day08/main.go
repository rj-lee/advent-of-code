package main

import (
	"fmt"

	"in.thewardro.be/rjlee/advent"
)

func main() {
	grid := [][]int{}
	for l := range advent.GetLines() {
		grid = append(grid, make([]int, len(l)))
		for i, c := range l {
			grid[len(grid)-1][i] = int(c - '0')
		}
	}

	visibleCount := 0
	maxScenicScore := 0
	for i := range grid {
		for j := range grid[i] {
			leftVisible := true
			leftDistance := 0
			for k := j - 1; k >= 0; k-- {
				leftDistance++
				if grid[i][k] >= grid[i][j] {
					leftVisible = false
					break
				}
			}

			rightVisible := true
			rightDistance := 0
			for k := j + 1; k < len(grid[i]); k++ {
				rightDistance++
				if grid[i][k] >= grid[i][j] {
					rightVisible = false
					break
				}
			}

			topVisible := true
			topDistance := 0
			for k := i - 1; k >= 0; k-- {
				topDistance++
				if grid[k][j] >= grid[i][j] {
					topVisible = false
					break
				}
			}

			bottomVisible := true
			bottomDistance := 0
			for k := i + 1; k < len(grid); k++ {
				bottomDistance++
				if grid[k][j] >= grid[i][j] {
					bottomVisible = false
					break
				}
			}

			if leftVisible || rightVisible || topVisible || bottomVisible {
				visibleCount++
			}

			scenicScore := leftDistance * rightDistance * topDistance * bottomDistance
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println(visibleCount)
	fmt.Println(maxScenicScore)
}
