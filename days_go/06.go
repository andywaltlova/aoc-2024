package main

import (
	"fmt"
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
	position  coordinate
	direction int
	visited   map[coordinate]struct{}
	grid      []string
}

func newGuard(x, y int, direction int, grid []string) *guard {
	position := coordinate{x: x, y: y}
	g := guard{
		position:  position,
		direction: direction,
		visited:   make(map[coordinate]struct{}),
		grid:      grid,
	}
	g.visited[position] = struct{}{}
	return &g
}

func (g *guard) turnClockWise() {
	if g.direction == 4 {
		g.direction = 1
	} else {
		g.direction += 1
	}
}

func (g *guard) move() bool {
	direction, _ := directions[g.direction]
	newCoordinates := coordinate{
		x: g.position.x + direction[0],
		y: g.position.y + direction[1],
	}

	if newCoordinates.x < 0 || newCoordinates.x == len(g.grid) || newCoordinates.y < 0 || newCoordinates.y == len(g.grid) {
		return false
	}

	if string(g.grid[newCoordinates.x][newCoordinates.y]) == "#" {
		g.turnClockWise()
		return true
	}

	g.position = newCoordinates
	g.visited[g.position] = struct{}{}
	return true
}

func getGuard(grid []string) *guard {
	var guard *guard
	for x, line := range grid {
		for y, char := range line {
			if string(char) == "^" {
				guard = newGuard(x, y, 1, grid)
				break
			}

		}
	}
	return guard
}

func part1(guard *guard) int {
	guardMoved := guard.move()
	for guardMoved {
		guardMoved = guard.move()
	}
	return len(guard.visited)
}

func part2(guard []string) int {
	loops := 0
	return loops
}

func main() {
	filepath := "../data/06.txt"
	content := getInputLines(filepath)
	guard := getGuard(content)
	fmt.Println(part1(guard))
	fmt.Println(part2(content))
}
