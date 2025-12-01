package main

import (
	"AoC/utils"
	"fmt"
	"math"
	"regexp"
)

func main() {
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day1Part1("day1/day1-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day1Part2("day1/day1-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", day1Part1("day1/day1.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", day1Part2("day1/day1.txt")))
}

func day1Part1(path string) int {
	lines := utils.ReadLines(path)
	dial := 50
	answer := 0

	var re = regexp.MustCompile(`(?m)([LR])(\d*)`)

	for _, line := range lines {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			number := utils.MustParseStringToInt(match[2])
			if match[1] == "L" {
				dial -= number
			}
			if match[1] == "R" {
				dial += number
			}
			if dial%100 == 0 {
				answer++
			}
		}
	}

	return answer
}

func day1Part2(path string) int {
	lines := utils.ReadLines(path)
	dial := 50
	answer := 0

	var re = regexp.MustCompile(`(?m)([LR])(\d*)`)

	for _, line := range lines {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			number := utils.MustParseStringToInt(match[2])
			prevDial := dial
			if match[1] == "L" {
				dial -= number
			}
			if match[1] == "R" {
				dial += number
			}

			// Calculate the number of steps based on the total value (so no % here)
			if dial > prevDial {
				answer += int(math.Floor(float64(dial)/100)) - int(math.Floor(float64(prevDial)/100))
			} else if dial < prevDial {
				answer += int(math.Floor(float64(prevDial-1)/100)) - int(math.Floor(float64(dial-1)/100))
			}
		}
	}

	return answer
}
