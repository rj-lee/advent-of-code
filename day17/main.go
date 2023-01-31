package main

import (
	"fmt"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type coord struct{ x, y int }
type rock []coord
type chamber map[coord]bool
type key struct {
	gas       int
	shape     int
	signature [100]coord // Use an arbitrary size so that it is hashable
}
type state struct {
	shape     int
	maxHeight int
}
type cache map[key]state

var shapes = []rock{
	{{2, 0}, {3, 0}, {4, 0}, {5, 0}},
	{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}},
	{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},
	{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
	{{2, 0}, {3, 0}, {2, 1}, {3, 1}},
}

func newRock(shape int, y int) *rock {
	r := rock{}
	for _, c := range shapes[shape] {
		r = append(r, coord{c.x, c.y + y})
	}
	return &r
}

func (r *rock) moveLeft() {
	for _, c := range *r {
		if c.x == 0 {
			return
		}
	}
	for i := range *r {
		(*r)[i].x -= 1
	}
}

func (r *rock) moveRight() {
	for _, c := range *r {
		if c.x == 6 {
			return
		}
	}
	for i := range *r {
		(*r)[i].x += 1
	}
}

func (r *rock) moveVertical(dy int) {
	for i := range *r {
		(*r)[i].y += dy
	}
}

func (r *rock) intersects(c chamber) bool {
	for _, rc := range *r {
		if c[rc] {
			return true
		}
	}
	return false
}

func (c *chamber) add(r rock) {
	for _, rc := range r {
		(*c)[rc] = true
	}
}

func (c *chamber) signature() [100]coord {
	maxY := 0
	for cc := range *c {
		if cc.y > maxY {
			maxY = cc.y
		}
	}
	sig := [100]coord{}
	for cc := range *c {
		if maxY-cc.y < 30 { // Use an arbitrary number that is large enough
			sig[len(sig)-1] = coord{cc.x, maxY - cc.y}
		}
	}
	return sig
}

func (c *chamber) print() {
	maxY := 0
	for cc := range *c {
		if cc.y > maxY {
			maxY = cc.y
		}
	}
	for y := maxY; y > 0; y-- {
		for x := 0; x < 7; x++ {
			if (*c)[coord{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

const rounds = 1000000000000

func main() {
	input := strings.Split(advent.GetInput(), "")
	chamber := chamber{}
	for i := 0; i < 7; i++ {
		chamber[coord{i, 0}] = true
	}
	cache := cache{}
	t, i, top, cycles := 0, 0, 0, 0
	for t < rounds {
		rock := newRock(t%5, top+4)
		for {
			if input[i] == "<" {
				rock.moveLeft()
				if rock.intersects(chamber) {
					rock.moveRight()
				}
			} else {
				rock.moveRight()
				if rock.intersects(chamber) {
					rock.moveLeft()
				}
			}
			i = (i + 1) % len(input)
			rock.moveVertical(-1)
			if rock.intersects(chamber) {
				rock.moveVertical(1)
				chamber.add(*rock)
				// chamber.print()
				for _, c := range *rock {
					if c.y > top {
						top = c.y
					}
				}
				key := key{i, t % 5, chamber.signature()}
				if state, ok := cache[key]; ok && t >= 2022 {
					dy := top - state.maxHeight
					dt := t - state.shape
					amount := (rounds - t) / dt
					cycles += amount * dy
					t += amount * dt
				}
				cache[key] = state{t, top}
				break
			}
		}
		t++
		if t == 2022 {
			fmt.Println(top)
		}
	}
	fmt.Println(top + cycles)
}
