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
	// ans2 := 0

	type Interval struct {
		From int64
		To   int64
	}
	ids := make([]int64, 0)

	ranges := make([]Interval, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			first, _ := strconv.ParseInt(parts[0], 10, 64)
			last, _ := strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, Interval{first, last})
		} else {
			v, _ := strconv.ParseInt(line, 10, 64)
			ids = append(ids, v)
		}
	}

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.From && id <= r.To {
				ans++
				break
			}
		}
	}

	fmt.Println(ans)

}
