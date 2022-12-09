package main

import (
	"fmt"
	"math"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func getTreeId(row int, column int) string {
	return fmt.Sprintf("%d-%d", row, column)
}

func countVisibleTrees(lines []string) int {
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

func calculateScore(target string, lines []string, row int, column int, direction int, bearing string) int {
	score := 0

	for row >= 0 && row < len(lines) && column >= 0 && column < len(lines[0]) {
		score += 1

		if string(lines[row][column]) >= target {
			break
		}

		if bearing == "vertical" {
			row += direction
		} else {
			column += direction
		}

	}

	return score
}

func findMostScenicTree(lines []string) int {
	rows := len(lines)
	columns := len(lines[0])
	maxScore := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < columns-1; j++ {
			target := string(lines[i][j])
			up := calculateScore(target, lines, i-1, j, -1, "vertical")
			down := calculateScore(target, lines, i+1, j, 1, "vertical")
			left := calculateScore(target, lines, i, j-1, -1, "horizontal")
			right := calculateScore(target, lines, i, j+1, 1, "horizontal")
			maxScore = int(math.Max(float64(maxScore), float64(up*down*left*right)))
		}
	}

	return maxScore
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	fmt.Println(countVisibleTrees(lines))
	fmt.Println(findMostScenicTree(lines))
}
