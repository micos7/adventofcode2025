package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func floorDiv(a, b int) int {
	// Mathematical floor(a / b) for integers
	// Works for positive and negative a.
	if b <= 0 {
		panic("b must be positive in floorDiv")
	}
	if a >= 0 {
		return a / b
	}
	// For negative a, (a / b) in Go truncates toward zero.
	// floor(a/b) = - ((-a + b - 1) / b)
	return -(((-a) + b - 1) / b)
}

func countMultiplesInRange(lower, upper, m int) int {
	// Count integers k such that m*k is in [lower, upper].
	// If lower>upper return 0.
	if lower > upper {
		return 0
	}
	// count = floor(upper/m) - floor((lower-1)/m)
	return floorDiv(upper, m) - floorDiv(lower-1, m)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: Please make sure 'input.txt' is in the current directory.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rotations []string
	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Part 1 (unchanged)
	posP1 := 50
	part1Password := 0
	for _, line := range rotations {
		if line == "" {
			continue
		}
		dir := line[0]
		stepsVal, _ := strconv.Atoi(line[1:])
		steps := stepsVal
		if dir == 'L' {
			steps = -steps
		}
		fullDisplacement := posP1 + steps
		posP1 = fullDisplacement % 100
		if posP1 < 0 {
			posP1 += 100
		}
		if posP1 == 0 {
			part1Password++
		}
	}

	// Part 2 (fixed)
	posP2 := 50
	part2Password := 0
	for _, line := range rotations {
		if line == "" {
			continue
		}
		dir := line[0]
		stepsVal, _ := strconv.Atoi(line[1:])
		steps := stepsVal
		if dir == 'L' {
			steps = -steps
		}

		start := posP2
		end := start + steps

		var lower, upper int
		if steps > 0 {
			// Moving right: we travel (start, end] => integers start+1 ... end
			lower = start + 1
			upper = end
		} else if steps < 0 {
			// Moving left: we travel [end, start) => integers end ... start-1
			lower = end
			upper = start - 1
		} else {
			// zero steps => nothing crossed
			lower = 1
			upper = 0 // ensures lower>upper => count 0
		}

		// Count multiples of 100 in [lower, upper]
		part2Password += countMultiplesInRange(lower, upper, 100)

		// Update position modulo 100
		posP2 = end % 100
		if posP2 < 0 {
			posP2 += 100
		}
	}

	fmt.Println("--- Advent of Code Day 1 ---")
	fmt.Println("Password (Part 1 - Final Landing on 0):", part1Password)
	fmt.Println("Password (Part 2 - Total Clicks on 0):", part2Password)
}
