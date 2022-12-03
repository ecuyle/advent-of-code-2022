package main

import (
    "fmt"
    "os"
    "bufio"
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

func readfile(path string) *os.File {
    file, err := os.Open(path)

    if err != nil {
        panic(err)
    }

    return file
}

func main() {
    file := readfile("./input.txt")
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
