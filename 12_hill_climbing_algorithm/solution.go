package main

import (
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type Point struct {
	x     int
	y     int
	value rune
	id    string
}

type TopologyMap struct {
	graph   [][]*Point
	low     *Point
	high    *Point
	visited *map[string]bool
}

type PointDistance struct {
	point    *Point
	distance int
}

func getEligiblePoints(topologyMap *TopologyMap, pointDistance *PointDistance, direction string) []*PointDistance {
	point := pointDistance.point
	distance := pointDistance.distance
	points := []*PointDistance{}

	for _, instruction := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
		newX := point.x + instruction[0]
		newY := point.y + instruction[1]

		if newY < 0 || newY >= len(topologyMap.graph) || newX < 0 || newX >= len(topologyMap.graph[0]) {
			continue
		}

		newPoint := topologyMap.graph[newY][newX]

		if _, ok := (*topologyMap.visited)[newPoint.id]; ok {
			continue
		}

		if (direction == "up" && (newPoint.value-point.value) > 1) || (direction == "down" && ((point.value - newPoint.value) > 1)) {
			continue
		}

		points = append(points, &PointDistance{newPoint, distance + 1})
		(*topologyMap.visited)[newPoint.id] = true
	}

	return points
}

func getFewestSteps(topologyMap *TopologyMap, direction string) int {
	start := topologyMap.low
	target := topologyMap.high

	if direction == "down" {
		start = topologyMap.high
		target = topologyMap.low
	}

	queue := []*PointDistance{{start, 0}}

	for len(queue) != 0 {
		popped := queue[0]
		point := popped.point

		if point.value == target.value {
			return popped.distance
		}

		queue = queue[1:]
		eligiblePoints := getEligiblePoints(topologyMap, popped, direction)
		queue = append(queue, eligiblePoints...)
	}

	return 0
}

func getPointID(x int, y int) string {
	return fmt.Sprint(x) + "-" + fmt.Sprint(y)
}

func makeTopologyMap(lines []string) *TopologyMap {
	startChar := 'S'
	endChar := 'E'
	topologyMap := &TopologyMap{graph: [][]*Point{}, visited: &map[string]bool{}}

	for y, line := range lines {
		for x, char := range line {
			point := &Point{x, y, char, getPointID(x, y)}

			switch char {
			case startChar:
				point.value = 'a'
				topologyMap.low = point
			case endChar:
				point.value = 'z' + 1
				topologyMap.high = point
			}

			if len(topologyMap.graph) <= y {
				topologyMap.graph = append(topologyMap.graph, []*Point{})
			}

			topologyMap.graph[y] = append(topologyMap.graph[y], point)
		}
	}

	return topologyMap
}

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	fmt.Println(getFewestSteps(makeTopologyMap(lines), "up"))
	fmt.Println(getFewestSteps(makeTopologyMap(lines), "down"))
}
