package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getNumsOnLine(filename string) []int {
	var num []int
	for _, s := range getInputLines(filename) {
		for _, c := range s {
			num = append(num, int(c-'0'))
		}
	}
	return num
}

func getNumberInput(filename string) []int {
	var result []int
	for _, s := range getInputLines(filename) {
		num, _ := strconv.Atoi(s)
		result = append(result, num)
	}
	return result
}

func getInputLines(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()
	return result
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
