package main

import (
	"AoC/utils"
	"fmt"
)

func main() {
	var count int
	count = day10Part1("day10/day10-test-1.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day10Part2("day10/day10-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day10Part1("day10/day10.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day10Part2("day10/day10.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func getNeighbourIfExists(grid [][]string, y int, x int) string {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return ""
	}
	id := fmt.Sprintf("%d,%d", y, x)
	return id
	//return grid[y][x]
}

func day10Part1(path string) int {
	answer := 0

	var grid [][]string
	startPos := ""

	lines := utils.ReadLines(path)
	for y, line := range lines {
		var row []string
		for x, char := range line {
			row = append(row, string(char))
			if char == 'S' {
				startPos = fmt.Sprintf("%d,%d", y, x)
			}
		}
		grid = append(grid, row)
	}

	neighbours := make(map[string][]string)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			cell := grid[y][x]
			id := fmt.Sprintf("%d,%d", y, x)

			switch cell {
			case ".":
				// Do nothing
			case "|":
				// Neighbour above and below (north and south)
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y-1, x))
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y+1, x))
			case "-":
				// Neighbour left and right (east and west)
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y, x-1))
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y, x+1))
			case "L":
				// Neighbour north and east
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y-1, x))
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y, x+1))
			case "J":
				// Neighbour north and west
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y-1, x))
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y, x-1))
			case "7":
				// Neighbour south and west
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y+1, x))
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y, x-1))
			case "F":
				// Neighbour south and east
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y+1, x))
				neighbours[id] = append(neighbours[id], getNeighbourIfExists(grid, y, x+1))
			default:
				// Do nothing
			}

		}
	}

	// Find all the cells that have our starting position as neighbour and add those to the list
	for k, neighbour := range neighbours {
		for _, cell := range neighbour {
			if cell == startPos {
				neighbours[startPos] = append(neighbours[startPos], k)
			}
		}
	}

	visited := make(map[string]int)
	visited[startPos] = 0
	toBeChecked := []string{startPos}
	for len(toBeChecked) > 0 {
		// Pop the first element
		current := toBeChecked[0]
		toBeChecked = toBeChecked[1:]

		// Add all the neighbours to the list of cells to be checked
		for _, neighbour := range neighbours[current] {
			// Check if we've already visited this cell
			if _, ok := visited[neighbour]; ok {
				continue
			}

			toBeChecked = append(toBeChecked, neighbour)
			// Mark the cell as visited
			visited[neighbour] += visited[current] + 1

		}
	}

	//printGrid(grid)
	//fmt.Println("")
	//printVisited(grid, visited)

	// Find the cell with the longest distance
	longest := 0
	for _, distance := range visited {
		if distance > longest {
			longest = distance
		}
	}
	answer = longest

	return answer
}

func printGrid(grid [][]string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}

func printVisited(grid [][]string, visited map[string]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if distance, ok := visited[fmt.Sprintf("%d,%d", y, x)]; ok {
				fmt.Print(distance)
			} else {
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println()
	}
}

func day10Part2(path string) int {
	answer := 0

	lines := utils.ReadLines(path)
	for _, line := range lines {
		_ = line
	}

	return answer
}
