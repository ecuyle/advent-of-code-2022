package main

import (
  "bufio"
  "fmt"
  "os"
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
}
