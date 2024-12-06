package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var grid [][]rune

	var startPosition Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, v := range line {
			if v == '^' {
				fmt.Printf("Found the start at line %d, index %d\n", len(grid), i)
				startPosition.X = len(grid)
				startPosition.Y = i

			}
		}
		grid = append(grid, line)
	}
	fmt.Println(grid)
	fmt.Println(startPosition)
	fmt.Println(grid[startPosition.X][startPosition.Y])
	fmt.Println(checkDirections(grid, startPosition))
}

func checkDirections(grid [][]rune, start Point) int {
	cardinalDirections := [][2]int{
		{-1, 0},        //  Up,
		{0, 1}, {1, 0}, // Right,       Down
		{0, -1}, //  Left

	}

	currentDir := 0
	dirRow := cardinalDirections[currentDir][0]
	dirCol := cardinalDirections[currentDir][1]

	newRow, newCol := start.X+dirRow, start.Y+dirCol

	prevRow, prevCol := newRow, newCol
	visitedSquares := make(map[Point]bool)

	fmt.Printf("New row: %d, new col: %d\n", newRow, newCol)

	fmt.Printf("Hello from checkDirection\n")

	fmt.Println("Grid row len: ", len(grid[newRow]))

	for {
		if newRow < len(grid) && newRow >= 0 && newCol < len(grid[newRow]) && newCol >= 0 {
			if (visitedSquares[Point{X: newRow, Y: newCol}]) {
				prevRow, prevCol = newRow, newCol
				newRow += dirRow
				newCol += dirCol

				continue
			}
			if grid[newRow][newCol] == '#' {
				if currentDir+1 >= len(cardinalDirections) {
					currentDir = 0

				} else {
					currentDir++
				}

				newRow, newCol = prevRow, prevCol

				fmt.Println("New direction: ", cardinalDirections[currentDir])
				dirRow = cardinalDirections[currentDir][0]
				dirCol = cardinalDirections[currentDir][1]
				newRow += dirRow
				newCol += dirCol

			} else {
				visitedSquares[Point{X: newRow, Y: newCol}] = true
				prevRow, prevCol = newRow, newCol
				newRow += dirRow
				newCol += dirCol

			}
		} else {
			break
		}

	}
	return len(visitedSquares)
}
