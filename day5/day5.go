package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code")
	content, _ := ioutil.ReadFile("day5/input.txt")
	lines := strings.Split(string(content), "\n")
	crates := map[int][]string{
		1: {"L", "N", "W", "T", "D"},
		2: {"C", "P", "H"},
		3: {"W", "P", "H", "N", "D", "G", "M", "J"},
		4: {"C", "W", "S", "N", "T", "Q", "L"},
		5: {"P", "H", "C", "N"},
		6: {"T", "H", "N", "D", "M", "W", "Q", "B"},
		7: {"M", "B", "R", "J", "G", "S", "L"},
		8: {"Z", "N", "W", "G", "V", "B", "R", "T"},
		9: {"W", "G", "D", "N", "P", "L"},
	}
	for _, command := range lines {
		if len(command) > 0 && command[0] == 'm' {
			count, starIdx, endIdx, err := extractCommand(command)
			if err != nil {
				fmt.Println("Error:", err)
				return
			} else {
				moveCrates2(crates, count, starIdx, endIdx)
			}
		}
	}
	fmt.Println(crates)
}

func extractCommand(input string) (int, int, int, error) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	match := re.FindStringSubmatch(input)
	if len(match) != 4 {
		return 0, 0, 0, fmt.Errorf("invalid input format")
	}

	count, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, 0, 0, err
	}

	startIdx, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, 0, 0, err
	}

	endIdx, err := strconv.Atoi(match[3])
	if err != nil {
		return 0, 0, 0, err
	}

	return count, startIdx, endIdx, err
}

func moveCrates(crates map[int][]string, count int, startIdx int, endIdx int) {
	movesLeft := count
	for movesLeft > 0 {
		n := len(crates[startIdx]) - 1
		elem := crates[startIdx][n]                   // retrieve element at top of stack
		crates[endIdx] = append(crates[endIdx], elem) // append element to top of destination stack
		crates[startIdx] = crates[startIdx][:n]       //remove element from top of origin stack
		movesLeft -= 1
	}
}

func moveCrates2(crates map[int][]string, count int, startIdx int, endIdx int) {
	n := len(crates[startIdx]) - count
	slice := crates[startIdx][n:]
	for _, elem := range slice {
		crates[endIdx] = append(crates[endIdx], elem)
	}
	crates[startIdx] = crates[startIdx][:n]
}
