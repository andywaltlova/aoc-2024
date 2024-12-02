package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readLevels(filename string) [][]int {
	var levels [][]int
	lines := getInputLines(filename)
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var level []int
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			level = append(level, num)
		}
		levels = append(levels, level)
	}
	return levels
}

func isLevelSafe(level []int) bool {
	isAscending := level[0] < level[1]
	for i, num := range level {
		if i == len(level)-1 {
			break
		}
		nextNum := level[i+1]
		if (!isAscending && nextNum >= num) || (isAscending && nextNum <= num) || absDiffInt(num, nextNum) > 3 {
			return false
		}
	}
	return true
}

func main() {
	levels := readLevels("../data/02.txt")
	part1 := 0
	for _, level := range levels {
		if isLevelSafe(level) {
			part1++
		}
	}
	fmt.Println(part1)
}
