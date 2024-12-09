package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	numbers := []int{}

	currDataIndex := 0
	compactedData := []int{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for i, num := range numbers {
		//If even number then it is a file, not a space
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				compactedData = append(compactedData, currDataIndex)
			}
			currDataIndex++
		} else {

			for j := 0; j < num; j++ {
				compactedData = append(compactedData, -1)
			}

		}
	}

	fmt.Println("Disk data: ", compactedData)

	fmt.Println("Compacted string: ", compact(compactedData))
	fmt.Println("Checksum: ", getChecksum(compact(compactedData)))

}

func compact(runes []int) []int {
	lastIndex := len(runes) - 1

	fmt.Println("Last element: ", runes[lastIndex])

	for i := 0; i < lastIndex; i++ {
		fmt.Println("Index: ", i)
		if runes[i] == -1 {
			for runes[lastIndex] == -1 && lastIndex > i {
				lastIndex--
			}
			if lastIndex <= i {
				fmt.Println("Break")
				return runes[:lastIndex]
			}
			runes[i], runes[lastIndex] = runes[lastIndex], runes[i]
		}
	}

	return runes
}

func getChecksum(runes []int) int {
	checkSum := 0
	for i, num := range runes {
		checkSum += i * num
	}
	return checkSum
}
