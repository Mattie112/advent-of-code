package main

import (
	"AoC/utils"
	"fmt"
)

func main() {
	var count int
	count = day13Part1("day13/day13-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day13Part2("day13/day13-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = day13Part1("day13/day13.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day13Part2("day13/day13.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day13Part1(path string) int {
	answer := 0
	gridList := make([][][]string, 0)
	grid := make([][]string, 0)

	// First make a grid of the input
	lines := utils.ReadLines(path)
	for _, line := range lines {
		if line == "" {
			gridList = append(gridList, grid)
			grid = make([][]string, 0)
			continue
		}

		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	// Also add the last one
	gridList = append(gridList, grid)

	// Now we have a list of grids, we can loop through them
	for _, g := range gridList {
		printGrid(g)
		fmt.Println()

	}

	return answer
}

func day13Part2(path string) int {
	return 0
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
