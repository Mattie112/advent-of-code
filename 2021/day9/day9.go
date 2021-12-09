package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

// Yeah hackyhacky I know but hey if it is stupid and works it's not stupid right?
// used for part 2 to keep the points we have already visited
var skipGrid [1000][1000]bool

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day9/day9-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day9/day9.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day9/day9-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day9/day9.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	grid := make([][]int, len(lines))
	for i, line := range lines {
		numbers := utils.StrArrToIntArr(strings.Split(line, ""))
		grid[i] = numbers
	}

	maxY := len(grid)
	maxX := len(grid[0])
	minY := 0
	minX := 0
	lowPoints := 0

	for y, row := range grid {
		for x, cell := range row {

			// Check above
			if y-1 >= minY && y-1 < maxY {
				if cell >= grid[y-1][x] {
					continue
				}
			}

			// Check right
			if x+1 >= minX && x+1 < maxX {
				if cell >= grid[y][x+1] {
					continue
				}
			}

			// Check below
			if y+1 >= minY && y+1 < maxY {
				if cell >= grid[y+1][x] {
					continue
				}
			}

			// Check left
			if x-1 >= minX && x-1 < maxX {
				if cell >= grid[y][x-1] {
					continue
				}
			}

			log.Debugf("Found low point Y:%d X:%d with value %d (so %d points)", y, x, cell, 1+cell)
			// We are the lowest point
			lowPoints += 1 + cell

		}
	}

	return lowPoints
}

func part2(path string) int {
	lines := utils.ReadLines(path)

	grid := make([][]int, len(lines))
	for i, line := range lines {
		numbers := utils.StrArrToIntArr(strings.Split(line, ""))
		grid[i] = numbers
	}

	maxY := len(grid)
	maxX := len(grid[0])
	minY, minX := 0, 0
	lowPoints := make([]string, 0)

	for y, row := range grid {
		for x, cell := range row {
			// Check above
			if y-1 >= minY && y-1 < maxY {
				if cell >= grid[y-1][x] {
					continue
				}
			}

			// Check below
			if y+1 >= minY && y+1 < maxY {
				if cell >= grid[y+1][x] {
					continue
				}
			}

			// Check right
			if x+1 >= minX && x+1 < maxX {
				if cell >= grid[y][x+1] {
					continue
				}
			}

			// Check left
			if x-1 >= minX && x-1 < maxX {
				if cell >= grid[y][x-1] {
					continue
				}
			}

			log.Debugf("Found low point Y:%d X:%d with value %d (so %d points)", y, x, cell, 1+cell)
			// We are the lowest point (now store it, so we can use it in the 2nd loop
			lowPoints = append(lowPoints, fmt.Sprintf("%d,%d", y, x))
		}
	}

	// no comment
	high1, high2, high3 := 0, 0, 0

	// Loop through our lowPoints and then try to find more
	for _, lowPoint := range lowPoints {
		t := strings.Split(lowPoint, ",")
		lowY := utils.StringToInt(t[0])
		lowX := utils.StringToInt(t[1])
		lowPointValue := grid[lowY][lowX]

		// Check for LowerPoints (or really higher points? As it flows towards this point)
		basinCount := checkLowerPoints(grid, lowY, lowX)

		log.Debugf("Checking point Y:%d X:%d with value %d basinCount: %d", lowY, lowX, lowPointValue, basinCount)

		// Yeah as I said: no comment :)
		if basinCount > high1 {
			high3 = high2
			high2 = high1
			high1 = basinCount
		} else if basinCount > high2 {
			high3 = high2
			high2 = basinCount
		} else if basinCount > high3 {
			high3 = basinCount
		}
	}

	log.Debugf("%d * %d * %d", high1, high2, high3)
	return high1 * high2 * high3
}

func checkLowerPoints(grid [][]int, lowY int, lowX int) int {
	if skipGrid[lowY][lowX] {
		return 0
	}
	skipGrid[lowY][lowX] = true

	maxY := len(grid)
	maxX := len(grid[0])
	minY, minX, basinCount := 0, 0, 0
	// Now look into 4 directions until we find a 9
	// Check above and continue if we have a 9 or already checked that cell
	for y := lowY - 1; y >= minY; y-- {
		if grid[y][lowX] == 9 || skipGrid[y][lowX] {
			break
		}
		basinCount += checkLowerPoints(grid, y, lowX)
	}

	// Check below and continue if we have a 9 or already checked that cell
	for y := lowY + 1; y < maxY; y++ {
		if grid[y][lowX] == 9 || skipGrid[y][lowX] {
			break
		}
		basinCount += checkLowerPoints(grid, y, lowX)
	}

	// Check right and continue if we have a 9 or already checked that cell
	for x := lowX + 1; x < maxX; x++ {
		if grid[lowY][x] == 9 || skipGrid[lowY][x] {
			break
		}
		basinCount += checkLowerPoints(grid, lowY, x)
	}

	// Check left and continue if we have a 9 or already checked that cell
	for x := lowX - 1; x >= minX; x-- {
		if grid[lowY][x] == 9 || skipGrid[lowY][x] {
			break
		}
		basinCount += checkLowerPoints(grid, lowY, x)
	}

	basinCount++
	return basinCount
}
