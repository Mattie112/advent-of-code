package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	var count int
	count = day6Part1("day6/day6-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day6Part2("day6/day6-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day6Part1("day6/day6.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day6Part2("day6/day6.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func parse(path string) [][]int {
	lines := utils.ReadLines(path)
	// Parse the races and add it to the list
	regex := regexp.MustCompile(`\d+`)
	timeList := regex.FindAllString(lines[0], -1)
	distanceList := regex.FindAllString(lines[1], -1)
	var races [][]int
	for i := 0; i < len(timeList); i++ {
		races = append(races, []int{utils.MustParseStringToInt(timeList[i]), utils.MustParseStringToInt(distanceList[i])})
	}
	return races
}

func simulateRace(races [][]int) int {
	answer := 1
	for _, race := range races {
		raceTime := race[0]
		recordDistance := race[1]

		speed := 0
		winningDistance := 0
		distances := make([]int, raceTime+1)
		for i := 0; i <= raceTime; i++ {
			distance := speed * (raceTime - i)
			if distance > recordDistance {
				winningDistance++
			}
			distances[i] = distance
			speed++
		}
		answer *= winningDistance
	}
	return answer
}

func day6Part1(path string) int {
	start := time.Now()

	races := parse(path)
	answer := simulateRace(races)

	fmt.Println(time.Since(start))
	return answer
}

func day6Part2(path string) int {
	start := time.Now()

	races := parse(path)

	// For part 2 instead of having multiple races we just merge them into a single race
	newTime := ""
	newDistance := ""
	for _, race := range races {
		newTime += strconv.Itoa(race[0])
		newDistance += strconv.Itoa(race[1])
	}
	newRace := [][]int{{utils.MustParseStringToInt(newTime), utils.MustParseStringToInt(newDistance)}}

	answer := simulateRace(newRace)

	fmt.Println(time.Since(start))
	return answer
}
