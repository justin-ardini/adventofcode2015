package main

import (
	"fmt"
  "strconv"
  s "strings"
	"../util"
)

type Instruction struct {
  name string
  reg string
  offset int
}

func main() {
	path := "./day23/input.txt"
	lines, _ := util.ReadLines(path)
  instrs := parseLines(lines)
  reg := map[string]int{"a": 0, "b": 0}
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(instrs, reg))
	fmt.Println("-- Part 2 --")
  reg = map[string]int{"a": 1, "b": 0}
	fmt.Println(part1(instrs, reg))
}

func parseLines(lines []string) []Instruction {
  instrs := []Instruction{}
  for _, line := range lines {
    name := line[:3]
    reg := line[4:5]
    offset := 0
    if name == "jmp" {
      parts := s.Split(line, " ")
      n, _ := strconv.Atoi(parts[1])
      offset = n
      reg = ""
    } else if name == "jio" || name == "jie" {
      parts := s.Split(line, ", ")
      n, _ := strconv.Atoi(parts[1])
      offset = n
    }
    instrs = append(instrs, Instruction{name: name, reg: reg, offset: offset})
  }
  return instrs
}

func part1(instrs []Instruction, reg map[string]int) int {
  for i := 0;; {
    if i >= len(instrs) {
      return reg["b"]
    }
    instr := instrs[i]
    switch instr.name {
    case "hlf":
      reg[instr.reg] /= 2
      i++
    case "tpl":
      reg[instr.reg] *= 3
      i++
    case "inc":
      reg[instr.reg] += 1
      i++
    case "jmp":
      i += instr.offset
    case "jie":
      if reg[instr.reg] % 2 == 0 {
        i += instr.offset
      } else {
        i++
      }
    case "jio":
      if reg[instr.reg] == 1 {
        i += instr.offset
      } else {
        i++
      }
    }
  }
}
