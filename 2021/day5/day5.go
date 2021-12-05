package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day5/day5-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day5/day5.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day5/day5-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day5/day5.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		x1 := utils.StringToInt(res[1])
		y1 := utils.StringToInt(res[2])
		x2 := utils.StringToInt(res[3])
		y2 := utils.StringToInt(res[4])
		log.Debugf("x1: %d, x2: %d, y1: %d, y2: %d", x1, x2, y1, y2)

		// Now add this to the grid
		// Skip diagonal for part 1
		if x1 != x2 && y1 != y2 {
			continue
		}

		if y1 >= y2 && x1 >= x2 {
			for y := y2; y <= y1; y++ {
				for x := x2; x <= x1; x++ {
					grid[y][x]++
				}
			}
		} else if y2 >= y1 && x2 >= x1 {
			for y := y2; y >= y1; y-- {
				for x := x2; x >= x1; x-- {
					grid[y][x]++
				}
			}
		} else {
			panic("I should not get here")
		}
		if log.GetLevel() == log.DebugLevel {
			fmt.Println(line)
			printGrid(grid)
			fmt.Println("")
		}
	}
	log.Debugln("Grid done")
	return countOverlaps(grid)
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(cell)
		}
		fmt.Println("")
	}
}
func countOverlaps(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell > 1 {
				count++
			}
		}
	}
	return count
}

func part2(path string) int {
	lines := utils.ReadLines(path)

	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		x1 := utils.StringToInt(res[1])
		y1 := utils.StringToInt(res[2])
		x2 := utils.StringToInt(res[3])
		y2 := utils.StringToInt(res[4])
		log.Debugf("x1: %d, x2: %d, y1: %d, y2: %d", x1, x2, y1, y2)
		// Now add this to the grid

		// Do diagonal for part 2 (rest of the code is equal to part 1)
		if x1 != x2 && y1 != y2 {
			if x1 < x2 && y1 < y2 {
				// From up left to down right
				log.Debug("case 1")
				log.Debug(line)
				steps := x2 - x1
				for i := steps; i >= 0; i-- {
					grid[y1+i][x1+i]++
				}
			} else if x1 < x2 && y1 > y2 {
				// From down left to up right
				log.Debug("case 2")
				log.Debug(line)
				steps := x2 - x1
				for i := steps; i >= 0; i-- {
					grid[y1-i][x1+i]++
				}
			} else if x1 > x2 && y1 < y2 {
				// From up right to left down
				log.Debug("case 3")
				log.Debug(line)
				steps := y2 - y1
				for i := steps; i >= 0; i-- {
					grid[y1+i][x1-i]++
				}
			} else if x1 > x2 && y1 > y2 {
				// From down right to left up
				log.Debug("case 4")
				log.Debug(line)
				steps := x1 - x2
				for i := steps; i >= 0; i-- {
					grid[y1-i][x1-i]++
				}
			} else {
				panic("I should not get here")
			}
		} else if y1 >= y2 && x1 >= x2 {
			for y := y2; y <= y1; y++ {
				for x := x2; x <= x1; x++ {
					grid[y][x]++
				}
			}
		} else if y2 >= y1 && x2 >= x1 {
			for y := y2; y >= y1; y-- {
				for x := x2; x >= x1; x-- {
					grid[y][x]++
				}
			}
		} else {
			panic("I should not get here")
		}
		if log.GetLevel() == log.DebugLevel {
			fmt.Println(line)
			printGrid(grid)
			fmt.Println("")
		}
	}
	log.Debugln("Grid done")
	return countOverlaps(grid)
}
