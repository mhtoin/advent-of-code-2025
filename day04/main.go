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
	lines := common.Read2DLines(4)
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalAccessibleItems := 0

	for i := range lines {
		for j := range lines[i] {
			item := lines[i][j]
			if item == "@" {
				amountOfInvalidNeighbors := 0

				for _, dir := range directions {
					if amountOfInvalidNeighbors >= 4 {
						break
					}
					ni, nj := i+dir[0], j+dir[1]
					if ni < 0 || ni >= len(lines) || nj < 0 || nj >= len(lines[i]) {
						continue
					}
					neighbor := lines[ni][nj]
					if neighbor == "@" {
						amountOfInvalidNeighbors++
					}
				}
				if amountOfInvalidNeighbors < 4 {
					totalAccessibleItems++
				}
			}
		}
	}
	fmt.Printf("Total accessible items: %d\n", totalAccessibleItems)
}

func checkGrid(lines [][]string) ([][2]int, int) {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	markedPositions := make([][2]int, 0)

	totalAccessibleItems := 0

	for i := range lines {
		for j := range lines[i] {
			item := lines[i][j]
			if item == "@" {
				amountOfInvalidNeighbors := 0

				for _, dir := range directions {
					if amountOfInvalidNeighbors >= 4 {
						break
					}
					ni, nj := i+dir[0], j+dir[1]
					if ni < 0 || ni >= len(lines) || nj < 0 || nj >= len(lines[i]) {
						continue
					}
					neighbor := lines[ni][nj]
					if neighbor == "@" {
						amountOfInvalidNeighbors++
					}
				}
				if amountOfInvalidNeighbors < 4 {
					totalAccessibleItems++
					markedPositions = append(markedPositions, [2]int{i, j})

				}
			}
		}
	}
	return markedPositions, totalAccessibleItems
}

func removeMarkedPositions(lines [][]string, markedPositions [][2]int) int {
	removedCount := 0
	for _, pos := range markedPositions {
		i, j := pos[0], pos[1]
		lines[i][j] = "x"
		removedCount++
	}
	return removedCount
}

func solvePart2() {
	lines := common.Read2DLines(4)
	totalRemoved := 0

	for {
		markedPositions, totalAccessibleItems := checkGrid(lines)

		if totalAccessibleItems == 0 {
			break
		}

		removedCount := removeMarkedPositions(lines, markedPositions)
		totalRemoved += removedCount
	}
	fmt.Printf("Total removed items: %d\n", totalRemoved)
}
