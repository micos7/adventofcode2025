package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var dataLines []string
	var opLine string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.ContainsAny(line, "+*") {
			opLine = line
		} else {
			dataLines = append(dataLines, line)
		}
	}

	part1Numbers := make([][]int, len(dataLines))
	for i, line := range dataLines {
		fields := strings.Fields(line)
		nums := []int{}
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			nums = append(nums, n)
		}
		part1Numbers[i] = nums
	}

	part1Total := 0
	if len(part1Numbers) > 0 && len(part1Numbers[0]) > 0 {
		cols := len(part1Numbers[0])
		opFields := strings.Fields(opLine)

		currentOps := make([]string, len(opFields))
		copy(currentOps, opFields)

		if len(currentOps) < cols {
			for i := len(currentOps); i < cols; i++ {
				currentOps = append(currentOps, "+")
			}
		}

		for c := 0; c < cols; c++ {
			val := part1Numbers[0][c]
			operator := currentOps[c]

			for r := 1; r < len(part1Numbers); r++ {
				if operator == "+" {
					val += part1Numbers[r][c]
				} else if operator == "*" {
					val *= part1Numbers[r][c]
				}
			}
			part1Total += val
		}
	}
	fmt.Println("Part 1 total:", part1Total)

	maxLen := 0
	for _, line := range dataLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	if len(opLine) > maxLen {
		maxLen = len(opLine)
	}

	grid := make([][]rune, len(dataLines))
	for r, line := range dataLines {
		runes := []rune(line)
		for len(runes) < maxLen {
			runes = append(runes, ' ')
		}
		grid[r] = runes
	}

	opRunes := []rune(opLine)
	for len(opRunes) < maxLen {
		opRunes = append(opRunes, ' ')
	}

	ans2 := 0
	currentGroupNums := []int{}

	for col := maxLen - 1; col >= 0; col-- {
		numStr := ""
		for r := 0; r < len(grid); r++ {
			char := grid[r][col]
			if char >= '0' && char <= '9' {
				numStr += string(char)
			}
		}

		opChar := opRunes[col]
		hasNumber := len(numStr) > 0
		isOpColumn := (opChar == '+' || opChar == '*')

		if hasNumber {
			num, _ := strconv.Atoi(numStr)
			currentGroupNums = append(currentGroupNums, num)
		}

		if isOpColumn {
			if len(currentGroupNums) > 0 {
				val := currentGroupNums[0]
				for i := 1; i < len(currentGroupNums); i++ {
					if opChar == '+' {
						val += currentGroupNums[i]
					} else if opChar == '*' {
						val *= currentGroupNums[i]
					}
				}
				ans2 += val
				currentGroupNums = []int{}
			}
		}
	}

	fmt.Println("Part 2 ", ans2)
}
