package main

import (
	"fmt"
	"math"
)

// TODO: Currently is for each proportion to find n, but we can for each n to find proportion, if any proportion not fit than pass
func GuessNumbers(
	lowerBound int,
	upperBound int,
	proportions [][]float64,
) ([]int, error) {
	var proportionPossibleResults [][]int
	for _, proportion := range proportions {
		if !isValidProportion(proportion) {
			return nil, fmt.Errorf("invalid proportion: %v", proportion)
		}

		proportionPossibleResults = append(
			proportionPossibleResults,
			findPossibleResults(lowerBound, upperBound, proportion),
		)
	}

	return intersectResults(proportionPossibleResults), nil
}

func isValidProportion(proportion []float64) bool {
	var sum float64
	for _, v := range proportion {
		sum += v
	}
	return sum >= 99 && sum <= 101
}

func HasCompositionOfN(n int, propMap map[float64]int) bool {
	// We only need to check whether each v is near round(v)
	for prob := range propMap {
		v := float64(prob) * float64(n) / 100
		roundV := math.Round(v)
		if math.Abs(v-roundV) > 0.1 {
			return false
		}
	}

	return true
}

func findPossibleResults(lowerBound int, upperBound int, proportion []float64) []int {
	propMap := make(map[float64]int, len(proportion))
	for _, v := range proportion {
		if _, ok := propMap[v]; !ok {
			propMap[v] = 0
		}
		propMap[v]++
	}

	var possibleResults []int
	for n := lowerBound; n <= upperBound; n++ {
		if HasCompositionOfN(n, propMap) {
			possibleResults = append(possibleResults, n)
		}
	}
	return possibleResults
}

func intersectResults(proportionPossibleResults [][]int) []int {
	if len(proportionPossibleResults) == 0 {
		return nil
	}

	countMap := make(map[int]int)

	for _, possibleResults := range proportionPossibleResults {
		for _, possibleResult := range possibleResults {
			countMap[possibleResult]++
		}
	}

	// Find possible result that match all proportions
	var intersection []int
	for k, c := range countMap {
		if c == len(proportionPossibleResults) {
			intersection = append(intersection, k)
		}
	}

	return intersection
}
