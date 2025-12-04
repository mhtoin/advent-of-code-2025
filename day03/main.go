package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	file, err := common.GetInputFile(3)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSumOfJolts := 0

	for scanner.Scan() {
		line := scanner.Text()
		highestFound := 0
		secondHighestFound := 0

		for i := 0; i < len(line); i++ {
			char, _ := strconv.Atoi(string(line[i]))

			if highestFound == 0 || char > highestFound && i != len(line)-1 {
				highestFound = char
				secondHighestFound = 0
				continue

			}

			if (secondHighestFound == 0 && highestFound > 0) || (char > secondHighestFound) {
				secondHighestFound = char
				continue
			}

		}
		combinedNumber := highestFound*10 + secondHighestFound
		totalSumOfJolts += combinedNumber
	}

	fmt.Printf("Total Sum of Jolts: %d\n", totalSumOfJolts)

}

func solvePart2() {
	file, err := common.GetInputFile(3)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalSumOfJolts int64 = 0

	for scanner.Scan() {
		chosen := []int{}
		line := scanner.Text()
		start := 0

		for {
			windowSize := len(line) - (12 - len(chosen)) + 1

			slice := line[start:windowSize]
			ints := make([]int, len(slice))
			for i := 0; i < len(slice); i++ {
				ints[i] = int(slice[i] - '0')
			}
			maximum := slices.Max(ints)
			chosen = append(chosen, maximum)
			start += slices.Index(ints, maximum) + 1

			if len(chosen) == 12 {
				break
			}
		}

		var combinedNumber int64 = 0
		for _, digit := range chosen {
			combinedNumber = combinedNumber*10 + int64(digit)
		}
		totalSumOfJolts += int64(combinedNumber)
	}

	fmt.Printf("Total Sum of Jolts: %d\n", totalSumOfJolts)

}
