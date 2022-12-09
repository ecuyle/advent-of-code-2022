package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func aggregateCaloriesPerElf(lines []string) []int {
	caloriesPerElf := []int{0}
	currentCalories := 0

	for _, line := range lines {
		if line == "" {
			caloriesPerElf = append(caloriesPerElf, currentCalories)
			currentCalories = 0
			continue
		}

		caloriesFromLine, err := strconv.Atoi(line)
		utils.CheckError(err)
		currentCalories += caloriesFromLine
	}

	return caloriesPerElf
}

func calculateTotalCaloriesForNElves(elves int, lines []string) int {
	caloriesPerElf := aggregateCaloriesPerElf(lines)
	sort.Slice(caloriesPerElf, func(a, b int) bool {
		return caloriesPerElf[a] > caloriesPerElf[b]
	})

	if len(caloriesPerElf) < elves {
		return utils.Sum(caloriesPerElf)
	}

	return utils.Sum(caloriesPerElf[:elves])
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	fmt.Println("Max calories carried by an elf is", calculateTotalCaloriesForNElves(1, lines))
	fmt.Println("Total calories of top three elves are:", calculateTotalCaloriesForNElves(3, lines))
}
