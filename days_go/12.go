package main

import (
	"fmt"
)

type coordinate struct {
	x, y int
}

type plot struct {
	cells     []coordinate
	perimeter int
	label     string
}

var directions = []coordinate{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func isInBounds(grid []string, coord coordinate) bool {
	return coord.x >= 0 && coord.x < len(grid) && coord.y >= 0 && coord.y < len(grid[0])
}

func findPlots(grid []string) []plot {
	visited := make(map[coordinate]bool)
	var plots []plot

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			coord := coordinate{row, col}
			if !visited[coord] {
				plot := explorePlot(grid, coord, visited)
				plots = append(plots, plot)
			}
		}
	}
	return plots
}

func explorePlot(grid []string, start coordinate, visited map[coordinate]bool) plot {
	var plot plot
	stack := []coordinate{start}
	char := grid[start.x][start.y]
	plot.label = string(char)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[curr] {
			continue
		}
		visited[curr] = true
		plot.cells = append(plot.cells, curr)

		for _, dir := range directions {
			neighbor := coordinate{curr.x + dir.x, curr.y + dir.y}
			if isInBounds(grid, neighbor) && grid[neighbor.x][neighbor.y] == char {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			} else {
				// increment perimeter if out of bounds or different character
				plot.perimeter++
			}
		}
	}
	return plot
}

func part1(plots []plot) int {
	result := 0
	for _, plot := range plots {
		result += plot.perimeter * len(plot.cells)
	}
	return result
}

func main() {
	grid := getInputLines("../data/12.txt")
	plots := findPlots(grid)

	fmt.Println(part1(plots))
}
