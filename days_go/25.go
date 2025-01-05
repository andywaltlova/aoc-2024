package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lock struct {
	heights []int
}

type Key struct {
	heights []int
}

func (lock Lock) fits(key Key, totalHeight int) bool {
	for i := 0; i < len(lock.heights); i++ {
		if lock.heights[i]+key.heights[i] > totalHeight {
			return false
		}
	}
	return true
}

func parseHeights(schematic []string, isLock bool) []int {
	columns := len(schematic[0])
	heights := make([]int, columns)

	for col := 0; col < columns; col++ {
		height := 0
		for row := 0; row < len(schematic); row++ {
			if isLock {
				if schematic[row][col] == '#' {
					height++
				} else {
					break
				}
			} else {
				if schematic[len(schematic)-1-row][col] == '#' {
					height++
				} else {
					break
				}
			}
		}
		heights[col] = height
	}
	return heights
}

func readLocksAndKeys(filename string) ([]Lock, []Key, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var locks []Lock
	var keys []Key
	var currentSchematic []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(currentSchematic) > 0 {
				isLock := strings.Contains(currentSchematic[0], "#")
				if isLock {
					locks = append(locks, Lock{heights: parseHeights(currentSchematic[1:], true)})
				} else {
					keys = append(keys, Key{heights: parseHeights(currentSchematic[:len(currentSchematic)-1], false)})
				}
				currentSchematic = []string{}
			}
		} else {
			currentSchematic = append(currentSchematic, line)
		}
	}

	if len(currentSchematic) > 0 {
		isLock := strings.Contains(currentSchematic[0], "#")
		if isLock {
			locks = append(locks, Lock{heights: parseHeights(currentSchematic[1:], true)})
		} else {
			keys = append(keys, Key{heights: parseHeights(currentSchematic[:len(currentSchematic)-1], false)})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return locks, keys, nil
}

func main() {
	filename := "../data/25.txt"

	locks, keys, err := readLocksAndKeys(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	totalHeight := 5
	count := 0

	for _, lock := range locks {
		for _, key := range keys {
			if lock.fits(key, totalHeight) {
				count++
			}
		}
	}

	fmt.Printf("Number of unique lock/key pairs that fit together: %d\n", count)
}
