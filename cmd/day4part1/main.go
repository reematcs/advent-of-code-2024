package main

import (
	"aoc/pkg/common"
	"bufio"
	"fmt"
)

func main() {
	resp := common.LoadClient("https://adventofcode.com/2024/day/4/input")
	scanner := bufio.NewScanner(resp.Body)
	//scanningstring := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	// Convert input to 2D grid
	var grid [][]rune
	//scanner := bufio.NewScanner(strings.NewReader(scanningstring))
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
	}

	count := 0
	dirs := [][2]int{
		{0, 1},  // right
		{1, 0},  // down
		{1, 1},  // diagonal down-right
		{1, -1}, // diagonal down-left
	}

	for i := range grid {
		for j := range grid[i] {
			for _, dir := range dirs {
				// Check forward
				if checkXMAS(grid, i, j, dir[0], dir[1]) {
					count++
				}
				// Check backward
				if checkXMAS(grid, i, j, -dir[0], -dir[1]) {
					count++
				}
			}
		}
	}

	fmt.Printf("Total count: %d\n", count)
}

func checkXMAS(grid [][]rune, row, col, dx, dy int) bool {
	target := "XMAS"
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < len(target); i++ {
		newRow := row + i*dx
		newCol := col + i*dy
		// Check if the pattern would go out of bounds
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}
		/* Check if the rune (int32) conversion of target character at index i in "XMAS"
		   doesn't equal the character at index+directional shift in grid.
		*/
		if grid[newRow][newCol] != rune(target[i]) {
			return false
		}
	}
	return true
}
