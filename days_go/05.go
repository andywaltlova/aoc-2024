package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func ContainsInt(slice []int, value int) bool {
	// NOTE: would be better to use map for faster lookup
	// but I need to keep order of ints in update
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func loadRulesAndUpdates(lines []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.ContainsAny(line, "|") {
			nums := strings.Split(line, "|")
			numKey, _ := strconv.Atoi(nums[0])
			numVal, _ := strconv.Atoi(nums[1])
			val, ok := rules[numKey]
			if !ok {
				rules[numKey] = []int{numVal}
			} else {
				rules[numKey] = append(val, numVal)
			}
		} else {
			numsStr := strings.Split(line, ",")
			var numsInt []int
			for _, num := range numsStr {
				numInt, _ := strconv.Atoi(num)
				numsInt = append(numsInt, numInt)
			}
			updates = append(updates, numsInt)
		}

	}
	return rules, updates
}

func validateUpdate(rules map[int][]int, update []int) bool {
	seen := make(map[int]int)

	for i := len(update) - 1; i >= 0; i-- {
		num := update[i]
		previousNumbers, rulesExists := rules[num]
		if !rulesExists {
			seen[num] = i
			continue
		}
		// For every number that should be seen, check that if it is in update we saw it
		for _, ruleNum := range previousNumbers {
			_, numInSeen := seen[ruleNum]
			if ContainsInt(update, ruleNum) && !numInSeen {
				// Number in rule should be already seen
				return false
			}
		}
		seen[num] = i
	}
	return true
}

func calculateInDegree(graph map[int][]int) map[int]int {
	// Compute incoming edges to the node = in degree
	inDegree := make(map[int]int)

	for node := range graph {
		if _, exists := inDegree[node]; !exists {
			inDegree[node] = 0
		}
		for _, neighbor := range graph[node] {
			inDegree[neighbor]++
		}
	}
	return inDegree
}

func initializeQueue(inDegree map[int]int) *list.List {
	// Start the topological sort with nodes that have zero dependencies
	queue := list.New()
	for node, degree := range inDegree {
		if degree == 0 {
			queue.PushBack(node)
		}
	}
	return queue
}

func topologicalSort(graph map[int][]int) []int {
	inDegree := calculateInDegree(graph)
	queue := initializeQueue(inDegree)

	result := []int{}
	for queue.Len() > 0 {
		front := queue.Front()
		node := front.Value.(int)
		queue.Remove(front)
		result = append(result, node)

		// Process neighbors and reduce their in degree
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue.PushBack(neighbor)
			}
		}
	}

	return result
}

func getOrderedUpdate(orderedRules []int, update []int) []int {
	// FIXME: It would be faster to have map of numbers in update, but sufficient for solution
	var newUpdate []int
	for _, num := range orderedRules {
		if ContainsInt(update, num) {
			newUpdate = append(newUpdate, num)
		}
	}
	// Now we could have numbers in original update(row) that are not in the rules at all
	// So we should insert them at the original positions because they can't break already ordered array
	for i, num := range update {
		if !ContainsInt(newUpdate, num) {
			newUpdate = append(newUpdate[:i], append([]int{num}, newUpdate[i:]...)...)
		}
	}
	return newUpdate
}

func solve(rules map[int][]int, updates [][]int) (int, int) {
	part1 := 0
	part2 := 0
	for _, update := range updates {
		if validateUpdate(rules, update) {
			part1 += update[len(update)/2]
		} else {
			// Ah for part2 my naive approach won't work :( let's do the work and create topological sort
			subGraph := make(map[int][]int)
			for _, num := range update {
				val, _ := rules[num]
				subGraph[num] = val
			}
			orderedRuleInts := topologicalSort(subGraph)
			newUpdate := getOrderedUpdate(orderedRuleInts, update)
			part2 += newUpdate[len(newUpdate)/2]
		}
	}
	return part1, part2
}

func main() {
	lines := getInputLines("../data/05.txt")
	rules, updates := loadRulesAndUpdates(lines)
	fmt.Println(solve(rules, updates))

}
