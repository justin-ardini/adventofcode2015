package main

import (
	"fmt"
	"../util"
)

func main() {
	path := "./day10/input.txt"
	lines, _ := util.ReadLines(path)
  n := []rune(lines[0])
  for i := 0; i < len(n); i++ {
    n[i] -= 48
  }
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(n))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(n))
}

func part1(n []rune) int {
  for i := 0; i < 40; i++ {
    n = step(n)
  }
  return len(n)
}

func step(n []rune) []rune {
  var prev rune
  var count int8
  count = 0
  res := []rune{}
  for i := 0; i < len(n); i++ {
    c := n[i]
    if count > 0 && c == prev {
      count++
    } else {
      if count > 0 {
        res = append(res, rune(count), prev)
      }
      prev = c
      count = 1
    }
  }
  res = append(res, rune(count), prev)
  return res
}

func part2(n []rune) int {
  for i := 0; i < 50; i++ {
    n = step(n)
  }
  return len(n)
}
