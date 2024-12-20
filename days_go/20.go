package main

import (
	"container/heap"
	"fmt"
	"math"
)

// State represents the current position in the grid and the cost to reach it
type State struct {
	f    int
	cost int
	x, y int
}

// PriorityQueue implements a min-heap for State
type PriorityQueue []State

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func aStar(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])

	start, end := [2]int{}, [2]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start[0], start[1] = i, j
			}
			if cell == 'E' {
				end[0], end[1] = i, j
			}
		}
	}

	sx, sy := start[0], start[1]
	ex, ey := end[0], end[1]

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	heuristic := func(x, y int) int {
		return int(math.Abs(float64(x-ex)) + math.Abs(float64(y-ey)))
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{f: heuristic(sx, sy), cost: 0, x: sx, y: sy})

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(State)
		x, y, cost := curr.x, curr.y, curr.cost

		if x == ex && y == ey {
			return cost
		}

		if visited[x][y] {
			continue
		}
		visited[x][y] = true

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				if grid[nx][ny] != '#' && !visited[nx][ny] {
					heap.Push(pq, State{
						f:    cost + 1 + heuristic(nx, ny),
						cost: cost + 1,
						x:    nx,
						y:    ny,
					})
				}
			}
		}
	}
	return -1
}

func main() {
	grid := getInputLines("../data/20_test.txt")
	result := aStar(grid)
	if result == -1 {
		fmt.Println("No path found")
	} else {
		fmt.Printf("Shortest path without wall-breaking: %d\n", result)
	}
}
