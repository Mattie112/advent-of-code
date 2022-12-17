package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
)

func main() {
	var count int
	count = day4Part1("day4/day4-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day4Part1("day4/day4.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day4Part2("day4/day4-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day4Part2("day4/day4.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day4Part1(path string) int {
	lines := utils.ReadLines(path)
	// Named capture groups: https://pkg.go.dev/regexp/syntax
	var re = regexp.MustCompile(`(?m)(?P<firstFrom>\d+)-(?P<firstTo>\d+),(?P<secondFrom>\d+)-(?P<secondTo>\d+)`)
	overlapCount := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		matchesInt := utils.StringSliceToIntSlice(matches[1:]) // First match (the line) cannot be parsed
		matchesInt = append([]int{-1}, matchesInt...)          // To make it easier I re-add the first element

		first := expandInts(matchesInt[re.SubexpIndex("firstFrom")], matchesInt[re.SubexpIndex("firstTo")])
		second := expandInts(matchesInt[re.SubexpIndex("secondFrom")], matchesInt[re.SubexpIndex("secondTo")])

		// Find total overlaps
		if second[0] >= first[0] && second[len(second)-1] <= first[len(first)-1] || first[0] >= second[0] && first[len(first)-1] <= second[len(second)-1] {
			overlapCount++
		}
	}

	return overlapCount
}

// Converts 4-8 into [4,5,6,7,8]
func expandInts(start int, end int) []int {
	output := make([]int, 0)
	for i := start; i <= end; i++ {
		output = append(output, i)
	}
	return output
}

func day4Part2(path string) int {
	lines := utils.ReadLines(path)
	// Named capture groups: https://pkg.go.dev/regexp/syntax
	var re = regexp.MustCompile(`(?m)(?P<firstFrom>\d+)-(?P<firstTo>\d+),(?P<secondFrom>\d+)-(?P<secondTo>\d+)`)
	overlapCount := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		matchesInt := utils.StringSliceToIntSlice(matches[1:]) // First match (the line) cannot be parsed
		matchesInt = append([]int{-1}, matchesInt...)          // To make it easier I re-add the first element

		first := expandInts(matchesInt[re.SubexpIndex("firstFrom")], matchesInt[re.SubexpIndex("firstTo")])
		second := expandInts(matchesInt[re.SubexpIndex("secondFrom")], matchesInt[re.SubexpIndex("secondTo")])

		// Don't find overlaps, just kick-out the things that DON'T overlap ;) (here we don't care about the total overlaps!)
		if second[0] > first[len(first)-1] || second[len(second)-1] < first[0] {
			continue
		}

		if first[0] > second[len(second)-1] || first[len(first)-1] < second[0] {
			continue
		}

		overlapCount++
	}

	return overlapCount
}
