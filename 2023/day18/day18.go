package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
)

func main() {
	var count int
	//count = day18Part1("day18/day18-test.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day18Part2("day18/day18-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day18Part1("day18/day18.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day18Part2("day18/day18.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

type job struct {
	direction string
	distance  int
	color     string
}

type cell struct {
	x, y  int
	color string
	char  string
}

func day18Part1(path string) int {
	answer := 0
	// D 5 (#0dc571)
	jobs := make([]job, 0)
	re := regexp.MustCompile(`([RLDU]).*(\d+).*(#[0-9a-f]{6})`)
	lines := utils.ReadLines(path)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		jobs = append(jobs, job{
			direction: matches[1],
			distance:  utils.MustParseStringToInt(matches[2]),
			color:     matches[3],
		})
	}

	size := 15000
	grid := make([][]cell, size)
	for y, row := range grid {
		row = make([]cell, size)
		grid[y] = row
		for x := range row {
			grid[y][x] = cell{
				x:     x,
				y:     y,
				color: "",
				char:  ".",
			}
		}
	}

	startX, startY := 0, 0
	for _, job := range jobs {
		grid, startY, startX = process(job, grid, startY, startX)
	}

	// Print the grid
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Print(grid[y][x].char)
			if grid[y][x].char == "#" {
				answer++
			}
		}
		fmt.Println()
	}

	// Now do a flood-fill for the grid

	return answer
}

func process(job job, grid [][]cell, startY int, startX int) (newGrid [][]cell, lastY int, lastX int) {
	switch job.direction {
	case "R":
		for i := 0; i < job.distance; i++ {
			grid[startY][startX+i].char = "#"
		}
		startX += job.distance
	case "L":
		for i := 0; i < job.distance; i++ {
			grid[startY][startX-i].char = "#"
		}
		startX -= job.distance
	case "U":
		for i := 0; i < job.distance; i++ {
			grid[startY-i][startX].char = "#"
		}
		startY -= job.distance

	case "D":
		for i := 0; i < job.distance; i++ {
			grid[startY+i][startX].char = "#"
		}
		startY += job.distance
	}
	return grid, startY, startX
}

func day18Part2(path string) int {

	return 0
}
