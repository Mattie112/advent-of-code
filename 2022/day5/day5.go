package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var count int
	count = day5Part1("day5/day5-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day5Part1("day5/day5.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day5Part2("day5/day5-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = day5Part2("day5/day5.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day5Part1(path string) int {
	lines := utils.ReadLines(path)

	stacks := make(map[int][]string)

	// First look for the line that has the amount of columns
	columnAmount := 0
	moves := false
	var regColumnNumbers = regexp.MustCompile(`(?m)\d+(?:   )+`)
	var regCrate = regexp.MustCompile(`(?m)([\w]| {3})\]*`)
	for _, line := range lines {
		if moves {
			fmt.Println(line)
		}

		if regColumnNumbers.MatchString(line) {
			columns := strings.Split(line, "   ")
			columnAmount = utils.MustParseStringToInt(columns[len(columns)-1])
			moves = true
		}

		if !moves {
			// We start with filling our columns
			matches := regCrate.FindAllStringSubmatch(line, -1)
			for i, match := range matches {
				fmt.Println(match, "found at index", i)
				stacks[i+1] = append(stacks[i+1], match[1])
			}
		}

	}

	_ = columnAmount
	printStacks(stacks)
	return 0
}

func printStacks(input map[int][]string) {

	// First find highest
	highest := 0
	for _, crates := range input {
		if len(crates) > highest {
			highest = len(crates)
		}
	}
	fmt.Println(fmt.Sprintf("highest: %d", highest))

	for i := 0; i < highest; i++ {
		for j := 0; j < len(input); j++ {
			crates := input[j+1]
			crate := crates[i]
			if crate == "   " {
				fmt.Print("---")
			} else {
				fmt.Print("[" + crate + "]")
			}
		}
		fmt.Println("")
	}

}

func day5Part2(path string) int {

	return 0
}
