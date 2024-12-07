package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findCombination(numbers []int, target int, index int, currentResult int, useConcatOperator bool) bool {
	if index == len(numbers)-1 {
		return currentResult == target
	}
	nextNum := numbers[index+1]

	if currentResult > target {
		return false
	}

	if findCombination(numbers, target, index+1, currentResult+nextNum, useConcatOperator) {
		return true
	}
	if findCombination(numbers, target, index+1, currentResult*nextNum, useConcatOperator) {
		return true
	}

	// Part2
	if useConcatOperator {
		newNumStr := fmt.Sprintf("%d%d", currentResult, nextNum)
		newNum, _ := strconv.Atoi(newNumStr)
		if findCombination(numbers, target, index+1, newNum, useConcatOperator) {
			return true
		}
	}
	return false
}

func part1(input [][]int) int {
	result := 0
	for _, equation := range input {
		target := equation[0]
		numbers := equation[1:]
		if findCombination(numbers, target, 0, numbers[0], false) {
			result += target
		}
	}
	return result
}

func part2(input [][]int) int {
	result := 0
	for _, equation := range input {
		target := equation[0]
		numbers := equation[1:]
		if findCombination(numbers, target, 0, numbers[0], true) {
			result += target
		}
	}
	return result
}

func getTargetAndNumbers(line string) []int {
	chars := strings.Split(line, ": ")
	target, _ := strconv.Atoi(chars[0])
	numbers := strings.Split(chars[1], " ")
	result := []int{target}
	for _, n := range numbers {
		num, _ := strconv.Atoi(n)
		result = append(result, num)
	}
	return result
}

func main() {
	lines := getInputLines("../data/07.txt")
	var input [][]int
	for _, line := range lines {
		result := getTargetAndNumbers(line)
		input = append(input, result)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
