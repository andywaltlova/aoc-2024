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

type guard struct {
	startPos  coordinate
	startDir  int
	position  coordinate
	direction int
	visited   map[coordinate]struct{}
	grid      map[coordinate]string
}

func newGuard(start coordinate, direction int, grid map[coordinate]string) *guard {
	g := guard{
		startPos:  start,
		startDir:  direction,
		position:  start,
		direction: direction,
		visited:   make(map[coordinate]struct{}),
		grid:      grid,
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

func (g *guard) move() (moved bool, turned bool) {
	direction, _ := directions[g.direction]
	newCoordinates := coordinate{
		x: g.position.x + direction[0],
		y: g.position.y + direction[1],
	}

	gridLen := int(math.Sqrt(float64(len(g.grid))))
	if newCoordinates.x < 0 || newCoordinates.x == gridLen || newCoordinates.y < 0 || newCoordinates.y == gridLen {
		return false, false
	}

	if val, _ := g.grid[newCoordinates]; val == "#" {
		g.turnClockWise()
		return false, true
	}

	g.position = newCoordinates
	g.visited[g.position] = struct{}{}
	return true, false
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
		guardMoved, guardTurned := guard.move()
		guardFinished = !(guardMoved || guardTurned)
	}
	return len(guard.visited)
}

type seenCoordinate struct {
	coordinate coordinate
	direction  int
}

func part2(guard *guard, input []string) (int, int) {
	obstacles := make(map[coordinate]struct{})
	guardFinished := false
	previousPosition := guard.position
	for !guardFinished {
		guardMoved, guardTurned := guard.move()
		guardFinished = !(guardMoved || guardTurned)
		if guardMoved && previousPosition != guard.startPos {
			newGrid := getGrid(input)
			newGrid[guard.position] = "#"
			newGrid[guard.startPos] = "."
			newG := newGuard(previousPosition, guard.direction, newGrid)

			seen := make(map[seenCoordinate]struct{})
			seen[seenCoordinate{coordinate: previousPosition, direction: guard.direction}] = struct{}{}

			alternativeGuardFinished := false
			for !alternativeGuardFinished {
				guardMoved, guardTurned = newG.move()
				seenCoordinate := seenCoordinate{coordinate: newG.position, direction: newG.direction}
				if _, ok := seen[seenCoordinate]; ok && guardMoved {
					obstacles[guard.position] = struct{}{}
					break
				}
				seen[seenCoordinate] = struct{}{}
				alternativeGuardFinished = !(guardMoved || guardTurned)
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
