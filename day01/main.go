package main

import (
	"fmt"
	"../util"
)

func main() {
	path := "./day01/input.txt"
	lines, _ := util.ReadLines(path)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(lines[0]))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(lines[0]))
}

func part1(s string) int {
  floor := 0
  for i := 0; i < len(s); i++ {
    if s[i] == '(' {
      floor += 1
    } else {
      floor -= 1
    }
  }
	return floor
}

func part2(s string) int {
  floor := 0
  for i := 0; i < len(s); i++ {
    if s[i] == '(' {
      floor += 1
    } else {
      floor -= 1
    }
    if floor == -1 {
      return i + 1
    }
  }
	return -1
}
