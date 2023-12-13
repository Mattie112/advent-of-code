package main

import (
	"AoC/utils"
	"fmt"
)

func main() {
	var count int
	count = day13Part1("day13/day13-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day13Part2("day13/day13-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day13Part1("day13/day13.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day13Part2("day13/day13.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day13Part1(path string) int {
	answer := 0
	gridList := createGrid(path)

	for _, g := range gridList {
		tmp, _, _ := calculateGrid(g, -1, -1)
		answer += tmp
	}

	return answer
}

func day13Part2(path string) int {
	answer := 0
	gridList := createGrid(path)

	// Keep lists of reflections we have seen for a certain grid
	reflectionListRow := make(map[int]int)
	reflectionListColumn := make(map[int]int)

	for gridId, g := range gridList {
		_, reflectionRow, reflectionCol := calculateGrid(g, -1, -1)
		if reflectionCol != -1 {
			reflectionListColumn[gridId] = reflectionCol
		}
		if reflectionRow != -1 {
			reflectionListRow[gridId] = reflectionRow
		}
	}

	//fmt.Println("start with part 2")
	// Now we have to loop again but start changing characters until we have a different line that matches... wow...
gridList:
	for gridId, g := range gridList {
		for originalRowId, iHatePartTwo := range g {
			for originalColumnId := range iHatePartTwo {

				// Swap the letters
				if g[originalRowId][originalColumnId] == "#" {
					g[originalRowId][originalColumnId] = "."
				} else {
					g[originalRowId][originalColumnId] = "#"
				}

				//fmt.Println()
				//printGrid(g)
				//fmt.Printf("Changed row %d colum %d\n", originalRowId, originalColumnId)
				//fmt.Println()

				tmpAnswer, _, _ := calculateGrid(g, reflectionListRow[gridId], reflectionListColumn[gridId])

				// Change it back for the next iteration
				if g[originalRowId][originalColumnId] == "#" {
					g[originalRowId][originalColumnId] = "."
				} else {
					g[originalRowId][originalColumnId] = "#"
				}

				// If we have found an answer -> go to the next grid, if not loop back and change another letter
				if tmpAnswer != -1 {
					answer += tmpAnswer
					continue gridList
				}
			}
		}
	}

	return answer
}

// This will just create nice grids of the input, nothig special
func createGrid(path string) [][][]string {
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
	return gridList
}

func calculateGrid(g [][]string, skipRow int, skipCol int) (answer int, reflectionRow int, reflectionCol int) {
	answer = 0

	// Now we have a list of grids, we can loop through them to find reflections
	//printGrid(g)
	//fmt.Println()

	// Loop through each column to find vertical reflections
	for column := 1; column < len(g[0]); column++ {
		reflectionsThisColumn := 0
		for row := 0; row < len(g); row++ {
			// Get the first part of the row
			tmp := g[row][0:column]
			firstPart := make([]string, len(tmp))
			copy(firstPart, tmp)

			// Get the second part of the row
			secondPart := g[row][column:]

			// Mirror the first part
			for i, j := 0, len(firstPart)-1; i < j; i, j = i+1, j-1 {
				firstPart[i], firstPart[j] = firstPart[j], firstPart[i]
			}

			// Now compare the two parts
			if StringSliceEqual(firstPart, secondPart) {
				reflectionsThisColumn++
			}
		}

		if reflectionsThisColumn == len(g) {
			// For part 2: skip if we already found this reflection
			if skipCol == column {
				continue
			}

			return column, -1, column
		}
	}

	// Loop through each row to find horizontal reflections
	for row := 1; row < len(g); row++ {
		reflectionsThisRow := 0
		for column := 0; column < len(g[0]); column++ {
			// As this is now a slice instead of a string -> convert it into a string
			firstTmp := g[0:row]
			var firstPart []string
			for _, p := range firstTmp {
				firstPart = append(firstPart, p[column])
			}

			secondTmp := g[row:]
			var secondPart []string
			for _, p := range secondTmp {
				secondPart = append(secondPart, p[column])
			}

			// Mirror the first part
			for i, j := 0, len(firstPart)-1; i < j; i, j = i+1, j-1 {
				firstPart[i], firstPart[j] = firstPart[j], firstPart[i]
			}

			// Now compare the two parts
			if StringSliceEqual(firstPart, secondPart) {
				reflectionsThisRow++
			}
		}
		if reflectionsThisRow == len(g[0]) {
			// For part 2: skip if we already found this reflection
			if skipRow == row {
				continue
			}
			return row * 100, row, -1

		}
	}
	//fmt.Println()
	return -1, -1, -1
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func StringSliceEqual(a, b []string) bool {
	for i, v := range a {

		if i >= len(b) {
			continue
		}

		if b[i] != v {
			return false
		}
	}

	return true
}
