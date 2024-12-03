package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readFileContent(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read file")
	}
	return string(data)
}

func matchMulPattern(content string) []string {
	pattern := "mul\\(\\d+,\\d+\\)"
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(content, -1)
}

func matchEnabledMulPattern(content string) []string {
	pattern := "mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)"
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(content, -1)
}

func getNumsFromMatch(match string) (int, int) {
	r, _ := regexp.Compile("\\d+")
	nums := r.FindAllString(match, -1)
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return num1, num2
}

func part1(content string) int {
	result := 0
	matches := matchMulPattern(content)
	for _, match := range matches {
		num1, num2 := getNumsFromMatch(match)
		result += num1 * num2
	}
	return result
}

func part2(content string) int {
	result := 0
	enabled := true
	matches := matchEnabledMulPattern(content)
	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			num1, num2 := getNumsFromMatch(match)
			result += num1 * num2
		}
	}
	return result
}

func main() {
	filepath := "../data/03_test.txt"
	content := readFileContent(filepath)
	fmt.Println(part1(content))
	fmt.Println(part2(content))
}
