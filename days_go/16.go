package main

import (
	"container/heap"
	"fmt"
	"math"
)

var directions = []struct {
	dx, dy int
}{
	{1, 0},  // East
	{0, -1}, // North
	{0, 1},  // South
	{-1, 0}, // West
}

type state struct {
	x, y, dir, score int
}

type priorityQueue []*state

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*state))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(maze []string) ([]state, int) {
	rows, cols := len(maze), len(maze[0])
	startX, startY := 0, 0
	endX, endY := 0, 0

	// Locate start and end
	for y, row := range maze {
		for x, tile := range row {
			if tile == 'S' {
				startX, startY = x, y
			}
			if tile == 'E' {
				endX, endY = x, y
			}
		}
	}

	// Visited map to track (x, y, dir) states
	visited := make(map[[3]int]bool)

	distance := make(map[state]int)
	predecessor := make(map[state]*state) // To track the path
	pq := &priorityQueue{}
	heap.Init(pq)

	// facing East (0)
	initial := state{startX, startY, 0, 0}
	heap.Push(pq, &initial)
	distance[initial] = 0

	var endState *state

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*state)
		if curr.x == endX && curr.y == endY {
			endState = curr
			break
		}

		stateKey := [3]int{curr.x, curr.y, curr.dir}
		if visited[stateKey] {
			continue
		}
		visited[stateKey] = true

		for i := 0; i < 4; i++ {
			newDir := i
			rotationCost := 0
			if newDir != curr.dir {
				rotationCost = 1000
			}

			dx, dy := directions[newDir].dx, directions[newDir].dy
			nx, ny := curr.x+dx, curr.y+dy

			if nx >= 0 && ny >= 0 && nx < cols && ny < rows && maze[ny][nx] != '#' {
				newState := state{nx, ny, newDir, curr.score + 1 + rotationCost}
				if oldCost, ok := distance[newState]; !ok || newState.score <= oldCost {
					distance[newState] = newState.score
					heap.Push(pq, &newState)

					predecessor[newState] = curr // Track the predecessor
				}
			}
		}
	}

	// Reconstruct the path
	path := []state{}
	for s := endState; s != nil; s = predecessor[*s] {
		path = append([]state{*s}, path...) // Prepend to build path in order
	}

	if endState == nil {
		return nil, math.MaxInt32 // No path found
	}
	return path, endState.score
}

func main() {
	maze := getInputLines("../data/16_test.txt")
	path, cost := dijkstra(maze)
	fmt.Println("Cost:", cost)
	fmt.Println("Path:", len(path))

	// Print the path
	for x := 0; x < len(maze[0]); x++ {
		for y := 0; y < len(maze); y++ {
			tile := maze[x][y]
			for _, state := range path {
				if state.x == y && state.y == x {
					tile = 'o'
				}
			}
			fmt.Printf("%c", tile)
		}
		fmt.Println()
	}
}
