package main

import (
	"fmt"
	"math"
)

func GuessNumbers(lowerBound, upperBound int, proportions [][]float64) ([]int, error) {
	if lowerBound > upperBound {
		return nil, fmt.Errorf(
			"invalid bounds: lowerBound (%d) > upperBound (%d)",
			lowerBound,
			upperBound,
		)
	}

	propMaps := make([]map[float64]int, len(proportions))
	for i, proportion := range proportions {
		if !isValidProportion(proportion) {
			return nil, fmt.Errorf("invalid proportion: %v", proportion)
		}
		propMaps[i] = newProportionMap(proportion)
	}

	var possibleResults []int
	for n := lowerBound; n <= upperBound; n++ {
		if meetsAllProportions(n, propMaps) {
			possibleResults = append(possibleResults, n)
		}
	}

	return possibleResults, nil
}

// Helper to create a map from a proportion slice
func newProportionMap(proportion []float64) map[float64]int {
	propMap := make(map[float64]int)
	for _, v := range proportion {
		propMap[v]++
	}
	return propMap
}

// Helper to check if a number satisfies all proportion maps
func meetsAllProportions(n int, propMaps []map[float64]int) bool {
	for _, propMap := range propMaps {
		if !HasCompositionOfN(n, propMap) {
			return false
		}
	}
	return true
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
