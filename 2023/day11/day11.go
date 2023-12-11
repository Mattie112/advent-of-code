package main

import (
	"AoC/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	var count int
	count = day11Part1("day11/day11-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day11Part2("day11/day11-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day11Part1("day11/day11.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day11Part2("day11/day11.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day11Part1(path string) int {
	return run(path, false)
}

func day11Part2(path string) int {
	return run(path, true)
}

func run(path string, partTwo bool) int {
	growDistance := 1
	if partTwo {
		growDistance = 1000000 - 1
	}
	answer := 0

	grid := make([][]string, 0)

	// First make a grid of the input
	lines := utils.ReadLines(path)
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	//printGrid(grid)
	//fmt.Println("")

	// Store the IDs of the empty rows and columns
	emptyY := make(map[int]bool)
	emptyX := make(map[int]bool)

	// Find empty rows
	for y := 0; y < len(grid); y++ {
		empty := true
		for x := 0; x < len(grid[y]); x++ {
			char := grid[y][x]
			if char != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyY[y] = true
			// This was my solution for part 1 that is now no longer needed
			//grid = append(grid[:y+1], grid[y:]...)
			//y++
		}
	}

	// Find empty columns
	for x := 0; x < len(grid[0]); x++ {
		empty := true
		for y := 0; y < len(grid); y++ {
			char := grid[y][x]
			if char != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyX[x] = true
			// This was my solution for part 1 that is now no longer needed
			//for y := 0; y < len(grid); y++ {
			//	grid[y] = append(grid[y][:x+1], grid[y][x:]...)
			//}
			//x++
		}
	}

	//printGrid(grid)

	// Find galaxies in grid and put them in a list
	var galaxies []string
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "#" {
				galaxies = append(galaxies, fmt.Sprintf("%d,%d", y, x))
			}
		}
	}

	// List to keep track of which galaxies we've already checked
	done := make(map[string]bool)

	// Loop through all the galaxys and for each one loop through all the galaxys again and find the shortest distance between them
	for i, galaxy := range galaxies {
		// Grab coordinates from the first galaxy
		galaxyX := utils.MustParseStringToInt(strings.Split(galaxy, ",")[1])
		galaxyY := utils.MustParseStringToInt(strings.Split(galaxy, ",")[0])

		// Loop through all the galaxys again
		for j, otherGalaxy := range galaxies {
			// If the galaxy is the same as the other galaxy, skip it
			_, exists := done[fmt.Sprintf("%d,%d", i, j)]
			if galaxy == otherGalaxy || exists {
				continue
			}

			// Before we forget: mark both ways as done (we'll only count one)
			done[fmt.Sprintf("%d,%d", i, j)] = true
			done[fmt.Sprintf("%d,%d", j, i)] = true

			// Grab coordinates from the second galaxy
			otherX := utils.MustParseStringToInt(strings.Split(otherGalaxy, ",")[1])
			otherY := utils.MustParseStringToInt(strings.Split(otherGalaxy, ",")[0])

			// Calculate the distance between the two galaxies (dijkstra)
			distance := math.Abs(float64(galaxyX-otherX)) + math.Abs(float64(galaxyY-otherY))

			// Go through our empty rows/colums. If the two galaxies are on opposite sides of an empty row/column, add the growDistance to the answer
			for y, _ := range emptyY {
				if galaxyY < y && otherY > y || galaxyY > y && otherY < y {
					distance += float64(growDistance)
				}
			}
			for x, _ := range emptyX {
				if galaxyX < x && otherX > x || galaxyX > x && otherX < x {
					distance += float64(growDistance)
				}
			}

			answer += int(distance)
			//fmt.Println("The distance between", galaxy, "and", otherGalaxy, "with ID", i+1, "and", j+1, "is", distance)
		}
	}

	return answer
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
