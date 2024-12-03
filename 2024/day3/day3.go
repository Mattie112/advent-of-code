package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day1Part1("day3/day3-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day1Part2("day3/day3-test-part2.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", day1Part1("day3/day3.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", day1Part2("day3/day3.txt")))
}

func day1Part1(path string) int {
	return solve(path, regexp.MustCompile(`(?m)(mul\(((?P<x>\d+),(?P<y>\d+))\))`))
}

func day1Part2(path string) int {
	return solve(path, regexp.MustCompile(`(?m)(mul\(((?P<x>\d+),(?P<y>\d+))\)|(do\(\))|(don't\(\)))`))
}

func solve(path string, re *regexp.Regexp) int {
	lines := utils.ReadLines(path)
	answer := 0

	do := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "don't()" {
				do = false
				continue
			}
			if match[0] == "do()" {
				do = true
				continue
			}

			if do {
				answer += utils.MustParseStringToInt(match[re.SubexpIndex("x")]) * utils.MustParseStringToInt(match[re.SubexpIndex("y")])
			}
		}
	}

	return answer
}
