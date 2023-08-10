package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Advent of Code")
	datastream, _ := ioutil.ReadFile("day6/input.txt")
	fmt.Println(findStartOfPacket(datastream))
	fmt.Println(findStartOfMessage(datastream))
}

func findStartOfPacket(datastream []byte) int {
	charsProcessed := 3
	buffer := datastream[:3]
	for _, char := range datastream[3:] {
		charsProcessed += 1
		if len(buffer) == 3 {
			buffer = append(buffer, char)
		} else {
			buffer = append(buffer[1:], char)
		}
		if !hasDuplicates(buffer) {
			return charsProcessed
		}
	}
	return charsProcessed
}

func findStartOfMessage(datastream []byte) int {
	charsProcessed := 13
	buffer := datastream[:13]
	for _, char := range datastream[13:] {
		charsProcessed += 1
		if len(buffer) == 13 {
			buffer = append(buffer, char)
		} else {
			buffer = append(buffer[1:], char)
		}
		if !hasDuplicates(buffer) {
			return charsProcessed
		}
	}
	return charsProcessed
}
func hasDuplicates(slice []byte) bool {
	seen := make(map[byte]bool)

	for _, b := range slice {
		if seen[b] {
			return true // Duplicate found
		}
		seen[b] = true
	}

	return false // No duplicates found
}
