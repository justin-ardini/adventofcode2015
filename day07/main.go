package main

import (
	"fmt"
  "strconv"
	s "strings"
	"../util"
)

type Op int
const (
    None Op = iota
    Not
    Or
    And
    Rshift
    Lshift
)

type Instruction struct {
  op Op
  l1 string
  l2 string
  r string
}

func main() {
	path := "./day07/input.txt"
	lines, _ := util.ReadLines(path)
  instructions := parseLines(lines)
  a := run(instructions)

	fmt.Println("-- Part 1 --")
	fmt.Println(a)

  instructions[3].l1 = strconv.Itoa(a)
	fmt.Println("-- Part 2 --")
	fmt.Println(run(instructions))
}

func parseLines(lines []string) []Instruction {
  instructions := []Instruction{}
  for _, line := range lines {
    parts := s.Split(line, " -> ")
    left := s.Split(parts[0], " ")
    var op Op
    l1 := ""
    l2 := ""
    for _, l := range left {
      if l == "NOT" {
        op = Not
      } else if l == "OR" {
        op = Or
      } else if l == "AND" {
        op = And
      } else if l == "RSHIFT" {
        op = Rshift
      } else if l == "LSHIFT" {
        op = Lshift
      } else if l1 == "" {
        l1 = l
      } else {
        l2 = l
      }
    }
    r := parts[1]
    instructions = append(instructions, Instruction{op: op, l1: l1, l2: l2, r: r})
  }
  return instructions
}

func run(instructions []Instruction) int {
  wires := make(map[string]int)
  for {
    if a, ok := wires["a"]; ok {
      return a
    }

    for _, i := range instructions {
      n1, err := strconv.Atoi(i.l1)
      if err != nil {
        v, ok := wires[i.l1]
        if !ok {
          continue
        }
        n1 = v
      }
      n2 := 0
      if i.op == Or || i.op == And || i.op == Rshift || i.op == Lshift {
        n2, err = strconv.Atoi(i.l2)
        if err != nil {
          v, ok := wires[i.l2]
          if !ok {
            continue
          }
          n2 = v
        }
      }
      switch i.op {
        case None:
          wires[i.r] = n1
        case Not:
          wires[i.r] = ^n1
        case Or:
          wires[i.r] = n1 | n2
        case And:
          wires[i.r] = n1 & n2
        case Rshift:
          wires[i.r] = n1 >> n2
        case Lshift:
          wires[i.r] = n1 << n2
      }
    }
  }
}
