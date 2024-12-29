package main

import (
	"aoc/pkg/common"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	resp := common.LoadClient("https://adventofcode.com/2024/day/3/input")
	scanner := bufio.NewScanner(resp.Body)
	pattern := regexp.MustCompile(`mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)`)
	mul_sum := 0
	var err error
	for scanner.Scan() {

		line := scanner.Text()
		// -1 means fi. nd all matches (no limit)
		allMatches := pattern.FindAllStringSubmatch(line, -1)
		first_no := 0
		second_no := 0

		for _, matches := range allMatches {
			first_no, err = strconv.Atoi(matches[1])
			second_no, err = strconv.Atoi(matches[2])
			if err != nil {
				fmt.Println("Error converting ASCII to integer", err)
			}
			mul_sum += first_no * second_no
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
	}
	fmt.Printf("Total count %d\n", mul_sum)

}
