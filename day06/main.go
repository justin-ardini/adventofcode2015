package main

import (
	"fmt"
  "strconv"
  s "strings"
	"../util"
)

type Op int
const (
    TurnOn Op = iota
    TurnOff
    Toggle
)

type Vec2d struct {
  x int
  y int
}

type Instruction struct {
  op Op
  lo Vec2d
  hi Vec2d
}

func main() {
	path := "./day06/input.txt"
	lines, _ := util.ReadLines(path)
  instructions := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(instructions))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(instructions))
}

func parseLines(lines []string) []Instruction {
  instructions := []Instruction{}
  for _, line := range lines {
    parts := s.Split(line, " through ")
    var op Op
    if s.Contains(parts[0], "on") {
      op = TurnOn
    } else if s.Contains(parts[0], "off") {
      op = TurnOff
    } else {
      op = Toggle
    }
    sub := s.Split(parts[0], " ")
    lo := parseVec2d(sub[len(sub) - 1])
    hi := parseVec2d(parts[1])
    instructions = append(instructions, Instruction{op: op, lo: lo, hi: hi})
  }
  return instructions
}

func parseVec2d(str string) Vec2d {
  parts := s.Split(str, ",")
  x, _ := strconv.Atoi(parts[0])
  y, _ := strconv.Atoi(parts[1])
  return Vec2d{x: x, y: y}
}

func part1(instructions []Instruction) int {
  grid := [1000][1000]bool{}
  for _, i := range instructions {
    for x := i.lo.x; x <= i.hi.x; x++ {
      for y := i.lo.y; y <= i.hi.y; y++ {
        switch i.op {
        case TurnOn:
          grid[x][y] = true
        case TurnOff:
          grid[x][y] = false
        case Toggle:
          grid[x][y] = !grid[x][y]
        }
      }
    }
  }

  lit := 0
  for x := 0; x < 1000; x++ {
    for y := 0; y < 1000; y++ {
      if grid[x][y] {
        lit++
      }
    }
  }
  return lit
}

func part2(instructions []Instruction) int {
  grid := [1000][1000]int{}
  for _, i := range instructions {
    for x := i.lo.x; x <= i.hi.x; x++ {
      for y := i.lo.y; y <= i.hi.y; y++ {
        switch i.op {
        case TurnOn:
          grid[x][y]++
        case TurnOff:
          if grid[x][y] > 0 {
            grid[x][y]--
          }
        case Toggle:
          grid[x][y] += 2
        }
      }
    }
  }

  brightness := 0
  for x := 0; x < 1000; x++ {
    for y := 0; y < 1000; y++ {
      brightness += grid[x][y]
    }
  }
  return brightness
}
