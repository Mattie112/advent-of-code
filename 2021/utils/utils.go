package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file: %s", path))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(fmt.Sprintf("Error while reading file: %s", err))
	}

	return lines
}

func ReadLinesAsInteger(path string) []int {
	lines := ReadLines(path)
	var linesAsInt []int
	for _, l := range lines {
		number, _ := strconv.Atoi(l)
		linesAsInt = append(linesAsInt, number)
	}

	return linesAsInt
}
