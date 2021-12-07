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

// Simply storing all the fish in a slice and doing a count after 80 days
func part1(path string) int {
	lines := utils.ReadLines(path)
	maxPos := 0
	crabs := make([]int, 0)
	for _, line := range lines {
		crabArr := utils.StrArrToIntArr(strings.Split(line, ","))
		for _, crapPos := range crabArr {
			crabs = append(crabs, crapPos)
			if crapPos > maxPos {
				maxPos = crapPos
			}
		}
	}

	leastAmountOfFuel := 3.402823e+38

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

// Storing the fishAmount in a map by age(=timer) instead of storing all the fish
func part2(path string) int {
	lines := utils.ReadLines(path)
	maxPos := 0
	crabs := make([]int, 0)
	for _, line := range lines {
		crabArr := utils.StrArrToIntArr(strings.Split(line, ","))
		for _, crapPos := range crabArr {
			crabs = append(crabs, crapPos)
			if crapPos > maxPos {
				maxPos = crapPos
			}
		}
	}

	leastAmountOfFuel := 3.402823e+38

	for i := 0; i < maxPos; i++ {
		tmpFuel := 0.0
		for _, crab := range crabs {

			diff := math.Abs(float64(crab - i))
			tmpFuel += diff * (diff + 1) / 2

		}
		if tmpFuel < leastAmountOfFuel {
			leastAmountOfFuel = tmpFuel
		}
	}

	return int(leastAmountOfFuel)
}
