package main

import (
	"bufio"
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func parseGraphIntoArrays(lines []string) [][]string {
	parsed := [][]string{}
	lines = lines[:len(lines)-1]

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		spaceBetweenBoxValues := 4
		currentBox := 0

		for j := 1; j < len(line); j += spaceBetweenBoxValues {
			if len(parsed) <= currentBox {
				parsed = append(parsed, []string{})
			}

			value := string(line[j])

			if value != " " {
				parsed[currentBox] = append(parsed[currentBox], string(line[j]))
			}
			currentBox += 1
		}
	}

	return parsed
}

func parseCommand(command string) []int {
	var count, from, to int
	_, err := fmt.Sscanf(command, "move %d from %d to %d", &count, &from, &to)

	if err != nil {
		panic(err)
	}

	return []int{count, from - 1, to - 1}
}

func move(count int, from int, to int, graph [][]string) [][]string {
	for i := 0; i < count; i++ {
		fromBucket := graph[from]
		toBucket := graph[to]
		graph[to] = append(toBucket, fromBucket[len(fromBucket)-1])
		graph[from] = fromBucket[:len(fromBucket)-1]
	}

	return graph
}

func partOne() {
	file := utils.Readfile("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph := []string{}
	commands := []string{}
	isGraphComplete := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isGraphComplete = true
			continue
		}

		if isGraphComplete {
			commands = append(commands, line)
			continue
		}

		graph = append(graph, line)
	}

	parsedGraph := parseGraphIntoArrays(graph)

	for _, command := range commands {
		parsedCommand := parseCommand(command)
		parsedGraph = move(parsedCommand[0], parsedCommand[1], parsedCommand[2], parsedGraph)
	}

	result := ""

	for _, column := range parsedGraph {
		result += column[len(column)-1]
	}

	fmt.Println(result)
}

func main() {
	partOne()
}