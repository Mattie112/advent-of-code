package main

import (
	"AoC/utils"
	"fmt"
)

func main() {
	var count int
	count = day1Part1("day1/day1-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day1Part2("day1/day1-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day1Part1("day1/day1.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day1Part2("day1/day1.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day1Part1(path string) int {
	lines := utils.ReadLinesAsInteger(path)
	increasedCount := 0
	prevValue := -1
	for _, c := range lines {
		if prevValue != -1 && c > prevValue {
			increasedCount++
		}
		prevValue = c
	}

	return increasedCount
}

func day1Part2(path string) int {
	lines := utils.ReadLinesAsInteger(path)
	increasedCount := 0
	prevValue := -1
	for i := 0; i < len(lines); i++ {
		if i+1 >= len(lines) || i+2 >= len(lines) {
			// Break if we are at the end of our list
			break
		}
		// Sum the next 3 values
		count := lines[i] + lines[i+1] + lines[i+2]
		if prevValue != -1 && count > prevValue {
			increasedCount++
		}
		prevValue = count
	}

	return increasedCount
}
