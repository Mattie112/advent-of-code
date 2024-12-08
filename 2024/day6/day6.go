package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day6/day6-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day6/day6-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day6/day6.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", Part2("day6/day6.txt")))
}

const (
	FLOOR    = "."
	OBSTACLE = "#"
	UP       = "^"
	RIGHT    = ">"
	DOWN     = "v"
	LEFT     = "<"
)

type cell struct {
	xPos   int
	yPos   int
	letter string
}

func Part1(path string) int {
	lines := utils.ReadLines(path)
	grid := make([][]cell, len(lines))
	answer := 0
	guardPos := cell{0, 0, ""}

	// Loop through the lines and split them pages and oders
	for y, line := range lines {
		letters := strings.Split(line, "")
		for x, letter := range letters {
			grid[y] = append(grid[y], cell{x, y, letter})
			if letter == "^" {
				guardPos = cell{x, y, letter}
			}
		}
	}

	visited := getVisited(grid, guardPos)

	answer = len(visited)

	return answer
}

func getVisited(grid [][]cell, guardPos cell) map[string]bool {
	visited := make(map[string]bool)

	for guardPos.yPos > 0 && guardPos.yPos < len(grid) && guardPos.xPos > 0 && guardPos.xPos < len(grid[0]) {
		grid[guardPos.yPos][guardPos.xPos].letter = "X"
		visited[fmt.Sprintf("%d,%d", guardPos.xPos, guardPos.yPos)] = true

		switch guardPos.letter {
		case UP:
			if guardPos.yPos-1 >= 0 && grid[guardPos.yPos-1][guardPos.xPos].letter == OBSTACLE {
				guardPos.letter = RIGHT
			} else {
				guardPos.yPos--
			}
		case RIGHT:
			if guardPos.xPos+1 < len(grid[0]) && grid[guardPos.yPos][guardPos.xPos+1].letter == OBSTACLE {
				guardPos.letter = DOWN
			} else {
				guardPos.xPos++
			}
		case DOWN:
			if guardPos.yPos+1 < len(grid) && grid[guardPos.yPos+1][guardPos.xPos].letter == OBSTACLE {
				guardPos.letter = LEFT
			} else {
				guardPos.yPos++
			}
		case LEFT:
			if guardPos.xPos-1 >= 0 && grid[guardPos.yPos][guardPos.xPos-1].letter == OBSTACLE {
				guardPos.letter = UP
			} else {
				guardPos.xPos--
			}
		}
		if guardPos.yPos < 0 || guardPos.yPos > len(grid)-1 || guardPos.xPos < 0 || guardPos.xPos > len(grid[0])-1 {
			continue
		}
		grid[guardPos.yPos][guardPos.xPos].letter = guardPos.letter

		printGrid(grid)
	}

	return visited
}

func Part2(path string) int {
	lines := utils.ReadLines(path)
	grid := make([][]cell, len(lines))
	answer := 0
	guardPos := cell{0, 0, ""}

	// Loop through the lines and split them pages and oders
	for y, line := range lines {
		letters := strings.Split(line, "")
		for x, letter := range letters {
			grid[y] = append(grid[y], cell{x, y, letter})
			if letter == "^" {
				guardPos = cell{x, y, letter}
			}
		}
	}

	// For each loop we need to remember the original grid and the visited cells
	var originalGrid [][]cell
	originalGrid = copyGrid(grid)
	part1Visited := getVisited(grid, guardPos)
	part1Count := 0
	log.Debugf("Part 1 visited %d\n", len(part1Visited))
	originalGuardPos := guardPos

	validObstructionPlaces := 0

	// Loop through all the cells that we visited in part 1
	// these are the only places wher placing an obstruction will have any effect
	for k, _ := range part1Visited {
		log.Debugf("Counter  %d/%d\n", part1Count, len(part1Visited))
		visited := make(map[string]string)

		grid = copyGrid(originalGrid)
		printGrid(grid)
		coords := strings.Split(k, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		// Skip the guard pos itself
		if y == originalGuardPos.yPos && x == originalGuardPos.xPos {
			continue
		}

		originalLetter := grid[y][x].letter
		grid[y][x].letter = OBSTACLE
		log.Debugf("Testing %d, %d", y, x)
		printGrid(grid)
		guardPos = originalGuardPos
		visited[fmt.Sprintf("%d,%d, %s", guardPos.xPos, guardPos.yPos, guardPos.letter)] = guardPos.letter

		for guardPos.yPos > 0 && guardPos.yPos < len(grid) && guardPos.xPos > 0 && guardPos.xPos < len(grid[0]) {
			// If we visited this cell AND we are in the same direction we are done
			if _, ok := visited[fmt.Sprintf("%d,%d,%s", guardPos.xPos, guardPos.yPos, guardPos.letter)]; ok {
				log.Debugf("Endless loop found %d, %d", y, x)
				validObstructionPlaces++
			}

			grid[guardPos.yPos][guardPos.xPos].letter = "X"

			visited[fmt.Sprintf("%d,%d,%s", guardPos.xPos, guardPos.yPos, guardPos.letter)] = guardPos.letter

			switch guardPos.letter {
			case UP:
				if guardPos.yPos-1 >= 0 && grid[guardPos.yPos-1][guardPos.xPos].letter == OBSTACLE {
					guardPos.letter = RIGHT
				} else {
					guardPos.yPos--
				}
			case RIGHT:
				if guardPos.xPos+1 < len(grid[0]) && grid[guardPos.yPos][guardPos.xPos+1].letter == OBSTACLE {
					guardPos.letter = DOWN
				} else {
					guardPos.xPos++
				}
			case DOWN:
				if guardPos.yPos+1 < len(grid) && grid[guardPos.yPos+1][guardPos.xPos].letter == OBSTACLE {
					guardPos.letter = LEFT
				} else {
					guardPos.yPos++
				}
			case LEFT:
				if guardPos.xPos-1 >= 0 && grid[guardPos.yPos][guardPos.xPos-1].letter == OBSTACLE {
					guardPos.letter = UP
				} else {
					guardPos.xPos--
				}
			}

			printGrid(grid)
		}

		grid[y][x].letter = originalLetter
		part1Count++
	}

	answer = validObstructionPlaces

	return answer
}

func printGrid(grid [][]cell) {
	if log.GetLevel() == log.DebugLevel {
		for _, row := range grid {
			for _, c := range row {
				fmt.Print(c.letter)
			}
			fmt.Println()
		}
		fmt.Print("\n\n\n")
	}
}

func copyGrid(grid [][]cell) [][]cell {
	newGrid := make([][]cell, len(grid))
	for y, row := range grid {
		for x, c := range row {
			newGrid[y] = append(newGrid[y], cell{x, y, c.letter})
		}
	}
	return newGrid
}
