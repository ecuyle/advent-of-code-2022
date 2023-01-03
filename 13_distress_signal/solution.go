package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type Packet []interface{}

func isInRightOrder(left, right Packet) int {
	for i := 0; i < len(left); i++ {
		if i >= len(right) {
			return -1
		}

		lFloat, isLeftFloat := left[i].(float64)
		rFloat, isRightFloat := right[i].(float64)
		lArray, isLeftArray := left[i].([]interface{})
		rArray, isRightArray := right[i].([]interface{})

		if isLeftArray && isRightArray {
			res := isInRightOrder(lArray, rArray)
			if res != 0 {
				return res
			}
		} else if isLeftFloat && isRightArray {
			res := isInRightOrder(append(Packet{}, lFloat), rArray)
			if res != 0 {
				return res
			}
		} else if isRightFloat && isLeftArray {
			res := isInRightOrder(lArray, append(Packet{}, rFloat))
			if res != 0 {
				return res
			}
			continue
		} else if lFloat < rFloat {
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

func partOne(packets []Packet) int {
	sum := 0
	current := 1

	for len(packets) > 0 {
		left := packets[0]
		right := packets[1]
		packets = packets[2:]

		res := isInRightOrder(left, right)

		if res == 1 {
			sum += current
		}

		current += 1
	}

	return sum
}

func isDivider(packet Packet) bool {
	if len(packet) != 1 {
		return false
	}

	arr, isArray := packet[0].([]interface{})

	if !isArray || len(arr) != 1 {
		return false
	}

	value, isFloat := arr[0].(float64)

	return isFloat && (value == 2 || value == 6)
}

func partTwo(packets []Packet) int {
	dividers := []Packet{parseLine([]byte("[[2]]")), parseLine([]byte("[[6]]"))}
	packets = append(packets, dividers...)
	product := 1
	sort.Slice(packets, func(i, j int) bool {
		return isInRightOrder(packets[i], packets[j]) == 1
	})

	for i, packet := range packets {
		if isDivider(packet) {
			product *= (i + 1)
		}
	}

	return product
}

func getPackets(input string) []Packet {
	strs := strings.Split(input, "\n\n")
	packets := []Packet{}

	for _, str := range strs {
		rawPair := strings.Split(str, "\n")
		left := parseLine([]byte(rawPair[0]))
		right := parseLine([]byte(rawPair[1]))
		packets = append(packets, left, right)
	}

	return packets
}

func main() {
	// str := utils.ReadFileIntoString("./input_test.txt")
	str := utils.ReadFileIntoString("./input.txt")
	fmt.Println(partOne(getPackets(str)))
	fmt.Println(partTwo(getPackets(str)))
}
