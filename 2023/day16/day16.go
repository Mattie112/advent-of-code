package main

import (
	"AoC/utils"
	"fmt"
	"math"
)

func main() {
	var count int
	count = day16Part1("day16/day16-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day16Part2("day16/day16-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day16Part1("day16/day16.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day16Part2("day16/day16.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

const RIGHT = ">"
const DOWN = "v"
const LEFT = "<"
const UP = "^"

type coordinateStruct struct {
	x         int
	y         int
	direction string
}

func day16Part1(path string) int {
	grid := makeGrid(path)
	answer := letThereBeLight(grid, coordinateStruct{0, 0, RIGHT})

	return answer
}

func day16Part2(path string) int {
	grid := makeGrid(path)

	maxTilesEnergized := math.MinInt

	// Get the dimensions of the grid, so we can loop over the edges
	numRows := len(grid)
	numCols := len(grid[0])

	// Loop over the top/bottom edge (y = 0)
	for x := 0; x < numCols; x++ {
		maxTilesEnergized = max(maxTilesEnergized, letThereBeLight(grid, coordinateStruct{x, 0, DOWN}))
		maxTilesEnergized = max(maxTilesEnergized, letThereBeLight(grid, coordinateStruct{x, numRows - 1, UP}))
	}

	// Loop over the left/right edge (x = 0)
	for y := 1; y < numRows-1; y++ {
		maxTilesEnergized = max(maxTilesEnergized, letThereBeLight(grid, coordinateStruct{0, y, RIGHT}))
		maxTilesEnergized = max(maxTilesEnergized, letThereBeLight(grid, coordinateStruct{numCols - 1, y, LEFT}))
	}

	return maxTilesEnergized
}

func makeGrid(path string) [][]string {
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
	return grid
}

func decideNextTile(grid [][]string, todo coordinateStruct) (newTodos []coordinateStruct) {
	cell := grid[todo.y][todo.x]
	switch cell {
	case ".":
		switch todo.direction {
		case RIGHT:
			newTodos = append(newTodos, coordinateStruct{todo.x + 1, todo.y, RIGHT})
		case DOWN:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y + 1, DOWN})
		case LEFT:
			newTodos = append(newTodos, coordinateStruct{todo.x - 1, todo.y, LEFT})
		case UP:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y - 1, UP})
		}
	case "/":
		switch todo.direction {
		case RIGHT:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y - 1, UP})
		case DOWN:
			newTodos = append(newTodos, coordinateStruct{todo.x - 1, todo.y, LEFT})
		case LEFT:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y + 1, DOWN})
		case UP:
			newTodos = append(newTodos, coordinateStruct{todo.x + 1, todo.y, RIGHT})
		}
	case "\\":
		switch todo.direction {
		case RIGHT:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y + 1, DOWN})
		case DOWN:
			newTodos = append(newTodos, coordinateStruct{todo.x + 1, todo.y, RIGHT})
		case LEFT:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y - 1, UP})
		case UP:
			newTodos = append(newTodos, coordinateStruct{todo.x - 1, todo.y, LEFT})
		}
	case "|":
		switch todo.direction {
		case RIGHT:
			fallthrough
		case LEFT:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y - 1, UP})
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y + 1, DOWN})
		case DOWN:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y + 1, DOWN})
		case UP:
			newTodos = append(newTodos, coordinateStruct{todo.x, todo.y - 1, UP})
		}
	case "-":
		switch todo.direction {
		case RIGHT:
			newTodos = append(newTodos, coordinateStruct{todo.x + 1, todo.y, RIGHT})
		case LEFT:
			newTodos = append(newTodos, coordinateStruct{todo.x - 1, todo.y, LEFT})
		case DOWN:
			fallthrough
		case UP:
			newTodos = append(newTodos, coordinateStruct{todo.x - 1, todo.y, LEFT})
			newTodos = append(newTodos, coordinateStruct{todo.x + 1, todo.y, RIGHT})
		}
	}

	// Remove any todos that are out of bounds
	for i, newTodo := range newTodos {
		if newTodo.x < 0 || newTodo.y < 0 || newTodo.x >= len(grid[0]) || newTodo.y >= len(grid) {
			//fmt.Println("Removing todo y:", newTodo.y, " ", "x:", newTodo.x, " direction:", newTodo.direction)
			newTodos = append(newTodos[:i], newTodos[i+1:]...)
		}
	}

	return newTodos
}

func letThereBeLight(grid [][]string, start coordinateStruct) int {
	todos := []coordinateStruct{start}
	tilesVisited := make(map[string]string) // I had this first for debug purposes (to print) but it is also needed to prevent endless loops :)
	for len(todos) > 0 {
		// Get the first todo and store it in the visited map (with the direction)
		todo := todos[0]
		tilesVisited[fmt.Sprintf("%d,%d", todo.x, todo.y)] = todo.direction

		// Remove the first todo from the list
		todos = todos[1:]

		// Get our new todos based in the action of this tile
		newTodos := decideNextTile(grid, todo)

		//for y, row := range grid {
		//	for x, char := range row {
		//		if tilesVisited[fmt.Sprintf("%d,%d", x, y)] != "" && char == "." {
		//			fmt.Print(tilesVisited[fmt.Sprintf("%d,%d", x, y)])
		//		} else {
		//			fmt.Print(char)
		//		}
		//	}
		//	fmt.Println()
		//}
		//fmt.Println()

		// Any place we already visited in the same direction -> remove it (to prevent endless loops)
		for i, newTodo := range newTodos {
			if tilesVisited[fmt.Sprintf("%d,%d", newTodo.x, newTodo.y)] == newTodo.direction {
				if len(newTodos) == 1 {
					newTodos = []coordinateStruct{}
				} else {
					newTodos = append(newTodos[:i], newTodos[i+1:]...)
				}
			}
		}

		todos = append(todos, newTodos...)
	}
	return len(tilesVisited)
}
