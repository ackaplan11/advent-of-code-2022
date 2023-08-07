package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Advent of Code")
	content, _ := ioutil.ReadFile("day3/input.txt")
	rucksacks := strings.Split(string(content), "\n")
	prioritySum := int32(0)
	for _, rucksack := range rucksacks {
		prioritySum += ruckSackPriority(rucksack)

	}
	fmt.Println(prioritySum)
	prioritySum = 0
	for i := 2; i < len(rucksacks); i += 3 {
		rucksack1 := rucksacks[i-2]
		rucksack2 := rucksacks[i-1]
		rucksack3 := rucksacks[i]
		prioritySum += ruckSackPriority2(rucksack1, rucksack2, rucksack3)
	}
	fmt.Println(prioritySum)
}

func determinePriority(r rune) int32 {
	// Perform some operations with the rune
	intValue := int32(r)
	priority := int32(0)
	if intValue >= 65 && intValue <= 90 { // Uppercase letters (A-Z)
		priority = intValue - 38
	} else if intValue >= 97 && intValue <= 122 { // Lowercase letters (a-z)
		priority = intValue - 96
	} else {
		priority = 0 // No transformation for non-alphabet characters
	}
	return priority
}

func ruckSackPriority(input string) int32 {
	str1 := input[:len(input)/2]
	str2 := input[len(input)/2:]
	priority := int32(0)
	for _, char1 := range str1 {
		if strings.ContainsRune(str2, char1) {
			priority += determinePriority(char1)
			break
		}
	}
	return priority
}

func ruckSackPriority2(group1 string, group2 string, group3 string) int32 {
	priority := int32(0)
	for _, char1 := range group1 {
		if strings.ContainsRune(group2, char1) && strings.ContainsRune(group3, char1) {
			priority += determinePriority(char1)
			break
		}
	}
	return priority
}
