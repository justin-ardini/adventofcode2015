package main

import (
	"fmt"
	"../util"
)

func main() {
	path := "./day08/input.txt"
	lines, _ := util.ReadLines(path)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(lines))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
  total := 0
  for _, line := range lines {
    total += len(line)
    inner := 0
    for i := 1; i < len(line) - 1; i++ {
      c := string(line[i])
      if c == "\\" {
        if (string(line[i + 1]) == "x") {
          i += 3 // \xab
        } else {
          i++;  // \\ or \"
        }
      }
      inner += 1
    }
    total -= inner
  }
	return total
}

func part2(lines []string) int {
  total := 0
  for _, line := range lines {
    encoded := 2
    for i := 0; i < len(line); i++ {
      c := string(line[i])
      if c == "\\" || c == "\"" {
        encoded += 2
      } else {
        encoded += 1
      }
    }
    total += encoded
    total -= len(line)
  }
	return total
}
