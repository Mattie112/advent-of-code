package main

import (
	"AoC/utils"
	"fmt"
	"github.com/solarlune/paths"
)

func main() {
	var count int
	count = day17Part1("day17/day17-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day17Part2("day17/day17-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = day17Part1("day17/day17.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day17Part2("day17/day17.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

type Node struct {
	X, Y int
	cost int
	prev *Node
}

const (
	LEFT  = "<"
	RIGHT = ">"
	UP    = "^"
	DOWN  = "v"
)

func day17Part1(path string) int {
	answer := 0
	grid := make([][]rune, 0)

	// First make a grid of the input
	lines := utils.ReadLines(path)
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}

	firstMap := paths.NewGridFromRuneArrays(grid, 1, 1)
	for i := 0; i <= 9; i++ {
		firstMap.SetCost([]rune(fmt.Sprintf("%d", i))[0], float64(i))
	}

	firstMap.Get(0, 0).Cost = 0

	firstPath := firstMap.GetPath(0, 0, float64(len(grid[0])-1), float64(len(grid)-1), false, false)

	fmt.Println(firstPath)

	sameDirCounter := 0
	lastDir := RIGHT
	var resetNode paths.Cell
	resetNode.Cost = -1
	for {
		cur := firstPath.Current()

		if resetNode.Cost != -1 {
			firstPath = firstMap.GetPath(float64(cur.X), float64(cur.Y), float64(len(grid[0])-1), float64(len(grid)-1), false, false)
			firstMap.Get(resetNode.X, resetNode.Y).Cost = resetNode.Cost
		}

		answer += int(cur.Cost)

		// What would be our next direction?
		next := firstPath.Next()
		if cur.Y == next.Y {
			if cur.X < next.X {
				if lastDir == RIGHT {
					sameDirCounter++
				} else {
					sameDirCounter = 0
				}
				lastDir = RIGHT
			} else {
				if lastDir == LEFT {
					sameDirCounter++
				} else {
					sameDirCounter = 0
				}
				lastDir = LEFT
			}
		} else {
			if cur.Y < next.Y {
				if lastDir == DOWN {
					sameDirCounter++
				} else {
					sameDirCounter = 0
				}
				lastDir = DOWN
			} else {
				if lastDir == UP {
					sameDirCounter++
				} else {
					sameDirCounter = 0
				}
				lastDir = UP
			}
		}

		// If the directoin would be 3 in a row, then we need to turn
		if sameDirCounter == 3 {
			fmt.Println("same direction")
			resetNode = *next
			resetNode.Cost = next.Cost
			firstPath.Next().Cost = 999
			firstMap.Get(resetNode.X, resetNode.Y).Cost = 999
			continue
		}

		firstPath.Advance()

		fmt.Println("Going from Y:", cur.Y, "X:", cur.X, "to Y:", next.Y, "X:", next.X, "Cost:", cur.Cost, "Next Cost:", next.Cost, "Same Dir Counter:", sameDirCounter, "Last Dir:", lastDir)

		if firstPath.AtEnd() {
			break
		}
	}
	fmt.Println(answer)

	fmt.Println(answer)
	return answer
}

func day17Part2(path string) int {
	return 0
}
