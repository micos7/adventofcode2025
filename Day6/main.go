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
		fmt.Println("Error opening file: Please make sure 'input.txt' is in the current directory.")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans := 0
	var rows [][]int
	var ops []rune

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// fmt.Println(line)
		fields := strings.Fields(line)

		if len(fields) > 0 && (fields[0] == "+" || fields[0] == "*") {
			for _, o := range fields {
				ops = append(ops, rune(o[0]))
			}
			continue
		}

		var nums []int
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println("Error parsing ", err)
				return
			}
			nums = append(nums, n)
		}
		rows = append(rows, nums)

	}

	cols := len(rows[0])

	for c := range cols {
		operation := ops[c]
		value := rows[0][c]
		for _, row := range rows[1:] {
			if operation == '+' {
				value += row[c]
			} else if operation == '*' {
				value *= row[c]
			}
		}
		ans += value
	}

	fmt.Println(ans)

}
