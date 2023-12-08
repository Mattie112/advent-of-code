package main

import (
	"AoC/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var count int
	count = day7Part1("day7/day7-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day7Part2("day7/day7-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day7Part1("day7/day7.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day7Part2("day7/day7.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

type handTypeType int

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPairs     = 3
	OnePair      = 2
	HighCard     = 1
)

func (h handTypeType) String() string {
	return [...]string{"HighCard", "OnePair", "TwoPairs", "ThreeOfAKind", "FullHouse", "FourOfAKind", "FiveOfAKind"}[h-1]
}

type handTypeStruct struct {
	hand     string
	bid      int
	handType handTypeType
}

// This method will convert a 'hand' to a score (using hexadecimals)
func handToScore(hand handTypeStruct, partTwo bool) int64 {
	score := "0x" + strconv.Itoa(int(hand.handType))
	for _, symbol := range hand.hand {
		switch symbol {
		case 'A':
			score += "E"
		case 'K':
			score += "D"
		case 'Q':
			score += "C"
		case 'J':
			if !partTwo {
				score += "B"
			} else {
				score += "0"
			}
		case 'T':
			score += "A"
		default:
			score += string(symbol)
		}
	}
	n, err := strconv.ParseInt(score, 0, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func getType(input string, partTwo bool) handTypeType {

	// Calculate the card and the amount of each card
	cards := utils.CountChars(input)

	// Golang maps (and slices only have an int as index) are not ordered, so we need to sort them by making a slice for the keys and sorting that
	keys := make([]string, 0, len(cards))
	for key := range cards {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return cards[keys[i]] > cards[keys[j]] })

	// Our first key is now the letter of the card with the highest amount
	keyToCheck := 0

	// For part 2 we need to handle the joker a bit differently
	if partTwo {
		// Get the amount of Joker cards
		jAmount := cards["J"]

		// The first card that is not a Joker will get the amount of Joker cards added to it (does not really matter which one)
		for i, key := range keys {
			if key != "J" {
				cards[keys[i]] += jAmount
				break
			}
		}

		// If the Joker has the largest amount of cards we will check the second card instead (except when all the cards are Jokers)
		if keys[0] == "J" && input != "JJJJJ" && len(keys) > 1 {
			keyToCheck = 1
		}
	}

	// Now see how we score and return that enum
	if cards[keys[keyToCheck]] == 5 {
		return FiveOfAKind
	}

	if cards[keys[keyToCheck]] == 4 {
		return FourOfAKind
	}

	if cards[keys[keyToCheck]] == 3 && cards[keys[keyToCheck+1]] == 2 {
		return FullHouse
	}

	if cards[keys[keyToCheck]] == 3 {
		return ThreeOfAKind
	}

	if cards[keys[keyToCheck]] == 2 && cards[keys[keyToCheck+1]] == 2 {
		return TwoPairs
	}

	if cards[keys[keyToCheck]] == 2 {
		return OnePair
	}

	return HighCard
}

func run(path string, partTwo bool) int {
	lines := utils.ReadLines(path)
	answer := 0

	hands := make(map[handTypeType][]handTypeStruct)

	for _, line := range lines {
		tmp := strings.Split(line, " ")
		hand := tmp[0]
		bid := utils.MustParseStringToInt(tmp[1])
		handType := getType(hand, partTwo)

		hands[handType] = append(hands[handType], handTypeStruct{hand, bid, handType})
	}

	scoreSlice := make(map[int64]handTypeStruct)

	for _, handsOfASpecificType := range hands {
		for _, hand := range handsOfASpecificType {
			// Convert the hand to a score (using hexadecimals)
			scoreSlice[handToScore(hand, partTwo)] = hand
		}
	}

	keys := make([]int64, 0, len(scoreSlice))

	// Populate the slice with keys from the map
	for key := range scoreSlice {
		keys = append(keys, key)
	}

	// Sort the keys in descending order
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Iterate through the sorted keys
	for i, key := range keys {
		tmp := scoreSlice[key]
		answer += tmp.bid * (i + 1)
		//fmt.Printf("Key: %v, Hand: %v, Bid: %v, HandType: %s (%d), Rank:%d \n", key, scoreSlice[key].hand, scoreSlice[key].bid, scoreSlice[key].handType, scoreSlice[key].handType, i+1)
	}

	return answer
}

func day7Part1(path string) int {
	answer := run(path, false)
	return answer
}

func day7Part2(path string) int {
	answer := run(path, true)
	return answer
}
