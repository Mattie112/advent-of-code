package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day6/day6-test.txt")))
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day6/day6-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day6/day6.txt")))
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", Part2("day6/day6.txt")))
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
	visited := make(map[string]bool)

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
		//grid[guardPos.yPos][guardPos.xPos].letter = guardPos.letter
		//for _, row := range grid {
		//	for _, c := range row {
		//		fmt.Print(c.letter)
		//	}
		//	fmt.Println()
		//}
		//fmt.Println("\n\n\n")
	}

	answer = len(visited)

	return answer
}

func Part2(path string) int {
	return 1
}
