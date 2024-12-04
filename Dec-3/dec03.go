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
	// mulRegex := `mul\((\d{1,3}),(\d{1,3})\)`
	// mulReg := regexp.MustCompile(mulRegex)

	combiRegex := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	combiReg := regexp.MustCompile(combiRegex)

	shouldMultiply := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := combiReg.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fmt.Println(match)
			fmt.Println(shouldMultiply)
			if len(match) > 1 && match[1] != "" && match[2] != "" && shouldMultiply {
				first, _ := strconv.Atoi(match[1])
				second, _ := strconv.Atoi(match[2])
				totalMultiplied += first * second
			} else {
				fmt.Printf("Matched pattern: %s\n", match[0])
				if match[0] == "do()" {
					shouldMultiply = true
				} else {
					shouldMultiply = false
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(totalMultiplied)
}
