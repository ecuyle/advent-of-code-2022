package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func atoi(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return value
}

func splitStringPairIntoIntPair(str string, delimiter string) []int {
	pair := strings.Split(str, delimiter)

	return []int{atoi(pair[0]), atoi(pair[1])}
}

func main() {
	file := utils.Readfile("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		firstPair := splitStringPairIntoIntPair(assignments[0], "-")
		secondPair := splitStringPairIntoIntPair(assignments[1], "-")

		// if either sides are equal, the one contains the other
		if firstPair[0] == secondPair[0] || firstPair[1] == secondPair[1] {
			result += 1
			continue
		}

		// normalize such that firstPair starts earlier than second
		if secondPair[0] < firstPair[0] {
			firstPair, secondPair = secondPair, firstPair
		}

		if firstPair[1] > secondPair[1] {
			fmt.Println(firstPair, secondPair)
			result += 1
		}
	}

	fmt.Println(result)
}
