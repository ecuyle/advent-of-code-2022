package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculateTotalCalories(places int, filepath string) int {
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	caloricHistory := []int{0}
	calories := 0

	for scanner.Scan() {
		raw := scanner.Text()

		if raw != "" {
			value, err := strconv.Atoi(raw)
			check(err)
			calories += value
			continue
		}

		caloricHistory = append(caloricHistory, calories)
		calories = 0
	}

	sort.Ints(caloricHistory)
	totalCalories := 0

	for i := 0; i < places; i++ {
		if len(caloricHistory) <= i {
			break
		}

		totalCalories += caloricHistory[len(caloricHistory)-i-1]
	}

	return totalCalories
}

func main() {
	fmt.Println("Max calories carried by an elf is", calculateTotalCalories(1, "./input.txt"))
	fmt.Println("Total calories of top three elves are:", calculateTotalCalories(3, "./input.txt"))
}
