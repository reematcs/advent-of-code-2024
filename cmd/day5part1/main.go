package main

import (
	"aoc/pkg/common"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Edge struct {
	From, To int
}

func main() {
	resp := common.LoadClient("https://adventofcode.com/2024/day/5/input")
	scanner := bufio.NewScanner(resp.Body)
	// sample input
	// scanningstring := "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47"

	// scanner := bufio.NewScanner(strings.NewReader(scanningstring))

	adjacencyList := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // Empty line separates sections
			break
		}

		parts := strings.Split(line, "|")
		dependency, _ := strconv.Atoi(parts[0])
		dependsOn, _ := strconv.Atoi(parts[1])

		adjacencyList[dependsOn] = append(adjacencyList[dependsOn], dependency)
	}
	// debug
	// fmt.Println("Processed adjacency list")
	// debug
	// fmt.Printf("%v\n", adjacencyList)
	var sum_mid int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		sequence := []int{}
		for _, part := range parts {
			element, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting ASCII sequence element to int", err)
			}
			sequence = append(sequence, element)
		}

		if validateSequence(adjacencyList, sequence) {
			// debug
			// fmt.Printf("Valid sequence:%v\n", sequence)
			sum_mid += sequence[len(sequence)>>1]
		} /* else {
			fmt.Printf("Invalid sequence:%v\n", sequence)
		} */

	}
	fmt.Println("Sum of middle page", sum_mid)
}
func validateSequence(adjacencyList map[int][]int, sequence []int) bool {
	pos := make(map[int]int)
	for i, node := range sequence {
		pos[node] = i
	}
	// debug
	// fmt.Printf("sequence positions: %v\n", pos)
	// For each node, check only its actual dependencies
	for _, node := range sequence {
		// debug
		// fmt.Printf("Checking validity of node %d\n", node)
		// fmt.Println(adjacencyList[node])
		for _, dep := range adjacencyList[node] {
			if pos[node] < pos[dep] {
				// debug
				// fmt.Printf("pos[%d] = %d < pos[%d] = %d\n", node, pos[node], dep, pos[dep])
				return false
			}
		}
	}
	return true
}
