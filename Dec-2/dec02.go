package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var safeRoutesCount int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue // Skip blank lines
		}
		fmt.Println(line)
		vals := strings.Split(line, " ")
		if checkIfSafe(vals, true) {
			safeRoutesCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(safeRoutesCount)
}

func checkIfSafe(vals []string, firstTime bool) bool {
	var increasing bool = true
	for i, v := range vals {
		fmt.Println("vals len ", len(vals))
		if i+1 == len(vals) {
			return true
		}
		curr, _ := strconv.Atoi(v)
		next, _ := strconv.Atoi(vals[i+1])

		if i == 0 && next < curr {
			increasing = false
		}
		fmt.Println(increasing)

		if increasing && next < curr {
			fmt.Println("Here when increasing fucks up")
			if firstTime {
				if checkIfSafeWithRemoved(vals) {
					return true
				}
			}
			return false
		} else if !increasing && curr < next {
			fmt.Println("Was decreasing, now increasing ", curr, next)
			if firstTime {
				if checkIfSafeWithRemoved(vals) {
					return true
				}
			}
			return false
		} else if 0 == int(math.Abs(float64(curr-next))) {
			fmt.Println("Same number twice %n, %n", curr, next)
			if firstTime {
				if checkIfSafeWithRemoved(vals) {
					return true
				}
			}
			return false
		} else if 3 < int(math.Abs(float64(curr-next))) {
			fmt.Println("Difference larger than 3 ", curr, next)
			if firstTime {
				if checkIfSafeWithRemoved(vals) {
					return true
				}
			}
			return false
		}
	}
	return true
}

func checkIfSafeWithRemoved(slice []string) bool {
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		// Create a trimmed version of the slice
		trimmed := remove(slice, i)
		fmt.Printf("Removed index %d (%s): %v\n", i, slice[i], trimmed)

		// Check the trimmed slice recursively
		if checkIfSafe(trimmed, false) {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(append([]string{}, slice[:s]...), slice[s+1:]...)
}
