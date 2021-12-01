package main

import (
	"AoC/utils"
	"fmt"
	"strconv"
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
	lines, err := utils.ReadLines(path)
	if err != nil {
		fmt.Println("No lines")
	}

	increasedCount := 0
	prevValue := -1
	for _, s := range lines {
		c, _ := strconv.Atoi(s)
		if prevValue != -1 && c > prevValue {
			increasedCount++
		}
		prevValue = c
	}

	return increasedCount
}

func day1Part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		fmt.Println("No lines")
	}

	increasedCount := 0
	prevValue := -1
	for i := 0; i < len(lines); i++ {
		if i+1 >= len(lines) || i+2 >= len(lines) {
			break
		}
		first, _ := strconv.Atoi(lines[i])
		second, _ := strconv.Atoi(lines[i+1])
		third, _ := strconv.Atoi(lines[i+2])

		count := first + second + third

		if prevValue != -1 && count > prevValue {
			increasedCount++
		}

		prevValue = count
	}

	return increasedCount
}
