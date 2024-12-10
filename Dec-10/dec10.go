package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	// Read the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var grid [][]int
	var startingPoints []Point

	// Read the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for y, char := range line {
			if len(grid) <= y {
				grid = append(grid, []int{})
			}
			if char == '0' {
				startingPoints = append(startingPoints, Point{len(grid[y]), y})
			}
			parsedChar, _ := strconv.Atoi(string(char))
			grid[y] = append(grid[y], parsedChar)
		}
	}

	fmt.Println(grid)
	fmt.Println(startingPoints)

	// Part 1
	fmt.Println(part1(grid, startingPoints))

	// Part 2
	// fmt.Println(part2(grid))
}

func part1(grid [][]int, startingPoints []Point) int {
	total := 0

	for _, point := range startingPoints {
		// Check if point can be reached
		total += checkIfReachable(point, grid)
		fmt.Println("Total", total)
	}
	return total
}

func checkIfReachable(point Point, grid [][]int) int {
	cardinalDirections := [][2]int{
		{-1, 0},        // Up
		{0, 1}, {1, 0}, // Right, Down
		{0, -1}, // Left
	}

	// fmt.Println("Checking point", point)

	queue := make([]Point, 0)
	visited := make(map[Point]bool)
	peaks := 0

	queue = append(queue, point)
	visited[point] = true

	for len(queue) > 0 {
		currPoint := queue[0]
		queue = queue[1:]
		currentValue := grid[currPoint.y][currPoint.x]

		for _, dir := range cardinalDirections {
			newPoint := Point{currPoint.x + dir[0], currPoint.y + dir[1]}
			if newPoint.x >= 0 && newPoint.x < len(grid[0]) && newPoint.y >= 0 && newPoint.y < len(grid) {
				if newPoint.x >= 0 && newPoint.x < len(grid[0]) && newPoint.y >= 0 && newPoint.y < len(grid) {
					// Skip already visited points within the current path
					if visited[newPoint] {
						continue
					}
					// Check if the path increments by 1
					if grid[newPoint.y][newPoint.x] == currentValue+1 {
						if grid[newPoint.y][newPoint.x] == 9 {
							// Found a peak; count it
							peaks++
						} else {
							// Add to queue for further exploration
							queue = append(queue, newPoint)
						}
						// Mark as visited within this path
						visited[newPoint] = true
					}
				}
			}
		}
	}

	return peaks
}

func part2(grid [][]int) int {
	return 0
}
