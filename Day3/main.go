package main

import (
	"bufio"
	"fmt"
	"os"
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
	ans_ar := make([]int, 12)
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, len(line))

		for i, c := range line {
			digits[i] = int(c - '0')
		}
		//part1
		max1 := digits[0]
		maxIdx := 0
		for i := 0; i < len(digits)-1; i++ {
			if digits[i] > max1 {
				max1 = digits[i]
				maxIdx = i
			}
		}

		max2 := 0
		for i := maxIdx + 1; i < len(digits); i++ {
			max2 = max(max2, digits[i])
		}
		ans += max1*10 + max2

		//part2

		current_search_start_idx := 0

		for ans_id := range 12 {
			last_valid_idx := len(digits) - (12 - ans_id)

			max_digit := -1
			foundIdx := -1

			for curIdx := current_search_start_idx; curIdx <= last_valid_idx; curIdx++ {
				if digits[curIdx] > max_digit {
					max_digit = digits[curIdx]
					foundIdx = curIdx
				}
				if max_digit == 9 {
					break
				}
			}

			ans_ar[ans_id] = max_digit

			current_search_start_idx = foundIdx + 1
		}
		fmt.Println(ans_ar)
		ans2 += DigitsToInt(ans_ar)
	}
	fmt.Println(ans)
	fmt.Println(ans2)

}

func DigitsToInt(digits []int) int {
	n := 0
	for _, d := range digits {
		n = n*10 + d
	}
	return n
}
