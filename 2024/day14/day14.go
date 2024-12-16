package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"regexp"
	"time"
)

func main() {
	log.SetLevel(log.InfoLevel)
	start := time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day14/day14-test.txt", 11, 7, 100), time.Since(start).String()))
	//start = time.Now()
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d, took %s", Part2("day14/day14-test.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Answer %d, took %s", Part1("day14/day14.txt", 101, 103, 100), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Answer %d, took %s", Part2("day14/day14.txt", 101, 103, 100), time.Since(start).String()))
}

type pos struct {
	x int
	y int
}
type robot struct {
	id       int
	position pos
	velocity pos
}

func Part1(path string, width int, height int, seconds int) int {
	lines := utils.ReadLines(path)
	answer := 0
	robotIds := 0
	robots := make([]robot, 0)
	grid := map[pos][]*robot{}
	var re = regexp.MustCompile(`(?m)p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	// Loop through the lines and split them into left and right
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		r := robot{
			id: robotIds,
			position: pos{
				x: utils.MustParseStringToInt(matches[0][1]),
				y: utils.MustParseStringToInt(matches[0][2]),
			},
			velocity: pos{
				x: utils.MustParseStringToInt(matches[0][3]),
				y: utils.MustParseStringToInt(matches[0][4]),
			},
		}

		robots = append(robots, r)
		grid[r.position] = append(grid[r.position], &r)
		robotIds++
	}

	log.Debugln("Initial state:")
	printGrid(grid, width, height)

	robotOrder := make([]int, len(robots))
	for k, v := range robots {
		robotOrder[v.id] = k
	}

	// Now move the robots
	for i := 1; i <= seconds; i++ {
		for k := range robotOrder {
			r := &robots[k]
			newXPos := (r.position.x + r.velocity.x) % width
			newYPos := (r.position.y + r.velocity.y) % height
			if newXPos < 0 {
				newXPos = width + newXPos
			}
			if newYPos < 0 {
				newYPos = height + newYPos
			}
			for rInCellId, rInCell := range grid[r.position] {
				if rInCell.id == r.id {
					grid[r.position] = append(grid[r.position][:rInCellId], grid[r.position][rInCellId+1:]...)
				}
			}
			r.position = pos{newXPos, newYPos}
			//robots[k] = r
			grid[r.position] = append(grid[r.position], r)
		}
		log.Debugf("After %d seconds:\n", i)
		printGrid(grid, width, height)
	}

	// Devide the grid in 4 squares, count all the robot in each square
	halfX := int(math.Floor(float64(width / 2)))
	halfY := int(math.Floor(float64(height / 2)))

	gridTopLeft := 0
	// Now count the number of robots in each square
	for y := 0; y < halfY; y++ {
		for x := 0; x < halfX; x++ {
			if _, ok := grid[pos{x, y}]; ok {
				gridTopLeft += len(grid[pos{x, y}])
			}
		}
	}

	gridTopRight := 0
	// Now count the number of robots in each square
	for y := 0; y < halfY; y++ {
		for x := halfX + 1; x < width; x++ {
			if _, ok := grid[pos{x, y}]; ok {
				gridTopRight += len(grid[pos{x, y}])
			}
		}
	}

	gridBottomLeft := 0
	// Now count the number of robots in each square
	for y := halfY + 1; y < height; y++ {
		for x := 0; x < halfX; x++ {
			if _, ok := grid[pos{x, y}]; ok {
				gridBottomLeft += len(grid[pos{x, y}])
			}
		}
	}

	gridBottomRight := 0
	// Now count the number of robots in each square
	for y := halfY + 1; y < height; y++ {
		for x := halfX + 1; x < width; x++ {
			if _, ok := grid[pos{x, y}]; ok {
				gridBottomRight += len(grid[pos{x, y}])
			}
		}
	}

	answer = gridTopLeft * gridTopRight * gridBottomLeft * gridBottomRight

	return answer
}

func Part2(path string, width int, height int, seconds int) int {
	lines := utils.ReadLines(path)
	robotIds := 0
	robots := make([]robot, 0)
	grid := map[pos][]*robot{}
	var re = regexp.MustCompile(`(?m)p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	// Loop through the lines and split them into left and right
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		r := robot{
			id: robotIds,
			position: pos{
				x: utils.MustParseStringToInt(matches[0][1]),
				y: utils.MustParseStringToInt(matches[0][2]),
			},
			velocity: pos{
				x: utils.MustParseStringToInt(matches[0][3]),
				y: utils.MustParseStringToInt(matches[0][4]),
			},
		}

		robots = append(robots, r)
		grid[r.position] = append(grid[r.position], &r)
		robotIds++
	}

	log.Debugln("Initial state:")
	printGrid(grid, width, height)

	robotOrder := make([]int, len(robots))
	for k, v := range robots {
		robotOrder[v.id] = k
	}

	// Now move the robots
	for i := 1; i <= 99999+seconds; i++ {
		for k := range robotOrder {
			r := &robots[k]
			newXPos := (r.position.x + r.velocity.x) % width
			newYPos := (r.position.y + r.velocity.y) % height
			if newXPos < 0 {
				newXPos = width + newXPos
			}
			if newYPos < 0 {
				newYPos = height + newYPos
			}
			for rInCellId, rInCell := range grid[r.position] {
				if rInCell.id == r.id {
					grid[r.position] = append(grid[r.position][:rInCellId], grid[r.position][rInCellId+1:]...)
				}
			}
			r.position = pos{newXPos, newYPos}
			//robots[k] = r
			grid[r.position] = append(grid[r.position], r)
		}

		// try to find grouped robots (this is for part 2)
		for y := 0; y < height; y++ {
		OUTER:
			for x := 0; x < width; x++ {
				groupedRobots := 0

				if _, ok := grid[pos{x, y}]; ok {
					for newY := y - 4; newY < y+4; newY++ {
						for newX := x - 4; newX < x+4; newX++ {
							if tmp, tmpOk := grid[pos{newX, newY}]; tmpOk {
								if len(tmp) > 0 {
									groupedRobots++
								}
							} else {
								continue OUTER
							}
						}
					}
				}
				if groupedRobots > 25 {
					log.Debugf("After %d seconds:\n", i)
					printGrid(grid, width, height)
					log.Debugf("Found at %d, %d - step %d \n", x, y, i)
					return i
				}
			}
		}
		log.Debugf("checked %d\n", i)
	}

	return -1
}

func printGrid(robots map[pos][]*robot, width int, height int) {
	if log.GetLevel() == log.DebugLevel {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if _, ok := robots[pos{x, y}]; ok {
					if len(robots[pos{x, y}]) != 0 {
						fmt.Print(len(robots[pos{x, y}]))
					} else {
						fmt.Print(".")
					}
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()

	}
}
