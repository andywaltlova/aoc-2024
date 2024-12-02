package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func readNumberPairs(filename string) ([]int, []int) {
	var left []int
	var right []int
	lines := getInputLines(filename)
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(numbers[0])
		righNum, _ := strconv.Atoi(numbers[1])
		left = append(left, leftNum)
		right = append(right, righNum)
	}
	slices.Sort(left)
	slices.Sort(right)
	return left, right
}

func getDistance(left []int, right []int) int {
	var distance int
	for i := 0; i < len(left); i++ {
		distance += absDiffInt(left[i], right[i])
	}
	return distance
}

func getOccurencesDict(numbers []int) map[int]int {
	occurences := make(map[int]int)
	for _, num := range numbers {
		occurences[num]++
	}
	return occurences
}

func getSimilarityScore(left []int, occurences map[int]int) int {
	var score int
	for _, num := range left {
		numOccurences, exists := occurences[num]
		if !exists {
			continue
		}
		score += num * numOccurences
	}
	return score
}

func main() {
	left, right := readNumberPairs("../data/01.txt")
	part1 := getDistance(left, right)
	fmt.Println(part1)

	// Part 2
	occurences := getOccurencesDict(left)
	part2 := getSimilarityScore(right, occurences)
	fmt.Println(part2)

}
