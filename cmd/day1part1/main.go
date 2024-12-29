package main

import (
	"aoc/pkg/common"
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {

	resp := common.LoadClient("https://adventofcode.com/2024/day/1/input")
	defer resp.Body.Close()
	var left, right []int
	scanner := bufio.NewScanner(resp.Body)
	lineno := 0
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		lineno++
		if len(nums) != 2 {
			fmt.Printf("Faulty line %d\n", lineno)
			fmt.Printf("len(nums) = %d\n", len(nums))
			continue
		}

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Error parsing left number at line %d: %v\n", lineno, err)
			continue
		}

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Error parsing right number at line %d: %v\n", lineno, err)
			continue
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Calculate total distance
	total := 0
	for i := 0; i < len(left); i++ {
		total += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println("Total distance:", total)
}
