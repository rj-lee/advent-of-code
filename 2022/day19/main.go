package main

import (
	"fmt"
	"regexp"

	"in.thewardro.be/rjlee/advent"
)

var regex = regexp.MustCompile(`\d+`)

type blueprint struct {
	ore      int
	clay     int
	obsidian struct{ ore, clay int }
	geode    struct{ ore, obsidian int }
}

type resources [4]int // ore, clay, obsidian, geode

func (b blueprint) getCost(robotType int) resources {
	costs := []resources{
		{b.ore, 0, 0, 0},
		{b.clay, 0, 0, 0},
		{b.obsidian.ore, b.obsidian.clay, 0, 0},
		{b.geode.ore, 0, b.geode.obsidian, 0},
	}
	return costs[robotType]
}

// func (b blueprint) need(robotType int, robots resources) bool {
// 	if robotType == 0 {
// 		if robots[0] >= b.ore {
// 			return false
// 		}
// 	} else if robotType == 1 {
// 		if robots[1] >= b.clay {
// 			return false
// 		}
// 	} else if robotType == 2 {
// 	}
// 	return false
// }

func (r *resources) subtract(r2 resources) {
	for i, c := range r2 {
		r[i] -= c
	}
}

func main() {
	blueprints := []blueprint{}
	for l := range advent.GetLines() {
		f := regex.FindAllString(l, -1)
		b := blueprint{
			ore:      advent.ParseInt(f[1]),
			clay:     advent.ParseInt(f[2]),
			obsidian: struct{ ore, clay int }{advent.ParseInt(f[3]), advent.ParseInt(f[4])},
			geode:    struct{ ore, obsidian int }{advent.ParseInt(f[5]), advent.ParseInt(f[6])},
		}
		blueprints = append(blueprints, b)
	}
	quality := 0
	for i, b := range blueprints {
		robots := resources{1, 0, 0, 0}
		ores := resources{0, 0, 0, 0}
		efficiency := solve(b, robots, ores, 24)
		fmt.Println(i+1, efficiency)
		q := (i + 1) * efficiency
		quality += q
	}
	fmt.Println(quality)
}

var cache = map[string]int{}

func solve(blueprint blueprint, robots resources, ores resources, time int) int {
	key := fmt.Sprintf("%v%v%v%v", blueprint, robots, ores, time)
	if g, ok := cache[key]; ok {
		return g
	}

	if time <= 0 {
		return ores[3] // return the geodes
	}

	maxGeodes := 0

	for robotType := range robots {
		cost := blueprint.getCost(robotType)
		cost.subtract(ores)
		timeNeeded := 1
		for _, c := range cost {
			if c > timeNeeded {
				timeNeeded = c
			}
		}

		nextOres := ores
		for i, robotCount := range robots {
			nextOres[i] += robotCount * timeNeeded
		}
		nextOres[robotType] -= cost[robotType]

		nextRobots := robots
		nextRobots[robotType]++

		geodes := solve(blueprint, nextRobots, nextOres, time-timeNeeded)
		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	cache[key] = maxGeodes
	return maxGeodes
}
