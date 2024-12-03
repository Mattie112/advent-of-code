package main

import (
	"AoC/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day1Part1("day2/day2-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day1Part2("day2/day2-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", day1Part1("day2/day2.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", day1Part2("day2/day2.txt")))
}

func day1Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	for _, line := range lines {
		levelsStr := strings.Split(line, " ")
		levelsInt := make([]int, len(levelsStr))

		for i, levelStr := range levelsStr {
			level, _ := strconv.Atoi(levelStr)
			levelsInt[i] = level
		}

		if isValid(levelsInt) {
			answer++
		}
	}

	return answer
}

func day1Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	linesCount := len(lines)
	tryAgain := make([][]int, 1)

linesForEach:
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineCount := i
		levelsStr := strings.Split(line, " ")
		levelsInt := make([]int, len(levelsStr))

		for j, levelStr := range levelsStr {
			level, _ := strconv.Atoi(levelStr)
			levelsInt[j] = level
		}

		if !isValid(levelsInt) {
			if lineCount < linesCount-1 {
				// Only do this the first time, I am lazy so I fill another array with these
				tryAgain = append(tryAgain, levelsInt)
			}
			continue linesForEach
		}

		answer++
	}

	// Go through this list and only for these try to get rid of one element to see if it is valid then
	for _, levels := range tryAgain {
		for j := range levels {
			// Remove current element from levels and check isValid again
			levelsTmp := make([]int, len(levels))
			copy(levelsTmp, levels)
			levelsTmp = append(levelsTmp[:j], levelsTmp[j+1:]...)
			if isValid(levelsTmp) {
				//log.Printf("Found a valid line %+v", levelsTmp)
				answer++
				break
			}
		}
	}

	return answer
}

func isValid(levelsInt []int) bool {

	// check for the sorting
	if levelsInt[0] > levelsInt[1] {
		prevLevel := levelsInt[0]
		for _, level := range levelsInt {
			if level > prevLevel {
				return false
			}
			prevLevel = level
		}
	} else if levelsInt[0] < levelsInt[1] {
		prevLevel := levelsInt[0]
		for _, level := range levelsInt {
			if level < prevLevel {
				return false
			}
			prevLevel = level
		}
	} else {
		return false
	}

	// Check for the max deviation of 3 (and not equal)
	prevLevel := levelsInt[0]
	for i := 1; i < len(levelsInt); i++ {
		level := levelsInt[i]
		x := float64(level) - float64(prevLevel)
		if math.Abs(x) > 3 || level == prevLevel {
			return false
		}
		prevLevel = level
	}
	return true
}
