package main

import (
	"AoC/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	//fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day3Part1("day3/day3-test.txt")))
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", day3Part1("day3/day3.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day3Part2("day3/day3-test.txt")))
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", day3Part2("day3/day3.txt")))
}

func day3Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	for _, line := range lines {
		batteries := strings.Split(line, "")

		// First find the position (and value) of the highest joltage
		highestJoltagePosition := 0
		highestJoltage := 0
		for i := 0; i < len(batteries)-1; i++ {
			joltage := utils.MustParseStringToInt(batteries[i])
			if joltage > highestJoltage {
				highestJoltage = joltage
				highestJoltagePosition = i
			}
			if highestJoltage == 9 {
				break
			}
		}

		secondHighestJoltage := 0
		// Now start there and search for the next highest value
		for i := highestJoltagePosition + 1; i < len(batteries); i++ {
			joltage := utils.MustParseStringToInt(batteries[i])
			if joltage > secondHighestJoltage {
				secondHighestJoltage = joltage
			}
			if secondHighestJoltage == 9 {
				break
			}
		}
		answer += utils.MustParseStringToInt(fmt.Sprintf("%d%d", highestJoltage, secondHighestJoltage))
	}

	return answer
}

func day3Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	for _, line := range lines {
		batteries := strings.Split(line, "")

		turnOn := make(map[int]int)
		amount := 0
		lowestPosInUse := math.MaxInt32
		valueOfLowestPos := 0
		// First; find the position (and value) of the highest joltage
		for amount < 12 {
			highestJoltagePosition := 0
			highestJoltage := 0

			for i := 0; i < len(batteries); i++ {
				if turnOn[i] != 0 {
					// We already have this one turned on
					continue
				}
				joltage := utils.MustParseStringToInt(batteries[i])
				if joltage > highestJoltage {
					// Make sure that we never select something with a LOWER value than the value of the first element as that will always lower the number

					if i > lowestPosInUse || joltage > valueOfLowestPos {
						highestJoltage = joltage
						highestJoltagePosition = i
					}

				}
			}
			turnOn[highestJoltagePosition] = highestJoltage
			if highestJoltagePosition < lowestPosInUse {
				lowestPosInUse = highestJoltagePosition
				valueOfLowestPos = highestJoltage
			}
			amount++
		}

		// Now create the number of all turned on items
		turnOnString := ""
		keys := make([]int, 0, len(turnOn))
		for k := range turnOn {
			keys = append(keys, k)
		}

		sort.Ints(keys)

		for _, k := range keys {
			v := turnOn[k]
			if v != 0 {
				turnOnString += fmt.Sprintf("%d", v)
			}
		}
		fmt.Println(turnOnString)
		answer += utils.MustParseStringToInt(turnOnString)
		turnOn = make(map[int]int)
	}

	return answer
}
