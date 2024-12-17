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

func (r *robot) movePart1AndPart2LeftAndRight(direction string) (moved bool) {
	// Move the robot in the direction if possible
	// Return true if robot moved, false otherwise
	empty, exists := r.g.canMoveBoxesPart1AndPart2LeftAndRight(r.c, direction)
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

func (g *grid) canMoveBoxesPart1AndPart2LeftAndRight(c coordinate, d string) (empty coordinate, exists bool) {
	// Search for the first empty cell in the direction
	for {
		newC := coordinate{c.x + directionMap[d].x, c.y + directionMap[d].y}
		if g.content[newC] == "#" {
			return c, false
		}
		if g.content[newC] == "." {
			return newC, true
		}
		c = newC
	}
}

func (r *robot) canMoveBoxesPart2UpAndDown(direction string) (bool, []coordinate) {
	// Trace the path of the robot
	dc := directionMap[direction]
	checkForEmpty := make(map[coordinate]struct{})
	checkForEmpty[coordinate{r.c.x + dc.x, r.c.y + dc.y}] = struct{}{}
	var toMoveMap = make(map[coordinate]struct{})
	var toMove []coordinate
	for len(checkForEmpty) > 0 {
		newCheckForEmpty := make(map[coordinate]struct{})
		for c, _ := range checkForEmpty {
			switch r.g.content[c] {
			case "]":
				left := coordinate{c.x, c.y - 1}
				leftNext := coordinate{left.x + dc.x, left.y + dc.y}
				currentNext := coordinate{c.x + dc.x, c.y + dc.y}
				newCheckForEmpty[leftNext] = struct{}{}
				newCheckForEmpty[currentNext] = struct{}{}

				if _, ok := toMoveMap[left]; !ok {
					toMoveMap[left] = struct{}{}
					toMove = append(toMove, left)
				}
				if _, ok := toMoveMap[c]; !ok {
					toMoveMap[c] = struct{}{}
					toMove = append(toMove, c)
				}
			case "[":
				right := coordinate{c.x, c.y + 1}
				rightNext := coordinate{right.x + dc.x, right.y + dc.y}
				currentNext := coordinate{c.x + dc.x, c.y + dc.y}
				newCheckForEmpty[rightNext] = struct{}{}
				newCheckForEmpty[currentNext] = struct{}{}

				if _, ok := toMoveMap[right]; !ok {
					toMoveMap[right] = struct{}{}
					toMove = append(toMove, right)
				}
				if _, ok := toMoveMap[c]; !ok {
					toMoveMap[c] = struct{}{}
					toMove = append(toMove, c)
				}
			case "#":
				return false, nil
			}
		}
		checkForEmpty = newCheckForEmpty
	}
	return true, toMove
}

func (r *robot) movePart2UpAndDown(direction string) (moved bool) {
	// Move the robot in the direction if possible
	// Return true if robot moved, false otherwise
	canBeMoved, toMove := r.canMoveBoxesPart2UpAndDown(direction)
	if !canBeMoved {
		return false
	}
	// Everything between the robot and the empty cell is moved
	dc := directionMap[direction]
	for i := len(toMove) - 1; i >= 0; i-- {
		c := toMove[i]
		empty := coordinate{c.x + dc.x, c.y + dc.y}
		r.g.content[empty] = r.g.content[c]
		r.g.content[c] = "."
	}
	r.g.content[r.c] = "."
	r.c = coordinate{r.c.x + directionMap[direction].x, r.c.y + directionMap[direction].y}
	r.g.content[r.c] = "@"
	return true
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
		r.movePart1AndPart2LeftAndRight(direction)
	}
	boxesGPS := 0
	for c, value := range r.g.content {
		if value == "O" {
			boxesGPS += 100*c.x + c.y
		}
	}
	return boxesGPS
}

func (r *robot) createExtendedGrid() {
	extendedContent := make(map[coordinate]string)
	for x := 0; x < r.g.xSize; x++ {
		newY := 1
		for y := 0; y < r.g.ySize; y++ {
			originalContent := r.g.content[coordinate{x, y}]

			newContent := "."
			if originalContent == "#" {
				newContent = "#"
			} else if originalContent == "O" {
				originalContent = "["
				newContent = "]"
			}

			if originalContent == "@" {
				r.c = coordinate{x, y + newY - 1}
			}

			extendedContent[coordinate{x, y + newY - 1}] = originalContent
			extendedContent[coordinate{x, y + newY}] = newContent
			newY++
		}
	}
	r.g = grid{extendedContent, r.g.xSize, r.g.ySize * 2}
}

func part2(r robot, instructions string) int {
	// Move the robot according to the instructions
	r.createExtendedGrid()

	for _, instruction := range instructions {
		direction := string(instruction)
		if direction == ">" || direction == "<" {
			r.movePart1AndPart2LeftAndRight(direction)
		} else {
			r.movePart2UpAndDown(direction)
		}
	}
	boxesGPS := 0
	for c, value := range r.g.content {
		if value == "[" {
			boxesGPS += 100*c.x + c.y
		}
	}
	return boxesGPS
}

func main() {
	robot, instructions := parseInput("../data/15.txt")
	// fmt.Println(part1(robot, instructions))
	fmt.Println(part2(robot, instructions))
}
