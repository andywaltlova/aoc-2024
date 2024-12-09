package main

import "fmt"

func createDiskMap(nums []int) []int {
	// This is memory inefficient, but it's easier to understand
	// It could probably be replaced by computing the result directly

	var numsWithSpaces []int
	fileID := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if i%2 == 0 || i == 0 {
			// File
			for j := 0; j < num; j++ {
				numsWithSpaces = append(numsWithSpaces, fileID)
			}
			fileID++
		} else {
			if i == len(nums)-1 {
				break
			}
			// Free space
			for j := 0; j < num; j++ {
				numsWithSpaces = append(numsWithSpaces, -1)
			}
		}
	}
	return numsWithSpaces
}

func part1(nums []int) int {
	result := 0

	numsWithSpaces := createDiskMap(nums)

	// Fill free space from start with fileIDs from the end
	endIndex := len(numsWithSpaces) - 1
	for i := 0; i < len(numsWithSpaces); i++ {
		for numsWithSpaces[endIndex] == -1 {
			endIndex--
		}
		if numsWithSpaces[i] == -1 {
			numsWithSpaces[i] = numsWithSpaces[endIndex]
			if i < endIndex {
				endIndex--
			}
		}
	}

	for i, num := range numsWithSpaces[:endIndex+1] {
		result += i * num
	}
	return result
}

func part2(nums []int) int {
	result := 0
	numsWithSpaces := createDiskMap(nums)

	fileIDToMove := numsWithSpaces[len(numsWithSpaces)-1]
	for fileIDToMove > 0 {
		fileIDIndex := len(numsWithSpaces) - 1
		for numsWithSpaces[fileIDIndex] != fileIDToMove {
			fileIDIndex--
		}
		size := 0
		for numsWithSpaces[fileIDIndex] == fileIDToMove {
			size++
			fileIDIndex--
		}

		freeSpaceIndex := 0
		for numsWithSpaces[freeSpaceIndex] != -1 {
			freeSpaceIndex++
		}

		capacity := 0
		for freeSpaceIndex < len(numsWithSpaces) {
			if freeSpaceIndex > fileIDIndex {
				break
			}
			if numsWithSpaces[freeSpaceIndex] == -1 {
				capacity++
				if capacity >= size {
					// Move file
					for i := 0; i < size; i++ {
						numsWithSpaces[freeSpaceIndex-i] = fileIDToMove
						numsWithSpaces[fileIDIndex+1+i] = -1
					}
					break
				}
			} else {
				capacity = 0
			}
			freeSpaceIndex++
		}
		fileIDToMove -= 1
	}
	for i, num := range numsWithSpaces {
		if num == -1 {
			continue
		}
		result += i * num
	}
	return result
}

func main() {
	nums := getNumsOnLine("../data/09.txt")

	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}
