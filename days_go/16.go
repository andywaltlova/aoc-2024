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

func dijkstra(maze []string) ([][]state, int) {
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

	visited := make(map[[3]int]bool)
	distance := make(map[state]int)
	predecessors := make(map[state][]state) // Support multiple predecessors
	pq := &priorityQueue{}
	heap.Init(pq)

	initial := state{startX, startY, 0, 0}
	heap.Push(pq, &initial)
	distance[initial] = 0

	allGoalStates := []state{}
	minGoalCost := math.MaxInt32

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*state)

		if curr.x == endX && curr.y == endY {
			if curr.score < minGoalCost {
				minGoalCost = curr.score
				allGoalStates = []state{*curr}
			} else if curr.score == minGoalCost {
				allGoalStates = append(allGoalStates, *curr)
			}
			continue
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
				newScore := curr.score + 1 + rotationCost
				newState := state{nx, ny, newDir, newScore}

				oldCost, exists := distance[newState]

				if !exists || newScore < oldCost {
					distance[newState] = newScore
					heap.Push(pq, &state{nx, ny, newDir, newScore})
					predecessors[newState] = []state{*curr}
				} else if newScore == oldCost {
					// Add as an alternative predecessor
					predecessors[newState] = append(predecessors[newState], *curr)
				}
			}
		}
	}

	// Reconstruct all paths
	var allPaths [][]state
	for _, goalState := range allGoalStates {
		var paths [][]state
		reconstructPaths(&paths, []state{}, goalState, predecessors)
		allPaths = append(allPaths, paths...)
	}

	return allPaths, minGoalCost
}

func reconstructPaths(paths *[][]state, currentPath []state, curr state, predecessors map[state][]state) {
	currentPath = append([]state{curr}, currentPath...)
	if preds, ok := predecessors[curr]; ok {
		for _, pred := range preds {
			reconstructPaths(paths, currentPath, pred, predecessors)
		}
	} else {
		// start
		*paths = append(*paths, currentPath)
	}
}

func main() {
	maze := getInputLines("../data/16.txt")
	paths, cost := dijkstra(maze)
	fmt.Println("Cost:", cost)

	uniqueCoords := make(map[[2]int]bool)
	for _, path := range paths {
		for _, state := range path {
			uniqueCoords[[2]int{state.x, state.y}] = true
		}
	}
	fmt.Println("Unique coordinates:", len(uniqueCoords))

	// Print the path
	// for x := 0; x < len(maze[0]); x++ {
	// 	for y := 0; y < len(maze); y++ {
	// 		tile := maze[x][y]
	// 		if uniqueCoords[[2]int{y, x}] {
	// 			tile = 'O'
	// 		}
	// 		fmt.Printf("%c", tile)
	// 	}
	// 	fmt.Println()
	// }
}
