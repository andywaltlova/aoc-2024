package main

import "fmt"

type sequence struct {
	value        [4]int
	worthBananas int
}

func nextSecret(secret int) int {
	// multiply by 64, mix, and prune
	secret ^= secret * 64
	secret %= 16777216

	// divide by 32, round down, mix, and prune
	secret ^= secret / 32
	secret %= 16777216

	// multiply by 2048, mix, and prune
	secret ^= secret * 2048
	secret %= 16777216

	return secret
}

func simulateNthSecret(initialSecrets []int, nth int) int {
	total := 0

	for _, secret := range initialSecrets {
		for i := 0; i < nth; i++ {
			secret = nextSecret(secret)
		}
		total += secret
	}
	return total
}

func chunkChangesToSequences(sequences map[[4]int]sequence, changes []int, prices []int, buyerID int) {
	chunkSize := 4
	seenSequences := make(map[[4]int]struct{})

	for i := 0; i <= len(changes)-chunkSize; i++ {
		chunk := changes[i : i+chunkSize]
		price := prices[i+chunkSize]
		sequenceValue := [4]int{chunk[0], chunk[1], chunk[2], chunk[3]}
		if _, ok := seenSequences[sequenceValue]; !ok {
			// Sequence is seen first time for this buyer

			seq, ok := sequences[sequenceValue]
			if !ok {
				// Sequence is seen first time overall
				seq = sequence{
					value:        sequenceValue,
					worthBananas: price,
				}
			} else {
				seq.worthBananas += price
			}
			sequences[sequenceValue] = seq
			seenSequences[sequenceValue] = struct{}{}
		}
	}
}

func getPriceChanges(prices []int) []int {
	changes := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		changes[i-1] = prices[i] - prices[i-1]
	}
	return changes
}

func getBestSequence(buyers []int, n int) int {
	bestBananaCount := 0
	sequences := make(map[[4]int]sequence)
	for _, b := range buyers {
		var prices []int
		secret := b
		prices = append(prices, secret%10)
		for i := 0; i < n; i++ {
			secret = nextSecret(secret)
			prices = append(prices, secret%10)
		}
		changes := getPriceChanges(prices)
		chunkChangesToSequences(sequences, changes, prices, b)
	}
	// get the best sequence
	for _, seq := range sequences {
		if seq.worthBananas > bestBananaCount {
			bestBananaCount = seq.worthBananas
		}
	}
	return bestBananaCount
}

func main() {
	buyers := getNumberInput("../data/22.txt")
	fmt.Println(simulateNthSecret(buyers, 2000))
	fmt.Println(getBestSequence(buyers, 2000))
}
