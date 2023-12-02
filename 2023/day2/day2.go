package main

import (
	"AoC/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	var count int
	count = day2Part1("day2/day2-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day2Part2("day2/day2-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day2Part1("day2/day2.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day2Part2("day2/day2.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

var (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type game struct {
	id    int
	blue  []int
	red   []int
	green []int
}

func day2Part1(path string) int {
	lines := utils.ReadLines(path)
	gameList := parseGames(lines)

	possibleGameCounter := 0

gameLoop:
	for _, game := range gameList {
		// For each game see if any of the color amounts is larger than our maxColors, if not this game is possible!
		for _, red := range game.red {
			if red > maxRed {
				continue gameLoop
			}
		}
		for _, green := range game.green {
			if green > maxGreen {
				continue gameLoop
			}
		}
		for _, blue := range game.blue {
			if blue > maxBlue {
				continue gameLoop
			}
		}

		possibleGameCounter += game.id
	}

	return possibleGameCounter
}

func day2Part2(path string) int {
	lines := utils.ReadLines(path)
	gameList := parseGames(lines)
	answer := 0

	for _, game := range gameList {
		minRed := math.MinInt
		minGreen := math.MinInt
		minBlue := math.MinInt

		// Loop through all colors and store the largest amount
		for _, red := range game.red {
			if red > minRed {
				minRed = red
			}
		}
		for _, green := range game.green {
			if green > minGreen {
				minGreen = green
			}
		}
		for _, blue := range game.blue {
			if blue > minBlue {
				minBlue = blue
			}
		}

		answer += minRed * minGreen * minBlue
	}

	return answer
}

func parseGames(lines []string) []game {
	re := regexp.MustCompile(`(Game (?P<gameid>\d*):)? (?P<amount>\d+) (?P<color>blue|red|green)`)
	var gameList []game

	for _, line := range lines {
		// First split the line by ';' as I can't get it to work in a single regex with repeated capture groups
		game := game{}
		for _, splitLine := range strings.Split(line, ";") {
			// With named capture groups we need a place to store those tempoary
			matches := re.FindAllStringSubmatch(splitLine, -1)

			// For each (partial) line put the data in the game object
			for _, match := range matches {
				result := utils.GetNamedRegexCaptureGroups(re, match)

				if result["gameid"] != "" {
					game.id = utils.MustParseStringToInt(result["gameid"])
				}
				switch result["color"] {
				case "red":
					game.red = append(game.red, utils.MustParseStringToInt(result["amount"]))
				case "green":
					game.green = append(game.green, utils.MustParseStringToInt(result["amount"]))
				case "blue":
					game.blue = append(game.blue, utils.MustParseStringToInt(result["amount"]))
				}
			}
		}
		gameList = append(gameList, game)
	}
	return gameList
}
