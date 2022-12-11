package main

import (
	"fmt"
	"math"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func calculateSignalStrength(xValue int, cycle int) int {
	return xValue * cycle
}

type instruction struct {
	command string
	value   int
}

func execute(instruction instruction, cycle int, x int, midCycleCallback func(cycle int)) (int, int) {
	commandsToCycleLength := map[string]int{
		"noop": 1,
		"addx": 2,
	}

	requiredCycles := commandsToCycleLength[instruction.command]

	for i := 0; i < requiredCycles; i++ {
		cycle += 1
		midCycleCallback(cycle)

		if i == requiredCycles-1 {
			x += instruction.value
		}

	}

	return cycle, x
}

func sumSignalStrengths(instructions []instruction) int {
	sum := 0
	x := 1
	cycle := 0

	for _, instruction := range instructions {
		cycle, x = execute(instruction, cycle, x, func(cycle int) {
			const SIGNAL_STRENTGH_START = 20
			const SIGNAL_INTERVAL = 40

			if ((cycle - SIGNAL_STRENTGH_START) % SIGNAL_INTERVAL) == 0 {
				sum += calculateSignalStrength(x, cycle)
			}

		})
	}

	return sum
}

func draw(crt *[]string, cycle int, x int) {
	cycle -= 1 // 0 index cycle
	const SIGNAL_INTERVAL = 40
	row := int(cycle / SIGNAL_INTERVAL)
	column := cycle - (row * SIGNAL_INTERVAL)

	if len(*crt) <= row {
		*crt = append(*crt, "")
	}

	pixel := "."

	if math.Abs(float64(column-x)) < 2 {
		pixel = "#"
	}

	(*crt)[row] += pixel
}

func visualizeCrt(crt *[]string) {
	for _, line := range *crt {
		fmt.Println(line)
	}
}

func renderImage(instructions []instruction) {
	x := 1
	cycle := 0
	crt := []string{}

	for _, instruction := range instructions {
		cycle, x = execute(instruction, cycle, x, func(cycle int) {
			draw(&crt, cycle, x)
		})

	}

	visualizeCrt(&crt)
}

func normalizeLines(lines []string) []instruction {
	instructions := []instruction{}

	for _, line := range lines {
		var command string
		value := 0
		fmt.Sscanf(line, "%s %d", &command, &value)

		instructions = append(instructions, instruction{command, value})
	}

	return instructions
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	instructions := normalizeLines(lines)

	fmt.Println(sumSignalStrengths(instructions))
	renderImage(instructions)
}
