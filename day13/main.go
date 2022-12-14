package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"in.thewardro.be/rjlee/advent"
)

func main() {
	lines := []string{}
	for l := range advent.GetLines() {
		if l == "" {
			continue
		}
		lines = append(lines, l)
	}

	packets := []any{}
	sum := 0
	for i := 0; i < len(lines); i += 2 {
		first, second := parse(lines[i]), parse(lines[i+1])
		packets = append(packets, first, second)
		if compare(first, second) <= 0 {
			sum += (i+1)/2 + 1
		}
	}
	fmt.Println(sum)

	packets = append(packets, parse("[[2]]"), parse("[[6]]"))
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) <= 0
	})

	decoderKey := 1
	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			decoderKey *= i + 1
		}
	}
	fmt.Println(decoderKey)
}

func parse(s string) any {
	var v any
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return v
}

func compare(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	if !lok && !rok {
		return int(left.(float64)) - int(right.(float64))
	} else if !lok {
		l = []any{left}
	} else if !rok {
		r = []any{right}
	}

	for i := 0; i < len(l) && i < len(r); i++ {
		if c := compare(l[i], r[i]); c != 0 {
			return c
		}
	}
	return len(l) - len(r)
}
