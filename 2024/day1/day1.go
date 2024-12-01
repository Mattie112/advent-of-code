package main

import (
	"AoC/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day1Part1("day1/day1-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day1Part2("day1/day1-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", day1Part1("day1/day1.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", day1Part2("day1/day1.txt")))
}

func day1Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	left := make([]int, 1)
	right := make([]int, 1)

	// Loop through the lines and split them into left and right
	for _, line := range lines {
		a := strings.Split(line, "   ")
		left = append(left, utils.MustParseStringToInt(a[0]))
		right = append(right, utils.MustParseStringToInt(a[1]))
	}

	// Sort all the numbers
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the answer using the sorted arrays
	for i := 0; i < len(left); i++ {
		answer += int(math.Abs(float64(left[i] - right[i])))
	}

	return answer
}

func day1Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	left := make([]int, 1)
	numberMap := make(map[int]int)

	// Loop through the lines, save the left part for later (so we only need one split) and count the numbers for the right part
	for _, line := range lines {
		a := strings.Split(line, "   ")
		left = append(left, utils.MustParseStringToInt(a[0]))
		numberMap[utils.MustParseStringToInt(a[1])] += 1
	}

	// Use the left part to calculate the answer with the map of the right part
	for i := 0; i < len(left); i++ {
		answer += left[i] * numberMap[left[i]]
	}

	return answer
}
