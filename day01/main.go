package main

import (
	"bufio"
	"strconv"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func main() {
	solvePart1()
	solvePart2()

}

func solvePart1() {
	file, err := common.GetInputFile(1)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	dial := 50
	timesAtZero := 0

	println(dial, timesAtZero)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		direction, value := runes[0], string(runes[1:])
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		if direction == 'R' {
			dial = (dial + valueInt) % 100
		} else if direction == 'L' {
			dial = ((dial-valueInt)%100 + 100) % 100
		}

		if dial == 0 {
			timesAtZero++
		}

		println(direction, valueInt, "->", dial, timesAtZero)
	}

	println("Times at 0:", timesAtZero)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func solvePart2() {
	file, err := common.GetInputFile(1)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	dial := 50
	timesAtZero := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		valueInt, _ := strconv.Atoi(line[1:])

		if direction == 'R' {
			distToZero := (100 - dial) % 100
			if distToZero == 0 {
				distToZero = 100
			}
			if valueInt >= distToZero {
				timesAtZero += 1 + (valueInt-distToZero)/100
			}
		} else {
			if valueInt >= dial && dial > 0 {
				timesAtZero += 1 + (valueInt-dial)/100
			} else if dial == 0 && valueInt >= 100 {
				timesAtZero += valueInt / 100
			}
		}

		dial = ((dial+(map[byte]int{'R': 1, 'L': -1}[direction]*valueInt))%100 + 100) % 100
	}
	println("Times at 0:", timesAtZero)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
