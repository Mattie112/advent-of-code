package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.DebugLevel)
	var count int
	//count = part1("day8/day8-test.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Test2 Answer %d", count))
	//count = part1("day8/day8-test2.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Test2 Answer %d", count))
	//count = part1("day8/day8.txt")
	//log.Infoln(fmt.Sprintf("Part 1 Answer %d", count))
	count = part2("day8/day8-test.txt")
	log.Infoln(fmt.Sprintf("Part 2 Test Answer %d", count))
	//count = part2("day8/day8-test2.txt")
	//log.Infoln(fmt.Sprintf("Part 2 Test2 Answer %d", count))
	//count = part2("day8/day8.txt")
	//log.Infoln(fmt.Sprintf("Part 2 Answer %d", count))
}

func part1(path string) int {
	lines := utils.ReadLines(path)

	simpleDigitCount := 0
	for _, line := range lines {
		split := strings.Split(line, " | ")
		signal := split[0]
		output := split[1]
		log.Debugf("Signal: %s", signal)
		log.Debugf("Output: %s", output)

		//signalPatterns := strings.Split(signal, " ")
		signalPatterns := strings.Split(output, " ")

		for _, pattern := range signalPatterns {
			length := len(pattern)
			switch length {
			case 2:
				log.Debugf("%s == %d", pattern, 1)
				simpleDigitCount++
			case 3:
				log.Debugf("%s == %d", pattern, 7)
				simpleDigitCount++
			case 4:
				log.Debugf("%s == %d", pattern, 4)
				simpleDigitCount++
			case 7:
				log.Debugf("%s == %d", pattern, 8)
				simpleDigitCount++
			}
		}

	}
	return simpleDigitCount
}

type diagram struct {
	Top         string
	Topleft     string
	Topright    string
	Mid         string
	Bottomleft  string
	Bottomright string
	Bottom      string
}

func part2(path string) int {
	lines := utils.ReadLines(path)

	simpleDigitCount := 0
	for _, line := range lines {
		split := strings.Split(line, " | ")
		signal := split[0]
		output := split[1]
		log.Debugf("Signal: %s", signal)
		log.Debugf("Output: %s", output)

		//diagram := map[string]string{"top": "", "topleft": "", "topright": "", "mid": "", "bottomleft": "", "bottomright": "", "bottom": ""}
		d := diagram{}
		signalPatterns := strings.Split(signal, " ")
		//signalPatterns := strings.Split(output, " ")

		for _, pattern := range signalPatterns {
			p := strings.Split(pattern, "")
			length := len(pattern)
			switch length {
			case 2:
				log.Debugf("%s == %d", pattern, 1)
				simpleDigitCount++
				d.Topright = p[0]
				d.Bottomright = p[1]
			case 3:
				log.Debugf("%s == %d", pattern, 7)
				simpleDigitCount++
				d.Top = p[0]
				d.Topright = p[1]
				d.Bottomright = p[2]
			case 4:
				log.Debugf("%s == %d", pattern, 4)
				simpleDigitCount++
				d.Topleft = p[0]
				d.Topright = p[1]
				d.Mid = p[2]
				d.Bottomright = p[3]
			case 7:
				log.Debugf("%s == %d", pattern, 8)
				simpleDigitCount++
			}
			log.Debugf("%+v", d)
			drawSegment(d)
		}

		// Now todo are 0/2/3/5/6/9

	}
	return simpleDigitCount
}

func drawSegment(d diagram) {
	//log.Debugf(" %s%s%s%s \n%s    %s\n%s    %s", d.Top, d.Top, d.Top, d.Top, d.Topleft, d.Topright, d.Topleft, d.Topright)
	log.Debugln(fmt.Sprintf(" %s%s%s%s ", d.Top, d.Top, d.Top, d.Top))
	log.Debugln(fmt.Sprintf("%s    %s", d.Topleft, d.Topright))
	log.Debugln(fmt.Sprintf("%s    %s", d.Topleft, d.Topright))
	log.Debugln(fmt.Sprintf(" %s%s%s%s ", d.Mid, d.Mid, d.Mid, d.Mid))
	log.Debugln(fmt.Sprintf("%s    %s", d.Bottomleft, d.Bottomright))
	log.Debugln(fmt.Sprintf("%s    %s", d.Bottomleft, d.Bottomright))
	log.Debugln(fmt.Sprintf(" %s%s%s%s ", d.Bottom, d.Bottom, d.Bottom, d.Bottom))
}
