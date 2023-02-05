package main

import (
	"fmt"
	"../util"
)

func main() {
	path := "./day03/input.txt"
	lines, _ := util.ReadLines(path)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(lines[0]))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(lines[0]))
}

func part1(moves string) int {
  visited := make(map[int]bool)
  x := 0
  y := 0
  visited[0] = true
  for i := 0; i < len(moves); i++ {
    switch moves[i] {
    case '^':
      y += 1
    case '>':
      x += 1
    case 'v':
      y -= 1
    case '<':
      x -= 1
    }
    visited[x + y * 100000] = true
  }
	return len(visited)
}

func part2(moves string) int {
  visited := make(map[int]bool)
  sx := 0
  sy := 0
  rx := 0
  ry := 0
  visited[0] = true
  for i := 0; i < len(moves); i++ {
    if i % 2 == 0 {
      switch moves[i] {
      case '^':
        sy += 1
      case '>':
        sx += 1
      case 'v':
        sy -= 1
      case '<':
        sx -= 1
      }
      visited[sx + sy * 100000] = true
    } else {
      switch moves[i] {
      case '^':
        ry += 1
      case '>':
        rx += 1
      case 'v':
        ry -= 1
      case '<':
        rx -= 1
      }
      visited[rx + ry * 100000] = true
    }
  }
	return len(visited)
}
