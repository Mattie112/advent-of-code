package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

func main() {
	var count int
	count = day4Part1("day4/day4-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day4Part2("day4/day4-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day4Part1("day4/day4.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day4Part2("day4/day4.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day4Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	for _, line := range lines {
		line = line[strings.Index(line, ": ")+2:]

		winningStr := strings.Split(line, " | ")[0]
		winning := strings.Split(winningStr, " ")
		numbersStr := strings.Split(line, " | ")[1]
		numbers := strings.Split(numbersStr, " ")

		linePoints := 0
		for _, number := range numbers {
			for _, win := range winning {
				if number == win && number != "" && win != "" {
					if linePoints == 0 {
						linePoints = 1
					} else {
						linePoints *= 2
					}
				}
			}
		}
		answer += linePoints
	}

	return answer
}

func day4Part2(path string) int {
	answer := 0

	return answer
}
