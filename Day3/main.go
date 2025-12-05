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
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, len(line))

		for i, c := range line {
			digits[i] = int(c - '0')
		}

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
	}

	fmt.Println(ans)

	// fmt.Println(sum)

}
