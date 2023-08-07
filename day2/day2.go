package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Advent of Code")
	content, _ := ioutil.ReadFile("day2/input.txt")
	bouts := strings.Split(string(content), "\n")
	fmt.Println(playStrategy1(bouts))
	fmt.Println(playStrategy2(bouts))
}

func playStrategy1(bouts []string) int {
	pointsPerBout := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	score := 0
	for _, bout := range bouts {
		score += pointsPerBout[bout]
	}
	return score
}

func playStrategy2(bouts []string) int {
	pointsPerBout := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}
	score := 0
	for _, bout := range bouts {
		score += pointsPerBout[bout]
	}
	return score
}

func transformThrow(throw string) string {
	verbose := ""
	if throw == "A" || throw == "X" {
		verbose = "ROCK"
	} else if throw == "B" || throw == "Y" {
		verbose = "PAPER"
	} else if throw == "C" || throw == "Z" {
		verbose = "SCISSORS"
	}
	return verbose
}

func transformOutcome(outcome string) string {
	verbose := ""
	if outcome == "X" {
		verbose = "LOSE"
	} else if outcome == "Y" {
		verbose = "DRAW"
	} else if outcome == "Z" {
		verbose = "WIN"
	}
	return verbose
}

func determineBoutScore(playerThrow string, outcome string) int {
	pointsPerThrow := map[string]int{
		"ROCK":     1,
		"PAPER":    2,
		"SCISSORS": 3,
	}
	pointsPerOutcome := map[string]int{
		"LOSE": 0,
		"DRAW": 3,
		"WIN":  6,
	}
	return pointsPerThrow[playerThrow] + pointsPerOutcome[outcome]
}
