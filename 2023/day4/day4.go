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
	count = day4Part2("day4/day4-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day4Part1("day4/day4.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day4Part2("day4/day4.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

// Normally I keep my part1 and part2 separate, but this one is so similar that I just combined them
func day4Part1(path string) int {
	part1, _ := both(path)
	return part1
}

func day4Part2(path string) int {
	_, part2 := both(path)
	return part2
}

func both(path string) (int, int) {
	lines := utils.ReadLines(path)
	part1 := 0
	part2 := 0

	cardsIHave := make([]int, len(lines)+1)

	for i, line := range lines {
		line = line[strings.Index(line, ": ")+2:]

		winningStr := strings.Split(line, " | ")[0]
		winning := strings.Split(winningStr, " ")
		numbersStr := strings.Split(line, " | ")[1]
		numbers := strings.Split(numbersStr, " ")

		i = i + 1 // Just to make it more readable (with the AoC input)

		winningNumberPoints := 0 // Part 1

		cardsIHave[i]++ // Part 2: We always have the card itself
		winningId := i  // Part 2: Our current card (we will win the next x cards)

		for _, number := range numbers {
			// First check the winning numbers
			for _, win := range winning {
				if number == win && number != "" && win != "" {
					// For part 1 add the points
					if winningNumberPoints == 0 {
						winningNumberPoints = 1
					} else {
						winningNumberPoints *= 2
					}
					// Part 2 add the cards based on the amound we already have
					winningId++
					cardsIHave[winningId] += cardsIHave[i]
				}
			}
		}
		part1 += winningNumberPoints
	}

	// Count part 2
	for _, amount := range cardsIHave {
		part2 += amount
	}

	return part1, part2
}
