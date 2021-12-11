package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.DebugLevel)
	var count int
	count = part1("day11/day11-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day11/day11.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day11/day11-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day11/day11.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

var (
	maxX        = 10
	maxY        = 10
	flashCount  = 0
	stepsAmount = 100
)

type octopus struct {
	energy        int
	lastFlashStep int
}

func part1(path string) int {
	lines := utils.ReadLines(path)
	grid := make([][]*octopus, 10)
	flashCount = 0
	for i := range grid {
		grid[i] = make([]*octopus, 10)
	}

	for y, line := range lines {
		row := strings.Split(line, "")
		for x, energy := range row {
			oct := octopus{energy: utils.StringToInt(energy)}
			grid[y][x] = &oct
		}
	}

	log.Debugf("Before any steps")
	draw(grid)

	for step := 1; step <= stepsAmount; step++ {
		// First increase the energy of each octopus
		for _, row := range grid {
			for _, octopus := range row {
				octopus.energy++
			}
		}
		// Then flash octopi with an energy level of > 9
		for y, row := range grid {
			for x, octopus := range row {
				if octopus.energy > 9 && octopus.lastFlashStep < step {
					flash(grid, y, x, step)
				}
			}
		}
		log.Debugf("After step %d", step)
		draw(grid)
	}

	return flashCount
}

func draw(grid [][]*octopus) {
	if log.GetLevel() != log.DebugLevel {
		return
	}
	for _, row := range grid {
		rowStr := ""
		for _, octopus := range row {
			rowStr += fmt.Sprintf("%d", octopus.energy)
		}
		log.Debugf(rowStr)
	}
}

func flash(grid [][]*octopus, y int, x int, step int) {
	// Flash this octopus
	octopus := grid[y][x]
	octopus.energy = 0
	octopus.lastFlashStep = step
	flashCount++
	// Now go and look for other octopi and flash those
	moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, elem := range moves {
		newY := y - elem[0]
		newX := x - elem[1]
		if newY < 0 || newY >= maxY || newX < 0 || newX >= maxX {
			continue
		}
		newOctopus := grid[newY][newX]
		// Only increase energy if not flashed this step (poorly explained in the assignment)
		if newOctopus.lastFlashStep < step {
			newOctopus.energy++
			if newOctopus.energy > 9 {
				flash(grid, newY, newX, step)
			}
		}
	}
}

func part2(path string) int {
	lines := utils.ReadLines(path)
	grid := make([][]*octopus, 10)
	for i := range grid {
		grid[i] = make([]*octopus, 10)
	}

	for y, line := range lines {
		row := strings.Split(line, "")
		for x, energy := range row {
			oct := octopus{energy: utils.StringToInt(energy)}
			grid[y][x] = &oct
		}
	}

	log.Debugf("Before any steps")
	draw(grid)

ContinueHere:
	for step := 1; step <= stepsAmount*99999; step++ {
		// First increase the energy of each octopus
		for _, row := range grid {
			for _, octopus := range row {
				octopus.energy++
			}
		}
		// Then flash octopi with an energy level of > 9
		for y, row := range grid {
			for x, octopus := range row {
				if octopus.energy > 9 && octopus.lastFlashStep < step {
					flash(grid, y, x, step)
				}
			}
		}

		log.Debugf("After step %d", step)
		draw(grid)

		for _, row := range grid {
			for _, octopus := range row {
				if octopus.energy != 0 {
					continue ContinueHere
				}
			}
		}
		return step
	}

	panic("Did not find the answer :( but hey Greetings to Stefan!")
}
