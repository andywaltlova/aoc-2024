package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readReports(filename string) [][]int {
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

func isReportSafe(level []int) bool {
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

func checkLevelsWithDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		// Make copy without current number
		newReport := make([]int, len(report))
		copy(newReport, report)

		if i == len(report)-1 {
			newReport = newReport[:i]
		} else {
			newReport = append(newReport[:i], newReport[i+1:]...)
		}
		if isReportSafe(newReport) {
			return true
		}
	}
	return false
}

func main() {
	reports := readReports("../data/02.txt")
	part1, part2 := 0, 0
	for _, report := range reports {
		if isReportSafe(report) {
			part1++
			part2++
			fmt.Println(report)
		} else if checkLevelsWithDampener(report) {
			// Not best solution, thre is probably better solution checking just problematic numbers
			part2++
			fmt.Println(report)
		}
	}
	fmt.Println(part1, part2)
}
