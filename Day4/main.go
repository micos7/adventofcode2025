package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: Please make sure 'input.txt' is in the current directory.")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans := 0
	// ans2 := 0
	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		grid = append(grid, letters)
	}

	dirs := [][2]int{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
		{-1, -1}, {-1, 1},
		{1, -1}, {1, 1},
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {

			if grid[r][c] != "@" {
				continue
			}

			adj := 0

			for _, d := range dirs {
				nr := r + d[0]
				nc := c + d[1]

				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[r]) {
					continue
				}

				if grid[nr][nc] == "@" {
					adj++
				}
			}

			if adj < 4 {
				ans++
			}

		}
		// fmt.Println()
	}

	fmt.Println(ans)

}
