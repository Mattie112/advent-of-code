package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file: %s", path))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

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

func MustParseStringToInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

type IntOrString interface {
	~int | ~string
}

func SliceToBooleanMap[T IntOrString](input []T) map[T]bool {
	reflect.ValueOf(input[0]).Kind()
	m := map[T]bool{}
	for _, i := range input {
		m[i] = true
	}
	return m
}

func StringSliceToIntSlice(input []string) []int {
	ints := make([]int, len(input))
	for i, v := range input {
		ints[i] = MustParseStringToInt(v)
	}
	return ints
}
