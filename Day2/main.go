package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum int64 = int64(0)
var sumPart2 int64 = int64(0)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: Please make sure 'input.txt' is in the current directory.")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, ",")
		for _, id := range ids {
			rang := strings.Split(id, "-")
			start, _ := strconv.ParseInt(rang[0], 10, 64)
			end, _ := strconv.ParseInt(rang[1], 10, 64)
			checkRange(start, end)
			checkRangePart2(start, end)

		}
	}

	fmt.Println(sum)
	fmt.Println(sumPart2)

}

func checkRange(start, end int64) {
	for i := start; i <= end; i++ {
		s := fmt.Sprintf("%d", i)
		l := len(s)

		if l%2 != 0 {
			continue
		}

		size := l / 2
		pattern := s[:size]

		if s[size:] == pattern {
			fmt.Println("Found a valid number:", i)
			sum += int64(i)
		}
	}
}

func checkRangePart2(start, end int64) {
	for i := start; i <= end; i++ {
		s := fmt.Sprintf("%d", i)
		l := len(s)

		isInvalid := false

		for p := 1; p <= l/2; p++ {
			if l%p != 0 {
				continue
			}

			baseSequence := s[:p]

			isRepeating := true
			for j := p; j < l; j += p {
				currentBlock := s[j : j+p]

				if currentBlock != baseSequence {
					isRepeating = false
					break
				}
			}
			if isRepeating {
				isInvalid = true
				break
			}
		}

		if isInvalid {
			sumPart2 += int64(i)
		}
	}
}
