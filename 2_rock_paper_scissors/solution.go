package main

import (
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func rotate(choice string, count int) string {
	choices := []string{"A", "B", "C"}
	// get 0 indexed value
	choiceIndex := int(choice[0]) - 65
	total := choiceIndex + count

	if total < 0 {
		remainder := (len(choices) - total) % len(choices)
		return choices[len(choices)-remainder]
	}

	remainder := total % len(choices)
	return choices[remainder]
}

// returns 6 if win, 3 if tie, 0 if loss
func play(mine string, opponent string) int {
	if mine == opponent {
		return 3
	}

    if rotate(opponent, 1) == mine {
        return 6
    }

	return 0
}

func getPointsBetweenChoices(mine string, opponent string) int {
	pointsForChoice := int(mine[0]) - 65 + 1

	return pointsForChoice + play(mine, opponent)
}

func calculatePointsFromStrategyGuide(lines [][]string) int {
	results := 0

	for _, line := range lines {
		opponent := line[0]
		me := line[1]
		results += getPointsBetweenChoices(me, opponent)
	}

	return results
}

func getRequiredPlay(opponentChoice string, instruction string) string {
	instructions := map[string]int{
		"A": -1,
		"B": 0,
		"C": 1,
	}

	return rotate(opponentChoice, instructions[instruction])
}

func calculatePointsAfterUpdate(lines [][]string) int {
	points := 0

	for _, line := range lines {
		opponent := line[0]
		instruction := line[1]
		me := getRequiredPlay(opponent, instruction)
		points += getPointsBetweenChoices(me, opponent)
	}

	return points
}

func normalizeLinesForPlay(lines []string) [][]string {
	normalized := [][]string{}

	for _, line := range lines {
		var opponent, mine string
		fmt.Sscanf(line, "%s %s", &opponent, &mine)
		// converts x|y|z to a|b|c
		normalized = append(normalized, []string{opponent, string(int(mine[0]) - 23)})
	}

	return normalized
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	normalized := normalizeLinesForPlay(lines)
	fmt.Println(calculatePointsFromStrategyGuide(normalized))
	fmt.Println(calculatePointsAfterUpdate(normalized))
}
