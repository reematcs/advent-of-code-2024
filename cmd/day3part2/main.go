package main

import (
	"aoc/pkg/common"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	resp := common.LoadClient("https://adventofcode.com/2024/day/3/input")
	scanner := bufio.NewScanner(resp.Body)

	mul_sum := 0
	enabled := true // Initial state - multiplications are enabled

	mulPattern := regexp.MustCompile(`mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)`)

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		pos := 0

		for pos < len(line) {
			remainder := line[pos:]

			// Find positions of control operators
			doPos := strings.Index(remainder, "do()")
			dontPos := strings.Index(remainder, "don't()")
			mulMatch := mulPattern.FindStringSubmatchIndex(remainder)

			// Find the earliest operator
			nextPos := len(remainder)
			if doPos >= 0 && doPos < nextPos {
				nextPos = doPos
			}
			if dontPos >= 0 && dontPos < nextPos {
				nextPos = dontPos
			}
			if mulMatch != nil && mulMatch[0] < nextPos {
				nextPos = mulMatch[0]
			}

			// Process the earliest operator
			if nextPos == len(remainder) {
				break
			} else if doPos == nextPos {
				enabled = true
				pos += doPos + 4
			} else if dontPos == nextPos {
				enabled = false
				pos += dontPos + 7
			} else if mulMatch != nil && mulMatch[0] == nextPos {
				if enabled {
					num1, _ := strconv.Atoi(remainder[mulMatch[2]:mulMatch[3]])
					num2, _ := strconv.Atoi(remainder[mulMatch[4]:mulMatch[5]])
					mul_sum += num1 * num2
				}
				pos += mulMatch[1]
			} else {
				pos++
			}
		}
	}
	fmt.Printf("Sum: %d\n", mul_sum)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net/http"
// 	"regexp"
// 	"strconv"
// )

// func main() {
// 	url := "https://adventofcode.com/2024/day/3/input"
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	req.AddCookie(&http.Cookie{
// 		Name:  "session",
// 		Value: "53616c7465645f5f28c8f1667102836be687d0a6ea3e147d499293ff26bc6aedaddb780d27f6a1c13d4009747369b2e3ae24cb28da063b9b5805dee413ea3f82",
// 	})

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error fetching URL:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	scanner := bufio.NewScanner(resp.Body)

// 	mul_sum := 0

// 	doPattern := regexp.MustCompile(`do\(\)`)
// 	dontPattern := regexp.MustCompile(`don't\(\)`)
// 	mulPattern := regexp.MustCompile(`mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)`)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		enabled := true
// 		pos := 0
// 		for pos < len(line) {
// 			if loc := doPattern.FindStringIndex(line[pos:]); loc != nil {
// 				enabled = true
// 				pos += loc[1]
// 				continue
// 			}
// 			if loc := dontPattern.FindStringIndex(line[pos:]); loc != nil {
// 				enabled = false
// 				pos += loc[1]
// 				continue
// 			}

// 			// Look for multiplication
// 			if matches := mulPattern.FindStringSubmatch(line[pos:]); matches != nil {
// 				if enabled {
// 					num1, _ := strconv.Atoi(matches[1])
// 					num2, _ := strconv.Atoi(matches[2])
// 					mul_sum += num1 * num2
// 				}
// 				pos += len(matches[0])
// 			} else {
// 				pos++ // Move forward if no pattern matches
// 			}
// 		}

// 	}
// 	fmt.Printf("Sum: %d\n", mul_sum)
// }
