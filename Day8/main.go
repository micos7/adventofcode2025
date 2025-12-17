package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j int
	dist int
}

var parent []int
var rank []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := []Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y, z int
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		points = append(points, Point{x, y, z})
	}

	n := len(points)

	edges := []Edge{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, Edge{
				i:    i,
				j:    j,
				dist: squaredDistance(points[i], points[j]),
			})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].dist < edges[b].dist
	})

	parent = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	attempts := 0
	for _, e := range edges {
		if attempts == 1000 {
			break
		}
		union(e.i, e.j)
		attempts++
	}

	componentSize := map[int]int{}
	for i := 0; i < n; i++ {
		root := find(i)
		componentSize[root]++
	}

	sizes := []int{}
	for _, size := range componentSize {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	part1 := sizes[0] * sizes[1] * sizes[2]

	parent = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	components := n
	part2 := 0

	for _, e := range edges {
		if union(e.i, e.j) {
			components--
			if components == 1 {
				part2 = points[e.i].x * points[e.j].x
				break
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func squaredDistance(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func union(x, y int) bool {
	rootX := find(x)
	rootY := find(y)
	if rootX == rootY {
		return false
	}

	if rank[rootX] < rank[rootY] {
		parent[rootX] = rootY
	} else if rank[rootX] > rank[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rank[rootX]++
	}
	return true
}
