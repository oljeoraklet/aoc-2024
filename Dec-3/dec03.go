package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var totalMultiplied int = 0
	mulRegex := `mul\((\d{1,3}),(\d{1,3})\)`
	mulReg := regexp.MustCompile(mulRegex)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := mulReg.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fmt.Println(match)
			if len(match) >= 3 {
				first, _ := strconv.Atoi(match[1])
				second, _ := strconv.Atoi(match[2])
				totalMultiplied += first * second
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(totalMultiplied)
}
