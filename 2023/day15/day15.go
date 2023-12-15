package main

import (
	"AoC/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var count int
	count = day15Part1("day15/day15-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day15Part2("day15/day15-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day15Part1("day15/day15.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day15Part2("day15/day15.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

type lensStruct struct {
	label       string
	focalLength int
}

func day15Part1(path string) int {
	answer := 0

	// First make a grid of the input
	lines := utils.ReadLines(path)
	line := lines[0]

	sequences := strings.Split(line, ",")
	for _, sequence := range sequences {
		//fmt.Println(sequence, value)
		answer += calculateHash(sequence)
	}

	return answer
}

func day15Part2(path string) int {
	answer := 0
	boxes := make(map[int][]lensStruct) // Map so easy access by ID, slice because it is ordered

	// First make a grid of the input
	lines := utils.ReadLines(path)
	line := lines[0]

	sequences := strings.Split(line, ",")
	for _, sequence := range sequences {
		//fmt.Println(sequence)
		re := regexp.MustCompile(`([a-z]*)([=\-])(\d*)`)
		matches := re.FindStringSubmatch(sequence)
		//fmt.Println(matches)

		label := matches[1]
		operator := matches[2]
		focalLength := -1
		if matches[3] != "" {
			focalLength = utils.MustParseStringToInt(matches[3])
		}

		hash := calculateHash(label)

		if operator == "-" {
			//fmt.Println("Removing", label)
			for id, lensInBox := range boxes[hash] {
				if lensInBox.label == label {
					// If it is the last one, delete the entire box
					if len(boxes[hash]) == 1 {
						delete(boxes, hash)
						break
					}
					boxes[hash] = append(boxes[hash][:id], boxes[hash][id+1:]...)
					break
				}
			}
		}
		if operator == "=" {
			//fmt.Println("Adding", label)
			// Either replace or add, not both
			replaced := false
			for id, lensInBox := range boxes[hash] {
				if lensInBox.label == label {
					boxes[hash][id] = lensStruct{label: label, focalLength: focalLength}
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[hash] = append(boxes[hash], lensStruct{label: label, focalLength: focalLength})
			}
		}

		//fmt.Println("After", sequence)
		//for id, lenses := range boxes {
		//	fmt.Print("Box ", id, ": ")
		//	for _, lens := range lenses {
		//		fmt.Print("[", lens.label, " ", lens.focalLength, "]")
		//	}
		//	fmt.Println()
		//}
		//fmt.Println()
	}

	// Calculate the answer
	for boxId, lenses := range boxes {
		for lensId, lens := range lenses {
			answer += (boxId + 1) * (lensId + 1) * lens.focalLength
			//fmt.Println("For the answer: ", lens.label, boxId+1, lensId+1, lens.focalLength, (boxId+1)*(lensId+1)*lens.focalLength)
		}
	}

	return answer
}

func calculateHash(label string) int {
	hash := 0
	value := int32(0)
	for _, char := range label {
		value += char
		value *= 17
		value %= 256
	}
	hash = int(value)
	return hash
}
