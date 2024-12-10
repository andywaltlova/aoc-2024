package main

import (
	"fmt"
	"strconv"
)

type coordinate struct {
	x, y int
}

func getTopologyMap(grid []string) (map[coordinate]int, int, int) {
	gridMap := make(map[coordinate]int)
	for x, line := range grid {
		for y, char := range line {
			val, _ := strconv.Atoi(string(char))
			gridMap[coordinate{x: x, y: y}] = val
		}
	}
	return gridMap, len(grid), len(grid[0])
}

func getNineHeightPositions(
	trailhead coordinate,
	topologyMap map[coordinate]int,
	node coordinate,
	target int,
	trailheadScore map[coordinate]map[coordinate]int) {
	val, ok := topologyMap[node]
	if !ok {
		// Out of bounds
		return
	}
	if val == target {
		if target == 9 {
			val, ok := trailheadScore[trailhead][node]
			if !ok {
				trailheadScore[trailhead][node] = 1
			} else {
				trailheadScore[trailhead][node] = val + 1
			}
		}
		up := coordinate{x: node.x - 1, y: node.y}
		right := coordinate{x: node.x, y: node.y + 1}
		left := coordinate{x: node.x, y: node.y - 1}
		down := coordinate{x: node.x + 1, y: node.y}

		getNineHeightPositions(trailhead, topologyMap, up, target+1, trailheadScore)
		getNineHeightPositions(trailhead, topologyMap, right, target+1, trailheadScore)
		getNineHeightPositions(trailhead, topologyMap, left, target+1, trailheadScore)
		getNineHeightPositions(trailhead, topologyMap, down, target+1, trailheadScore)

	} else {
		// Not valid path
		return
	}
	return
}

func findPaths(topologyMap map[coordinate]int, maxX, maxY int) (int, int) {
	part1, part2 := 0, 0
	nineHeightPositions := make(map[coordinate]map[coordinate]int)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			trailhead := coordinate{x: x, y: y}
			nineHeightPositions[trailhead] = make(map[coordinate]int)

			getNineHeightPositions(trailhead, topologyMap, trailhead, 0, nineHeightPositions)

			score, _ := nineHeightPositions[trailhead]
			if len(score) >= 1 {
				fmt.Println(trailhead, nineHeightPositions[trailhead])
			}
			part1 += len(score)
			rating := 0
			for _, val := range score {
				rating += val
			}
			part2 += rating
		}
	}
	return part1, part2
}

func main() {
	grid := getInputLines("../data/10.txt")
	topologyMap, maxX, maxY := getTopologyMap(grid)
	fmt.Println(findPaths(topologyMap, maxX, maxY))
}
