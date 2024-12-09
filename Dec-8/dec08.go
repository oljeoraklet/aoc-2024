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
	// Path to your text file
	filePath := "test.txt"

	grid, err := ReadFileTo2DArray(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	antennas := make(map[rune][]Point)

	//Part 1
	for r, row := range grid {
		for c, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], Point{r, c})
			}
		}
	}

	fmt.Printf("Antennas: %v\n", antennas)

	antiNodes := getAntiNodes(antennas)
	// filteredAntiNodes := make(map[Point][]Point)
	filteredAntinodes := make(map[Point]bool)

	for point := range antiNodes {
		fmt.Printf("Point: %v\n", point)
		if point.X >= 0 && point.X < len(grid) && point.Y >= 0 && point.Y < len(grid[0]) {
			fmt.Printf("Point: %v\n", point)
			filteredAntinodes[point] = true
		}

	}

	fmt.Printf("Antinodes: %v\n", antiNodes)
	fmt.Printf("Number of antinodes: %d\n", len(filteredAntinodes))

}

func getAntiNodes(antennas map[rune][]Point) map[Point]bool {
	//map[Point][]Point
	//In a 2 dimensional plane, the distance between points (x1, y1) and (x2, y2)

	//Store the antinodes as set to avoid duplicates
	antiNodes := make(map[Point]bool)

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				var point1 Point = antenna[i]
				var point2 Point = antenna[j]

				fmt.Printf("Point 1: %v, Point 2: %v\n", point1, point2)

				var distanceUp Point = Point{point1.X - point2.X, point1.Y - point2.Y}

				var distanceDown Point = Point{point2.X - point1.X, point2.Y - point1.Y}

				//Add the distance to both points in either direction
				var antiPoint1 Point = Point{point1.X + distanceUp.X, point1.Y + distanceUp.Y}

				fmt.Printf("AntiPoint 1: %v\n", antiPoint1)

				var antiPoint2 Point = Point{point2.X + distanceDown.X, point2.Y + distanceDown.Y}

				fmt.Printf("AntiPoint 2: %v\n", antiPoint2)

				antiNodes[antiPoint1] = true
				antiNodes[antiPoint2] = true

			}
		}
	}
	return antiNodes

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
