package main

import (
	"AoC/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var count int
	count = day1Part1("day1/day1-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day1Part2("day1/day1-test-part2.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day1Part1("day1/day1.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day1Part2("day1/day1.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day1Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	for _, line := range lines {
		number := ""

		// Loop through line to find the first number
		for _, char := range line {
			if strings.ContainsAny(string(char), "0123456789") {
				number = string(char)
				break
			}
		}

		// Loop backwards through line to find the last number
		for i, char := range line {
			i = len(line) - 1 - i
			char = int32(line[i])
			if strings.ContainsAny(string(char), "0123456789") {
				number = number + string(char)
				break
			}
		}

		// If we have 2 numbers add it to the answer
		if len(number) == 2 {
			t, _ := strconv.Atoi(number)
			answer += t
		}
	}

	return answer
}

func day1Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0
	for _, line := range lines {
		// Array to add our numbers
		arr := make([]string, 0)

		// Go through each line
	lineForEach:
		for positionInLine := 0; positionInLine < len(line); {
			positionSearchString := 1 // Internal counter to keep track of the search part in this line
			for {
				// Start a for loop for as long as we need (per letter) but break (the outermost loop) if we try to go out of bounds
				if positionInLine+positionSearchString >= len(line)+1 {
					break lineForEach
				}

				// This is the part we want to look for numbers or textNumbers
				searchString := line[positionInLine : positionInLine+positionSearchString]

				// Try to findStringAsNumber a "one" == 1 ...
				indexText, numberString := findStringAsNumber(searchString)
				// ... or just a regular number
				indexNumber := strings.IndexAny(searchString, "0123456789")

				if indexText == -1 && indexNumber == -1 {
					// If we did not findStringAsNumber any number or complete text try to look one character further in this line
					positionSearchString++
				}
				if indexText > -1 {
					// We have found a textual number, add it to our list (so "one" => "1")
					arr = append(arr, numberString)
					// Subtract additional letter just on case we have overlap (eg `twone` == two & one)
					positionInLine += positionSearchString - 1
					break
				}
				if indexNumber > -1 {
					// Our search string might be `gdodjrghj3` when we trigger in here, so we just take the last character and ignore the rest
					arr = append(arr, string(searchString[indexNumber]))
					positionInLine += positionSearchString
					break
				}
			}
		}

		// Take the first and last number and add them together
		answer += utils.MustParseStringToInt(arr[0] + arr[len(arr)-1])
	}
	return answer
}

func findStringAsNumber(input string) (index int, numberAsString string) {
	find := strings.Index(input, "one")
	if find > -1 {
		return find, "1"
	}
	find = strings.Index(input, "two")
	if find > -1 {
		return find, "2"
	}
	find = strings.Index(input, "three")
	if find > -1 {
		return find, "3"
	}
	find = strings.Index(input, "four")
	if find > -1 {
		return find, "4"
	}
	find = strings.Index(input, "five")
	if find > -1 {
		return find, "5"
	}
	find = strings.Index(input, "six")
	if find > -1 {
		return find, "6"
	}
	find = strings.Index(input, "seven")
	if find > -1 {
		return find, "7"
	}
	find = strings.Index(input, "eight")
	if find > -1 {
		return find, "8"
	}
	find = strings.Index(input, "nine")
	if find > -1 {
		return find, "9"
	}
	return -1, ""
}
