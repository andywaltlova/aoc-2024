package main

import (
	"fmt"
	"os"
	"strings"
)

type coordinate struct {
	x, y int
}

type grid struct {
	content map[coordinate]string
	xSize   int
	ySize   int
}

var directionMap = map[string]coordinate{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

type robot struct {
	c coordinate
	g grid
}

func (r *robot) move(direction string) (moved bool) {
	// Move the robot in the direction if possible
	// Return true if robot moved, false otherwise
	empty, exists := r.g.searchDirection(r.c, direction)
	if !exists {
		return false
	}
	// Everything between the robot and the empty cell is moved
	oppositeDirection := map[string]string{
		">": "<",
		"<": ">",
		"^": "v",
		"v": "^",
	}
	next := coordinate{r.c.x + directionMap[direction].x, r.c.y + directionMap[direction].y}
	for empty != next {
		opDirection := oppositeDirection[direction]
		beforeEmpty := coordinate{empty.x + directionMap[opDirection].x, empty.y + directionMap[opDirection].y}
		r.g.content[empty] = r.g.content[beforeEmpty]
		r.g.content[beforeEmpty] = "."
		empty = beforeEmpty
	}
	r.g.content[r.c] = "."
	r.c = next
	r.g.content[r.c] = "@"
	return true
}

func (g *grid) isOutOfBounds(c coordinate) bool {
	return c.x < 0 || c.x >= g.xSize || c.y < 0 || c.y >= g.ySize
}

func (g *grid) searchDirection(c coordinate, d string) (empty coordinate, exists bool) {
	// Search for the first empty cell in the direction
	for {
		newC := coordinate{c.x + directionMap[d].x, c.y + directionMap[d].y}
		if g.content[newC] == "#" {
			return c, false
		}

		if g.isOutOfBounds(newC) {
			return c, false
		}
		if g.content[newC] == "." {
			return newC, true
		}
		c = newC
	}
}

func (g *grid) printGrid() {
	for x := 0; x < g.xSize; x++ {
		for y := 0; y < g.ySize; y++ {
			fmt.Print(g.content[coordinate{x, y}])
		}
		fmt.Println()
	}
}

func parseInput(filename string) (robot, string) {
	data, _ := os.ReadFile(filename)

	parts := strings.Split(string(data), "\n\n")
	instructions := strings.ReplaceAll(parts[1], "\n", "")

	var r robot

	// Grid
	gridMapData := make(map[coordinate]string)
	gridContent := strings.Split(parts[0], "\n")
	xSize := len(gridContent)
	ySize := len(string(gridContent[0]))
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			value := string(gridContent[x][y])
			if value == "@" {
				r.c = coordinate{x, y}
			}
			gridMapData[coordinate{x, y}] = value
		}
	}

	r.g = grid{gridMapData, xSize, ySize}

	return r, instructions
}

func part1(r robot, instructions string) int {
	// Move the robot according to the instructions
	for _, instruction := range instructions {
		direction := string(instruction)
		r.move(direction)
	}
	boxesGPS := 0
	for c, value := range r.g.content {
		if value == "O" {
			boxesGPS += 100*c.x + c.y
		}
	}
	return boxesGPS
}

func main() {
	robot, instructions := parseInput("../data/15.txt")
	fmt.Println(part1(robot, instructions))
}
