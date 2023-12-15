package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var count int
	//count = day15Part1("day15/day15-test.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day15Part2("day15/day15-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = day15Part1("day15/day15.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day15Part2("day15/day15.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day15Part1(path string) int {
	answer := 0

	// First make a grid of the input
	lines := utils.ReadLines(path)
	line := lines[0]

	value := int32(0)

	sequences := strings.Split(line, ",")
	for _, sequence := range sequences {
		for _, char := range sequence {
			value += char
			value *= 17
			value %= 256
		}
		//fmt.Println(sequence, value)
		answer += int(value)
		value = 0
	}

	return answer
}

type lens struct {
	label       string
	focalLength int
}

func day15Part2(path string) int {
	answer := 0
	boxes := make(map[int][]lens)

	// First make a grid of the input
	lines := utils.ReadLines(path)
	line := lines[0]

	value := int32(0)

	sequences := strings.Split(line, ",")
	for _, sequence := range sequences {
		re := regexp.MustCompile(`([a-z]+)(=|-)(\d+)`)
		matches := re.FindStringSubmatch(sequence)
		fmt.Println(matches)

		label := matches[1]
		operator := matches[2]
		focalLength := utils.MustParseStringToInt(matches[3])

		hash := 0
		for _, char := range label {
			value += char
			value *= 17
			value %= 256
		}

		if operator == "-" {
			for id, lensInBox := range boxes[hash] {
				if lensInBox.label == label {
					boxes[hash] = append(boxes[hash][:id], boxes[hash][id+1:]...)
					break
				}
			}
		}
		if operator == "=" {
			// todo if we have a lens with the same label -> replace it with the new one
			// if not -> add to back
			boxes[hash] = append(boxes[hash], lens{label: label, focalLength: focalLength})
		}

		fmt.Println(sequence, value)
		answer += int(value)
		value = 0
	}

	return answer
}
