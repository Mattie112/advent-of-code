package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"regexp"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day14/day14-test.txt")
	log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day14/day14.txt")
	log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day14/day14-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day14/day14.txt")
	log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

// Get the insertions (eg {AB => C} (as it becomes ACB)
func getInsertions(lines []string) map[string]string {
	insertions := map[string]string{}
	var re = regexp.MustCompile(`(?m)(\w+) -> (\w+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		parent := match[1]
		child := match[2]
		insertions[parent] = child
	}
	return insertions
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	// First get the current polymer and remove the empty line
	polymerStr := lines[0]
	polymers := strings.Split(polymerStr, "")
	lines = lines[2:]

	// Get the insertions (eg {AB => C} (as it becomes ACB)
	insertions := getInsertions(lines)

	// Loop over the polymers, get a pair, insert a letter and add everything to a new slice, then do it again for the next step
	// As this creates a large slice this is not suitable for part 2
	for step := 1; step <= 10; step++ {
		polymersCopy := make([]string, len(polymers)+1)
		copy(polymersCopy, polymers)
		insertCount := 0
		for i, polymer := range polymers {
			if i+1 > len(polymers)-1 {
				break
			}
			partStr := fmt.Sprintf("%s%s", polymer, polymers[i+1])
			child := insertions[partStr]
			polymersCopy = append(polymersCopy[:i+1+insertCount], append([]string{child}, polymersCopy[i+1+insertCount:]...)...)
			insertCount++
		}
		polymers = make([]string, len(polymersCopy))
		copy(polymers, polymersCopy)
		log.Debugf("After step %d: %s", step, strings.Join(polymersCopy, ""))
	}

	counter := map[string]int{}
	for _, p := range polymers {
		counter[p]++
	}
	max, min := countResult(counter)
	return max - min
}

func part2(path string) int {
	lines := utils.ReadLines(path)

	// First get the current polymer (as map this time) and remove the empty line
	polymerStr := lines[0]
	polymersTMP := strings.Split(polymerStr, "")
	polymerPairs := map[string]int{}
	// Group polymers per pair, could have also done this in the for loop with the steps by doing +1 there
	for i, p := range polymersTMP {
		if i+1 > len(polymersTMP)-1 {
			break
		}
		polymerPairs[p+polymersTMP[i+1]]++
	}
	lines = lines[2:]

	// Get the insertions (eg {AB => C} (as it becomes ACB)
	insertions := getInsertions(lines)

	// Count per letter (single polymer, not pairs) start with adding our input
	counter := map[string]int{}
	for _, p := range polymersTMP {
		counter[p]++
	}

	// Instead of keeping the entire slice in memory we only keep the count
	// For example:
	// ABCCC == AB=1, BC=1, CC=2
	// Then we can loop through these counts and see comments in the loop
	for step := 1; step <= 40; step++ {
		newPolymers := map[string]int{}
		for key, value := range polymerPairs {
			log.Debugf("%s - %d", key, value)

			// Spit our polymer, find the insertion letter and create 2 new elements (thus duplicating one letter
			// as AB becomes AC and CB if the C gets inserted
			splitPolymer := strings.Split(key, "")
			insertion := insertions[key]
			newA := splitPolymer[0] + insertion
			newB := insertion + splitPolymer[1]

			// Add new elements the value is the amount the original element
			newPolymers[newA] += value
			newPolymers[newB] += value

			// key == "AB" // insertion = "C"
			counter[insertion] += value
		}
		polymerPairs = utils.CopyMapStringInt(newPolymers)
		log.Debugf("After step %d: %+v", step, polymerPairs)
	}

	max, min := countResult(counter)
	return max - min
}

// Function to go from XY = 123 to X = 123 & Y = 123 and then return the min and max
func countResult(polymers map[string]int) (max int, min int) {
	// We don't care anymore for the key (polymers) just the values
	max = 0
	min = math.MaxInt64
	for _, v := range polymers {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
