package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day12/day12-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day12/day12.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day12/day12-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day12/day12.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

var (
	cave          = map[string][]string{}
	possiblePaths []string
)

func part1(path string) int {
	lines := utils.ReadLines(path)

	possiblePaths = []string{}
	cave = map[string][]string{}

	for _, line := range lines {
		l := strings.Split(line, "-")
		parent := l[0]
		child := l[1]
		cave[parent] = append(cave[parent], child)
		// Also store the 'reverse' (the way back to the parent)
		cave[child] = append(cave[child], parent)
	}

	findPath("start", map[string]bool{}, []string{})

	for _, path := range possiblePaths {
		log.Debugf(path)
	}

	return len(possiblePaths)
}

func findPath(parent string, visited map[string]bool, thisPath []string) {
	children := cave[parent]
	visited[parent] = true
	thisPath = append(thisPath, parent)
	// Copy visited, so I don't modify it for other routes
	visitedCopy := utils.CopyMap(visited)

	for _, c := range children {
		// Reset visited if I go to the next child (= new path)
		visited = utils.CopyMap(visitedCopy)
		if utils.IsLower(c) && visited[c] == true {
			continue
		}
		if c == "end" {
			thisPath = append(thisPath, c)
			possiblePaths = append(possiblePaths, strings.Join(thisPath, ","))
			continue
		}
		findPath(c, visited, thisPath)
	}
}

func part2(path string) int {
	lines := utils.ReadLines(path)
	possiblePaths = []string{}

	cave = map[string][]string{}

	for _, line := range lines {
		l := strings.Split(line, "-")
		parent := l[0]
		child := l[1]
		cave[parent] = append(cave[parent], child)
		// Also store the 'reverse' (the way back to the parent)
		cave[child] = append(cave[child], parent)
	}

	findPath2("start", map[string]bool{}, []string{}, true)

	for _, path := range possiblePaths {
		log.Debugf(path)
	}

	return len(possiblePaths)
}

// Same as part 1 with the addition of a mayVisitTwice bool
// Thought of storing it in the visited map but hey yolo sometimes work just as good
func findPath2(parent string, visited map[string]bool, thisPath []string, mayVisitTwice bool) {
	children := cave[parent]
	visited[parent] = true
	thisPath = append(thisPath, parent)
	// Copy visited, so I don't modify it for other routes
	visitedCopy := utils.CopyMap(visited)
	visitTwiceCopy := mayVisitTwice

	for _, c := range children {
		// Reset visited if I go to the next child (= new path)
		visited = utils.CopyMap(visitedCopy)
		mayVisitTwice = visitTwiceCopy

		if utils.IsLower(c) && visited[c] == true {
			if mayVisitTwice == false {
				continue
			}
			if c == "start" || c == "end" {
				// Start/end may only be visited once
				continue
			}
			mayVisitTwice = false
		}
		if c == "end" {
			thisPath = append(thisPath, c)
			possiblePaths = append(possiblePaths, strings.Join(thisPath, ","))
			continue
		}
		findPath2(c, visited, thisPath, mayVisitTwice)
	}
}
