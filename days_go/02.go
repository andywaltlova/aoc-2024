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

func checkLevelWithDampener(level []int) bool {
	for i := 0; i < len(level); i++ {
		// Make copy without current number
		newLevel := make([]int, len(level))
		copy(newLevel, level)

		if i == len(level)-1 {
			newLevel = newLevel[:i]
		} else {
			newLevel = append(newLevel[:i], newLevel[i+1:]...)
		}
		if isLevelSafe(newLevel) {
			return true
		}
	}
	return false
}

func main() {
	levels := readLevels("../data/02.txt")
	part1, part2 := 0, 0
	for _, level := range levels {
		if isLevelSafe(level) {
			part1++
			part2++
		} else if checkLevelWithDampener(level) {
			// Not best solution, thre is probably better solution checking just problematic numbers
			part2++
		}
	}
	fmt.Println(part1, part2)
}
