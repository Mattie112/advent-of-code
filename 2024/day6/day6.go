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
	//fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day6/day6-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day6/day6-test.txt")))
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day6/day6.txt")))
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

		//if log.GetLevel() == log.DebugLevel {
		//	for _, row := range grid {
		//		for _, c := range row {
		//			fmt.Print(c.letter)
		//		}
		//		fmt.Println()
		//	}
		//	fmt.Println("\n\n\n")
		//}

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

	var originalGrid [][]cell
	originalGrid = copyGrid(grid)

	part1Visited := getVisited(grid, guardPos)
	part1Count := 0
	fmt.Printf("Part 1 visited %d\n", len(part1Visited))
	//part1Visited = make(map[string]bool)
	//part1Visited["3,6"] = true
	//part1Visited["6,7"] = true
	originalGuardPos := guardPos

	validObstructionPlaces := 0

	// Loop through all the cells that we visited in part 1
	// these are the only places wher placing an obstruction will have any effect
	for k, _ := range part1Visited {
		//fmt.Printf("Counter  %d/%d\n", part1Count, len(part1Visited))
		visited := make(map[string]string)

		grid = copyGrid(originalGrid)
		printGrid(grid)
		coords := strings.Split(k, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		if y == originalGuardPos.yPos && x == originalGuardPos.xPos {
			continue
		}

		originalLetter := grid[y][x].letter
		grid[y][x].letter = OBSTACLE
		log.Debugf("Testing %d, %d", y, x)
		//fmt.Printf("Testing %d, %d\n", y, x)
		printGrid(grid)
		firstStep := true
		guardPos = originalGuardPos
		visited[fmt.Sprintf("%d,%d", guardPos.xPos, guardPos.yPos)] = guardPos.letter
		loops := 0

		for guardPos.yPos > 0 && guardPos.yPos < len(grid) && guardPos.xPos > 0 && guardPos.xPos < len(grid[0]) {
			loops++
			if !firstStep && guardPos.yPos == originalGuardPos.yPos && guardPos.xPos == originalGuardPos.xPos && guardPos.letter == originalGuardPos.letter {
				log.Debugf("Endless loop found %d, %d", y+1, x+1)
				validObstructionPlaces++
				// WE ARE DONE!
				break
			}
			if loops > 10000 {
				fmt.Println("Endless loop based on numbers")
				validObstructionPlaces++
				break
			}
			if _, ok := visited[fmt.Sprintf("%d,%d", guardPos.xPos, guardPos.yPos)]; ok {
				if !firstStep && visited[fmt.Sprintf("%d,%d", guardPos.xPos, guardPos.yPos)] == guardPos.letter {
					log.Debugf("Endless loop found %d, %d", y+1, x+1)
					validObstructionPlaces++
					// WE ARE DONE!
					break
				}
			}

			firstStep = false
			grid[guardPos.yPos][guardPos.xPos].letter = "X"

			visited[fmt.Sprintf("%d,%d", guardPos.xPos, guardPos.yPos)] = guardPos.letter

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

			//if guardPos.yPos < 0 || guardPos.yPos > len(grid)-1 || guardPos.xPos < 0 || guardPos.xPos > len(grid[0])-1 {
			//	continue
			//}

			printGrid(grid)
		}

		// The guard walked out of bounds, not a valid position
		//validObstructionPlaces++

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
