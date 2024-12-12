package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func main() {
	log.SetLevel(log.InfoLevel)
	start := time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day12/day12-test1.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day12/day12-test2.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day12/day12-test3.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d, took %s", Part2("day12/day12-test1.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Answer %d, took %s", Part1("day12/day12.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Answer %d, took %s", Part2("day12/day12.txt"), time.Since(start).String()))
}

type point struct {
	x, y int
}

func Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	grid := map[point]string{}

	for y, line := range lines {
		letters := strings.Split(line, "")
		for x, letter := range letters {
			grid[point{x, y}] = letter
		}
	}

	seen := map[point]bool{}

	for currentPoint := range grid {
		if seen[currentPoint] {
			continue
		}
		seen[currentPoint] = true

		area := 1
		perimeter := 0

		queue := []point{currentPoint}

		for len(queue) > 0 {
			pointToCheck := queue[0]
			queue = queue[1:]

			for _, directions := range []point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				nextPointToCheck := point{pointToCheck.x + directions.x, pointToCheck.y + directions.y}

				// If the next point does not match our current point we are at the edge of the area
				if grid[pointToCheck] != grid[nextPointToCheck] {
					perimeter++

					continue
				}

				// If we have not seen the current point before, add it to the queue and increment the area
				if !seen[nextPointToCheck] {
					seen[nextPointToCheck] = true
					queue = append(queue, nextPointToCheck)
					area++
				}

			}
		}
		answer += area * perimeter
	}

	return answer
}

// Part2 Could be merged with part 1, but I wanted to keep it separate so it is a bit more clear on my thought process
func Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	grid := map[point]string{}

	for y, line := range lines {
		letters := strings.Split(line, "")
		for x, letter := range letters {
			grid[point{x, y}] = letter
		}
	}

	seen := map[point]bool{}
	// Place to store all areas (that are continous)
	areas := map[string]map[point]bool{}

	for currentPoint := range grid {
		if seen[currentPoint] {
			continue
		}
		seen[currentPoint] = true

		area := 1
		perimeter := 0

		queue := []point{currentPoint}
		tmpAreas := map[point]bool{}

		for len(queue) > 0 {
			pointToCheck := queue[0]
			queue = queue[1:]

			for _, directions := range []point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				nextPointToCheck := point{pointToCheck.x + directions.x, pointToCheck.y + directions.y}

				// If the next point does not match our current point we are at the edge of the area
				if grid[pointToCheck] != grid[nextPointToCheck] {
					// (not used for part 2)
					perimeter++
					continue
				}

				// If we have not seen the current point before, add it to the queue and increment the area
				if !seen[nextPointToCheck] {
					seen[nextPointToCheck] = true
					queue = append(queue, nextPointToCheck)
					tmpAreas[nextPointToCheck] = true
					area++
				}

			}
		}

		tmpAreas[currentPoint] = true
		areas[fmt.Sprintf("%d-%d-%s", currentPoint.x, currentPoint.y, grid[currentPoint])] = tmpAreas
	}

	// Go throught all the areas...
	for k, _ := range areas {
		sides := 0
		// ... and then their members
		for v, _ := range areas[k] {
			// Now find the 'corners' of this area
			left := areas[k][point{x: v.x - 1, y: v.y}]
			right := areas[k][point{x: v.x + 1, y: v.y}]
			up := areas[k][point{x: v.x, y: v.y - 1}]
			down := areas[k][point{x: v.x, y: v.y + 1}]
			upleft := areas[k][point{x: v.x - 1, y: v.y - 1}]
			upright := areas[k][point{x: v.x + 1, y: v.y - 1}]
			downleft := areas[k][point{x: v.x - 1, y: v.y + 1}]
			downright := areas[k][point{x: v.x + 1, y: v.y + 1}]

			// For example if we have no members left and up we are in this situation:
			// 0 0 0
			// 0 1 1 <- we are checking the middle 1
			// 0 1 1
			if !left && !up {
				sides++
			}
			if !right && !up {
				sides++
			}
			if !left && !down {
				sides++
			}
			if !right && !down {
				sides++
			}
			if !upright && up && right {
				sides++
			}
			if !upleft && up && left {
				sides++
			}
			if !downleft && down && left {
				sides++
			}
			if !downright && down && right {
				sides++
			}

			log.Debugf("left: %t, right: %t, up: %t, down: %t", left, right, up, down)
			log.Debugf("upleft: %t, upright: %t, downleft: %t, downright: %t", upleft, upright, downleft, downright)
		}
		answer += len(areas[k]) * sides
	}

	return answer
}
