package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(filename string) (patterns []string, designs []string, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	parts := strings.Split(string(data), "\n\n")
	patterns = strings.Split(parts[0], ", ")
	designs = strings.Split(parts[1], "\n")
	return patterns, designs, nil
}

func isPossible(design string, patterns []string, cache map[string]bool) bool {
	b, ok := cache[design]
	if ok {
		return b
	}

	for _, p := range patterns {
		if p == design {
			// Exact match for the rest of the design
			return true
		} else if strings.HasPrefix(design, p) {
			// Cut the prefix and check if the rest is possible
			isPoss := isPossible(strings.TrimPrefix(design, p), patterns, cache)
			if isPoss {
				cache[design] = true
				return true
			}
		}
	}
	cache[design] = false
	return false
}

func countWays(design string, patterns []string, cache map[string]int) int {
	ways := 0
	way, ok := cache[design]
	if ok {
		return way
	}

	for _, p := range patterns {
		if p == design {
			ways++
		} else if strings.HasPrefix(design, p) {
			ways += countWays(strings.TrimPrefix(design, p), patterns, cache)
		}
	}
	cache[design] = ways
	return ways
}

func analyze(patterns []string, designs []string) (int, int) {
	possible := 0
	ways := 0
	cache := make(map[string]bool)
	for _, design := range designs {
		if isPossible(design, patterns, cache) {
			// or just omit isPossible and use countWays (0 = no way to make it)
			possible++
			ways += countWays(design, patterns, make(map[string]int))
		}
	}
	return possible, ways
}

func main() {
	patterns, designs, _ := parseInput("../data/19.txt")
	possible, ways := analyze(patterns, designs)
	fmt.Println(possible, ways)
}
