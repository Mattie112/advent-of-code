package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)
	start := time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day11/day11-test.txt", 25), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d, took %s", Part2("day11/day11-test.txt", 75), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Answer %d, took %s", Part1("day11/day11.txt", 25), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Answer %d, took %s", Part2("day11/day11.txt", 75), time.Since(start).String()))
}

func Part1(path string, blinks int) int {
	// Part 1 originaly was just doing it the array way :) But yes that was way too slow (17 sec)
	// So for part 2 I rewrote it as the ordering did not even matter!
	return P1andP2(path, blinks)
}

func Part2(path string, blinks int) int {
	return P1andP2(path, blinks)
}

func P1andP2(path string, blinks int) int {
	lines := utils.ReadLines(path)
	answer := 0
	stones := make(map[int]int)

	tmp := strings.Split(lines[0], " ")
	for _, v := range tmp {
		stoneId := utils.MustParseStringToInt(v)
		stones[stoneId] += 1
	}
	if log.GetLevel() == log.DebugLevel {
		log.Debugf("Stones: %v", stones)
	}

	for count := 0; count < blinks; count++ {
		newStones := make(map[int]int)

		for stoneId, stoneAmount := range stones {
			if stoneId == 0 {
				newStones[1] += stoneAmount
				continue
			}

			numberAsString := fmt.Sprintf("%d", stoneId)
			if len(numberAsString)%2 == 0 {
				firstHalf := numberAsString[:len(numberAsString)/2]
				secondHalf := numberAsString[len(numberAsString)/2:]
				newStones[utils.MustParseStringToInt(firstHalf)] += stoneAmount
				newStones[utils.MustParseStringToInt(secondHalf)] += stoneAmount
				continue
			}

			newStones[stoneId*2024] += stoneAmount
		}

		stones = newStones

		if log.GetLevel() == log.DebugLevel {
			log.Debugf("Stones: %v", stones)
			tmpCount := 0
			for _, v := range stones {
				tmpCount += v
			}
			log.Debugf("Count after %d steps: %d", count+1, tmpCount)
		}
	}

	// Count all the stones, that's our answer
	for _, v := range stones {
		answer += v
	}

	return answer
}
