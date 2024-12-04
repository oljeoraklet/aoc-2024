package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Path to your text file
	filePath := "input.txt"

	grid, err := ReadFileTo2DArray(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Grid:")
	for _, row := range grid {
		fmt.Println(string(row))
	}

	part := 1
	// Find and print X -> M -> A patterns
	FindConnections(grid, part)

}

func ReadFileTo2DArray(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func FindConnections(grid [][]rune, part int) {

	xmasCount := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top-left, Top, Top-right
		{0, -1}, {0, 1}, // Left,       Right
		{1, -1}, {1, 0}, {1, 1}, // Bottom-left, Bottom, Bottom-right
	}

	for row, rows := range grid {
		for col, char := range rows {
			if part == 2 {
				if char == 'A' {
					if CheckCorners(grid, row, col) {
						xmasCount++
						fmt.Printf("Current xmas count is: %d\n", xmasCount)
					}
				}
			} else {
				if char == 'X' {
					for _, dir := range directions {
						dirRow, dirCol := dir[0], dir[1]
						if CheckForConnection(grid, row, col, dirRow, dirCol) {
							xmasCount++
							fmt.Println(xmasCount)
						}
					}
				}
			}

		}
	}

}

func CheckCorners(grid [][]rune, row, col int) bool {
	// Define the corner pairs: top-right to bottom-left, top-left to bottom-right
	topRight := [2]int{-1, 1}
	bottomLeft := [2]int{1, -1}
	topLeft := [2]int{-1, -1}
	bottomRight := [2]int{1, 1}

	topRightRow, topRightCol := row+topRight[0], col+topRight[1]
	bottomLeftRow, bottomLeftCol := row+bottomLeft[0], col+bottomLeft[1]

	if topRightRow >= 0 && topRightRow < len(grid) && topRightCol >= 0 && topRightCol < len(grid[topRightRow]) &&
		bottomLeftRow >= 0 && bottomLeftRow < len(grid) && bottomLeftCol >= 0 && bottomLeftCol < len(grid[bottomLeftRow]) {
		topRightChar := grid[topRightRow][topRightCol]
		bottomLeftChar := grid[bottomLeftRow][bottomLeftCol]

		// Check if the first pair matches the condition
		if !((topRightChar == 'M' && bottomLeftChar == 'S') || (topRightChar == 'S' && bottomLeftChar == 'M')) {
			return false
		}
	} else {
		// Out of bounds for top-right or bottom-left
		return false
	}

	// Check top-left to bottom-right
	topLeftRow, topLeftCol := row+topLeft[0], col+topLeft[1]
	bottomRightRow, bottomRightCol := row+bottomRight[0], col+bottomRight[1]

	if topLeftRow >= 0 && topLeftRow < len(grid) && topLeftCol >= 0 && topLeftCol < len(grid[topLeftRow]) &&
		bottomRightRow >= 0 && bottomRightRow < len(grid) && bottomRightCol >= 0 && bottomRightCol < len(grid[bottomRightRow]) {
		topLeftChar := grid[topLeftRow][topLeftCol]
		bottomRightChar := grid[bottomRightRow][bottomRightCol]

		// Check if the second pair matches the condition
		if !((topLeftChar == 'M' && bottomRightChar == 'S') || (topLeftChar == 'S' && bottomRightChar == 'M')) {
			return false
		}
	} else {
		// Out of bounds for top-left or bottom-right
		return false
	}

	return true
}

func CheckForConnection(grid [][]rune, row, col int, dirRow, dirCol int) bool {
	// First step: Check for M in the direction
	newRow, newCol := row+dirRow, col+dirCol
	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[newRow]) || grid[newRow][newCol] != 'M' {
		return false
	}

	// Second step: Check for A in the same direction
	newRow += dirRow
	newCol += dirCol
	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[newRow]) || grid[newRow][newCol] != 'A' {
		return false
	}

	newRow += dirRow
	newCol += dirCol
	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[newRow]) || grid[newRow][newCol] != 'S' {
		return false
	}

	// If both conditions are met, return true
	return true
}
