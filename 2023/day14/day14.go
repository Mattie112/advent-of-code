package main

import (
	"AoC/utils"
	"fmt"
)

func main() {
	var count int
	//count = day14Part1("day14/day14-test.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day14Part2("day14/day14-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = day14Part1("day14/day14.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day14Part2("day14/day14.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day14Part1(path string) int {
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

	for y, row := range grid {
		for x, cell := range row {
			if cell == "O" {

				// Try to move up as much as possible
				for i := y - 1; i >= 0; i-- {
					cellToCheck := grid[i][x]

					if cellToCheck == "#" || cellToCheck == "O" {
						// At the position BEFORE we hit something -> place the rock
						grid[i+1][x] = "O"
						// If we have moved set our current position to be empty ground
						if i+1 != y {
							grid[y][x] = "."
						}
						break
					} else if i == 0 {
						grid[0][x] = "O"
						grid[y][x] = "."
					}
				}
			}
		}
	}

	//printGrid(grid)
	//fmt.Println("")

	// Now calculate the load
	southEnd := len(grid) - 1
	for y, row := range grid {
		for _, cell := range row {
			if cell == "O" {
				answer += southEnd - y + 1
			}
		}
	}

	return answer
}

// Part 2 contains the same code as part 1 just 3 times more for the other rotations
func day14Part2(path string) int {
	answer := 0
	//startTime := time.Now()
	// All the grids we have seen, key is hash, value is cycle
	seenGrids := make(map[string]int)
	// Grids that will repeat, key is cycle, value is grid
	repeatList := make(map[int][][]string)
	// The load per cycle, key is cycle, value is load
	answerList := make(map[int]int)
	// Cycle of the first grid repeat
	firstRepeat := 0

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

	for cycle := 1; cycle <= 1000000000; cycle++ {

		// Roll everything north (up)
		for y, row := range grid {
			for x, cell := range row {
				if cell == "O" {
					// Try to move UP as much as possible
					for i := y - 1; i >= 0; i-- {
						cellToCheck := grid[i][x]

						if cellToCheck == "#" || cellToCheck == "O" {
							// At the position BEFORE we hit something -> place the rock
							grid[i+1][x] = "O"
							// If we have moved set our current position to be empty ground
							if i+1 != y {
								grid[y][x] = "."
							}
							break
						} else if i == 0 {
							grid[0][x] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}

		// Roll everything west (left)
		for y, row := range grid {
			for x, cell := range row {
				if cell == "O" {
					// Try to move LEFT as much as possible
					for i := x - 1; i >= 0; i-- {
						cellToCheck := grid[y][i]

						if cellToCheck == "#" || cellToCheck == "O" {
							// At the position BEFORE we hit something -> place the rock
							grid[y][i+1] = "O"
							// If we have moved set our current position to be empty ground
							if i+1 != x {
								grid[y][x] = "."
							}
							break
						} else if i == 0 {
							grid[y][0] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}

		// Roll everything south (down)
		for y := len(grid) - 1; y >= 0; y-- {
			for x, cell := range grid[y] {
				if cell == "O" {
					// Try to move DOWN as much as possible
					for i := y + 1; i < len(grid); i++ {
						cellToCheck := grid[i][x]

						if cellToCheck == "#" || cellToCheck == "O" {
							// At the position BEFORE we hit something -> place the rock
							grid[i-1][x] = "O"
							// If we have moved set our current position to be empty ground
							if i-1 != y {
								grid[y][x] = "."
							}
							break
						} else if i == len(grid)-1 {
							grid[len(grid)-1][x] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}

		// Roll everything east (right)
		for y := len(grid) - 1; y >= 0; y-- {
			for x := len(grid[y]) - 1; x >= 0; x-- {
				cell := grid[y][x]
				if cell == "O" {
					// Try to move RIGHT as much as possible
					for i := x + 1; i < len(grid[y]); i++ {
						cellToCheck := grid[y][i]

						if cellToCheck == "#" || cellToCheck == "O" {
							// At the position BEFORE we hit something -> place the rock
							grid[y][i-1] = "O"
							// If we have moved set our current position to be empty ground
							if i-1 != x {
								grid[y][x] = "."
							}
							break
						} else if i == len(grid[y])-1 {
							grid[y][len(grid[y])-1] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}

		gridHash := getHash(grid)
		_, ok := seenGrids[gridHash]
		if ok {
			//fmt.Println("Cycle", cycle, "is a repeat of cycle", seenGrids[getHash(grid)])
			// Only stop of we have a "full" repeat, not just a partial one (so we are starting to repeat for the second time
			if _, ok2 := repeatList[seenGrids[gridHash]]; ok2 {
				break
			}
			repeatList[seenGrids[gridHash]] = grid

			if firstRepeat == 0 {
				firstRepeat = cycle
			}

			// Now calculate the load value for this cycle
			southEnd := len(grid) - 1

			for y, row := range grid {
				for _, cell := range row {
					if cell == "O" {
						answerList[cycle] += southEnd - y + 1
					}
				}
			}
		} else {
			seenGrids[gridHash] = cycle
		}

		//if cycle%100000 == 0 {
		//	fmt.Println("The grid after cycle", cycle, "is:")
		//	printGrid(grid)
		//	fmt.Println("Runtime:", time.Since(startTime))
		//	fmt.Println("")
		//}
	}

	// I need the answer we had at the first repeat, plus the number of cycles left, modulo the length of the repeat list
	answer = answerList[firstRepeat+(1000000000-firstRepeat)%(len(answerList))]

	// Print runtime
	//fmt.Println("Runtime:", time.Since(startTime))
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

func getHash(grid [][]string) string {
	hash := ""
	for _, row := range grid {
		for _, cell := range row {
			hash += cell
		}
	}
	return hash
}
