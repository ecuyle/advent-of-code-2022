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

func partTwo() {
    file := utils.Readfile("./input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    sum := 0
    counter := 0
    seen := map[rune]int{}
    const ELVES_PER_GROUP = 3

    for scanner.Scan() {
        line := scanner.Text()
        localSeen := map[rune]int{}

        for _, char := range line {
            // if the character is unique in this line, add it to the overall seen map
            if localSeen[char] != 1 && seen[char] < counter + 1 {
                localSeen[char] = 1
                seen[char] += 1
            }

            if seen[char] == ELVES_PER_GROUP {
                sum += calculatePriorityFromUnicodeChar(char)
                break
            }
        }

        // we've hit the third elf in the group, reset
        if counter == 2 {
            counter = 0
            seen = map[rune]int{}
        } else {
            counter += 1
        }
    }

    fmt.Println(sum)
}

func main() {
    partOne()
    partTwo()
}

