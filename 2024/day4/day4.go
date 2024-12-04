package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day4/day4-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day4/day4-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day4/day4.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", Part2("day4/day4.txt")))
}

type cell struct {
	// I hoped that this was useful for part 2 but no :)
	xPos   int
	yPos   int
	letter string
}

func Part1(path string) int {
	return P1andP2(path, false)
}

func Part2(path string) int {
	return P1andP2(path, true)
}

// P1andP2 Part 1 & 2 share a lot of code so yeah this saves some lines
func P1andP2(path string, part2 bool) int {
	lines := utils.ReadLines(path)
	answer := 0
	grid := make([][]cell, len(lines))

	// Loop through the lines and split them into left and right
	for y, line := range lines {
		letters := strings.Split(line, "")
		for x, letter := range letters {
			grid[y] = append(grid[y], cell{x, y, letter})
		}
	}

	firstLetter := "X"
	if part2 {
		firstLetter = "A"
	}

	// Now go and find the first letter (our startingpoint, X)
	for y, row := range grid {
		for x, c := range row {
			if c.letter == firstLetter {
				// Now that we have found our first letter we need to search in all directions for the next one
				log.Debugf("Found A at %d, %d", y+1, x+1)
				foundXmas := 0
				if part2 {
					foundXmas = searchInGridAroundPosition2(grid, x, y)
				} else {
					foundXmas = searchInGridAroundPosition(grid, x, y, []string{"M", "A", "S"})
				}
				answer += foundXmas
			}
		}
	}

	if log.GetLevel() == log.DebugLevel {
		printGrid(grid)
	}

	return answer
}

func searchInGridAroundPosition(grid [][]cell, x, y int, wantedLetters []string) int {
	counter := 0
	// Search for the wanted letter in all directions
	// Up
	// Down
	// Left
	// Right
	// Up-Left
	// Up-Right
	// Down-Left
	// Down-Right
	lookHere := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, lookOption := range lookHere {
		lookXPos := x + lookOption[0]
		lookYPos := y + lookOption[1]

		// Check if we are out of bounds
		if lookXPos < 0 || lookXPos >= len(grid[0]) || lookYPos < 0 || lookYPos >= len(grid) {
			continue
		}

		// First check the "direction" of the first wanted letter (xMas)
		wantedLetter := wantedLetters[0]
		if grid[lookYPos][lookXPos].letter == wantedLetter {
			log.Debugf("Found %s at %d, %d", wantedLetter, lookYPos+1, lookXPos+1)

			// Found the letter we were looking for
			// Now we need to look in the same direction for the next letter
			for i := 1; i < len(wantedLetters); i++ {
				wantedLetter = wantedLetters[i]

				// Check if we are out of bounds
				nextLetterY := lookYPos + (i * lookOption[1])
				nextLetterX := lookXPos + (i * lookOption[0])
				log.Debugf("Looking for %s at %d, %d", wantedLetter, nextLetterY+1, nextLetterX+1)

				if nextLetterX < 0 || nextLetterX >= len(grid[0]) || nextLetterY < 0 || nextLetterY >= len(grid) {
					log.Debugf("Out of bounds")
					break
				}

				// Check of the letter in our direction is the one we are looking for
				if grid[nextLetterY][nextLetterX].letter != wantedLetter {
					log.Debugf("Not the letter we are looking for want: %s but got %s", wantedLetter, grid[nextLetterY][nextLetterX].letter)
					break
				}

				log.Debugf("Found %s at %d, %d", wantedLetter, nextLetterY+1, nextLetterX+1)

				// End of the word: we found it all!
				if i == len(wantedLetters)-1 {
					log.Debugf("Found XMAS! Adding 1 to answer")
					counter++
				}
			}
		}
	}
	return counter
}

func searchInGridAroundPosition2(grid [][]cell, x, y int) int {
	counter := 0
	// Search for the wanted letter in not all directions :)
	// Up-Left
	// Up-Right
	// Down-Left
	// Down-Right

	// First check if any of our options are out of bounds -> no X-MAS found
	lookHere := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, lookOption := range lookHere {
		lookXPos := x + lookOption[0]
		lookYPos := y + lookOption[1]

		// Check if we are out of bounds
		if lookXPos < 0 || lookXPos >= len(grid[0]) || lookYPos < 0 || lookYPos >= len(grid) {
			return 0
		}
	}

	log.Debugf("we are in bounds")

	log.Debugf("%s %s %s", grid[y-1][x-1].letter, grid[y-1][x].letter, grid[y-1][x+1].letter)
	log.Debugf("%s %s %s", grid[y][x-1].letter, grid[y][x].letter, grid[y][x+1].letter)
	log.Debugf("%s %s %s", grid[y+1][x-1].letter, grid[y+1][x].letter, grid[y+1][x+1].letter)

	// Then just a nice hardcoded search as there really are only 4 options, strange that part 2 is easier than part 1 in that regard :)
	if (grid[y-1][x-1].letter == "M" && grid[y+1][x+1].letter == "S") || (grid[y+1][x+1].letter == "M" && grid[y-1][x-1].letter == "S") {
		log.Debugf("Found bottom-left to top-right or top-right to bottom-left")
		if (grid[y-1][x+1].letter == "M" && grid[y+1][x-1].letter == "S") || (grid[y+1][x-1].letter == "M" && grid[y-1][x+1].letter == "S") {
			log.Debugf("Found top-left to bottom-right or bottom-right to top-left")
			counter++
		}
	}

	return counter
}

func printGrid(grid [][]cell) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(c.letter)
		}
		fmt.Println()
	}
}
