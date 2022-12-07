package main

import (
	"bufio"
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func hasSignalLock(stream []string) bool {
	seen := map[string]bool{}

	for _, char := range stream {
		if seen[char] {
			return false
		}

		seen[char] = true
	}

	return true
}

func findSignalLockMarker(inputPath string, signalLockLength int) int {
	file := utils.Readfile(inputPath)
	defer file.Close()
	reader := bufio.NewReader(file)
	seen := []string{}

	for {
		rune, _, err := reader.ReadRune()

		if err != nil {
			break
		}

		char := string(rune)
		seen = append(seen, char)

		if len(seen) < signalLockLength {
			continue
		}

		if hasSignalLock(seen[len(seen)-signalLockLength : len(seen)]) {
			return len(seen)
		}

	}

	return -1
}

func findSignalLockMarkerOptimized(inputPath string, signalLockLength int) int {
	file := utils.Readfile(inputPath)
	defer file.Close()
	reader := bufio.NewReader(file)
	seen := map[string]int{}
	p1 := 0
	p2 := 0
	successes := 0

	for {
		rune, _, err := reader.ReadRune()

		if err != nil {
			break
		}

		char := string(rune)

		if index, ok := seen[char]; ok && index >= p1 {
			p1 = index + 1
			successes = p2 - p1
		}

		successes += 1
		seen[char] = p2
		p2 += 1

		if successes == signalLockLength {
			return p2
		}
	}

	return -1
}

func main() {
	fmt.Println("Part 1 answer is:", findSignalLockMarker("./input.txt", 4))
	fmt.Println("Part 2 answer is:", findSignalLockMarker("./input.txt", 14))

	fmt.Println("Part 1 answer is:", findSignalLockMarkerOptimized("./input.txt", 4))
	fmt.Println("Part 2 answer is:", findSignalLockMarkerOptimized("./input.txt", 14))
}
