package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func splitStone(stone int) (int, int) {
	numDigits := int(math.Log10(float64(stone)) + 1)

	mid := numDigits / 2
	divisor := int(math.Pow10(mid))

	left := stone / divisor
	right := stone % divisor

	return left, right
}

func blink(stones []int) []int {
	var newStones []int
	for _, stone := range stones {
		numDigits := int(math.Log10(float64(stone)) + 1)
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if numDigits%2 == 0 {
			left, right := splitStone(stone)
			newStones = append(newStones, left, right)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}

func part1(nums []int, n int) int {
	result := nums
	for i := 0; i < n; i++ {
		result = blink(result)
	}
	return len(result)
}

func main() {
	numsLine := getInputLines("../data/11.txt")
	var nums []int
	for _, s := range strings.Split(numsLine[0], " ") {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	fmt.Println(part1(nums, 25))
	// Well .. lets change the approach before my memory explodes

}
