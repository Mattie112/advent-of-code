package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day8/day8-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test2 Answer %d", count))
	count = part1("day8/day8-test2.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test2 Answer %d", count))
	count = part1("day8/day8.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day8/day8-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day8/day8-test2.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test2 Answer %d", count))
	count = part2("day8/day8.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	simpleDigitCount := 0
	for _, line := range lines {
		split := strings.Split(line, " | ")
		signal := split[0]
		output := split[1]
		log.Debugf("Signal: %s", signal)
		log.Debugf("Output: %s", output)
		signalPatterns := strings.Split(output, " ")
		for _, pattern := range signalPatterns {
			length := len(pattern)
			switch length {
			case 2:
				log.Debugf("%s == %d", pattern, 1)
				simpleDigitCount++
			case 3:
				log.Debugf("%s == %d", pattern, 7)
				simpleDigitCount++
			case 4:
				log.Debugf("%s == %d", pattern, 4)
				simpleDigitCount++
			case 7:
				log.Debugf("%s == %d", pattern, 8)
				simpleDigitCount++
			}
		}

	}
	return simpleDigitCount
}

func part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	for _, line := range lines {
		split := strings.Split(line, " | ")
		signal := split[0]
		output := split[1]
		log.Debugf("Signal: %s", signal)
		log.Debugf("Output: %s", output)
		signalPatterns := strings.Split(signal, " ")
		onePattern := []string{"", ""}          // This is "1"
		fourPattern := []string{"", "", "", ""} // This is "4"
		lPattern := make([]string, 0)           // This is "4" - "1" just draw it, and you will see
		simpleNumberCount := 0
		tryAgain := make([]string, 0)
		solvedPatterns := map[string]int{}

		// Partly the same as part 1, first search and find simpleNumbers
		// Then keep looping until we have found all numbers
	RunAgain:
		for _, pattern := range signalPatterns {
			// We do sort the pattern here (a, b, c, ...) to make it easier to parse in the end
			p := strings.Split(pattern, "")
			sort.Strings(p)
			pattern = strings.Join(p, "")
			switch len(pattern) {
			case 2:
				log.Debugf("%s == %d", pattern, 1)
				onePattern = []string{p[0], p[1]}
				simpleNumberCount++
				solvedPatterns[pattern] = 1
			case 3:
				log.Debugf("%s == %d", pattern, 7)
				simpleNumberCount++
				solvedPatterns[pattern] = 7
			case 4:
				log.Debugf("%s == %d", pattern, 4)
				fourPattern = []string{p[0], p[1], p[2], p[3]}
				simpleNumberCount++
				solvedPatterns[pattern] = 4
			case 7:
				log.Debugf("%s == %d", pattern, 8)
				simpleNumberCount++
				solvedPatterns[pattern] = 8
			case 5:
				if simpleNumberCount != 4 {
					tryAgain = append(tryAgain, pattern)
					continue
				}
				// This can be 2 || 3 || 5
				// The only other digit that utils.Contains all of 1 (and more) is 3
				if utils.Contains(onePattern[0], p) && utils.Contains(onePattern[1], p) {
					log.Debugf("Found number 3 I think %s", pattern)
					solvedPatterns[pattern] = 3
				} else if len(lPattern) > 0 && utils.Contains(lPattern[0], p) && utils.Contains(lPattern[1], p) {
					// The number 5 has the L pattern in it
					log.Debugf("Found number 5 I think %s", pattern)
					solvedPatterns[pattern] = 5
				} else {
					// And the number 2 has nothing
					log.Debugf("Found number 2 I think %s", pattern)
					solvedPatterns[pattern] = 2
				}
			case 6:
				if simpleNumberCount != 4 {
					tryAgain = append(tryAgain, pattern)
					continue
				}
				// This can be 6 || 9 || 0
				if utils.Contains(fourPattern[0], p) && utils.Contains(fourPattern[1], p) && utils.Contains(fourPattern[2], p) && utils.Contains(fourPattern[3], p) {
					// The number 9 has the 4 pattern in it
					log.Debugf("Found number 9 I think %s", pattern)
					solvedPatterns[pattern] = 9
				} else if len(lPattern) > 0 && utils.Contains(lPattern[0], p) && utils.Contains(lPattern[1], p) {
					// The number 6 has the L pattern in it
					log.Debugf("Found number 6 I think %s", pattern)
					solvedPatterns[pattern] = 6
				} else {
					// And the number 0 has nothing
					log.Debugf("Found number 0 I think %s", pattern)
					solvedPatterns[pattern] = 0
				}
			default:
				tryAgain = append(tryAgain, pattern)
			}

			// If we have 1 and 4 we can get the lPattern
			if simpleNumberCount == 4 && onePattern[0] != "" && fourPattern[0] != "" && len(lPattern) == 0 {
				for _, p := range fourPattern {
					if !utils.Contains(p, onePattern) {
						lPattern = append(lPattern, p)
					}
				}
				log.Debugf("Found L pattern: %+v", lPattern)
			}
		}

		if len(solvedPatterns) != 10 {
			signalPatterns = tryAgain
			tryAgain = make([]string, 0)
			goto RunAgain
		} else {
			// We have 10 numbers, calculate the output
			log.Debugf("Found %d numbers", 10)
			log.Debugf("%+v", solvedPatterns)
			numberAsString := ""
			for _, o := range strings.Split(output, " ") {
				// Split the output, sort it and then join it together for an easy lookup
				sorted := strings.Split(o, "")
				sort.Strings(sorted)
				o = strings.Join(sorted, "")
				numberAsString += fmt.Sprintf("%d", solvedPatterns[o])
			}
			number := utils.StringToInt(numberAsString)
			answer += number
		}
	}

	return answer
}
