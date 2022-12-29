package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type Packet []interface{}

func isInRightOrder(left, right Packet) int {
	fmt.Println("Comparing", left, right)
	for i := 0; i < len(left); i++ {
		if i >= len(right) {
			return -1
		}

		lFloat, isLeftFloat := left[i].(float64)
		rFloat, isRightFloat := right[i].(float64)
		lArray, isLeftArray := left[i].([]interface{})
		rArray, isRightArray := right[i].([]interface{})

		fmt.Println(lFloat, rFloat)

		if isLeftArray && isRightArray {
			res := isInRightOrder(lArray, rArray)
			if res != 0 {
				return res
			}
		}

		if isLeftFloat && isRightArray {
			res := isInRightOrder(append(Packet{}, lFloat), rArray)
			if res != 0 {
				return res
			}
		}

		if isRightFloat && isLeftArray {
			res := isInRightOrder(lArray, append(Packet{}, rFloat))
			if res != 0 {
				return res
			}
		}

		if lFloat < rFloat {
			return 1
		} else if lFloat > rFloat {
			return -1
		}
	}

	if len(right) > len(left) {
		return 1
	}

	return 0
}

func parseLine(line []byte) Packet {
	packet := Packet{}
	json.Unmarshal(line, &packet)

	return packet
}

func partOne(lines []string) int {
	sum := 0
	current := 1

	for len(lines) > 0 {
		pair := strings.Split(lines[0], "\n")
		left := parseLine([]byte(pair[0]))
		right := parseLine([]byte(pair[1]))
		lines = lines[1:]

		res := isInRightOrder(left, right)
		fmt.Println("After comparing:", res)
		fmt.Println(left, right)
		if res == 1 {
			fmt.Println("Adding current:", current)
			sum += current
		}
		fmt.Println("-----------------------------------------------------------------------")
		current += 1
	}

	return sum
}

func main() {
	// str := utils.ReadFileIntoString("./input_test.txt")
	str := utils.ReadFileIntoString("./input.txt")
	lines := strings.Split(str, "\n\n")
	fmt.Println(partOne(lines))
}
