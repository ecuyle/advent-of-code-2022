package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type monkey struct {
	items          *[]int
	itemsInspected int
	operation      func(int) int
	test           func(int) int
}

func decreaseWorryLevel(worry int) int {
	return int(worry / 3)
}

func simulateRound(monkeys *[]*monkey, shouldReduceWorry bool) *[]*monkey {
	for _, m := range *monkeys {
		for _, item := range *m.items {
			m.itemsInspected += 1
			level := m.operation(item)

			if shouldReduceWorry {
				level = decreaseWorryLevel(level)
            }

			toMonkey := (*monkeys)[m.test(level)]
			*toMonkey.items = append(*toMonkey.items, level)
		}

		*m.items = []int{}
	}
	return monkeys
}

func printMonkeys(monkeys *[]*monkey) {
	for i, monkey := range *monkeys {
		fmt.Println("Monkey", i, ":", *monkey.items, monkey.itemsInspected)
	}
}

func getMonkeyBusinessLevel(monkeys *[]*monkey, rounds int, shouldReduceWorry bool) int {
	for i := 0; i < rounds; i++ {
		monkeys = simulateRound(monkeys, shouldReduceWorry)
	}
	printMonkeys(monkeys)

	levels := []int{}

	for _, monkey := range *monkeys {
		levels = append(levels, monkey.itemsInspected)
	}

	sort.Slice(levels, func(a, b int) bool {
		return levels[a] > levels[b]
	})

	return levels[0] * levels[1]
}

func partTwo(monkeys []monkey) {
}

func getItems(str string) *[]int {
	values := str[len("Starting items:"):]
	split := strings.Split(values, ", ")
	items := []int{}

	for _, item := range split {
		converted, err := strconv.Atoi(strings.Trim(item, " "))
		utils.CheckError(err)
		items = append(items, converted)
	}

	return &items
}

func getTargetMonkey(str string) int {
	var condition string
	var target int
	_, err := fmt.Sscanf(str, "If %s throw to monkey %d", &condition, &target)
	utils.CheckError(err)

	return target
}

func makeOperation(str string) func(int) int {
	operations := map[string]func(int, int) int{
		"+": func(a int, b int) int {
			return a + b
		},
		"-": func(a int, b int) int {
			return a - b
		},
		"*": func(a int, b int) int {
			return a * b
		},
		"/": func(a int, b int) int {
			return a / b
		},
	}

	var a, b, operator string
	fmt.Sscanf(str, "Operation: new = %s %s %s", &a, &operator, &b)

	return func(value int) int {
		operand1 := value
		operand2 := value

		if a != "old" {
			operand, err := strconv.Atoi(a)
			utils.CheckError(err)
			operand1 = operand
		}

		if b != "old" {
			operand, err := strconv.Atoi(b)
			utils.CheckError(err)
			operand2 = operand
		}

		return operations[operator](operand1, operand2)
	}
}

func makeTest(conditionStr string, trueStr string, falseStr string) func(int) int {
	var divisor int
	fmt.Sscanf(conditionStr, "Test: divisible by %d", &divisor)

	return func(value int) int {
		if (value % divisor) == 0 {
			return getTargetMonkey(trueStr)
		}

		return getTargetMonkey(falseStr)
	}
}

func makeMonkey(lines []string) *monkey {
	return &monkey{
		items:          getItems(lines[1]),
		itemsInspected: 0,
		operation:      makeOperation(lines[2]),
		test:           makeTest(lines[3], lines[4], lines[5]),
	}
}

func normalizeFile(file string) *[]*monkey {
	lines := strings.Split(file, "\nMonkey")
	monkeys := []*monkey{}

	for _, line := range lines {
		monkeyData := []string{}
		parts := strings.Split(line, "\n")
		for _, part := range parts {
			monkeyData = append(monkeyData, strings.Trim(part, " "))
		}

		monkeys = append(monkeys, makeMonkey(monkeyData))
	}

	return &monkeys
}

func main() {
	file := utils.ReadFileIntoString("./input_test.txt")
	monkeys := normalizeFile(file)

	// fmt.Println(getMonkeyBusinessLevel(monkeys, 20, true))
	fmt.Println(getMonkeyBusinessLevel(monkeys, 1000, false))
}
