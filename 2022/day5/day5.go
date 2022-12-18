package main

import (
	"AoC/utils"
	"fmt"
	"github.com/gammazero/deque"
	"regexp"
	"strings"
)

func main() {
	var count int
	count = day5Part1("day5/day5-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	//count = day5Part1("day5/day5.txt")
	//fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	//count = day5Part2("day5/day5-test.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = day5Part2("day5/day5.txt")
	//fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

type move struct {
	amount int
	from   int
	to     int
}

func day5Part1(path string) int {
	lines := utils.ReadLines(path)

	stacks := make(map[int]*deque.Deque[string])
	moves := make([]move, 0)

	// First look for the line that has the amount of columns
	columnAmount := 0
	parseMoves := false
	var regColumnNumbers = regexp.MustCompile(`(?m)\d+(?:   )+`)
	var regCrate = regexp.MustCompile(`(?m)(.){3} ?`)
	var regMove = regexp.MustCompile(`(?m)move (?P<move>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	for _, line := range lines {
		if parseMoves {
			fmt.Println(line)
			matches := regMove.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				move := move{
					amount: utils.MustParseStringToInt(match[regMove.SubexpIndex("move")]),
					from:   utils.MustParseStringToInt(match[regMove.SubexpIndex("from")]),
					to:     utils.MustParseStringToInt(match[regMove.SubexpIndex("to")]),
				}
				moves = append(moves, move)
				fmt.Println(move)
			}
		}

		if regColumnNumbers.MatchString(line) {
			columns := strings.Split(line, "   ")
			columnAmount = utils.MustParseStringToInt(columns[len(columns)-1])
			parseMoves = true
		}

		if !parseMoves {
			// We start with filling our columns
			matches := regCrate.FindAllStringSubmatch(line, -1)
			for i, match := range matches {
				fmt.Println(match[0], "found at index", i)
				token := match[0]
				if token == "    " {
					token = " "
				} else {
					token = strings.Trim(token, "[] ")
					// So now we either have a " " or a letter
					_, ok := stacks[i+1]
					if !ok {
						stacks[i+1] = deque.New[string]()
					}
					stacks[i+1].PushFront(token)
				}
			}
		}
	}
	printStacks(stacks)

	stacks = executeMove(stacks, moves[0])
	printStacks(stacks)

	_ = columnAmount
	_ = moves
	return 0
}

func executeMove(stacks map[int]*deque.Deque[string], move move) map[int]*deque.Deque[string] {

	fromColumn := stacks[move.from]
	toColumn := stacks[move.to]

	for cratesMoved := 0; cratesMoved < move.amount; cratesMoved++ {
		crate := fromColumn.PopFront()
		toColumn.PushFront(crate)
	}

	return stacks
}

func printStacks(input map[int]*deque.Deque[string]) {
	input = utils.CopyMap(input)
	// First find highest
	highest := 0
	for _, crates := range input {
		if crates.Len() > highest {
			highest = crates.Len()
		}
	}
	fmt.Println(fmt.Sprintf("highest: %d", highest))

	for i := 0; i < highest; i++ {
		for j := 0; j < len(input); j++ {
			crates := input[j+1]

			// Empty places are not stored, print a "---" for them
			if crates.Len() == 0 || i+crates.Len() < highest {
				fmt.Print("---")
				continue
			}

			crate := crates.PopBack()
			if crate == " " || crate == "" {
				fmt.Print("---")
			} else {
				fmt.Print("[" + crate + "]")
			}
		}
		fmt.Println("")
	}
	fmt.Println("++++++++++++++++++++++")

}

func day5Part2(path string) int {

	return 0
}
