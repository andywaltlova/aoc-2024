package main

import (
	"fmt"
	"math"
)

var directions = map[int][]int{
	1: []int{-1, 0}, // Up
	2: []int{0, 1},  // Right
	3: []int{1, 0},  // Down
	4: []int{0, -1}, // Left
}

type coordinate struct {
	x, y int
}

type dirCoordinate struct {
	coordinate coordinate
	direction  int
}
type guard struct {
	startPos   coordinate
	startDir   int
	position   coordinate
	direction  int
	visited    map[coordinate]struct{}
	visitedDir map[dirCoordinate]struct{}
	grid       map[coordinate]string
}

func newGuard(start coordinate, direction int, grid map[coordinate]string) *guard {
	g := guard{
		startPos:   start,
		startDir:   direction,
		position:   start,
		direction:  direction,
		visited:    make(map[coordinate]struct{}),
		visitedDir: make(map[dirCoordinate]struct{}),
		grid:       grid,
	}
	g.visited[start] = struct{}{}
	return &g
}

func (g *guard) turnClockWise() {
	if g.direction == 4 {
		g.direction = 1
	} else {
		g.direction += 1
	}
}

func (g *guard) move() (moved bool, turned bool, loop bool) {
	direction, _ := directions[g.direction]
	newCoordinates := coordinate{
		x: g.position.x + direction[0],
		y: g.position.y + direction[1],
	}

	if val, _ := g.grid[newCoordinates]; val == "#" {
		g.turnClockWise()
		return false, true, false
	}

	g.position = newCoordinates

	visitedDir := dirCoordinate{coordinate: g.position, direction: g.direction}
	_, alreadyVisited := g.visitedDir[visitedDir]
	if alreadyVisited {
		return false, false, true
	}

	gridLen := int(math.Sqrt(float64(len(g.grid))))
	if g.position.x < 0 || g.position.x == gridLen || g.position.y < 0 || g.position.y == gridLen {
		return false, false, false
	}

	g.visitedDir[visitedDir] = struct{}{}
	g.visited[g.position] = struct{}{}

	return true, false, false
}

func getGuard(grid map[coordinate]string) *guard {
	var guard *guard
	for coordinate, value := range grid {
		if value == "^" {
			guard = newGuard(coordinate, 1, grid)
			break
		}
	}
	return guard
}

func getGrid(grid []string) map[coordinate]string {
	gridMap := make(map[coordinate]string)
	for x, line := range grid {
		for y, char := range line {
			gridMap[coordinate{x: x, y: y}] = string(char)
		}
	}
	return gridMap
}

func part1(guard *guard) int {
	guardFinished := false
	for !guardFinished {
		guardMoved, guardTurned, _ := guard.move()
		guardFinished = !(guardMoved || guardTurned)
	}
	return len(guard.visited)
}

func part2(guard *guard, input []string) (int, int) {
	obstacles := make(map[coordinate]struct{})
	guardFinished := false
	previousPosition := guard.position
	for !guardFinished {
		guardMoved, guardTurned, _ := guard.move()
		guardFinished = !(guardMoved || guardTurned)
		_, alreadyTried := obstacles[guard.position]
		if guardMoved && previousPosition != guard.position && !alreadyTried {
			newGrid := getGrid(input)
			newGrid[guard.position] = "#"
			newG := newGuard(guard.startPos, guard.startDir, newGrid)

			alternativeGuardFinished := false
			for !alternativeGuardFinished {
				guardMoved, guardTurned, loop := newG.move()
				alternativeGuardFinished = !(guardMoved || guardTurned)

				if loop {
					obstacles[guard.position] = struct{}{}
					break
				}
			}
		}
		previousPosition = guard.position
	}
	return len(guard.visited), len(obstacles)
}

func main() {
	filepath := "../data/06.txt"

	input := getInputLines(filepath)
	grid := getGrid(input)

	guard1 := getGuard(grid)
	fmt.Println(part1(guard1))

	guard2 := getGuard(grid)
	fmt.Println(part2(guard2, input))
}
