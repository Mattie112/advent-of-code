package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var count int
	count = part1("day4/day4-test.txt")
	log.Info(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = part1("day4/day4.txt")
	log.Info(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day4/day4-test.txt")
	log.Info(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = part2("day4/day4.txt")
	log.Info(fmt.Sprintf("Part 2 Answer %d", count))
}

type bingoCard struct {
	numbers [][]int
}

func part1(path string) int {
	lines := utils.ReadLines(path)
	lines = append(lines, "") // Just add a new line, so I don't need extra checks in my loop

	numbersDrawn := utils.StrArrToIntArr(strings.Split(lines[0], ","))
	lines = utils.RemoveFromSlice(lines, 0)

	log.Debug(fmt.Sprintf("Numbers drawn: %+v", numbersDrawn))

	bingoCards := getBingoCards(lines)
	winningCardPoints := 0
	lastNumberDrawn := 0
	log.Debug(bingoCards)

	// Now we are going to draw number
BreakHere:
	for _, drawnNumber := range numbersDrawn {
		log.Debug(fmt.Sprintf("Drawing number %d", drawnNumber))

		// Loop through all boards and remove (set to -1) the winning numbers
		for cardId, bingoCard := range bingoCards {
			for rowId, row := range bingoCard.numbers {
				for columnId, numberOnCard := range row {
					if numberOnCard == drawnNumber {
						log.Debug(fmt.Sprintf("Found number %d on card %d | row %d | column %d", drawnNumber, cardId, rowId, columnId))
						bingoCard.numbers[rowId][columnId] = -1
						if checkWinningCard(bingoCard) {
							log.Debug(fmt.Sprintf("BINGO! CardID %d, cardNumbers %+v", cardId, bingoCard))
							winningCardPoints = calculateWinningCardPoints(bingoCard)
							lastNumberDrawn = drawnNumber
							break BreakHere
						}
					}
				}
			}
		}
		// Then check if we have a completed row or column

	}

	return winningCardPoints * lastNumberDrawn
}

func part2(path string) int {
	lines := utils.ReadLines(path)
	lines = append(lines, "") // Just add a new line, so I don't need extra checks in my loop

	numbersDrawn := utils.StrArrToIntArr(strings.Split(lines[0], ","))
	lines = utils.RemoveFromSlice(lines, 0)

	log.Debug(fmt.Sprintf("Numbers drawn: %+v", numbersDrawn))

	bingoCards := getBingoCards(lines)
	winningCardPoints := 0
	lastNumberDrawn := 0
	log.Debug(bingoCards)
	reDoCards := false

	// Now we are going to draw number
BreakHere:
	for _, drawnNumber := range numbersDrawn {
		log.Debug(fmt.Sprintf("Drawing number %d", drawnNumber))
		// Loop through all boards and remove (set to -1) the winning numbers
	BreakReDoCards:
		for cardId, bingoCard := range bingoCards {
		BreakAndContinue:
			for rowId, row := range bingoCard.numbers {
				for columnId, numberOnCard := range row {
					if numberOnCard == drawnNumber {
						//log.Debug(fmt.Sprintf("Found number %d on card %d | row %d | column %d", drawnNumber, cardId, rowId, columnId))
						bingoCard.numbers[rowId][columnId] = -1
						if checkWinningCard(bingoCard) {
							//log.Debug(fmt.Sprintf("BINGO! CardID %d, cardNumbers %+v", cardId, bingoCard))
							winningCardPoints = calculateWinningCardPoints(bingoCard)
							lastNumberDrawn = drawnNumber
							bingoCards = removeFromSlice(bingoCards, cardId)
							if len(bingoCards) == 0 {
								break BreakHere
							}
							reDoCards = true
							break BreakAndContinue
						}
					}
				}
			}
			if reDoCards {
				reDoCards = false
				goto BreakReDoCards
			}
		}
		// Then check if we have a completed row or column

	}

	return winningCardPoints * lastNumberDrawn
}

func getBingoCards(lines []string) []bingoCard {
	bingoCards := make([]bingoCard, 0)
	card := bingoCard{}
	cardLineCount := 0

	log.Debug(bingoCards)
	log.Debug(card)

	for _, line := range lines {
		if line == "" {
			// Store completed card
			if len(card.numbers) == 5 {
				bingoCards = append(bingoCards, card)
			}
			// Make a new card
			t := make([][]int, 5)
			for i := range t {
				t[i] = make([]int, 5)
			}
			card = bingoCard{numbers: t}
			cardLineCount = 0
			continue
		}
		// Parse and store our card
		rowNumbers := utils.StrArrToIntArr(strings.Split(line, " "))
		card.numbers[cardLineCount] = rowNumbers
		cardLineCount++
	}
	return bingoCards
}

func checkWinningCard(bingoCard bingoCard) bool {
	for _, row := range bingoCard.numbers {
		for columnId, numberOnCard := range row {
			if numberOnCard == -1 {
				// Now check for a valid row
				if row[0] == -1 && row[1] == -1 && row[2] == -1 && row[3] == -1 && row[4] == -1 {
					return true
				}
				// Or check for a valid column
				if bingoCard.numbers[0][columnId] == -1 && bingoCard.numbers[1][columnId] == -1 && bingoCard.numbers[2][columnId] == -1 && bingoCard.numbers[3][columnId] == -1 && bingoCard.numbers[4][columnId] == -1 {
					return true
				}
			}
		}
	}

	return false
}

func calculateWinningCardPoints(bingoCard bingoCard) int {
	points := 0
	for _, row := range bingoCard.numbers {
		for _, numberOnCard := range row {
			if numberOnCard != -1 {
				points += numberOnCard
			}
		}
	}
	return points
}

func removeFromSlice(slice []bingoCard, s int) []bingoCard {
	return append(slice[:s], slice[s+1:]...)
}
