package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var parent []int
var rank []int

var part1 = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := []Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		points = append(points, Point{x, y})
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := (abs(points[i].x-points[j].x) + 1) * (abs(points[i].y-points[j].y) + 1)
			part1 = max(part1, dist)
		}
	}

	fmt.Println("Part 1:", part1)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
