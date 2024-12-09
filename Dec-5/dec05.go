package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var content string = ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	middleTotal := 0

	sections := strings.Split(strings.TrimSpace(content), "\n\n")

	rules := make(map[int][]int)
	for _, line := range strings.Split(sections[0], "\n") {
		var a, b int
		fmt.Sscanf(line, "%d|%d", &a, &b)
		rules[a] = append(rules[a], b)
	}

	var sequences [][]int
	for _, line := range strings.Split(sections[1], "\n") {
		var sequence []int
		for _, num := range strings.Split(line, ",") {
			var n int
			fmt.Sscanf(num, "%d", &n)
			sequence = append(sequence, n)
		}
		sequences = append(sequences, sequence)
	}

	for _, sequence := range sequences {
		if validSequence(sequence, rules) {
			// middleIndex := len(sequence) / 2
			// middleTotal += sequence[middleIndex]
			// fmt.Printf("Middle total is %d\n", middleTotal)
		} else {
			// fmt.Println("Sequence: ", sequence)
			correctedSequence := reorderSequence(sequence, rules)
			// fmt.Println("Corrected sequence: ", correctedSequence)
			middleIndex := len(correctedSequence) / 2
			middleTotal += correctedSequence[middleIndex]
			fmt.Printf("Middle total is %d\n", middleTotal)

		}
	}

}

func validSequence(sequence []int, rules map[int][]int) bool {
	// Create a map for the position of each number in the sequence
	position := make(map[int]int)
	for i, num := range sequence {
		position[num] = i
	}

	// Check the rules
	for a, dependents := range rules {
		for _, b := range dependents {
			// If both a and b are in the sequence, a must come before b
			if posA, existsA := position[a]; existsA {
				if posB, existsB := position[b]; existsB {
					if posA > posB {
						return false
					}
				}
			}
		}
	}

	return true
}

func reorderSequence(sequence []int, rules map[int][]int) []int {
	// Create a reordered slice
	reordered := []int{}
	seen := make(map[int]bool)

	// Helper function to add a number and its dependencies
	var addNumber func(int)
	addNumber = func(num int) {
		if seen[num] {
			return
		}
		for _, dep := range rules[num] {
			if contains(sequence, dep) {
				addNumber(dep)
			}
		}
		reordered = append(reordered, num)
		seen[num] = true
	}

	// Process each number in the sequence
	for _, num := range sequence {
		addNumber(num)
	}

	return reordered
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}
