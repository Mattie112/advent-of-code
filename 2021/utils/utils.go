package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func RemoveFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func StrArrToIntArr(strings []string) []int {
	output := make([]int, 0)
	for _, s := range strings {
		if s == " " || s == "" {
			continue
		}
		number, _ := strconv.Atoi(s)
		output = append(output, number)
	}
	return output
}

func StringToInt(string string) int {
	i, err := strconv.Atoi(string)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse %s as integer", string))
	}
	return i
}

func GetMaxFromArr(arr []int) int {
	max := 0
	for _, a := range arr {
		if a > max {
			max = a
		}
	}
	return max
}

func Contains(needle string, haystack []string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

// IsLower https://stackoverflow.com/a/59293875/2451037
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func CopyMap(o map[string]bool) map[string]bool {
	c := map[string]bool{}
	for k, v := range o {
		c[k] = v
	}
	return c
}
