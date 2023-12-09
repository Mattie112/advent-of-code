package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

func main() {
	var count int
	count = day9Part1("day9/day9-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day9Part2("day9/day9-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day9Part1("day9/day9.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day9Part2("day9/day9.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day9Part1(path string) int {
	answer := 0

	lines := utils.ReadLines(path)
	for _, line := range lines {
		values := strings.Split(line, " ")
		//fmt.Println("")
		//fmt.Println(values)

		var sequences [][]int
		var numbers []int
		for _, v := range values {
			numbers = append(numbers, utils.MustParseStringToInt(v))
		}

		sequences = append(sequences, numbers)
		for {
			differences, isZero := getDifferences(numbers, false)
			//fmt.Println(differences)
			sequences = append(sequences, differences)
			numbers = differences
			if isZero {
				break
			}
		}
		//
		//fmt.Println("")

		// Now we have a zero sequence

		// First add a zero to the last sequence
		sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)
		//fmt.Println(sequences[len(sequences)-1])

		// Then loop through the sequences backwards and add the last number from the previous sequence
		lastAddedNumber := 0
		for i := len(sequences) - 2; i >= 0; i-- {
			// This is just the last number from the previous sequence and adding it to the last number of this sequence
			numberToAdd := sequences[i+1][len(sequences[i+1])-1] + sequences[i][len(sequences[i])-1]
			lastAddedNumber = numberToAdd
			sequences[i] = append(sequences[i], numberToAdd)
			//fmt.Println(sequences[i])
		}
		answer += lastAddedNumber
	}

	return answer
}

func getDifferences(input []int, partTwo bool) (differences []int, isZero bool) {
	isZero = true
	for i, value := range input {
		if i > 0 {
			difference := value - input[i-1]
			if partTwo {
				difference = input[i-1] - value
			}
			differences = append(differences, difference)
			if difference != 0 {
				isZero = false
			}
		}
	}
	return differences, isZero
}

// Part 2 is the same as part 1, but we need to add to the numbers to the front of the sequence instead of the back
func day9Part2(path string) int {
	answer := 0

	lines := utils.ReadLines(path)
	for _, line := range lines {
		values := strings.Split(line, " ")
		//fmt.Println("")
		//fmt.Println(values)

		var sequences [][]int
		var numbers []int
		for _, v := range values {
			numbers = append(numbers, utils.MustParseStringToInt(v))
		}

		sequences = append(sequences, numbers)
		for {
			differences, isZero := getDifferences(numbers, true)
			//fmt.Println(differences)
			sequences = append(sequences, differences)
			numbers = differences
			if isZero {
				break
			}
		}

		//fmt.Println("")

		// Now we have a zero sequence

		// First add a zero to the front of the last sequence
		sequences[len(sequences)-1] = append([]int{0}, sequences[len(sequences)-1]...)
		//fmt.Println(sequences[len(sequences)-1])

		// Then loop through the sequences backwards and add the last number from the previous sequence
		lastAddedNumber := 0
		for i := len(sequences) - 2; i >= 0; i-- {
			// This is just the FIRST number from the previous sequence and adding it to the FIRST number of this sequence
			numberToAdd := sequences[i+1][0] + sequences[i][0]
			lastAddedNumber = numberToAdd
			// Also append it to the front, not back
			sequences[i] = append([]int{numberToAdd}, sequences[i]...)
			//fmt.Println(sequences[i])
		}
		answer += lastAddedNumber
	}

	return answer
}
