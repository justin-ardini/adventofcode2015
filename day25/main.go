package main

import (
	"fmt"
)

func main() {
  row := 2981
  col := 3075
  start := 20151125
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(start, row, col))
	fmt.Println("-- Part 2 --")
	fmt.Println("Merry Christmas!!")
}

func part1(start, row, col int) int {
  n := start
  r := 1
  c := 1
  for {
    n = next(n)
    c += 1
    r -= 1
    if r == 0 {
      r = c
      c = 1
    }
    if r == row && c == col {
      return n
    }
  }
}

func next(prev int) int {
  return (prev * 252533) % 33554393
}
