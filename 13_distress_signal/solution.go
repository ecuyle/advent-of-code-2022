package main

import (
	"bufio"
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

func partOne(inputPath string) {
	file := utils.Readfile(inputPath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}

func partTwo(inputPath string) {
	file := utils.Readfile(inputPath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	partOne("./input_test.txt")
	// partTwo("./input_test.txt")
}
