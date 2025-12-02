package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day2Part1("day2/day2-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", day2Part1("day2/day2.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day2Part2("day2/day2-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", day2Part2("day2/day2.txt")))
}

type ID struct {
	first int
	last  int
}

func day2Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	idList := getIdListFromInput(lines)

	for _, id := range idList {
		start := id.first
		end := id.last
		for i := start; i <= end; i++ {
			// Simply cut the number in half and check if they are equal
			originalAsString := fmt.Sprintf("%d", i)
			firstHalf := originalAsString[:len(originalAsString)/2]
			secondHalf := originalAsString[len(originalAsString)/2:]
			if firstHalf == secondHalf {
				answer += i
			}
		}
	}

	return answer
}

func getIdListFromInput(lines []string) []ID {
	idList := make([]ID, 0)

	for _, line := range lines {
		splittedList := strings.Split(line, ",")
		for _, splitted := range splittedList {
			numbers := strings.Split(splitted, "-")
			id := ID{first: utils.MustParseStringToInt(numbers[0]), last: utils.MustParseStringToInt(numbers[1])}
			idList = append(idList, id)
		}
	}
	return idList
}

func day2Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	idList := getIdListFromInput(lines)

	for _, id := range idList {
		start := id.first
		end := id.last
		for i := start; i <= end; i++ {
			originalAsString := fmt.Sprintf("%d", i)
			splitted := strings.Split(originalAsString, "")
			// Loop through the first half of splitted
			for j := 0; j < len(splitted)/2; j++ {
				partA := strings.Join(splitted[0:j+1], "")
				partB := strings.Join(splitted[len(splitted)-j-1:], "")

				if partA == partB {
					// It seems that "1001" -> 1 - 1 is not a valid match but "111" -> 1 - 1 is, so all characters must be used
					// In that case we can just remove the part we found to see if anything is left over, if it is, then we don't count it as a match!
					leftover := strings.ReplaceAll(originalAsString, partA, "")
					if len(leftover) == 0 {
						answer += i
						//fmt.Printf("Found match %d = %s - %s\n", i, partA, partB)
						break
					}
				}
			}
		}
	}

	return answer
}
