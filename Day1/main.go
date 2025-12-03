package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

		move := 1
		if dir == 'L' {
			move = -1
		}

		for i := 0; i < stepsVal; i++ {
			posP2 += move

			if posP2%100 == 0 {
				part2Password++
			}
		}

		posP2 = posP2 % 100
		if posP2 < 0 {
			posP2 += 100
		}
	}

	fmt.Println("Password (Part 1 - Final Landing on 0):", part1Password)
	fmt.Println("Password (Part 2 - Total Clicks on 0):", part2Password)
}
