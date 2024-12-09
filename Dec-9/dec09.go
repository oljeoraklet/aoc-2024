package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DiskSpace struct {
	Position int
	Length   int
}

func main() {

	// partOne()
	partTwo()

}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	files := make(map[int]DiskSpace)
	freeSpaces := []DiskSpace{}

	currFileIndex := 0
	currPosition := 0

	for i, number := range numbers {
		if i%2 == 0 {
			//Save the file position and length under the file index
			files[currFileIndex] = DiskSpace{Position: currPosition, Length: number}
			currFileIndex++
		} else {
			freeSpaces = append(freeSpaces, DiskSpace{Position: currPosition, Length: number})
		}
		//Increase the current position by the number
		currPosition += number
	}

	fmt.Println("Files: ", files)
	fmt.Println("Free spaces: ", freeSpaces)

	for currFileIndex > 1 {
		currFileIndex--
		file := files[currFileIndex]

		for i, freeSpace := range freeSpaces {
			//If the free space is after the file, we don't need to consider it
			if freeSpace.Position >= file.Position {
				break
			}
			if file.Length <= freeSpace.Length {
				//If the file fits, move the position to the free space start
				files[currFileIndex] = DiskSpace{Position: freeSpace.Position, Length: file.Length}

				if file.Length == freeSpace.Length {
					//If the free space and file are the same length, the gap no longer exists
					freeSpaces = remove(freeSpaces, i)
				} else {
					freeSpaces[i] = DiskSpace{Position: freeSpace.Position + file.Length, Length: freeSpace.Length - file.Length}
				}
				break
			}
		}
	}

	fmt.Println("Files: ", files)

	checkSum := getChecksumTwo(files)

	fmt.Println("Checksum: ", checkSum)

}

func remove(slice []DiskSpace, s int) []DiskSpace {
	return append(slice[:s], slice[s+1:]...)
}

func getChecksumTwo(files map[int]DiskSpace) int {
	checkSum := 0

	for key, file := range files {

		fmt.Printf("Key: %d, File: %v\n", key, file)

		for j := file.Position; j < file.Position+file.Length; j++ {
			checkSum += key * j
		}
	}

	return checkSum
}

func partOne() {

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
	fixedData := compact(compactedData)
	fmt.Println("Checksum: ", getChecksum(fixedData))
}

func compact(numbers []int) []int {
	lastIndex := len(numbers) - 1

	for i := 0; i < lastIndex; i++ {
		if numbers[i] == -1 {
			for numbers[lastIndex] == -1 && lastIndex > i {
				lastIndex--
			}
			if lastIndex <= i {
				return numbers[:lastIndex]
			}
			numbers[i], numbers[lastIndex] = numbers[lastIndex], numbers[i]
		}
	}

	return numbers
}

func compactTwo(numbers []int) []int {
	lastIndex := len(numbers) - 1

	fmt.Println("Last element: ", numbers[lastIndex])

	for i := 0; i < lastIndex; i++ {
		fmt.Println("Index: ", i)
		if numbers[i] == -1 {
			for numbers[lastIndex] == -1 && lastIndex > i {
				lastIndex--
			}
			if lastIndex <= i {
				fmt.Println("Break")
				return numbers[:lastIndex]
			}
			numbers[i], numbers[lastIndex] = numbers[lastIndex], numbers[i]
		}
	}

	return numbers
}

func getChecksum(runes []int) int {
	checkSum := 0
	for i, num := range runes {
		checkSum += i * num
	}
	return checkSum
}
