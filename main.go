package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Input struct {
	LowerBound  int         `json:"lowerBound"`
	UpperBound  int         `json:"upperBound"`
	Proportions [][]float64 `json:"proportions"`
}

func main() {
	file, err := os.Open("input.json")
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %s", err.Error()))
	}

	var input Input
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&input); err != nil {
		panic(fmt.Sprintf("Error decoding JSON: %s", err.Error()))
	}

	answers, err := GuessNumbers(
		input.LowerBound,
		input.UpperBound,
		input.Proportions,
	)
	if err != nil {
		panic(fmt.Sprintf("Error guessing numbers: %s", err.Error()))
	}

	fmt.Println("Possible numbers:", answers)
}
