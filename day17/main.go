package main

import (
	"fmt"
	"strconv"
	"../util"
)

func main() {
	path := "./day17/input.txt"
	lines, _ := util.ReadLines(path)
  ns := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(ns))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(ns))
}

func parseLines(lines []string) []int {
  ns := []int{}
  for _, v := range lines {
    n, _ := strconv.Atoi(v)
    ns = append(ns, n)
  }
  return ns
}

func part1(ns []int) int {
	return combos(0, ns)
}

func combos(n int, ns []int) int {
  if n == 150 {
    return 1
  }
  if n > 150 || len(ns) == 0 {
    return 0
  }
  return combos(n, ns[1:]) + combos(n + ns[0], ns[1:])
}

func part2(ns []int) int {
	return combos2(0, 0, ns)
}

func combos2(total, n int, ns []int) int {
  if total == 150  && n == 4 {
    return 1
  }
  if total > 150 || len(ns) == 0 || n > 4 {
    return 0
  }
  return combos2(total, n, ns[1:]) + combos2(total + ns[0], n + 1, ns[1:])
}
