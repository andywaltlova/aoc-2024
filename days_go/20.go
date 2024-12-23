package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Grid struct {
	cells []string
	rows  int
	cols  int
}

func NewGrid(cells []string) *Grid {
	return &Grid{
		cells: cells,
		rows:  len(cells),
		cols:  len(cells[0]),
	}
}

func (g *Grid) IsWall(x, y int) bool {
	return g.cells[x][y] == '#'
}

func (g *Grid) IsInBounds(x, y int) bool {
	return x >= 0 && x < g.rows && y >= 0 && y < g.cols
}

func (g *Grid) FindStartEnd() ([2]int, [2]int) {
	var start, end [2]int
	for i, row := range g.cells {
		for j, cell := range row {
			if cell == 'S' {
				start = [2]int{i, j}
			} else if cell == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	return start, end
}

type State struct {
	x, y, cost, estimatedTotal int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].estimatedTotal < pq[j].estimatedTotal
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

func findShortcutsOptimized(grid *Grid, start, end [2]int, maxInvincibility, threshold int) map[[4]int]int {
	// Compute the shortest path and record path cells
	shortestPath, pathCells := aStarAndRecordPath(grid, start, end)
	if shortestPath == -1 {
		fmt.Println("No valid path found")
		return nil
	}

	// Simulate invincibility (cheat) activation from each cell on the original path
	shortcuts := make(map[[4]int]int)
	for _, cell := range pathCells {
		x, y, costToCell := cell[0], cell[1], cell[2]
		simulatedShortcuts := simulateUniquePaths(grid, pathCells, x, y, costToCell, maxInvincibility, shortestPath, threshold)
		for startEnd, timesaved := range simulatedShortcuts {
			shortcuts[startEnd] = timesaved
		}
	}

	return shortcuts
}

func simulateUniquePaths(grid *Grid, pathCells [][3]int, activationX, activationY, costToActivation, maxInvincibility, originalPath, threshold int) map[[4]int]int {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][][]bool, grid.rows)
	for i := range visited {
		visited[i] = make([][]bool, grid.cols)
		for j := range visited[i] {
			visited[i][j] = make([]bool, maxInvincibility+1)
		}
	}

	type State struct {
		x, y, cost, timeLeft int
	}

	// Start exploring from the activation point
	queue := []State{
		{x: activationX, y: activationY, cost: costToActivation, timeLeft: maxInvincibility},
	}

	shortcuts := make(map[[4]int]int) // Key: [startX, startY, endX, endY], Value: time saved

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If invincibility ends in a wall, skip this path
		if curr.timeLeft == 0 && grid.IsWall(curr.x, curr.y) {
			continue
		}

		// Skip if already visited with the same remaining invincibility time
		if visited[curr.x][curr.y][curr.timeLeft] {
			continue
		}
		visited[curr.x][curr.y][curr.timeLeft] = true

		// Check if we're back on the shortest path closer to the end
		if curr.timeLeft == 0 {
			if !grid.IsWall(curr.x, curr.y) {
				for _, cell := range pathCells {
					if cell[0] == curr.x && cell[1] == curr.y {
						remainingCost := pathCells[len(pathCells)-1][2] - cell[2]
						newPathLength := curr.cost + remainingCost

						if originalPath-newPathLength >= threshold {
							timeSaved := originalPath - newPathLength
							shortcuts[[4]int{activationX, activationY, curr.x, curr.y}] = timeSaved
							// fmt.Printf("From activation point %d,%d, to %d,%d, time saved: %d\n", activationX, activationY, curr.x, curr.y, timeSaved)
						}
						break
					}
				}
			}
			continue
		}

		// Explore neighbors
		for _, dir := range directions {
			nx, ny := curr.x+dir[0], curr.y+dir[1]
			if !grid.IsInBounds(nx, ny) {
				continue
			}

			queue = append(queue, State{
				x:        nx,
				y:        ny,
				cost:     curr.cost + 1,
				timeLeft: max(0, curr.timeLeft-1),
			})
		}
	}

	return shortcuts
}

func aStarAndRecordPath(grid *Grid, start, end [2]int) (int, [][3]int) {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, grid.rows)
	for i := range visited {
		visited[i] = make([]bool, grid.cols)
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{
		x:              start[0],
		y:              start[1],
		cost:           0,
		estimatedTotal: heuristic(start[0], start[1], end[0], end[1]),
	})

	var pathCells [][3]int
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(State)
		x, y, cost := curr.x, curr.y, curr.cost

		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		pathCells = append(pathCells, [3]int{x, y, cost})

		if x == end[0] && y == end[1] {
			return cost, pathCells
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if grid.IsInBounds(nx, ny) && !grid.IsWall(nx, ny) && !visited[nx][ny] {
				heap.Push(pq, State{
					x:              nx,
					y:              ny,
					cost:           cost + 1,
					estimatedTotal: cost + 1 + heuristic(nx, ny, end[0], end[1]),
				})
			}
		}
	}

	return -1, nil
}

func heuristic(x, y, ex, ey int) int {
	return int(math.Abs(float64(x-ex)) + math.Abs(float64(y-ey)))
}

func main() {
	grid := getInputLines("../data/20.txt")

	g := NewGrid(grid)
	start, end := g.FindStartEnd()

	shortcuts := findShortcutsOptimized(g, start, end, 2, 100)
	fmt.Println("Total shortcuts part 1:", len(shortcuts))

	allShortcuts := make(map[[4]int]int)
	for invincibility := 1; invincibility <= 20; invincibility++ {
		fmt.Println("Simulating invincibility:", invincibility)
		shortcuts := findShortcutsOptimized(g, start, end, invincibility, 100)
		for k, v := range shortcuts {
			allShortcuts[k] = v
		}
	}
	fmt.Println("Total shortcuts part 2:", len(allShortcuts))
}
