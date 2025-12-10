package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	From int64
	To   int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: Please make sure 'input.txt' is in the current directory.")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans := 0
	ans2 := int64(0)

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

	merged := mergeIntervals(ranges)
	for _, r := range merged {
		ans2 += r.To - r.From + 1
	}

	fmt.Println(ans, ans2)

}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].From < intervals[j].From
	})

	merged := []Interval{intervals[0]}

	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]

		if curr.From <= last.To {
			if curr.To > last.To {
				last.To = curr.To
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
