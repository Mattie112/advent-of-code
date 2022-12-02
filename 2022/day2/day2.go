package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

const WIN = 6
const LOSE = 0
const DRAW = 3

func main() {
	var count int
	count = day1Part1("day2/day2-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day1Part2("day2/day2-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day1Part1("day2/day2.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day1Part2("day2/day2.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day1Part1(path string) int {
	lines := utils.ReadLines(path)
	score := 0

	for _, line := range lines {
		items := strings.Split(line, " ")
		outcome := check(items[0], items[1])
		modifier := 0
		if items[1] == "X" {
			modifier = 1
		}
		if items[1] == "Y" {
			modifier = 2
		}
		if items[1] == "Z" {
			modifier = 3
		}
		score += outcome + modifier
	}

	return score
}

func check(a string, b string) int {
	if a == "A" && b == "X" {
		return DRAW
	}
	if a == "A" && b == "Y" {
		return WIN
	}
	if a == "B" && b == "Y" {
		return DRAW
	}
	if a == "B" && b == "Z" {
		return WIN
	}
	if a == "C" && b == "Z" {
		return DRAW
	}
	if a == "C" && b == "X" {
		return WIN
	}

	return LOSE
}

func day1Part2(path string) int {
	lines := utils.ReadLines(path)
	score := 0
	options := []string{"Y", "X", "Z"}

	for _, line := range lines {
		items := strings.Split(line, " ")

		expectedOutcome := 0
		if items[1] == "X" {
			// lose
			expectedOutcome = LOSE
		}
		if items[1] == "Y" {
			// draw
			expectedOutcome = DRAW
		}
		if items[1] == "Z" {
			// win
			expectedOutcome = WIN
		}
		letter := ""
		for _, l := range options {
			outcome := check(items[0], l)
			if outcome == expectedOutcome {
				letter = l
				break
			}
		}
		modifier := 0
		if letter == "X" {
			modifier = 1
		}
		if letter == "Y" {
			modifier = 2
		}
		if letter == "Z" {
			modifier = 3
		}
		score += expectedOutcome + modifier
	}

	return score
}
