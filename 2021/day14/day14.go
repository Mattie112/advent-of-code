package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"sort"
	"strings"
)

func main() {
	log.SetLevel(log.DebugLevel)
	var count int
	//count = part1("day14/day14-test.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = part1("day14/day14.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day14/day14-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = part2("day14/day14.txt")
	//log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	polymerStr := lines[0]
	polymers := strings.Split(polymerStr, "")
	lines = lines[2:]

	insertions := map[string]string{}

	var re = regexp.MustCompile(`(?m)(\w+) -> (\w+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		parent := match[1]
		child := match[2]
		insertions[parent] = child
	}

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
		//log.Debugf("After step %d", step)
	}

	counter := map[string]int{}
	for _, p := range polymers {
		counter[p]++
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range counter {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss[0].Value - ss[len(ss)-1].Value
}

func part2(path string) int {
	lines := utils.ReadLines(path)

	polymerStr := lines[0]
	polymers := strings.Split(polymerStr, "")
	lines = lines[2:]

	insertions := map[string]string{}

	var re = regexp.MustCompile(`(?m)(\w+) -> (\w+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		parent := match[1]
		child := match[2]
		insertions[parent] = child
	}

	for step := 1; step <= 40; step++ {
		polymersCopy := make([]string, len(polymers))
		insertCount := 0
		for i, polymer := range polymers {
			if i+1 > len(polymers)-1 {
				break
			}
			partStr := fmt.Sprintf("%s%s", polymer, polymers[i+1])
			child := insertions[partStr]
			//polymersCopy = append(polymersCopy[:i+1+insertCount], append([]string{child}, polymersCopy[i+1+insertCount:]...)...)
			polymersCopy = append(polymersCopy, polymers[i]+child+polymers[i+1])
			insertCount++
		}

		polymers = make([]string, len(polymersCopy))
		copy(polymers, polymersCopy)
		log.Debugf("After step %d: %s", step, strings.Join(polymersCopy, ""))
		nbbCount := 0
		log.Debugf("After step %d: NBBCount: %d, polymerLen: %d, answerCount :%d", step, nbbCount, len(polymers), 0)
		//log.Debugf("After step %d", step)
	}

	return countResult(polymers)
}

func countResult(polymers []string) int {
	counter := map[string]int{}
	for _, p := range polymers {
		counter[p]++
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range counter {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss[0].Value - ss[len(ss)-1].Value
}

func countNBB(input []string) int {
	count := 0
	for i, s := range input {
		if i+1 > len(input)-1 || i+2 > len(input)-1 {
			break
		}
		substr := s + input[i+1] + input[i+2]
		if substr == "NBB" {
			count++
		}
	}
	return count
}
