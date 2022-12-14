package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var count int
	count = day3Part1("day3/day3-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day3Part2("day3/day3-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day3Part1("day3/day3.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day3Part2("day3/day3.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day3Part1(path string) int {
	lines := utils.ReadLines(path)
	prioSum := 0
	for _, line := range lines {
		firstCompartiment := utils.SliceToBooleanMap(strings.Split(line[:len(line)/2], ""))
		secondCompartiment := utils.SliceToBooleanMap(strings.Split(line[len(line)/2:], ""))
		for f := range firstCompartiment {
			if secondCompartiment[f] == true {
				prioSum += getLetterPriority([]rune(f)[0])
			}
		}
	}
	return prioSum
}

func getLetterPriority(input rune) int {
	value := int(input - 'A' - 31)
	if value < 0 {
		value += 58
	}
	return value
}

func day3Part2(path string) int {
	lines := utils.ReadLines(path)
	prioSum := 0
	for i := 0; i < len(lines); i += 3 {
		first := utils.SliceToBooleanMap(strings.Split(lines[i], ""))
		second := utils.SliceToBooleanMap(strings.Split(lines[i+1], ""))
		third := utils.SliceToBooleanMap(strings.Split(lines[i+2], ""))

		for f := range first {
			if second[f] && third[f] {
				prioSum += getLetterPriority([]rune(f)[0])
			}
		}
	}

	return prioSum
}
