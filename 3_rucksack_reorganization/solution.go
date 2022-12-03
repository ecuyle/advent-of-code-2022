package main

import (
    "bufio"
    "fmt"
    "github.com/ecuyle/advent-of-code-2022/utils"
)

func calculatePriorityFromUnicodeChar(char rune) int {
    intChar := int(char)

    // if uppercase, unicode values are between 65 and 90, but need to get down between 27 - 52
    if intChar >= 65 && intChar <= 90 {
        return intChar - 65 + 27
    }

    // otherwise lowercase, unicode values are between 97 and 122, but need to get down to 1 - 26
    return intChar - 97 + 1
}

func partOne() {
    file := utils.Readfile("./input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    sum := 0

    for scanner.Scan() {
        seen := map[rune]int{}
        line := scanner.Text()
        firstCompartment := line[0:len(line) / 2]

        for _, char := range firstCompartment {
            seen[char] = 1
        }

        secondCompartment := line[len(line) / 2: len(line)]

        for _, char := range secondCompartment {
            if seen[char] == 1 {
                converted := calculatePriorityFromUnicodeChar(char)
                sum += converted
                break
            }
        }

    }

    fmt.Println(sum)
}

func main() {
    partOne()
}

