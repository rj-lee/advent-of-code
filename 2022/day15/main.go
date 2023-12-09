package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"

	"in.thewardro.be/rjlee/advent"
)

type coord struct {
	x, y int
}

func main() {
	lines := []string{}
	for l := range advent.GetLines() {
		lines = append(lines, l)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	y := 2000000
	yRange := [2]int{math.MaxInt, math.MinInt}
	beacons := map[int]struct{}{}
	for _, l := range lines {
		sensor, beacon := getSensorBeacon(l)
		if beacon.y == y {
			beacons[beacon.x] = struct{}{}
		}
		r := getRangeForY(sensor, beacon, y)
		if len(r) == 0 {
			continue
		}
		if r[0] < yRange[0] {
			yRange[0] = r[0]
		}
		if r[1] > yRange[1] {
			yRange[1] = r[1]
		}
	}
	covered := yRange[1] - yRange[0] + 1
	fmt.Println(covered - len(beacons))
}

func part2(lines []string) {
	for y := 0; y < 4000000; y++ {
		ranges := [][2]int{}
		for _, l := range lines {
			sensor, beacon := getSensorBeacon(l)
			r := getRangeForY(sensor, beacon, y)
			if len(r) == 0 {
				continue
			}
			ranges = append(ranges, [2]int{r[0], r[1]})
		}
		// Sort the ranges by lower bound
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i][0] < ranges[j][0]
		})
		// Consolidate the ranges
		consolidated := [][2]int{}
		for _, r := range ranges {
			length := len(consolidated)
			if length > 0 && r[0] <= consolidated[length-1][1]+1 {
				if r[1] > consolidated[length-1][1] {
					consolidated[len(consolidated)-1][1] = r[1]
				}
			} else {
				consolidated = append(consolidated, r)
			}
		}
		if len(consolidated) == 2 {
			x := consolidated[0][1] + 1
			frequency := x*4000000 + y
			fmt.Println(frequency)
			break
		}
	}
}

func getSensorBeacon(l string) (coord, coord) {
	sensorRegex := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	matches := sensorRegex.FindStringSubmatch(l)
	sensor := coord{advent.ParseInt(matches[1]), advent.ParseInt(matches[2])}
	beacon := coord{advent.ParseInt(matches[3]), advent.ParseInt(matches[4])}
	return sensor, beacon
}

func getRangeForY(sensor coord, beacon coord, y int) []int {
	distance := abs(sensor.x-beacon.x) + abs(sensor.y-beacon.y)
	top := sensor.y - distance
	bottom := sensor.y + distance
	if y < top || y > bottom {
		return []int{}
	}
	dy := abs(sensor.y - y)
	dx := abs(distance - dy)
	left := sensor.x - dx
	right := sensor.x + dx
	return []int{left, right}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
