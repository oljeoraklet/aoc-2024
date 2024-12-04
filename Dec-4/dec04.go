package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Path to your text file
	filePath := "input2.txt"

	grid, err := ReadFileTo2DArray(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Grid:")
	for _, row := range grid {
		fmt.Println(string(row))
	}

	part := 2
	// Find and print X -> M -> A patterns
	FindConnections(grid)

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

func FindConnections(grid [][]rune) {

	xmasCount := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top-left, Top, Top-right
		{0, -1}, {0, 1}, // Left,       Right
		{1, -1}, {1, 0}, {1, 1}, // Bottom-left, Bottom, Bottom-right
	}
	corners := [][2]int{
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for row, rows := range grid {
		for col, char := range rows {
			if char == 'A' {
				for _, dir := range directions {
					fmt.Println(dir)
					dirRow, dirCol := dir[0], dir[1]
					if CheckForConnection(grid, row, col, dirRow, dirCol) {
						fmt.Printf("Pattern found at [%d][%d] in direction (%d, %d)\n", row, col, dirRow, dirCol)
						xmasCount++
						fmt.Println(xmasCount)
					}
				}
			}
		}
	}

}

func CheckCorners(grid [][]rune, row, col int, dirRow, dirCol int) bool {
	newRow, newCol := row+dirRow, col+dirCol
	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[newRow]) || grid[newRow][newCol] != 'M' {
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
