package main

import (
	"AoC/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var count int
	count = day2Part1("day3/day3-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day2Part2("day3/day3-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day2Part1("day3/day3.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day2Part2("day3/day3.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day2Part1(path string) int {
	lines := utils.ReadLines(path)

	// First calculate the amount of 1 and 0 for each position
	bitCounts := getBitCounts(lines)

	// Now find the most common bits for the gamma rate
	gammaRate := getGammaRate(bitCounts)

	// And the other way around for the epsilon rate
	epsilonRate := getEpsilonRate(bitCounts)

	g, _ := strconv.ParseInt(gammaRate, 2, 64)
	e, _ := strconv.ParseInt(epsilonRate, 2, 64)

	return int(g * e)
}

func day2Part2(path string) int {
	lines := utils.ReadLines(path)

	// GAMMA
	// First calculate the amount of 1 and 0 for each position
	bitmap := getBitCounts(lines)
	// Now find the most common bits for the gamma rate
	gammaRate := getGammaRate(bitmap)
	// Filter-out all matching lines (so removing the ones that don't match my requested bit)
	gammaArr := strings.Split(gammaRate, "")
	linesToSearchThrough := lines

	for pos, bit := range gammaArr {
		bit = gammaArr[pos] // Not sure why, I do overwrite my gammaArr at the end of this loop but `bit` is still the 'old' value
		newSlice := make([]string, 0)

		for _, line := range linesToSearchThrough {
			strArr := strings.Split(line, "")
			if bit == strArr[pos] {
				newSlice = append(newSlice, line)
			}
		}

		// Get our new bitCounts
		bitmap = getBitCounts(newSlice)
		// Update our lines to search
		linesToSearchThrough = newSlice
		// Get our new gamma rate (aka find the most common bit)
		gammaRate = getGammaRate(bitmap)
		gammaArr = strings.Split(gammaRate, "")
		if len(newSlice) == 1 {
			// We are now done, no need to continue
			gammaRate = newSlice[0]
			break
		}
	}
	// END GAMMA

	// Now I have the oxygen generator rating
	oxyRating, _ := strconv.ParseInt(gammaRate, 2, 64)
	fmt.Println(fmt.Sprintf("Oxygen rating: %d", oxyRating))

	// EPSILON
	// And the other way around for the epsilon rate
	// First calculate the amount of 1 and 0 for each position
	bitmap = getBitCounts(lines)
	// Now find the most common bits for the gamma rate
	epsilonRate := getEpsilonRate(bitmap)
	// Filter-out all matching lines (so removing the ones that don't match my requested bit)
	epsilonArr := strings.Split(epsilonRate, "")
	linesToSearchThrough = lines

	for pos, bit := range epsilonArr {
		bit = epsilonArr[pos] // Not sure why, I do overwrite my epsilonArr at the end of this loop but `bit` is still the 'old' value
		newSlice := make([]string, 0)

		for _, line := range linesToSearchThrough {
			strArr := strings.Split(line, "")
			if bit == strArr[pos] {
				newSlice = append(newSlice, line)
			}
		}

		// Get our new bitCounts
		bitmap = getBitCounts(newSlice)
		// Update our lines to search
		linesToSearchThrough = newSlice
		// Get our new epsilon rate (aka find the least common bit)
		epsilonRate = getEpsilonRate(bitmap)
		epsilonArr = strings.Split(epsilonRate, "")
		if len(newSlice) == 1 {
			// We are now done, no need to continue
			epsilonRate = newSlice[0]
			break
		}
	}
	// END EPSILON

	// Now I have the CO2 scrubber rating
	co2Rating, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Println(fmt.Sprintf("CO2 Scrubber rating: %d", co2Rating))

	return int(oxyRating * co2Rating)
}

func getBitCounts(lines []string) [][]int {
	bitCounts := make([][]int, len(lines[0]))
	for _, line := range lines {
		strArr := strings.Split(line, "")
		for i, char := range strArr {
			bit, _ := strconv.Atoi(char)

			if len(bitCounts[i]) <= bit {
				bitCounts[i] = make([]int, 2)
			}

			bitCounts[i][bit] = bitCounts[i][bit] + 1
		}
	}
	return bitCounts
}

func getGammaRate(bitmap [][]int) string {
	gammaRate := ""
	for _, bit := range bitmap {
		zeroCount := bit[0]
		oneCount := bit[1]
		if zeroCount > oneCount {
			gammaRate += "0"
		} else if oneCount >= zeroCount {
			gammaRate += "1"
		} else {
			panic("yeah no idea here")
		}
	}
	return gammaRate
}

func getEpsilonRate(bitmap [][]int) string {
	epsilonRate := ""
	for _, bit := range bitmap {
		zeroCount := bit[0]
		oneCount := bit[1]
		if zeroCount <= oneCount {
			epsilonRate += "0"
		} else if oneCount < zeroCount {
			epsilonRate += "1"
		} else {
			panic("yeah no idea here")
		}
	}
	return epsilonRate
}
