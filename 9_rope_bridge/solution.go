package main

import (
	"fmt"
	"math"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type node struct {
	x float64
	y float64
}

func moveHead(node node, moveSet node) node {
	node.x += moveSet.x
	node.y += moveSet.y

	fmt.Println("H:", node)
	return node
}

func moveTail(head node, tail node) node {
	for math.Abs(head.x-tail.x) >= 2 || math.Abs(head.y-tail.y) >= 2 {
		diffX := head.x - tail.x
		diffY := head.y - tail.y
		fmt.Println(diffX, diffY)

		if math.Abs(diffX) > 0 {
			tail.x += math.Copysign(1, diffX)
		}

		if math.Abs(diffY) > 0 {
			tail.y += math.Copysign(1, diffY)
		}
	}

	fmt.Println("T:", tail)
	return tail
}

func getPosition(node node) string {
	return fmt.Sprintf("%f-%f", node.x, node.y)
}

func uniqueTailVisits(directions []string) int {
	tVisited := map[string]bool{}
	h := node{0, 0}
	t := node{0, 0}
	moveSets := map[string]node{
		"U": {0, 1},
		"R": {1, 0},
		"D": {0, -1},
		"L": {-1, 0},
	}

	for _, direction := range directions {
		fmt.Println(direction)
		h = moveHead(h, moveSets[direction])
		t = moveTail(h, t)
		tVisited[getPosition(t)] = true
	}

	return len(tVisited)
}

func partTwo(lines []string) int {
	return 0
}

// converts `["R 4"]` to `["R", "R", "R", "R"]`
func normalizeDirections(lines []string) []string {
	normalized := []string{}
	for _, line := range lines {
		var direction string
		var count int
		fmt.Sscanf(line, "%s %d", &direction, &count)

		for i := 0; i < count; i++ {
			normalized = append(normalized, direction)
		}
	}

	return normalized
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	lines = normalizeDirections(lines)
	fmt.Println(uniqueTailVisits(lines))
	fmt.Println(partTwo(lines))
}
