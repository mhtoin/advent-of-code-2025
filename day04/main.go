package main

import (
	"fmt"

	"github.com/mhtoin/advent-of-code-2025/common"
)

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	lines := common.Read2DLines(4)
	_, count := checkGrid(lines, false)
	fmt.Printf("Total accessible items: %d\n", count)
}

func checkGrid(lines [][]string, collectPositions bool) ([][2]int, int) {
	rows := len(lines)
	if rows == 0 {
		return nil, 0
	}
	cols := len(lines[0])

	var markedPositions [][2]int
	if collectPositions {
		markedPositions = make([][2]int, 0, rows)
	}

	totalAccessibleItems := 0

	for i := range rows {
		for j := range cols {
			if lines[i][j] != "@" {
				continue
			}

			invalidNeighbors := 0
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols && lines[ni][nj] == "@" {
					invalidNeighbors++
					if invalidNeighbors >= 4 {
						break
					}
				}
			}

			if invalidNeighbors < 4 {
				totalAccessibleItems++
				if collectPositions {
					markedPositions = append(markedPositions, [2]int{i, j})
				}
			}
		}
	}
	return markedPositions, totalAccessibleItems
}

func solvePart2() {
	lines := common.Read2DLines(4)
	totalRemoved := 0

	for {
		markedPositions, count := checkGrid(lines, true)
		if count == 0 {
			break
		}

		for _, pos := range markedPositions {
			lines[pos[0]][pos[1]] = "x"
		}
		totalRemoved += count
	}
	fmt.Printf("Total removed items: %d\n", totalRemoved)
}
