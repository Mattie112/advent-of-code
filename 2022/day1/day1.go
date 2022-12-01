package main

import (
	"AoC/utils"
	"fmt"
	"math"
)

func main() {
	var count int
	count = day1Part1("day1/day1-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day1Part2("day1/day1-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day1Part1("day1/day1.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day1Part2("day1/day1.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day1Part1(path string) int {
	lines := utils.ReadLines(path)

	maxValue := 0
	counter := 0
	for _, line := range lines {
		if line == "" {
			maxValue = int(math.Max(float64(maxValue), float64(counter)))
			counter = 0
			continue
		}
		counter += utils.MustParseStringToInt(line)
	}

	return maxValue
}

func day1Part2(path string) int {
	lines := utils.ReadLines(path)
	lines = append(lines, "")

	maxValue := []int{0, 0, 0}
	counter := 0
	for _, line := range lines {
		if line == "" {

			for i, value := range maxValue {
				if counter > value {
					maxValue = append(maxValue[:i+1], maxValue[i:]...)
					maxValue[i] = counter
					maxValue = maxValue[0:3]
					counter = 0
				}
			}

			counter = 0
			continue
		}
		counter += utils.MustParseStringToInt(line)
	}

	return maxValue[0] + maxValue[1] + maxValue[2]
}
