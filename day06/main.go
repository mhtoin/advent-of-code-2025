package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func getColumn(grid [][]string, colIndex int) []string {
	column := make([]string, len(grid))
	for i := range grid {
		column[i] = grid[i][colIndex]
	}
	return column
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	grid := make([][]string, 0)
	common.ForEachLine(6, func(line string) {
		parts := strings.Fields(line)
		grid = append(grid, parts)
	})
	total := 0

	for col := 0; col < len(grid[0]); col++ {
		column := getColumn(grid, col)

		columnTotal := 0
		columnNumbers := column[:len(column)-1]

		if column[len(column)-1] == "*" {
			columnTotal = 1
			for _, val := range columnNumbers {
				num := common.MustAtoi(val)
				columnTotal *= num
			}
			total += columnTotal
		}

		if column[len(column)-1] == "+" {
			for _, val := range columnNumbers {
				num := common.MustAtoi(val)
				columnTotal += num
			}
			total += columnTotal
		}
	}
	fmt.Printf("Part 1: %d\n", total)

}

func solvePart2() {
	total := 0
	grid := make([][]string, 0)
	common.ForEachLine(6, func(line string) {
		parts := strings.Split(line, "")
		grid = append(grid, parts)
	})

	currentOperation := make([][]string, 0)

	processOperation := func() {
		if len(currentOperation) == 0 {
			return
		}

		operationItems := make([]int, 0)
		operator := ""
		operationTotal := 0
		for _, row := range currentOperation {
			numberStr := ""
			for _, val := range row {
				if val == "*" || val == "+" {
					operator = val
				} else if val != " " {
					numberStr += val
				}
			}
			if numberStr != "" {
				num := common.MustAtoi(numberStr)
				operationItems = append(operationItems, num)
			}
		}

		for _, item := range operationItems {
			switch operator {
			case "*":
				if operationTotal == 0 {
					operationTotal = 1
				}
				operationTotal *= item
			case "+":
				operationTotal += item
			}
		}
		total += operationTotal
	}

	for col := 0; col < len(grid[0]); col++ {
		column := getColumn(grid, col)
		allEmpty := !slices.ContainsFunc(column, func(s string) bool {
			return s != " "
		})

		if !allEmpty {
			currentOperation = append(currentOperation, column)
		} else {
			processOperation()
			currentOperation = make([][]string, 0)
		}
	}

	// Process the last operation
	processOperation()

	fmt.Println("Total: ", total)
}
