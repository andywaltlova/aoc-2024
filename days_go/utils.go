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

// Set implementation using map[T]struct{} for memory efficiency
type Set[T comparable] struct {
	elements map[T]struct{}
}

// Create a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// Add an element to the set
func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// Remove an element from the set
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Check if the set contains an element
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Get the size of the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Convert the set to a slice
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for key := range s.elements {
		slice = append(slice, key)
	}
	return slice
}

// Union: combines two sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

// Intersection: returns common elements between two sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

// Difference: returns elements in the current set but not in the other set
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

// Symmetric Difference: returns elements in either set, but not in both
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	for key := range other.elements {
		if !s.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
