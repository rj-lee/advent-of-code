package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"in.thewardro.be/rjlee/advent"
)

type device struct {
	cycle     int
	register  int
	signalSum int
}

func (d *device) tick() {
	d.cycle++
	if d.cycle%40 == 20 {
		d.signalSum += d.register * d.cycle
	}
	d.draw()
}

func (d *device) draw() {
	c := d.cycle % 40
	pos := c - 1
	if abs(d.register-pos) < 2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if c == 0 {
		fmt.Println()
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	device := device{register: 1}
	for l := range advent.GetLines() {
		f := strings.Fields(l)
		device.tick()

		if f[0] == "addx" {
			device.tick()
			v, err := strconv.Atoi(f[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			device.register += v
		}
	}

	fmt.Println(device.signalSum)
}
