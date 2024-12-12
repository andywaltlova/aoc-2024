package main

import (
	"fmt"
)

type coordinate struct {
	x, y int
}

type plot struct {
	cells     []coordinate
	sides     int
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

func countSides(plot *plot) int {
	// Oh wow, this took me a while, but the trick might be counting the corners
	plotCellMap := make(map[coordinate]bool)
	for _, cell := range plot.cells {
		plotCellMap[cell] = true
	}

	s := 0
	for _, cell := range plot.cells {
		up := coordinate{cell.x, cell.y - 1}
		down := coordinate{cell.x, cell.y + 1}
		left := coordinate{cell.x - 1, cell.y}
		right := coordinate{cell.x + 1, cell.y}

		if !plotCellMap[up] && !plotCellMap[left] {
			s++
		}
		if !plotCellMap[up] && !plotCellMap[right] {
			s++
		}
		if !plotCellMap[down] && !plotCellMap[left] {
			s++
		}
		if !plotCellMap[down] && !plotCellMap[right] {
			s++
		}
		// Inner corners
		diagonalDownRight := coordinate{cell.x + 1, cell.y + 1}
		diagonalDownLeft := coordinate{cell.x - 1, cell.y + 1}
		diagonalUpRight := coordinate{cell.x + 1, cell.y - 1}
		diagonalUpLeft := coordinate{cell.x - 1, cell.y - 1}

		if !plotCellMap[diagonalUpLeft] && plotCellMap[left] && plotCellMap[up] {
			s++
		}
		if !plotCellMap[diagonalUpRight] && plotCellMap[right] && plotCellMap[up] {
			s++
		}
		if !plotCellMap[diagonalDownLeft] && plotCellMap[left] && plotCellMap[down] {
			s++
		}
		if !plotCellMap[diagonalDownRight] && plotCellMap[right] && plotCellMap[down] {
			s++
		}
	}
	plot.sides = s
	return s
}

func part2(plots []plot) int {
	result := 0
	for _, plot := range plots {
		countSides(&plot)
		result += plot.sides * len(plot.cells)
	}
	return result
}

func main() {
	grid := getInputLines("../data/12.txt")
	plots := findPlots(grid)

	fmt.Println(part1(plots))
	fmt.Println(part2(plots))

}
