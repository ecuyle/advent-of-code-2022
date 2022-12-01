package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  file, err := os.Open("./input.txt")
  check(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  maxCalories := 0
  calories := 0

  for scanner.Scan() {
    raw := scanner.Text()

    if raw != "" {
      value, err := strconv.Atoi(raw)
      check(err)
      calories += value
      continue
    }

    if calories > maxCalories {
      maxCalories = calories
    }

    calories = 0
  }

  if calories > maxCalories {
    maxCalories = calories
  }

  fmt.Println("Max calories carried by an elf is", maxCalories)

  file, err = os.Open("./input.txt")
  check(err)
  defer file.Close()

  scanner = bufio.NewScanner(file)
  caloricHistory := []int{0}
  calories = 0

  for scanner.Scan() {
    raw := scanner.Text()

    if raw != "" {
      value, err := strconv.Atoi(raw)
      check(err)
      calories += value
      continue
    }

    caloricHistory = append(caloricHistory, calories) 
    calories = 0
  }

  sort.Ints(caloricHistory)
  totalCalories := 0

  for i := 0; i < 3; i++ {
    if len(caloricHistory) <= i {
      break
    }

    totalCalories += caloricHistory[len(caloricHistory) - i - 1]
  }

  fmt.Println("Total calories are:", totalCalories)
}
