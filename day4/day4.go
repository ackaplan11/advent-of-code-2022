package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code")
	content, _ := ioutil.ReadFile("day4/input.txt")
	lines := strings.Split(string(content), "\n")
	boundsOverlapCount := 0
	anyOverlapCount := 0
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println(boundsOverlapCount)
			fmt.Println(anyOverlapCount)
			return
		}
		pairs := strings.Split(line, ",")
		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")
		low1, errLow1 := parseIntFromString(pair1[0])
		low2, errLow2 := parseIntFromString(pair2[0])
		high1, errHigh1 := parseIntFromString(pair1[1])
		high2, errHigh2 := parseIntFromString(pair2[1])
		if errLow1 != nil || errLow2 != nil || errHigh1 != nil || errHigh2 != nil {
			fmt.Println("Error Parsing Strings")
		} else {
			if boundsOverlap(low1, low2, high1, high2) {
				boundsOverlapCount += 1
			}
			if anyOverlap(low1, low2, high1, high2) {
				anyOverlapCount += 1
			}
		}
	}
}

func parseIntFromString(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to parse string to integer: %w", err)
	}
	return num, nil
}

func boundsOverlap(low1 int, low2 int, high1 int, high2 int) bool {
	return (low1 <= low2 && high1 >= high2) || (low1 >= low2 && high1 <= high2)
}

func anyOverlap(low1 int, low2 int, high1 int, high2 int) bool {
	// 5, 7, 7, 9
	// 2, 3, 8, 7
	// 6, 6, 4, 6
	// 2, 4, 6, 8
	return (low1 >= low2 && low1 <= high2) || (high1 >= low2 && high1 <= high2) ||
		(low2 >= low1 && low2 <= high1) || (high2 >= low1 && high2 <= high1)
}
