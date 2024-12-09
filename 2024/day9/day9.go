package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.SetLevel(log.InfoLevel)
	start := time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day9/day9-test.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day9/day9-test2.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d, took %s", Part2("day9/day9-test.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Answer %d, took %s", Part1("day9/day9.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Answer %d, took %s", Part2("day9/day9.txt"), time.Since(start).String()))
}

type file struct {
	ID   int
	char string
}

func Part1(path string) int {
	lines := utils.ReadLines(path)
	line := lines[0]
	answer := 0

	diskMap := strings.Split(line, "")
	visual := make([]string, 0)  // To help debugging
	nonVisual := make([]file, 0) // Actual values we can use for our calculation

	fileID := 0
	for i := 0; i < len(diskMap); i++ {
		length := utils.MustParseStringToInt(diskMap[i])
		switch i % 2 {
		case 0:
			for j := 0; j < length; j++ {
				nonVisual = append(nonVisual, file{ID: fileID, char: "x"})
			}
			// Note that values > 10 are truncated :)
			visual = append(visual, strings.Split(strings.Repeat(strconv.Itoa(fileID%10), length), "")...)
			fileID++
		case 1:
			for j := 0; j < length; j++ {
				nonVisual = append(nonVisual, file{ID: -1, char: "."})
			}
			visual = append(visual, strings.Split(strings.Repeat(".", length), "")...)
		}
	}

	if log.GetLevel() == log.DebugLevel {
		log.Debugf("Visual: %s", strings.Join(visual, ""))
	}

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
					nonVisual[i] = nonVisual[j]

					// Remove the element we just moved (we don't really care about these)
					visual = append(visual[:j])
					nonVisual = append(nonVisual[:j])

					lastPos = j - 1
					if log.GetLevel() == log.DebugLevel {
						// Leaving this as-is with the log level set to info has a runtime of 1m15s for part1
						// with this one extra IF it is 31ms
						log.Debugf("Visual after step %d: %s", i, strings.Join(visual, ""))
					}
					continue OUTER
				}
			}
		}
	}

	if log.GetLevel() == log.DebugLevel {
		log.Debugf("Visual after sorting: %s", strings.Join(visual, ""))
	}

	// Calculate the checksum (we could do this in one go in the loop above but this is easier to read)
	for i := 0; i < len(nonVisual); i++ {
		if nonVisual[i].char == "." {
			break
		}
		answer += i * (nonVisual[i].ID)
	}

	return answer
}

func Part2(path string) int {
	return 0
}
