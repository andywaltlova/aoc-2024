package main

import "fmt"

var directions = [][]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{-1, 0},  // Up
	{1, 0},   // Down
	{1, 1},   // Diagonal down-right
	{1, -1},  // Diagonal down-left
	{-1, 1},  // Diagonal up-right
	{-1, -1}, // Diagonal up-left
}

var target = "XMAS"

func part1(grid []string) int {
	count := 0
	for x, line := range grid {
		for y := 0; y < len(line); y++ {
			if string(line[y]) == "X" {
				for _, dir := range directions {
					count += move(grid, x+dir[0], y+dir[1], dir, 1, target)
				}
			}
		}
	}
	return count
}

func move(grid []string, x int, y int, dir []int, i int, target string) int {
	// NOTE: having index was enough for part1, added target to reuse function for part2
	if x < 0 || y < 0 || x == len(grid[0]) || y == len(grid[0]) || grid[x][y] != target[i] {
		return 0
	}
	if i == len(target)-1 {
		return 1
	}
	return move(grid, x+dir[0], y+dir[1], dir, i+1, target)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func part2(grid []string) int {
	count := 0
	for x, line := range grid {
		for y := 0; y < len(line); y++ {
			if string(line[y]) == "A" {
				downRight := directions[4]
				downLeft := directions[5]
				upRight := directions[6]
				upLeft := directions[7]

				for _, c := range []string{"SMSM", "SSMM", "MMSS", "MSMS"} {
					uR := move(grid, x+upRight[0], y+upRight[1], upRight, 0, string(c[0]))
					uL := move(grid, x+upLeft[0], y+upLeft[1], upLeft, 0, string(c[1]))
					dR := move(grid, x+downRight[0], y+downRight[1], downRight, 0, string(c[2]))
					dL := move(grid, x+downLeft[0], y+downLeft[1], downLeft, 0, string(c[3]))

					if sum([]int{dR, uR, dL, uL}) == 4 {
						count += 1
						break
					}
				}

			}
		}
	}
	return count
}

func main() {
	grid := getInputLines("../data/04.txt")
	fmt.Println(part1(grid))
	fmt.Println(part2(grid))

}
