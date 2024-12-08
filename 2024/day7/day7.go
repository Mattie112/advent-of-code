package main

import (
	"AoC/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", Part1("day7/day7-test.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", Part2("day7/day7-test.txt")))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", Part1("day7/day7.txt")))
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", Part2("day7/day7.txt")))
}

type jobStruct struct {
	calibrationValue int
	values           []int
}

func Part1(path string) int {
	return P1andP2(path, false)
}

func Part2(path string) int {
	return P1andP2(path, true)
}

func P1andP2(path string, part2 bool) int {
	lines := utils.ReadLines(path)
	answer := 0
	jobs := make([]jobStruct, 0)

	for _, line := range lines {
		tmp := strings.Split(line, ":")
		values := strings.Split(strings.Trim(tmp[1], " "), " ")
		job := jobStruct{
			calibrationValue: utils.MustParseStringToInt(tmp[0]),
			values:           utils.StringSliceToIntSlice(values),
		}
		log.Debugf("Job: %v", job)
		jobs = append(jobs, job)
	}

	for _, job := range jobs {
		queue := [][]int{job.values}

		for len(queue) > 0 {
			// Grab the first job from the queue (unshift)
			equation := queue[0]
			queue = queue[1:]
			if len(equation) >= 2 {
				// Grab the first 2 numbers (again, unshift)
				valA := equation[0]
				valB := equation[1]
				equation = equation[2:]
				log.Debugf("ValA: %d, ValB: %d, Job: %v", valA, valB, equation)

				// Add new jobs for both operators
				queue = append(queue, append([]int{valA + valB}, equation...))
				queue = append(queue, append([]int{valA * valB}, equation...))
				if part2 {
					queue = append(queue, append([]int{utils.MustParseStringToInt(fmt.Sprintf("%d%d", valA, valB))}, equation...))
				}
			} else {
				// If we have a single element, go and check if it's the calibration value
				if job.calibrationValue == equation[0] {
					answer += job.calibrationValue
					break
				}
				continue
			}
		}
	}

	return answer
}
