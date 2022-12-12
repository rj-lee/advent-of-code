package main

import (
	"fmt"
	"sort"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type monkey struct {
	items []int

	operator string
	operand2 int
	useOld   bool

	divisor     int
	trueMonkey  int
	falseMonkey int

	inspections int
}

func (m *monkey) inspect(old int) int {
	o2 := m.operand2
	if m.useOld {
		o2 = old
	}
	switch m.operator {
	case "+":
		return old + o2
	case "*":
		return old * o2
	default:
		return old
	}
}

func (m *monkey) test(worryLevel int) bool {
	return worryLevel%m.divisor == 0
}

func (m *monkey) throw(otherMonkey *monkey, item int) {
	m.items = m.items[1:]
	otherMonkey.items = append(otherMonkey.items, item)
}

func monkeyBusiness(rounds int, ridiculous bool) {
	m := &monkey{}
	monkeys := []*monkey{m}
	for l := range advent.GetLines() {
		f := strings.Fields(l)
		if len(f) == 0 {
			m = &monkey{}
			monkeys = append(monkeys, m)
			continue
		}

		switch f[0] {
		case "Starting":
			for _, item := range f[2:] {
				m.items = append(m.items, advent.ParseInt(strings.Trim(item, ",")))
			}
		case "Operation:":
			m.operator = f[4]
			operand2 := f[5]
			if operand2 == "old" {
				m.useOld = true
			} else {
				m.operand2 = advent.ParseInt(operand2)
			}
		case "Test:":
			m.divisor = advent.ParseInt(f[3])
		case "If":
			if f[1] == "true:" {
				m.trueMonkey = advent.ParseInt(f[5])
			} else {
				m.falseMonkey = advent.ParseInt(f[5])
			}
		}
	}

	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= m.divisor
	}

	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				item := m.inspect(m.items[0])
				if !ridiculous {
					item /= 3
				}
				item %= commonDivisor
				m.inspections++
				if m.test(item) {
					m.throw(monkeys[m.trueMonkey], item)
				} else {
					m.throw(monkeys[m.falseMonkey], item)
				}
			}
		}
	}

	sort.Slice(monkeys, func(m1, m2 int) bool {
		return monkeys[m1].inspections > monkeys[m2].inspections
	})

	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)
}

func main() {
	monkeyBusiness(20, false)
	monkeyBusiness(10000, true)
}
