package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day7/day7-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day7/day7.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day7/day7-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day7/day7.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)
	crabs := utils.StrArrToIntArr(strings.Split(lines[0], ","))
	maxPos := utils.GetMaxFromArr(crabs)

	leastAmountOfFuel := math.MaxFloat64

	for i := 0; i < maxPos; i++ {
		tmpFuel := 0.0
		for _, crab := range crabs {
			tmpFuel += math.Abs(float64(crab - i))
		}
		if tmpFuel < leastAmountOfFuel {
			leastAmountOfFuel = tmpFuel
		}
	}

	return int(leastAmountOfFuel)
}

func part2(path string) int {
	lines := utils.ReadLines(path)
	crabs := utils.StrArrToIntArr(strings.Split(lines[0], ","))
	maxPos := utils.GetMaxFromArr(crabs)

	leastAmountOfFuel := math.MaxFloat64

	for i := 0; i < maxPos; i++ {
		tmpFuel := 0.0
		for _, crab := range crabs {
			// For example 5 steps = 1+2+3+4+5 == 15 And that is 5 * (5+1) / 2
			steps := math.Abs(float64(crab - i))
			tmpFuel += steps * (steps + 1) / 2
		}
		if tmpFuel < leastAmountOfFuel {
			leastAmountOfFuel = tmpFuel
		}
	}

	return int(leastAmountOfFuel)
}
