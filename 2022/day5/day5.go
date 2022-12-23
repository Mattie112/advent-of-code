package main

import (
	"AoC/utils"
	"fmt"
	"github.com/gammazero/deque"
	"regexp"
	"strings"
)

func main() {
	var crateString string
	crateString = day5Part1("day5/day5-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %s", crateString))
	crateString = day5Part1("day5/day5.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %s", crateString))
	crateString = day5Part2("day5/day5-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %s", crateString))
	crateString = day5Part2("day5/day5.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %s", crateString))
}

type move struct {
	amount int
	from   int
	to     int
}

func day5Part1(path string) string {
	lines := utils.ReadLines(path)

	stacks, moves := prepare(lines)

	//fmt.Println("")
	//printStacks(stacks)
	for _, move := range moves {
		stacks = executeMove(stacks, move)
		//printStacks(stacks)
	}
	//printStacks(stacks)

	return getAnswer(stacks)
}

func prepare(lines []string) (map[int]*deque.Deque[string], []move) {
	stacks := make(map[int]*deque.Deque[string])
	moves := make([]move, 0)
	// First look for the line that has the amount of columns
	parseMoves := false
	var regColumnNumbers = regexp.MustCompile(`(?m)\d+(?:   )+`)
	var regCrate = regexp.MustCompile(`(?m)(.){3} ?`)
	var regMove = regexp.MustCompile(`(?m)move (?P<move>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	for _, line := range lines {
		if parseMoves {
			matches := regMove.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				move := move{
					amount: utils.MustParseStringToInt(match[regMove.SubexpIndex("move")]),
					from:   utils.MustParseStringToInt(match[regMove.SubexpIndex("from")]),
					to:     utils.MustParseStringToInt(match[regMove.SubexpIndex("to")]),
				}
				moves = append(moves, move)
			}
		}

		if regColumnNumbers.MatchString(line) {
			parseMoves = true
		}

		if !parseMoves {
			// We start with filling our columns
			matches := regCrate.FindAllStringSubmatch(line, -1)
			for i, match := range matches {
				token := match[0]
				token = strings.Trim(token, "[] ")
				if token == "" {
					token = " "
				} else {
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
	return stacks, moves
}

func getAnswer(stacks map[int]*deque.Deque[string]) string {
	crateString := ""
	for i := 0; i < len(stacks); i++ {
		crates := stacks[i+1]
		crateString += crates.PopBack()
	}
	return crateString
}

func executeMove(stacks map[int]*deque.Deque[string], move move) map[int]*deque.Deque[string] {
	fromColumn := stacks[move.from]
	toColumn := stacks[move.to]

	for cratesMoved := 0; cratesMoved < move.amount; cratesMoved++ {
		crate := fromColumn.PopBack()
		toColumn.PushBack(crate)
	}

	return stacks
}

func printStacks(input map[int]*deque.Deque[string]) {
	// First find highest
	highest := 0
	for _, crates := range input {
		if crates.Len() > highest {
			highest = crates.Len()
		}
	}

	// Could not get it to go nicely in the same orientation so this is rotated 90deg to the right
	for i := 0; i < len(input); i++ {
		crates := input[i+1]
		for j := 0; j < crates.Len(); j++ {
			crate := crates.At(j)
			fmt.Print("[" + crate + "]")
		}
		fmt.Println("")
	}

	fmt.Println("++++++++++++++++++++++")
}

func executeMovePart2(stacks map[int]*deque.Deque[string], move move) map[int]*deque.Deque[string] {
	fromColumn := stacks[move.from]
	toColumn := stacks[move.to]

	tmp := make([]string, 0)
	for cratesMoved := 0; cratesMoved < move.amount; cratesMoved++ {
		crate := fromColumn.PopBack()
		tmp = append([]string{crate}, tmp...)
	}
	for _, crate := range tmp {
		toColumn.PushBack(crate)
	}

	return stacks
}
func day5Part2(path string) string {
	lines := utils.ReadLines(path)

	stacks, moves := prepare(lines)

	//fmt.Println("")
	//printStacks(stacks)
	for _, move := range moves {
		stacks = executeMovePart2(stacks, move)
		//printStacks(stacks)
	}
	//printStacks(stacks)

	return getAnswer(stacks)
}
