package main

import (
	"aoc/pkg/common"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	resp := common.LoadClient("https://adventofcode.com/2024/day/2/input")

	safeLevels := 0
	scanner := bufio.NewScanner(resp.Body)
	lineno := 0

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		lineno++

		levels := make([]int, 0)
		for _, num := range nums {
			level_value, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error converting string to number:", err)
				return
			}
			levels = append(levels, level_value)
		}

		if isValidSequence(levels) {
			safeLevels++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
	}

	fmt.Println("Total safe levels:", safeLevels)

}

func isValidSequence(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Determine if sequence is increasing or decreasing based on first pair
	isIncreasing := levels[1] > levels[0]

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]

		// Check if difference is between 1 and 3
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}

		// For increasing sequence, all differences must be positive
		if isIncreasing && diff <= 0 {
			return false
		}

		// For decreasing sequence, all differences must be negative
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}
