package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

const DONTUP = 1
const DONTRIGHT = 2
const DONTLEFT = 3
const DONTDOWN = 4

func main() {
	log.SetLevel(log.DebugLevel)
	var count int
	//count = part1("day9/day9-test.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = part1("day9/day9.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day9/day9-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = part2("day9/day9.txt")
	//log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	grid := make([][]int, len(lines))
	for i, line := range lines {
		//grid = append(grid, make([]int, len(lines[0])))
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
		//grid = append(grid, make([]int, len(lines[0])))
		numbers := utils.StrArrToIntArr(strings.Split(line, ""))
		grid[i] = numbers
	}

	maxY := len(grid)
	maxX := len(grid[0])
	minY := 0
	minX := 0
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
			// We are the lowest point (now store it so we can use it in the 2nd loop
			lowPoints = append(lowPoints, fmt.Sprintf("%d,%d", y, x))
		}
	}

	high1 := 0
	high2 := 0
	high3 := 0

	for _, lowPoint := range lowPoints {
		t := strings.Split(lowPoint, ",")
		lowY := utils.StringToInt(t[0])
		lowX := utils.StringToInt(t[1])
		lowPointValue := grid[lowY][lowX]

		basinCount := 0

		basinCount = checkLowerPoints(grid, lowY, lowX, basinCount, -1)

		log.Debugf("Checking point Y:%d X:%d with value %d", lowY, lowX, lowPointValue)

		// Now look into 4 directions until either a 9 or a lower number (if that is possible?)

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

func checkLowerPoints(grid [][]int, lowY int, lowX int, basinCount int, dontCheck int) int {
	maxY := len(grid)
	maxX := len(grid[0])
	minY := 0
	minX := 0
	// Check above
	if dontCheck != DONTUP {
		for y := lowY - 1; y >= minY; y-- {
			if grid[y][lowX] == 9 {
				break
			}
			//basinCount++
			basinCount += checkLowerPoints(grid, y, lowX, basinCount, DONTDOWN)
		}
	}

	// Check below
	if dontCheck != DONTDOWN {
		for y := lowY + 1; y < maxY; y++ {
			if grid[y][lowX] == 9 {
				break
			}
			//basinCount++
			basinCount += checkLowerPoints(grid, y, lowX, basinCount, DONTUP)
		}
	}

	// Check right
	if dontCheck != DONTRIGHT {
		for x := lowX + 1; x < maxX; x++ {
			if grid[lowY][x] == 9 {
				break
			}
			//basinCount++
			basinCount += checkLowerPoints(grid, lowY, x, basinCount, DONTLEFT)
		}
	}

	// Check left
	if dontCheck != DONTLEFT {
		for x := lowX - 1; x >= minX; x-- {
			if grid[lowY][x] == 9 {
				break
			}
			//basinCount++
			basinCount += checkLowerPoints(grid, lowY, x, basinCount, DONTRIGHT)
		}
	}
	basinCount++
	return basinCount
}
