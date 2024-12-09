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
				startPosition = Point{X: len(grid), Y: i}
			}
		}
		grid = append(grid, line)
	}

	loopPositions := loopCheck(grid, startPosition)
	if len(loopPositions) > 0 {
		fmt.Println("Positions where placing a '#' will cause a loop:")
		for _, pos := range loopPositions {
			fmt.Printf("%v\n", pos)
		}
		fmt.Printf("Total positions: %d\n", len(loopPositions))
	} else {
		fmt.Println("No single '#' placement creates a loop.")
	}
}

func checkDirections(grid [][]rune, start Point) (int, bool) {
	cardinalDirections := [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	currentDir := 0
	visited := make(map[Point]map[int]bool)
	current := start
	visited[current] = make(map[int]bool)

	for {
		// Mark the current position and direction as visited
		visited[current][currentDir] = true

		// Calculate next position
		next := Point{X: current.X + cardinalDirections[currentDir][0], Y: current.Y + cardinalDirections[currentDir][1]}

		// Check bounds
		if next.X < 0 || next.X >= len(grid) || next.Y < 0 || next.Y >= len(grid[0]) {
			break // Out of bounds
		}

		if grid[next.X][next.Y] == '#' {
			// Turn right if hitting an obstacle
			currentDir = (currentDir + 1) % 4
		} else {
			current = next
		}

		// Check for a loop
		if visited[current] != nil && visited[current][currentDir] {
			return len(visited), true
		}

		// Initialize visited map for new positions
		if visited[current] == nil {
			visited[current] = make(map[int]bool)
		}
	}

	return len(visited), false
}

func loopCheck(grid [][]rune, start Point) []Point {
	loopPositions := []Point{}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '.' && !(row == start.X && col == start.Y) {
				// Temporarily place a '#'
				grid[row][col] = '#'

				_, isLoop := checkDirections(grid, start)
				if isLoop {
					loopPositions = append(loopPositions, Point{X: row, Y: col})
				}

				// Restore the grid
				grid[row][col] = '.'
			}
		}
	}

	return loopPositions
}
