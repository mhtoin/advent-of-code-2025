package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	file, err := common.GetInputFile(2)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, ",")

		for i := range ids {
			id := strings.TrimSpace(ids[i])
			parts := strings.Split(id, "-")
			if len(parts) != 2 {
				fmt.Printf("Invalid ID format: %s\n", id)
				continue
			}

			startStr, endStr := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			start, err1 := strconv.Atoi(startStr)
			end, err2 := strconv.Atoi(endStr)

			if err1 != nil || err2 != nil {
				fmt.Printf("Invalid number in ID: %s\n", id)
				continue
			}

			for num := start; num <= end; num++ {
				numStr := strconv.Itoa(num)
				if len(numStr)%2 != 0 {
					continue
				}

				mid := len(numStr) / 2
				numParts := []string{numStr[:mid], numStr[mid:]}

				if len(numParts) != 2 {
					continue
				}

				firstHalf, secondHalf := numParts[0], numParts[1]
				if firstHalf == secondHalf {
					totalSum += num
				}

			}

		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total sum: %d\n", totalSum)
}

func solvePart2() {
	file, err := common.GetInputFile(2)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, ",")

		for i := range ids {
			id := strings.TrimSpace(ids[i])
			parts := strings.Split(id, "-")
			if len(parts) != 2 {
				continue
			}

			startStr, endStr := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			start, err1 := strconv.Atoi(startStr)
			end, err2 := strconv.Atoi(endStr)

			if err1 != nil || err2 != nil {
				continue
			}

			for num := start; num <= end; num++ {
				numStr := strconv.Itoa(num)
				lenNum := len(numStr)

				for parts := 2; parts <= lenNum; parts++ {
					if lenNum%parts != 0 {
						continue
					}

					partLen := lenNum / parts
					pattern := numStr[:partLen]

					if strings.Repeat(pattern, parts) == numStr {
						totalSum += num
						break
					}
				}
			}
		}
	}
	fmt.Printf("Total sum: %d\n", totalSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
