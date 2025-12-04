package main

import (
	"bufio"
	"fmt"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	file, _ := common.GetInputFile(3)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		var num int = 0
		pos := 0

		for i := range 2 {
			remaining := 2 - i - 1
			end := len(line) - remaining
			maxDigit, maxIdx := byte('0'), pos

			for j := pos; j < end; j++ {
				if line[j] > maxDigit {
					maxDigit, maxIdx = line[j], j
				}
			}
			num = num*10 + int(maxDigit-'0')
			pos = maxIdx + 1
		}
		total += num
	}
	fmt.Printf("Total Sum of Jolts: %d\n", total)
}

func solvePart2() {
	file, _ := common.GetInputFile(3)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var total int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		var num int64 = 0
		pos := 0

		for i := range 12 {
			remaining := 12 - i - 1
			end := len(line) - remaining
			maxDigit, maxIdx := byte('0'), pos

			for j := pos; j < end; j++ {
				if line[j] > maxDigit {
					maxDigit, maxIdx = line[j], j
				}
			}
			num = num*10 + int64(maxDigit-'0')
			pos = maxIdx + 1
		}
		total += num
	}
	fmt.Printf("Total Sum of Jolts: %d\n", total)
}
