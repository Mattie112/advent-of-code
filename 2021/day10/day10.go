package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day10/day10-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day10/day10.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day10/day10-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day10/day10.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

var (
	openingBrackets = map[string]bool{
		"(": true,
		"<": true,
		"[": true,
		"{": true,
	}

	closingBrackets = map[string]string{
		")": "(",
		">": "<",
		"]": "[",
		"}": "{",
	}

	scoresPart1 = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	// Yes technically we score the closing brackets for part2 but this saves me a lookup (otherwise I first need to find the opposite bracket)
	scoresPart2 = map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
)

func part1(path string) int {
	lines := utils.ReadLines(path)
	score := 0

	for _, line := range lines {
		brackets := strings.Split(line, "")
		stack := make([]string, 0)

		for _, bracket := range brackets {
			if openingBrackets[bracket] {
				stack = append(stack, bracket)
			} else if openingBracketOfACertainClosingBracket, ok := closingBrackets[bracket]; ok {
				lastBracket := stack[len(stack)-1]
				if lastBracket != openingBracketOfACertainClosingBracket {
					// This is bad we have no match ;( (aka corrupted chunk)
					score += scoresPart1[bracket]
					log.Debugf("Found illegal bracket: %s , expected %s", bracket, openingBracketOfACertainClosingBracket)
				}
				// Remove the (in)valid chunk we never ever have to look at it again
				stack = stack[:len(stack)-1]
			} else {
				panic("This space is for rent!")
			}
		}
	}

	return score
}

func part2(path string) int {
	lines := utils.ReadLines(path)
	result := make([]int, 0)

ContinueHere:
	for _, line := range lines {
		brackets := strings.Split(line, "")
		stack := make([]string, 0)

		// This part is exactly the same as part 1, but if we found an illegal bracket just go to the next line
		for _, bracket := range brackets {
			if openingBrackets[bracket] {
				stack = append(stack, bracket)
			} else if openingBracketOfACertainClosingBracket, ok := closingBrackets[bracket]; ok {
				lastBracket := stack[len(stack)-1]
				if lastBracket != openingBracketOfACertainClosingBracket {
					// This is bad we have no match ;( (aka corrupted chunk)
					log.Debugf("Found illegal bracket: %s , expected %s", bracket, openingBracketOfACertainClosingBracket)
					continue ContinueHere // Invalid line do not continue
				}
				// Remove the (in)valid chunk we never ever have to look at it again
				stack = stack[:len(stack)-1]
			} else {
				panic("This space is for rent!")
			}
		}

		// Now: what I have left are unfinished / closed brackets, go through them and finish it
		// do this in reverse order as we want to start with the last (as that is the 'first' unfinished)
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			bracket := stack[i]
			score *= 5
			score += scoresPart2[bracket]
		}

		result = append(result, score)
	}

	// Sort so we can get the middle one (always an odd amount according to AoC)
	sort.Ints(result)
	return result[len(result)/2]
}
