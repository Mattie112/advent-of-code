package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day13/day13-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day13/day13.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	part2("day13/day13-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer ^"))
	part2("day13/day13.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer ^"))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	grid, folds := getGridAndFolds(lines)

	grid = fold(grid, folds, true)

	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				count++
			}
		}
	}

	return count
}

func part2(path string) string {
	lines := utils.ReadLines(path)

	grid, folds := getGridAndFolds(lines)

	grid = fold(grid, folds, false)

	output := ""
	// Print a subset so we can visually see the answer
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			if grid[y][x] != "#" {
				output += "."
				fmt.Print(".")
			} else {
				output += "#"
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println("")
		output += "\n"
	}

	return output
}

func fold(grid [][]string, folds []string, breakAtFirstFold bool) [][]string {
	// Start folding
	var re = regexp.MustCompile(`(?m)([x|y])=(\d+)`)
	for _, line := range folds {
		fold := re.FindStringSubmatch(line)
		axis := fold[1]
		amount := utils.StringToInt(fold[2])
		log.Debugf("Fold along %s=%d", axis, amount)

		yCount := amount - 1
		xCount := amount - 1
		switch axis {
		case "y":
			for y, row := range grid {
				if y <= amount {
					continue
				}
				for x, cell := range row {
					if cell == "#" {
						grid[yCount][x] = cell
					}
					grid[y][x] = ""
				}
				yCount--
			}

		case "x":
			for y, row := range grid {
				xCount = amount - 1
				for x, cell := range row {
					if x <= amount {
						continue
					}
					if cell == "#" {
						log.Debugf(fmt.Sprintf("%d, %d", y, xCount))
						grid[y][xCount] = cell
					}
					grid[y][x] = ""
					xCount--
				}
			}
		}

		// Part 1 just the first fold
		if breakAtFirstFold {
			break
		}
	}
	return grid
}

func getGridAndFolds(lines []string) ([][]string, []string) {
	grid := make([][]string, 1500)
	for i := range grid {
		grid[i] = make([]string, 1500)
	}

	folds := make([]string, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "fold") {
			folds = append(folds, line)
			continue
		}
		coordinates := strings.Split(line, ",")
		x := utils.StringToInt(coordinates[0])
		y := utils.StringToInt(coordinates[1])
		grid[y][x] = "#"
	}
	return grid, folds
}
