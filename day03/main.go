package main

import (
	"fmt"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	total := 0

	common.ForEachLine(3, func(line string) {
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
	})
	fmt.Printf("Total Sum of Jolts: %d\n", total)
}

func solvePart2() {
	var total int64 = 0

	common.ForEachLine(3, func(line string) {
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
	})
	fmt.Printf("Total Sum of Jolts: %d\n", total)
}
