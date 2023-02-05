package main

import (
	"fmt"
	"strconv"
  s "strings"
	"../util"
)

func main() {
	path := "./day02/input.txt"
	lines, _ := util.ReadLines(path)
  dims := toDims(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(dims))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(dims))
}

func toDims(lines []string) [][]int {
  var allDims [][]int
  for _, line := range lines {
    var dims []int
    for _, sn := range s.Split(line, "x") {
		  n, _ := strconv.Atoi(sn)
      dims = append(dims, n)
    }
    allDims = append(allDims, dims)
  }
  return allDims
}

func part1(dims [][]int) int {
  total := 0
  for _, p := range dims {
    a := p[0] * p[1]
    b := p[1] * p[2]
    c := p[2] * p[0]
    total += 2 * a + 2 * b + 2 * c + util.MinOf(a, b, c)
  }
	return total
}

func part2(dims [][]int) int {
  total := 0
  for _, p := range dims {
    a := p[0] + p[1]
    b := p[1] + p[2]
    c := p[2] + p[0]
    total += 2 * util.MinOf(a, b, c) + p[0] * p[1] * p[2]
  }
	return total
}
