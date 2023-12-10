package main

import (
	"fmt"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type coord struct{ x, y int }
type chamber map[coord]rune
type rock struct {
	coords []coord
	height int
}
type key struct {
	heights [7]int
	shape   int
	gas     int
}
type state struct {
	height  int
	stopped bool
}

var shapes = [][]coord{
	[]coord{{2, 0}, {3, 0}, {4, 0}, {5, 0}},
	[]coord{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}},
	[]coord{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},
	[]coord{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
	[]coord{{2, 0}, {3, 0}, {2, 1}, {3, 1}},
}
var moves = []string{}

func newRock(coords []coord) *rock {
	r := &rock{coords: coords}
	r.height = r.coords[0].y
	return r
}

func (r *rock) spawn(chamber *chamber) {
	for _, c := range r.coords {
		(*chamber)[c] = '@'
	}
}

func (r *rock) move(chamber *chamber, dx, dy int) bool {
	for _, c := range r.coords {
		newCoord := coord{c.x + dx, c.y + dy}
		occupied := (*chamber)[newCoord] == '#'
		outOfBounds := newCoord.y <= 0 || newCoord.x < 0 || newCoord.x > 6
		if occupied || outOfBounds {
			return false
		}
	}

	for i, c := range r.coords {
		delete(*chamber, c)
		newCoord := coord{c.x + dx, c.y + dy}
		(*chamber)[newCoord] = '@'
		r.coords[i] = newCoord
	}
	return true
}

func (r *rock) rest(chamber *chamber) {
	for _, c := range r.coords {
		(*chamber)[c] = '#'
	}
}

func (c *chamber) print(maxY int) {
	for i := maxY + 7; i >= 0; i-- {
		fmt.Printf("|")
		for x := 0; x < 7; x++ {
			if r, ok := (*c)[coord{x, i}]; ok {
				fmt.Printf("%c", r)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
	fmt.Println()
}

func solve(rounds int) {
	chamber := chamber{}
	maxY := 0

	cache := map[key]int{}
	shapeI := 0
	inputI := 0
	lenShapes := len(shapes)
	lenMoves := len(moves)

	for i := 0; i < rounds; i++ {
		shape := i % lenShapes
		rock := newRock()

		rock.spawn(&chamber)
		for {
			gas := rune(input[inputI%len(input)])
			inputI++
			dx := -1
			if gas == '>' {
				dx = 1
			}
			rock.move(&chamber, dx, 0)
			fall := rock.move(&chamber, 0, -1)
			if !fall {
				break
			}
		}
		rock.rest(&chamber)
		for _, c := range rock.coords {
			if c.y > maxY {
				maxY = c.y
			}
		}
	}
	fmt.Println(maxY)
}

func main() {
	input := advent.GetInput()
	moves = strings.Split(input, "")
	solve(2022)
}
