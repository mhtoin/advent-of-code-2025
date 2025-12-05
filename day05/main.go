package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mhtoin/advent-of-code-2025/common"
)

type Range struct {
	start int
	end   int
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	lines := common.ReadLines(5)

	// dataset includes ranges and product ids, separated by a blank space
	// so find the index of the blank space first
	blankIndex := slices.Index(lines, "")

	idRanges := lines[:blankIndex]
	productIDs := lines[blankIndex+1:]
	freshIds := make([]string, 0, len(productIDs))

	/*
	* Naive loop solution
	 */

	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for _, id := range productIDs {
			idNum, _ := strconv.Atoi(id)

			if idNum >= start && idNum <= end && !slices.Contains(freshIds, id) {
				freshIds = append(freshIds, id)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(freshIds))

}

func solvePart2() {
	file, err := common.GetInputFile(5)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ranges := make([]Range, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, Range{start: start, end: end})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})

	mergedRanges := make([]Range, 0)
	currentRange := ranges[0]

	for _, r := range ranges[1:] {
		if r.start <= currentRange.end {
			if r.end > currentRange.end {
				currentRange.end = r.end
			}
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = r
		}
	}
	mergedRanges = append(mergedRanges, currentRange)

	totalNumberOfIDsCovered := 0
	for _, r := range mergedRanges {
		totalNumberOfIDsCovered += (r.end - r.start + 1)
	}

	fmt.Printf("Part 2: %d\n", totalNumberOfIDsCovered)
}
