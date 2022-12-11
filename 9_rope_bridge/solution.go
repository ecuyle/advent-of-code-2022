package main

import (
	"fmt"
	"math"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type node struct {
	x    float64
	y    float64
	next *node
}

func makeLinkedList(x float64, y float64, length int) *node {
	head := node{x, y, nil}
	next := &head

	for i := 1; i < length; i++ {
		next.next = &node{x, y, nil}
		next = next.next
	}

	return &head
}

func moveTail(head *node, tail *node) {
	for math.Abs(head.x-tail.x) >= 2 || math.Abs(head.y-tail.y) >= 2 {
		diffX := head.x - tail.x
		diffY := head.y - tail.y

		if math.Abs(diffX) > 0 {
			tail.x += math.Copysign(1, diffX)
		}

		if math.Abs(diffY) > 0 {
			tail.y += math.Copysign(1, diffY)
		}
	}

	if tail.next != nil {
		moveTail(tail, tail.next)
	}
}

func getPosition(node *node) string {
	return fmt.Sprintf("%f-%f", node.x, node.y)
}

type moveSet struct {
	x float64
	y float64
}

func moveHead(node *node, moveSet moveSet) {
	node.x += moveSet.x
	node.y += moveSet.y

	if node.next != nil {
		moveTail(node, node.next)
	}
}

func getTail(node *node) *node {
	if node.next == nil {
		return node
	}

	return getTail(node.next)
}

func uniqueTailVisits(directions []string, head *node) int {
	tVisited := map[string]bool{}
	moveSets := map[string]moveSet{
		"U": {0, 1},
		"R": {1, 0},
		"D": {0, -1},
		"L": {-1, 0},
	}

	tail := getTail(head)

	for _, direction := range directions {
		moveHead(head, moveSets[direction])
		tVisited[getPosition(tail)] = true
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

	fmt.Println(uniqueTailVisits(lines, makeLinkedList(0, 0, 2)))
	fmt.Println(uniqueTailVisits(lines, makeLinkedList(0, 0, 10)))
}
