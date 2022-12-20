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

func (b blueprint) canBuild(robot int, ores resources) (bool, resources) {
	ore, clay, obsidian := ores[0], ores[1], ores[2]
	if robot == 0 {
		if ore >= b.ore {
			return true, resources{b.ore, 0, 0, 0}
		}
	} else if robot == 1 {
		if ore >= b.clay {
			return true, resources{b.clay, 0, 0, 0}
		}
	} else if robot == 2 {
		if ore >= b.obsidian.ore && clay >= b.obsidian.clay {
			return true, resources{ore - b.obsidian.ore, clay - b.obsidian.clay, 0, 0}
		}
	} else if robot == 3 {
		if ore >= b.geode.ore && obsidian >= b.geode.obsidian {
			return true, resources{ore - b.geode.ore, 0, obsidian - b.geode.obsidian, 0}
		}
	}
	return false, resources{0, 0, 0, 0}
}

func (r *resources) subtract(cost resources) {
	for i, c := range cost {
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

	if time == 0 {
		return ores[3] // return the geodes
	}

	maxGeodes := 0
	nextOres := resources{}
	for i, r := range robots {
		nextOres[i] = ores[i] + r
	}

	// try building a robot
	for i := range robots {
		if ok, cost := blueprint.canBuild(i, ores); ok {
			nextRobots := robots
			nextRobots[i]++
			nextOresMinusCost := nextOres
			nextOresMinusCost.subtract(cost)
			geodes := solve(blueprint, nextRobots, nextOresMinusCost, time-1)
			if geodes > maxGeodes {
				maxGeodes = geodes
			}
		}
	}

	// don't build any robots
	geodes := solve(blueprint, robots, nextOres, time-1)
	if geodes > maxGeodes {
		maxGeodes = geodes
	}

	cache[key] = maxGeodes
	return maxGeodes
}
