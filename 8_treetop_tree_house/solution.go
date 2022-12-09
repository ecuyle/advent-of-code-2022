package main

import (
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func getTreeId(row int, column int) string {
	return fmt.Sprintf("%d-%d", row, column)
}

func partOne(lines []string) int {
	visible := map[string]bool{}

	// two pointers horizontally
	for row, line := range lines {
		front := 0
		back := len(line) - 1
		highestFront := line[0] - 1
		highestBack := line[len(line)-1] - 1

		for front < back {
			if line[front] > highestFront {
				highestFront = line[front]
				visible[getTreeId(row, front)] = true
			}

			if line[back] > highestBack {
				highestBack = line[back]
				visible[getTreeId(row, back)] = true
			}

			if line[front] <= line[back] {
				front += 1
			} else {
				back -= 1
			}
		}
	}

	// two pointers vertically
	for column := 0; column < len(lines[0])-1; column++ {
		front := 0
		back := len(lines) - 1
		highestFront := lines[0][column] - 1
		highestBack := lines[len(lines)-1][column] - 1

		for front < back {
			if lines[front][column] > highestFront {
				highestFront = lines[front][column]
				visible[getTreeId(front, column)] = true
			}

			if lines[back][column] > highestBack {
				highestBack = lines[back][column]
				visible[getTreeId(back, column)] = true
			}

			if lines[front][column] <= lines[back][column] {
				front += 1
			} else {
				back -= 1
			}
		}
	}

	return len(visible)
}

func partTwo(lines []string) {
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	fmt.Println(partOne(lines))
	// partTwo("./input_test.txt")
}
