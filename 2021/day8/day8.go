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
	count = part1("day8/day8-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = part1("day8/day8.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	//count = part2("day8/day8-test.txt")
	//log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = part2("day8/day8.txt")
	//log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	for _, line := range lines {
		split := strings.Split(line, " | ")
		signal := split[0]
		output := split[1]
		log.Debugf("Signal: %s", signal)
		log.Debugf("Output: %s", output)

	}
	return 0
}

func part2(path string) int {

	return 0
}
