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

func partOne(patterns []string, designs []string) int {
	result := 0
	cache := make(map[string]bool)
	for _, design := range designs {
		if isPossible(design, patterns, cache) {
			result++
		}
	}
	return result
}

func main() {
	patterns, designs, _ := parseInput("../data/19.txt")
	possible := partOne(patterns, designs)
	fmt.Println(possible)
}
