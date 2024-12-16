package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func main() {
	log.SetLevel(log.InfoLevel)
	start := time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d, took %s", Part1("day13/day13-test.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d, took %s", Part2("day13/day13-test.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 1 Answer %d, took %s", Part1("day13/day13.txt"), time.Since(start).String()))
	start = time.Now()
	fmt.Println(fmt.Sprintf("Part 2 Answer %d, took %s", Part2("day13/day13.txt"), time.Since(start).String()))
}

type button struct {
	xOffset int
	yOffset int
}
type location struct {
	x int
	y int
}
type prizeThingyStruct struct {
	buttonA button
	buttonB button
	prize   location
}

func Part1(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	prizeThingys := make([]prizeThingyStruct, 0)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || i+2 >= len(lines) {
			continue
		}
		var re = regexp.MustCompile(`(?m)Button [AB]: X\+(\d*), Y\+(\d*)`)
		var re2 = regexp.MustCompile(`(?m)X=(\d*), Y=(\d*)`)

		match := re.FindAllStringSubmatch(lines[i], -1)
		buttonA := button{xOffset: utils.MustParseStringToInt(match[0][1]), yOffset: utils.MustParseStringToInt(match[0][2])}

		match = re.FindAllStringSubmatch(lines[i+1], -1)
		buttonB := button{xOffset: utils.MustParseStringToInt(match[0][1]), yOffset: utils.MustParseStringToInt(match[0][2])}

		match = re2.FindAllStringSubmatch(lines[i+2], -1)
		prize := location{x: utils.MustParseStringToInt(match[0][1]), y: utils.MustParseStringToInt(match[0][2])}

		log.Debugf("Button A: %v, Button B: %v, Prize: %v", buttonA, buttonB, prize)
		prizeThingys = append(prizeThingys, prizeThingyStruct{buttonA: buttonA, buttonB: buttonB, prize: prize})
		i += 2
	}

	for _, prizeThingy := range prizeThingys {

		buttonAPresses := 0
		buttonBPresses := 0
		clawLocation := location{x: 0, y: 0}

	buttonA:
		for buttonAPresses <= 100 {
			buttonAPresses++
			clawLocation.x += prizeThingy.buttonA.xOffset
			clawLocation.y += prizeThingy.buttonA.yOffset

			for buttonBPresses < 100 {
				buttonBPresses++
				clawLocation.x += prizeThingy.buttonB.xOffset
				clawLocation.y += prizeThingy.buttonB.yOffset

				if clawLocation.x == prizeThingy.prize.x && clawLocation.y == prizeThingy.prize.y {
					answer += (buttonAPresses * 3) + (buttonBPresses * 1)
					log.Debugf("Found prize at %v, %v, buttonA %d, buttonB: %d", clawLocation.x, clawLocation.y, buttonAPresses, buttonBPresses)
					break buttonA
				}
			}
			clawLocation.x -= 100 * prizeThingy.buttonB.xOffset
			clawLocation.y -= 100 * prizeThingy.buttonB.yOffset
			buttonBPresses = 0
		}
	}

	return answer
}

func Part2(path string) int {
	lines := utils.ReadLines(path)
	answer := 0

	prizeThingys := make([]prizeThingyStruct, 0)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || i+2 >= len(lines) {
			continue
		}
		var re = regexp.MustCompile(`(?m)Button [AB]: X\+(\d*), Y\+(\d*)`)
		var re2 = regexp.MustCompile(`(?m)X=(\d*), Y=(\d*)`)

		match := re.FindAllStringSubmatch(lines[i], -1)
		buttonA := button{xOffset: utils.MustParseStringToInt(match[0][1]), yOffset: utils.MustParseStringToInt(match[0][2])}

		match = re.FindAllStringSubmatch(lines[i+1], -1)
		buttonB := button{xOffset: utils.MustParseStringToInt(match[0][1]), yOffset: utils.MustParseStringToInt(match[0][2])}

		match = re2.FindAllStringSubmatch(lines[i+2], -1)
		prize := location{x: utils.MustParseStringToInt(match[0][1]) + 10000000000000, y: utils.MustParseStringToInt(match[0][2]) + 10000000000000}

		log.Debugf("Button A: %v, Button B: %v, Prize: %v", buttonA, buttonB, prize)
		prizeThingys = append(prizeThingys, prizeThingyStruct{buttonA: buttonA, buttonB: buttonB, prize: prize})
		i += 2
	}

	for _, prizeThingy := range prizeThingys {
		d := prizeThingy.buttonA.xOffset*prizeThingy.buttonB.yOffset - prizeThingy.buttonA.yOffset*prizeThingy.buttonB.xOffset
		d1 := prizeThingy.prize.x*prizeThingy.buttonB.yOffset - prizeThingy.prize.y*prizeThingy.buttonB.xOffset
		d2 := prizeThingy.prize.y*prizeThingy.buttonA.xOffset - prizeThingy.prize.x*prizeThingy.buttonA.yOffset

		if d1%d != 0 || d2%d != 0 {
			continue
		}

		answer += (3 * (d1 / d)) + (d2 / d)
	}

	return answer
}
