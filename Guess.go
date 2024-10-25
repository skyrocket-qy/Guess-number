package main

import (
	"fmt"
	"math"

	"github.com/mohae/deepcopy"
)

func GuessNumbers(
	lowerBound int,
	upperBound int,
	proportions [][]int,
) ([]int, error) {
	var possibleResult [][]int
	for _, proportion := range proportions {
		if !checkproportion(proportion) {
			return nil, fmt.Errorf("invalid proportion: %v", proportion)
		}

		possibleResult = append(
			possibleResult,
			findPossibleResults(lowerBound, upperBound, proportion),
		)
	}

	return Intersect(possibleResult), nil
}

func checkproportion(proportion []int) bool {
	sum := 0
	for _, v := range proportion {
		sum += v
	}
	return sum >= 98 && sum <= 102
}

func findPossibleResults(lowerBound int, upperBound int, proportion []int) []int {
	propMap := make(map[int]int, len(proportion))
	for _, v := range proportion {
		if _, ok := propMap[v]; !ok {
			propMap[v] = 0
		}
		propMap[v]++
	}

	var possibleResults []int
	for n := lowerBound; n <= upperBound; n++ {
		composition := []int{}
		copyPropMap := deepcopy.Copy(propMap).(map[int]int)
		curMembers := 0
		for curMembers < n {
			find := false
			for i := 0; i < n; i++ {
				p := int(math.Round(float64(i*100) / float64(n)))
				if _, ok := copyPropMap[p]; ok {
					copyPropMap[p]--
					if v := copyPropMap[p]; v == 0 {
						delete(copyPropMap, p)
					}
					composition = append(composition, i)
					find = true
					break
				}
			}
			if !find {
				break
			}
		}
		sum := 0
		for _, v := range composition {
			sum += v
		}
		if sum == n {
			possibleResults = append(possibleResults, n)
		}
	}
	return possibleResults
}

func Intersect(proportionPossibleResults [][]int) []int {
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
	var res []int
	for k, c := range countMap {
		if c == len(proportionPossibleResults) {
			res = append(res, k)
		}
	}

	return res
}
