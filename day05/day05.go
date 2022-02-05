package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const SHOW_HORIZONTAL_VERTICAL_LINES_ONLY = false // true for part1, false for part2

const PLANE_HEIGHT = 1000
const PLANE_WIDTH = 1000

type Point = struct {
	x int
	y int
}

type Line = struct {
	start Point
	end   Point
}

func createPointFromString(textLine string) Point {
	coords := strings.Split(textLine, ",")

	startX, _ := strconv.Atoi(coords[0])
	startY, _ := strconv.Atoi(coords[1])

	return Point{startX, startY}
}

func createLineFromString(textLine string) Line {
	stringCoords := strings.Split(textLine, " -> ")

	startPoint := createPointFromString(stringCoords[0])
	endPoint := createPointFromString(stringCoords[1])

	return Line{startPoint, endPoint}
}

func loadFileToSlice(path string) []Line {
	lines := make([]Line, 0, 500)

	file, error := os.Open(path)
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		textLine := scanner.Text()
		line := createLineFromString(textLine)
		lines = append(lines, line)
	}

	file.Close()

	return lines
}

func findLineOverlaps(plane [PLANE_HEIGHT][PLANE_WIDTH]int) int {
	overlappingPointsCount := 0
	for i := 0; i < PLANE_HEIGHT; i++ {
		for j := 0; j < PLANE_WIDTH; j++ {
			if plane[i][j] >= 2 {
				overlappingPointsCount++
			}
		}
	}
	return overlappingPointsCount
}

func main() {
	lines := loadFileToSlice("./input.txt")

	var planeIntersections [PLANE_HEIGHT][PLANE_WIDTH]int

	for _, line := range lines {
		if SHOW_HORIZONTAL_VERTICAL_LINES_ONLY {
			if line.start.x != line.end.x && line.start.y != line.end.y {
				continue
			}
		}

		iterator := Point{line.start.x, line.start.y}

		for iterator.x != line.end.x || iterator.y != line.end.y {
			planeIntersections[iterator.y][iterator.x]++
			if iterator.x != line.end.x { // horizontal line
				if iterator.x < line.end.x { // go up if end is higher
					iterator.x++
				} else { // go down
					iterator.x--
				}
			}
			if iterator.y != line.end.y { // vertical line
				if iterator.y < line.end.y {
					iterator.y++
				} else {
					iterator.y--
				}
			}

		}
		planeIntersections[iterator.y][iterator.x]++ // last point
	}

	fmt.Println("Number of overlapping points:", findLineOverlaps(planeIntersections))
}
