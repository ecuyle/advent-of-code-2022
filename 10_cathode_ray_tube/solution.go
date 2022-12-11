package main

import (
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func calculateSignalStrength(xValue int, cycle int) int {
	return xValue * cycle
}

type instruction struct {
	command string
	value   int
}

func execute(instruction instruction, cycle int, x int, sum int) (int, int, int) {
	commandsToCycleLength := map[string]int{
		"noop": 1,
		"addx": 2,
	}

	requiredCycles := commandsToCycleLength[instruction.command]
	const SIGNAL_STRENTGH_START = 20
    const SIGNAL_INTERVAL = 40

	for i := 0; i < requiredCycles; i++ {
		cycle += 1

		if ((cycle - SIGNAL_STRENTGH_START) % SIGNAL_INTERVAL) == 0 {
			sum += calculateSignalStrength(x, cycle)
		}

		if i == requiredCycles-1 {
			x += instruction.value
		}

	}

	return cycle, x, sum
}

func sumSignalStrengths(instructions []instruction) int {
	sum := 0
	x := 1
	cycle := 0

	for _, instruction := range instructions {
		cycle, x, sum = execute(instruction, cycle, x, sum)
	}

	return sum
}

func partTwo(inputPath string) {
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
	// partTwo("./input_test.txt")
}
