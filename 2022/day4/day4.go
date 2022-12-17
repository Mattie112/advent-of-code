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

		// Find total overlaps (split into 2 for readability)
		if matchesInt[re.SubexpIndex("secondFrom")] >= matchesInt[re.SubexpIndex("firstFrom")] && matchesInt[re.SubexpIndex("secondTo")] <= matchesInt[re.SubexpIndex("firstTo")] {
			overlapCount++
		} else if matchesInt[re.SubexpIndex("firstFrom")] >= matchesInt[re.SubexpIndex("secondFrom")] && matchesInt[re.SubexpIndex("firstTo")] <= matchesInt[re.SubexpIndex("secondTo")] {
			overlapCount++
		}
	}

	return overlapCount
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

		// Don't find overlaps, just kick-out the things that DON'T overlap ;) (here we don't care about the total overlaps!)
		if matchesInt[re.SubexpIndex("secondFrom")] > matchesInt[re.SubexpIndex("firstTo")] || matchesInt[re.SubexpIndex("secondTo")] < matchesInt[re.SubexpIndex("firstFrom")] {
			continue
		}

		if matchesInt[re.SubexpIndex("firstFrom")] > matchesInt[re.SubexpIndex("secondTo")] || matchesInt[re.SubexpIndex("firstTo")] < matchesInt[re.SubexpIndex("secondFrom")] {
			continue
		}

		overlapCount++
	}

	return overlapCount
}
