package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	coord    coordinate
	priority int
	steps    int
	index    int
}

type PriorityQueue []*item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type coordinate struct {
	x, y int
}

func readInputFromFile(filename string) ([]coordinate, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var coordinates []coordinate
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 == nil && err2 == nil {
			coordinates = append(coordinates, coordinate{x: x, y: y})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return coordinates, nil
}

func simulateFallingBytes(grid map[coordinate]bool, bytes []coordinate) {
	for _, b := range bytes {
		grid[b] = true
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func heuristic(a, b coordinate) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func aStar(grid map[coordinate]bool, start, goal coordinate) int {
	directions := []coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{coord: start, priority: 0, steps: 0})
	visited := make(map[coordinate]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*item)
		if current.coord == goal {
			return current.steps
		}
		if visited[current.coord] {
			continue
		}
		visited[current.coord] = true

		for _, d := range directions {
			neighbor := coordinate{x: current.coord.x + d.x, y: current.coord.y + d.y}
			if neighbor.x < start.x || neighbor.y < start.y || neighbor.x > goal.x || neighbor.y > goal.y {
				continue
			}
			if grid[neighbor] || visited[neighbor] {
				continue
			}
			heap.Push(pq, &item{
				coord:    neighbor,
				priority: current.steps + 1 + heuristic(neighbor, goal),
				steps:    current.steps + 1,
			})
		}
	}

	return -1
}

func main() {
	coordinates, _ := readInputFromFile("../data/18.txt")
	grid := make(map[coordinate]bool)
	simulateFallingBytes(grid, coordinates[:1024])

	start := coordinate{0, 0}
	goal := coordinate{70, 70}
	steps := aStar(grid, start, goal)
	fmt.Println(steps)
}
