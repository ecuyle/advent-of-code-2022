package main

import (
    "bufio"
    "fmt"
    "github.com/ecuyle/advent-of-code-2022/utils"
)

func getPointsBetweenChoices(choice string, opponentChoice string) int {
    points := map[string]int{
        "A": 1,
        "B": 2,
        "C": 3,
    }

    diff := int(choice[0]) - int(opponentChoice[0])

    if diff == 0 {
        return 3 + points[choice]
    } else if diff == 1 || diff == -2 {
        return 6 + points[choice]
    }

    return 0 + points[choice]
}

func partOne() {
    file := utils.Readfile("./input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    results := 0

    for scanner.Scan() {
        line := scanner.Text()
        opponent := string(line[0])
        me := string(int(string(line[2])[0] - 23))
        results += getPointsBetweenChoices(me, opponent)
    }

    fmt.Println(results)
}

func getRequiredPlay(opponentChoice string, target string) string {
    alphabetOffsetFromZero := 65
    if target == "Y" {
        return opponentChoice
    } else if target == "X" {
        play := int(opponentChoice[0]) - alphabetOffsetFromZero - 1

        if play < 0 {
            return "C"
        }

        return string(play + alphabetOffsetFromZero)
    }

    play := int(opponentChoice[0]) - alphabetOffsetFromZero + 1

    if play > 2 {
        return "A"
    }

    return string(play + alphabetOffsetFromZero)
}

func partTwo() {
    file := utils.Readfile("./input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    results := 0

    for scanner.Scan() {
        line := scanner.Text()
        opponent := string(line[0])
        target := string(line[2])
        me := getRequiredPlay(opponent, target)
        results += getPointsBetweenChoices(me, opponent)
    }

    fmt.Println(results)
}

func main() {
    partOne()
    partTwo()
}
