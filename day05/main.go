package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type state int

const (
	stateRead state = iota
	stateBlank
	stateMove
)

type stack []rune

func (s *stack) push(r ...rune) {
	*s = append(*s, r...)
}

func (s *stack) pop() rune {
	i := len(*s) - 1
	r := (*s)[i]
	*s = (*s)[:i]
	return r
}

func top(stacks []stack) string {
	t := []rune{}
	for _, stack := range stacks {
		t = append(t, stack[len(stack)-1])
	}
	return string(t)
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

func main() {
	state := stateRead
	stacks1 := []stack{}
	stacks2 := []stack{}
	for l := range advent.GetLines() {
		switch state {
		case stateRead:
			if len(stacks1) == 0 {
				stacks1 = make([]stack, len(l)/4+1)
				stacks2 = make([]stack, len(l)/4+1)
			}

			if rune(l[1]) == '1' {
				state = stateBlank
				break
			}

			for i := 0; i < len(stacks1); i++ {
				pos := i*4 + 1
				if r := rune(l[pos]); r != ' ' {
					stacks1[i] = append(stack{r}, stacks1[i]...)
					stacks2[i] = append(stack{r}, stacks2[i]...)
				}
			}
		case stateBlank:
			state = stateMove
		case stateMove:
			f := strings.Fields(l)
			moves, src, dest := parseInt(f[1]), parseInt(f[3])-1, parseInt(f[5])-1
			sl := len(stacks1[src])
			for i := 0; i < moves; i++ {
				stacks1[dest].push(stacks1[src].pop())
			}
			stack2Moves := stacks2[src][sl-moves : sl]
			stacks2[src] = stacks2[src][:sl-moves]
			stacks2[dest].push([]rune(stack2Moves)...)
		}
	}

	fmt.Println(top(stacks1))
	fmt.Println(top(stacks2))
}
