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

	for i := range grid[1:] {
		for j := range grid[i][1:] {

			// Check forward
			if checkXMAS(grid, i, j) {
				count++
			}
			// // Check backward
			// if checkXMAS(grid, i, j, -dir[0], -dir[1]) {
			// 	count++
			// }
		}
	}

	fmt.Printf("Total count: %d\n", count)
}

func checkXMAS(grid [][]rune, row, col int) bool {
	rows, cols := len(grid), len(grid[0])
	if grid[row][col] != 'A' {
		return false
	}

	if row-1 < 0 || row+1 >= rows || col-1 < 0 || col+1 >= cols {
		return false
	}

	if grid[row-1][col-1] == 'M' && grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M' && grid[row+1][col+1] == 'S' {
		return true
	} else if grid[row-1][col-1] == 'S' && grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S' && grid[row+1][col+1] == 'M' {
		return true
	} else if grid[row-1][col-1] == 'M' && grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S' && grid[row+1][col+1] == 'S' {
		return true
	} else if grid[row-1][col-1] == 'S' && grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M' && grid[row+1][col+1] == 'M' {
		return true
	} else {
		return false
	}

}
