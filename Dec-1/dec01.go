package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var listOne []int
	var listTwo []int

	//Split the list into the two arrays
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		var entry string = scanner.Text()
		val, _ := strconv.Atoi(entry)

		if count == 0 || count%2 == 0 {
			listOne = append(listOne, val)
		} else {
			listTwo = append(listTwo, val)
		}
		count++
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	//Part One
	// var totalDist int = 0

	//Part Two
	var simScore int = 0

	for _, v := range listOne {
		//Part One
		// dist := int(math.Abs(float64(v - listTwo[i])))
		// totalDist += dist

		//Part Two
		count := 0

		for _, number := range listTwo {
			if v == number {
				count++
			}
		}

		simScore += (count * v)
	}

	fmt.Println(simScore)
	// fmt.Println(totalDist)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
