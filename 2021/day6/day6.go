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
	count = part1("day6/day6-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day6/day6.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day6/day6-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day6/day6.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

// Simply storing all the fish in a slice and doing a count after 80 days
func part1(path string) int {
	lines := utils.ReadLines(path)

	school := make([]int, 0)
	for _, line := range lines {
		fish := utils.StrArrToIntArr(strings.Split(line, ","))
		for _, fish := range fish {
			school = append(school, fish)
		}
	}

	max := 80

	for i := 0; i < max; i++ {
		schoolCopy := school
		for j, fishTimer := range school {
			schoolCopy[j] = fishTimer - 1
			if schoolCopy[j] == -1 {
				schoolCopy = append(schoolCopy, 8)
				schoolCopy[j] = 6
			}
		}
		school = schoolCopy
	}
	count := 0
	for range school {
		count++
	}

	return count
}

// Storing the fishAmount in a map by age(=timer) instead of storing all the fish
func part2(path string) int {
	lines := utils.ReadLines(path)

	// Make a map for [fishTimer]fishAmount
	school := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, line := range lines {
		fish := utils.StrArrToIntArr(strings.Split(line, ","))
		for _, fishTimer := range fish {
			school[fishTimer]++
		}
	}

	max := 256

	for i := 0; i < max; i++ {
		// First 'age' the fish
		newFish := school[0]
		school[0] = school[1]
		school[1] = school[2]
		school[2] = school[3]
		school[3] = school[4]
		school[4] = school[5]
		school[5] = school[6]
		school[6] = school[7]
		school[7] = school[8]
		// Then give birth to new fish (and re-insert the parents)
		school[8] = newFish
		school[6] += newFish
	}

	count := 0
	for _, amount := range school {
		count += amount
	}

	return count
}
