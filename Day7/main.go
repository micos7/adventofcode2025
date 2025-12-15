package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	startR, startC := -1, -1
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'S' {
				startR = r
				startC = c
				break
			}
		}
		if startC != -1 {
			break
		}
	}

	active := map[int]bool{
		startC: true,
	}

	ans1 := 0
	ans2 := 0

	for r := startR + 1; r < len(grid); r++ {
		next := map[int]bool{}
		for c := range active {
			if c < 0 || c >= len(grid[r]) {
				continue
			}

			if grid[r][c] == '^' {
				ans1++
				next[c-1] = true
				next[c+1] = true

			} else {
				next[c] = true
			}
		}

		active = next
		if len(active) == 0 {
			break
		}
	}

	active2 := map[int]int{
		startC: 1,
	}

	for r := startR + 1; r < len(grid); r++ {
		next := map[int]int{}
		for c, count := range active2 {
			if c < 0 || c >= len(grid[r]) {
				continue
			}
			if grid[r][c] == '^' {
				next[c-1] += count
				next[c+1] += count
			} else {
				next[c] += count
			}
		}
		active2 = next
	}

	for _, count := range active2 {
		ans2 += count
	}

	fmt.Println(ans1, ans2)
}
