package main

import (
	"fmt"
	"strconv"
	"../util"
)

func main() {
	path := "./day20/input.txt"
	lines, _ := util.ReadLines(path)
  n, _ := strconv.Atoi(lines[0])
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(n))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(n))
}

func part1(n int) int {
  for i := 775000;; i = i + 2 {
	  if total(i) >= n {
      return i
    }
  }
  return -1
}

func total(n int) int {
  t := n + 1
  for i := 2; i <= n/2; i++ {
    if n % i == 0 {
      t += i
    }
  }
  return t * 10
}

func part2(n int) int {
  for i := 775000;; i = i + 2 {
	  if total2(i) >= n {
      return i
    }
  }
  return -1
}

func total2(n int) int {
  t := n
  for i := n/50; i <= n/2; i++ {
    if n % i == 0 {
      t += i
    }
  }
  return t * 11
}
