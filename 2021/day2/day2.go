package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var count int
	count = day2Part1("day2/day2-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day2Part2("day2/day2-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day2Part1("day2/day2.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day2Part2("day2/day2.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day2Part1(path string) int {
	lines := utils.ReadLines(path)

	re := regexp.MustCompile(`([a-z]+) (\d+)`)

	horizontal := 0
	depth := 0

	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		direction := res[1]
		amount, _ := strconv.Atoi(res[2])

		switch direction {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	return horizontal * depth
}

func day2Part2(path string) int {
	lines := utils.ReadLines(path)

	re := regexp.MustCompile(`([a-z]+) (\d+)`)

	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		direction := res[1]
		amount, _ := strconv.Atoi(res[2])

		switch direction {
		case "forward":
			horizontal += amount
			depth += amount * aim
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	return horizontal * depth
}
