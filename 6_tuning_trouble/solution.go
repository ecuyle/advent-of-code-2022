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

func partOne() int {
	file := utils.Readfile("./input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	const SIGNAL_LOCK_LEN = 4
	seen := []string{}

	for {
		rune, _, err := reader.ReadRune()

		if err != nil {
			break
		}

		char := string(rune)
		seen = append(seen, char)

		if len(seen) < SIGNAL_LOCK_LEN {
			continue
		}

		if hasSignalLock(seen[len(seen)-SIGNAL_LOCK_LEN : len(seen)]) {
			return len(seen)
		}

	}

	return -1
}

func main() {
	fmt.Println(partOne())
}
