package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

var parent []int
var rank []int

var part1 = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := []Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		points = append(points, Point{x, y})
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := (abs(points[i].x-points[j].x) + 1) * (abs(points[i].y-points[j].y) + 1)
			part1 = max(part1, dist)
		}
	}

	fmt.Println("Part 1:", part1)
	solvePart2(points)
	fmt.Println("Part 2:", part2)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var part2 = 0

func solvePart2(points []Point) {
	// ---- 1. Collect coordinates ----
	xs := []int{}
	ys := []int{}

	for _, p := range points {
		xs = append(xs, p.x)
		ys = append(ys, p.y)
	}

	// also include segment interiors
	for i := 0; i < len(points); i++ {
		a := points[i]
		b := points[(i+1)%len(points)]

		if a.x == b.x {
			ys = append(ys, a.y, b.y)
		} else {
			xs = append(xs, a.x, b.x)
		}
	}

	xs = compress(xs)
	ys = compress(ys)

	xIndex := map[int]int{}
	yIndex := map[int]int{}
	for i, v := range xs {
		xIndex[v] = i
	}
	for i, v := range ys {
		yIndex[v] = i
	}

	w, h := len(xs), len(ys)

	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}

	// ---- 2. Draw loop ----
	for i := 0; i < len(points); i++ {
		a := points[i]
		b := points[(i+1)%len(points)]

		x1, y1 := xIndex[a.x], yIndex[a.y]
		x2, y2 := xIndex[b.x], yIndex[b.y]

		if x1 == x2 {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				grid[y][x1] = 1
			}
		} else {
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				grid[y1][x] = 1
			}
		}
	}

	// ---- 3. Flood fill outside ----
	type P struct{ x, y int }
	queue := []P{{0, 0}}
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	visited[0][0] = true

	dirs := []P{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nx, ny := p.x+d.x, p.y+d.y
			if nx < 0 || ny < 0 || nx >= w || ny >= h {
				continue
			}
			if visited[ny][nx] || grid[ny][nx] == 1 {
				continue
			}
			visited[ny][nx] = true
			queue = append(queue, P{nx, ny})
		}
	}

	// ---- 4. Fill interior ----
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 0 && !visited[y][x] {
				grid[y][x] = 1
			}
		}
	}

	// ---- 5. Prefix sum ----
	ps := make([][]int, h+1)
	for i := range ps {
		ps[i] = make([]int, w+1)
	}

	for y := 1; y <= h; y++ {
		for x := 1; x <= w; x++ {
			ps[y][x] = grid[y-1][x-1] +
				ps[y-1][x] +
				ps[y][x-1] -
				ps[y-1][x-1]
		}
	}

	rectSum := func(x1, y1, x2, y2 int) int {
		return ps[y2+1][x2+1] -
			ps[y1][x2+1] -
			ps[y2+1][x1] +
			ps[y1][x1]
	}

	// ---- 6. Rectangle search ----
	part2 = 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			if a.x == b.x || a.y == b.y {
				continue
			}

			x1 := min(xIndex[a.x], xIndex[b.x])
			x2 := max(xIndex[a.x], xIndex[b.x])
			y1 := min(yIndex[a.y], yIndex[b.y])
			y2 := max(yIndex[a.y], yIndex[b.y])

			if rectSum(x1, y1, x2, y2) != (x2-x1+1)*(y2-y1+1) {
				continue
			}

			realArea := (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
			part2 = max(part2, realArea)
		}
	}
}

func compress(vals []int) []int {
	m := map[int]bool{}
	for _, v := range vals {
		m[v] = true
		m[v-1] = true
		m[v+1] = true
	}
	out := []int{}
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}
