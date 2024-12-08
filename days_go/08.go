package main

import "fmt"

type point struct {
	x, y int
}

func difference(p1, p2 point) point {
	return point{x: p2.x - p1.x, y: p2.y - p1.y}
}

func getNewAntinodesPart1(p1, p2, diff point) (point, point) {
	newpoint1 := point{x: p1.x + 2*diff.x, y: p1.y + 2*diff.y}
	newpoint2 := point{x: p2.x - 2*diff.x, y: p2.y - 2*diff.y}
	return newpoint1, newpoint2
}

func generateCombinations(points []point) [][]point {
	var combinations [][]point
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			combinations = append(combinations, []point{points[i], points[j]})
		}
	}
	return combinations
}

func getNewAntinodesPart2(p1, p2, diff point, grid []string) []point {
	var newPoints []point
	newAddAntinode := point{x: p1.x + diff.x, y: p1.y + diff.y}
	for pointInBounds(newAddAntinode, grid) {
		newPoints = append(newPoints, newAddAntinode)
		newAddAntinode = point{x: newAddAntinode.x + diff.x, y: newAddAntinode.y + diff.y}
	}

	newSubAntinode := point{x: p1.x - diff.x, y: p1.y - diff.y}
	for pointInBounds(newSubAntinode, grid) {
		newPoints = append(newPoints, newSubAntinode)
		newSubAntinode = point{x: newSubAntinode.x - diff.x, y: newSubAntinode.y - diff.y}
	}
	return newPoints
}

func pointInBounds(p point, grid []string) bool {
	return p.x >= 0 && p.x < len(grid) && p.y >= 0 && p.y < len(grid[0])
}

func printAntinodes(antinodes map[point]struct{}, grid []string) {
	for x := range grid {
		for y := range grid[x] {
			point := point{x: x, y: y}
			if _, ok := antinodes[point]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getAntennas(grid []string) map[string][]point {
	antennaMap := make(map[string][]point)
	for x, line := range grid {
		for y, char := range line {
			if string(char) == "." {
				continue
			}
			if existingAntennas, ok := antennaMap[string(char)]; ok {
				antennaMap[string(char)] = append(existingAntennas, point{x: x, y: y})
			} else {
				antennaMap[string(char)] = []point{point{x: x, y: y}}
			}
		}
	}
	return antennaMap
}

func part1(antennas map[string][]point, grid []string) int {
	newAntinodes := make(map[point]struct{})
	for _, points := range antennas {
		for _, combination := range generateCombinations(points) {
			diff := difference(combination[0], combination[1])
			newpoint1, newpoint2 := getNewAntinodesPart1(combination[0], combination[1], diff)
			if pointInBounds(newpoint1, grid) {
				newAntinodes[newpoint1] = struct{}{}
			}
			if pointInBounds(newpoint2, grid) {
				newAntinodes[newpoint2] = struct{}{}
			}
		}
	}
	// printAntinodes(newAntinodes, grid)
	return len(newAntinodes)
}

func part2(antennas map[string][]point, grid []string) int {
	newAntinodes := make(map[point]struct{})
	for _, points := range antennas {
		for _, combination := range generateCombinations(points) {
			diff := difference(combination[0], combination[1])
			newPoints := getNewAntinodesPart2(combination[0], combination[1], diff, grid)
			for _, newPoint := range newPoints {
				newAntinodes[newPoint] = struct{}{}
			}
			// Add antennas to antinodes
			newAntinodes[combination[0]] = struct{}{}
			newAntinodes[combination[1]] = struct{}{}
		}
	}
	// printAntinodes(newAntinodes, grid)
	return len(newAntinodes)
}

func main() {
	lines := getInputLines("../data/08.txt")
	antennas := getAntennas(lines)
	fmt.Println(part1(antennas, lines))
	fmt.Println(part2(antennas, lines))

}
