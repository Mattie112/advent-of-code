package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day9/day9-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day9/day9-test2.txt")))
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day9/day9-test.txt")))
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day9/day9.txt")))
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", Part2("day9/day9.txt")))
}

func Part1(path string) int {
	lines := utils.ReadLines(path)
	line := lines[0]
	answer := 0

	diskMap := strings.Split(line, "")
	visual := make([]string, 0)

	fileID := 0
	for i := 0; i < len(diskMap); i++ {
		length := utils.MustParseStringToInt(diskMap[i])
		switch i % 2 {
		case 0:
			visual = append(visual, strings.Split(strings.Repeat(strconv.Itoa(fileID), length), "")...)
			fileID++
		case 1:
			visual = append(visual, strings.Split(strings.Repeat(".", length), "")...)
		}
	}

	log.Debugf("Visual: %s", strings.Join(visual, ""))

	// Now defragment this 'disk' (we use the visual representation to make it easier)
	lastPos := len(visual) - 1
OUTER:
	for i := 0; i < len(visual); i++ {
		if visual[i] == "." {

			// Loop backwards starting from the last position and find any non-dot
			for j := lastPos; j >= 0; j-- {
				for visual[j] != "." && j > i {
					// Now grab the last element and put it here
					visual[i] = visual[j]
					visual[j] = "."
					lastPos = j - 1
					log.Debugf("Visual after step %d: %s", i, strings.Join(visual, ""))
					continue OUTER
				}
			}
		}
	}

	log.Debugf("Visual after sorting: %s", strings.Join(visual, ""))

	// Calculate the checksum

	for i := 0; i < len(visual); i++ {
		if visual[i] == "." {
			break
		}
		answer += i * utils.MustParseStringToInt(visual[i])
	}
	return answer
}

func Part2(path string) int {
	return 0
}
