package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", day6Part1("day6/day6-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", day6Part1("day6/day6.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", day6Part2("day6/day6-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", day6Part2("day6/day6.txt")))
}

func day6Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	grid := make([][]string, len(strings.Split(lines[0], " "))-1)

	for _, line := range lines {
		// Split each line on space (might be multiple)
		splitted := strings.Split(line, " ")
		itemPos := 0
		for _, item := range splitted {
			if item == "" {
				continue
			}
			grid[itemPos] = append(grid[itemPos], item)
			itemPos++
		}
	}

	for _, row := range grid {
		if row == nil {
			continue
		}
		tmpAnswer := 0
		// Pop the last item of the row and get the operator
		lastItem := row[len(row)-1]
		row = row[:len(row)-1]
		switch lastItem {
		case "*":
			{
				tmpAnswer = 1
				for _, item := range row {
					if item == "" {
						continue
					}
					tmpAnswer *= utils.MustParseStringToInt(item)
				}
			}
		case "+":
			{
				for _, item := range row {
					if item == "" {
						continue
					}
					tmpAnswer += utils.MustParseStringToInt(item)
				}
			}
		}
		answer += tmpAnswer
	}
	return answer
}

func day6Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	linesButSplitted := make([][]string, 0)
	largestLine := 0
	for _, line := range lines {
		linesButSplitted = append(linesButSplitted, strings.Split(line, ""))
		if len(line) > largestLine {
			largestLine = len(line)
		}
	}

	// Loop through the splitted lines but based on columns
	tmpNumbers := make([]string, 0)
	operator := ""
	for i := 0; i < largestLine; i++ {
		tmpString := ""
		emptyColumnCount := 0
		for j := 0; j < len(linesButSplitted); j++ {
			x := ""
			// Not all lines have equal length
			if len(linesButSplitted[j]) > i {
				x = linesButSplitted[j][i]
			}

			// We are now in the empty column
			if x == " " {
				emptyColumnCount++
			}

			// When we have a whole empty column -> do the calculation
			if emptyColumnCount == len(lines) {
				answer += calc(operator, tmpNumbers)
				tmpNumbers = make([]string, 0)
				operator = ""
				emptyColumnCount = 0
				continue
			}

			if x == " " {
				continue
			}

			if x == "*" || x == "+" {
				operator = x
				continue
			}

			tmpString += x
		}
		tmpNumbers = append(tmpNumbers, tmpString)
	}

	// Remember the last one!
	answer += calc(operator, tmpNumbers)

	return answer
}

// As go has no "fancy" stuff like PHP where you can just eval a string like "1*2*3*4" we need to do it all by ourselves
func calc(operator string, row []string) int {
	tmpAnswer := 0
	switch operator {
	case "*":
		{
			tmpAnswer = 1
			for _, item := range row {
				if item == "" {
					continue
				}
				tmpAnswer *= utils.MustParseStringToInt(item)
			}
		}
	case "+":
		{
			for _, item := range row {
				if item == "" {
					continue
				}
				tmpAnswer += utils.MustParseStringToInt(item)
			}
		}
	}
	return tmpAnswer
}
