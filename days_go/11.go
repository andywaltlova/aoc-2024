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

func mapBlink(stones map[int]int) map[int]int {
	newStones := map[int]int{}

	add := func(key, incr int) {
		if _, ok := newStones[key]; !ok {
			newStones[key] = 0
		}
		newStones[key] += incr
	}

	for stone, count := range stones {
		if stone == 0 {
			add(1, count)
		} else if digits := int(math.Log10(float64(stone)) + 1); digits%2 == 0 {
			left, right := splitStone(stone)
			add(left, count)
			add(right, count)
		} else {
			add(stone*2024, count)
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

func part2(nums []int, n int) int {
	cache := map[int]int{}

	for _, v := range nums {
		cache[v] = 1
	}

	for range n {
		cache = mapBlink(cache)
	}

	sum := 0
	for _, v := range cache {
		sum += v
	}
	return sum
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
	fmt.Println(part2(nums, 75))

}
