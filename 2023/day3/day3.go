package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

func main() {
	var count int
	count = day3Part1("day3/day3-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day3Part2("day3/day3-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day3Part1("day3/day3.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day3Part2("day3/day3.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

var (
	directions = [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	numbers    = "0123456789"
)

func getShematic(path string) [][]string {
	lines := utils.ReadLines(path)

	schematic := make([][]string, len(lines))

	for y, line := range lines {
		schematic[y] = make([]string, len(line))
		for x, char := range line {
			schematic[y][x] = string(char)
		}
	}
	return schematic
}

func day3Part1(path string) int {
	schematic := getShematic(path)
	answer := 0

	for yPos, y := range schematic {
		xPos := 0
		for ; xPos < len(y); xPos++ {
			char := schematic[yPos][xPos]
			if strings.ContainsAny(char, numbers) {
				// We have found a Number, now we have to find any engine part numbers that may surround it (including diagonals)
				for _, direction := range directions {
					xCheck := xPos + direction[0]
					yCheck := yPos + direction[1]

					if xCheck >= 0 && xCheck < len(schematic[yPos]) && yCheck >= 0 && yCheck < len(schematic) {
						// We are within the bounds of the schematic
						// Check if the character is NOT a number (so a symbol)
						if !strings.ContainsAny(schematic[yCheck][xCheck], numbers) && schematic[yCheck][xCheck] != "." {
							partNumber := schematic[yPos][xPos]
							// We have found the start of the engine part number
							// Now we have to find the complete part number (so look left and right untill we find a non-number)

							// First check left
							for partXCheck := xPos - 1; partXCheck >= 0; partXCheck-- {
								if strings.ContainsAny(schematic[yPos][partXCheck], numbers) {
									partNumber = schematic[yPos][partXCheck] + partNumber
								} else {
									break
								}
							}

							// Then check right
							for partXCheck := xPos + 1; partXCheck < len(schematic[yCheck]); partXCheck++ {
								if strings.ContainsAny(schematic[yPos][partXCheck], numbers) {
									partNumber = partNumber + schematic[yPos][partXCheck]
									xPos++
								} else {
									break
								}
							}
							// Now we have a complete part number, add it to our answer
							answer += utils.MustParseStringToInt(partNumber)
						}
					}
				}
			}
		}
	}
	return answer
}

// Roughly the same as part 1, but instead of finding numbers (and then symbols) we find symbols (and then numbers)
func day3Part2(path string) int {
	schematic := getShematic(path)
	answer := 0

	// loop through all characters and find any simbols (excluding the period)
	for yPos, y := range schematic {
		for xPos := range y {
			char := schematic[yPos][xPos]
			if char == "*" {
				// We have found a symbol, now we have to find any gear part numbers that may surround it (including diagonals)
				gearCount := 0
				var gearValues []string

			directionsLoop:
				for _, direction := range directions {
					xCheck := xPos + direction[0]
					yCheck := yPos + direction[1]

					if xCheck >= 0 && xCheck < len(schematic[yPos]) && yCheck >= 0 && yCheck < len(schematic) {
						// We are within the bounds of the schematic
						// Check if the character is a number
						if strings.ContainsAny(schematic[yCheck][xCheck], numbers) {
							partNumber := schematic[yCheck][xCheck]
							// We have found a gear part number
							// Now we have to find the complete part number (so look left and right untill we find a non-number)
							for partXCheck := xCheck - 1; partXCheck >= 0; partXCheck-- {
								// First check left
								if strings.ContainsAny(schematic[yCheck][partXCheck], numbers) {
									partNumber = schematic[yCheck][partXCheck] + partNumber
								} else {
									break
								}
							}
							for partXCheck := xCheck + 1; partXCheck < len(schematic[yCheck]); partXCheck++ {
								// Then check right
								if strings.ContainsAny(schematic[yCheck][partXCheck], numbers) {
									partNumber = partNumber + schematic[yCheck][partXCheck]
								} else {
									break
								}
							}

							// Let's hope we don't have any duplicate part numbers
							// Otherwise I could try to remove the found gear positions from the "directions" array (or possible use the same xPos++ stuff as in part 1)
							// For cases like:
							// ...*.......
							// ..35.......
							// I will first find the 3 (and add the 5), then find the 5 (and add the 3)
							for _, gearValue := range gearValues {
								if gearValue == partNumber {
									continue directionsLoop
								}
							}
							gearCount++
							gearValues = append(gearValues, partNumber)
						}
					}
				}
				if gearCount == 2 {
					answer += utils.MustParseStringToInt(gearValues[0]) * utils.MustParseStringToInt(gearValues[1])
				}
			}
		}
	}
	return answer
}
