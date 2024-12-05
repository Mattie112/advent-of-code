package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"slices"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day5/day5-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day5/day5-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day5/day5.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", Part2("day5/day5.txt")))
}

func Part1(path string) int {
	return P1andP2(path, false)
}

func Part2(path string) int {
	return P1andP2(path, true)
}

func P1andP2(path string, part2 bool) int {
	lines := utils.ReadLines(path)
	answer := 0
	orderList := make([][]int, 0)
	pagesList := make([][]int, 0)
	goodList := make([][]int, 0)
	badList := make([][]int, 0)

	// Loop through the lines and split them pages and oders
	firstPart := true
	for _, line := range lines {
		if line == "" {
			firstPart = false
			continue
		}
		if firstPart {
			tmp := strings.Split(line, "|")
			orderList = append(orderList, utils.StringSliceToIntSlice(tmp))
			continue
		} else {
			tmp := strings.Split(line, ",")
			pagesList = append(pagesList, utils.StringSliceToIntSlice(tmp))
			continue
		}
	}

	// Loop through the pages and find the bad ones
	for _, pages := range pagesList {
		isPageGood := true
		for _, order := range orderList {
			first := order[0]
			second := order[1]
			indexFirst := slices.Index(pages, first)
			indexLast := slices.Index(pages, second)

			// If one of the numbers is not found -> assume good
			if indexFirst == -1 || indexLast == -1 {
				continue
			}

			// This sorting is correct
			if indexFirst < indexLast {
				continue
			}

			isPageGood = false
		}
		if isPageGood {
			goodList = append(goodList, pages)
		} else {
			badList = append(badList, pages)
		}
	}

	if part2 {
		// Empty out the goodlist so we can put our re-sorted pages in there
		goodList = make([][]int, 0)

		for i := 0; i < len(badList); i++ {
			pages := badList[i]
			madeChange := false
			for _, order := range orderList {
				first := order[0]
				second := order[1]
				indexFirst := slices.Index(pages, first)
				indexLast := slices.Index(pages, second)

				// If one of the numbers is not found -> assume good
				if indexFirst == -1 || indexLast == -1 {
					continue
				}

				// This sorting is correct
				if indexFirst < indexLast {
					continue
				}

				// Sorting is wrong shall we just try to swap the numbers?
				pages[indexFirst], pages[indexLast] = pages[indexLast], pages[indexFirst]
				madeChange = true
			}

			if madeChange {
				// If we have made a change we need to re-check this page, it might be that we have swapped a number that is now in the wrong place
				// (or there might be other numbers)
				i--
			}

			if !madeChange {
				// If we have not made a change assume it is good now
				goodList = append(goodList, pages)
			}
		}
	}

	// Get the middle page and append them to the answer
	for _, good := range goodList {
		answer += good[len(good)/2]
	}

	return answer
}
