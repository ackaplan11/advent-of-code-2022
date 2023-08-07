package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code")
	content, _ := ioutil.ReadFile("day1/input.txt")
	lines := strings.Split(string(content), "\n")
	fmt.Println(content)
	fmt.Println(findTop3MostCalories(lines))
}

func findMostCalories(arr []string) float64 {
	var mostCalories float64 = 0
	var currentCalories float64 = 0
	for _, item := range arr {
		calories, _ := strconv.ParseFloat(item, 64)
		if calories == 0 {
			currentCalories = 0
		} else {
			currentCalories += calories
		}
		mostCalories = math.Max(currentCalories, mostCalories)
	}
	return mostCalories
}

func findTop3MostCalories(arr []string) float64 {
	var mostCalories float64 = 0
	secondMostCalories := 0.0
	thirdMostCalories := 0.0
	var currentCalories float64 = 0
	for _, item := range arr {
		calories, _ := strconv.ParseFloat(item, 64)
		if calories == 0 {
			currentCalories = 0
		} else {
			currentCalories += calories
		}
		if currentCalories > mostCalories {
			thirdMostCalories = secondMostCalories
			secondMostCalories = mostCalories
			mostCalories = currentCalories
		} else if currentCalories > secondMostCalories {
			thirdMostCalories = secondMostCalories
			secondMostCalories = currentCalories
		} else if currentCalories > thirdMostCalories {
			thirdMostCalories = currentCalories
		}

	}
	return mostCalories + secondMostCalories + thirdMostCalories
}
