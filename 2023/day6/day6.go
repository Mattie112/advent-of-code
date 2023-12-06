package main

import (
	"fmt"
)

func main() {
	var count int
	count = day6Part1("day6/day6-test.txt")
	fmt.Println(fmt.Sprintf("Part 1 Test Answer %d", count))
	count = day6Part2("day6/day6-test.txt")
	fmt.Println(fmt.Sprintf("Part 2 Test Answer %d", count))
	count = day6Part1("day6/day6.txt")
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", count))
	count = day6Part2("day6/day6.txt")
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", count))
}

func day6Part1(path string) int {
	//lines := utils.ReadLines(path)
	answer := 1

	//races := make([][]int, 100)
	//tmpTimes := strings.Split(lines[0], " ")
	//tmpDistances := strings.Split(lines[1], " ")
	//j := 0
	//for i := 0; i < len(tmpTimes); i++ {
	//	tmpTime := tmpTimes[i]
	//	time, err := strconv.Atoi(tmpTime)
	//	if err != nil {
	//		races[j] = []int{time, 0}
	//		j++
	//	}
	//}
	//
	//for i, tmpDistance := range tmpDistances {
	//	distance, err := strconv.Atoi(tmpDistance)
	//	if err != nil {
	//		races[i] = []int{races[i][0], distance}
	//	}
	//}

	//distances := lines[1]

	// Time & distance
	//input := [][]int{{7, 9}, {15, 40}, {30, 200}} // test
	//input := [][]int{{54, 446}, {81, 1292}, {70, 1035}, {88, 1007}} // correct! - 2065338
	//input := [][]int{{71530, 940200}} // test part 2 - ok works
	input := [][]int{{54817088, 446129210351007}} // test part 2 - ok works - 34934171

	for _, race := range input {

		time := race[0]
		recordDistance := race[1]

		speed := 0
		distances := make([]int, time+1)
		for i := 0; i <= time; i++ {

			distances[i] = speed * (time - i)

			speed++
		}
		winningDistance := 0
		for _, d := range distances {
			if d > recordDistance {
				winningDistance++
			}
		}
		answer *= winningDistance
	}

	return answer
}

func day6Part2(path string) int {
	//lines := utils.ReadLines(path)
	answer := 0
	//for _, line := range lines {
	//
	//}

	return answer
}
