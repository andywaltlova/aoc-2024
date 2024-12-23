package main

import (
	"fmt"
	"sort"
	"strings"
)

type computer struct {
	id    string
	links map[string]*computer
}

func (c *computer) addLink(computer *computer) {
	c.links[computer.id] = computer
}

func parseInput(filename string) map[string]*computer {
	lines := getInputLines(filename)
	computers := make(map[string]*computer)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		computerA, computerB := parts[0], parts[1]
		cA, exists := computers[computerA]
		if !exists {
			cA = &computer{id: computerA, links: make(map[string]*computer)}
			computers[computerA] = cA
		}
		cB, exists := computers[computerB]
		if !exists {
			cB = &computer{id: computerB, links: make(map[string]*computer)}
			computers[computerB] = cB
		}
		cB.addLink(cA)
		cA.addLink(cB)
	}
	return computers
}

func findThreeInterconnected(network map[string]*computer) map[string]bool {
	uniqueTriplets := make(map[string]bool)

	for _, comp := range network {
		for neighborA := range comp.links {
			for neighborB := range comp.links {
				if neighborA >= neighborB {
					continue // duplicates
				}
				if _, ok := network[neighborA].links[neighborB]; ok {
					ids := []string{comp.id, neighborA, neighborB}
					sort.Strings(ids)
					triplet := strings.Join(ids, "-")
					uniqueTriplets[triplet] = true
				}
			}
		}
	}

	filteredTriplets := make(map[string]bool)
	for triplet := range uniqueTriplets {
		parts := strings.Split(triplet, "-")
		for _, part := range parts {
			if part[0] == 't' {
				filteredTriplets[triplet] = true
				break
			}
		}
	}
	return filteredTriplets
}

func bronKerbosch(currentClique, candidates, processed map[string]bool, network map[string]*computer, cliques *[][]string) {
	if len(candidates) == 0 && len(processed) == 0 {
		// nothing left to process
		clique := []string{}
		for node := range currentClique {
			clique = append(clique, node)
		}
		*cliques = append(*cliques, clique)
		return
	}

	for c := range candidates {
		neighbors := network[c].links

		newClique := copySet(currentClique)
		newClique[c] = true
		newCandidates := intersectSets(candidates, neighbors)
		newProcessed := intersectSets(processed, neighbors)

		bronKerbosch(newClique, newCandidates, newProcessed, network, cliques)

		delete(candidates, c)
		processed[c] = true
	}
}

func copySet(set map[string]bool) map[string]bool {
	newSet := make(map[string]bool)
	for k := range set {
		newSet[k] = true
	}
	return newSet
}

func intersectSets(set1 map[string]bool, neighbors map[string]*computer) map[string]bool {
	intersection := make(map[string]bool)
	for k := range set1 {
		if _, exists := neighbors[k]; exists {
			intersection[k] = true
		}
	}
	return intersection
}

func findLargestCliqueBronKerbosch(network map[string]*computer) string {
	// I had to look up the Bron-Kerbosch algorithm to solve this problem, probably my favourite algorithm so far this year
	clique := make(map[string]bool)
	candidates := make(map[string]bool)
	processed := make(map[string]bool)
	for id := range network {
		candidates[id] = true
	}

	cliques := [][]string{}
	bronKerbosch(clique, candidates, processed, network, &cliques)

	// Find the largest clique
	var largestClique []string
	for _, clique := range cliques {
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}
	sort.Strings(largestClique)
	if len(largestClique) == 0 {
		return "No clique found"
	}
	return strings.Join(largestClique, ",")
}

func main() {
	network := parseInput("../data/23.txt")
	fmt.Println(len(findThreeInterconnected(network)))
	fmt.Println(findLargestCliqueBronKerbosch(network))
}
