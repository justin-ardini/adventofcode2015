package main

import (
	"fmt"
  "strconv"
  s "strings"
	"../util"
)
type Sue map[string]int

func main() {
	path := "./day16/input.txt"
	lines, _ := util.ReadLines(path)
  sues := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(sues))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(sues))
}

func parseLines(lines []string) []Sue {
  sues := []Sue{}
  for _, line := range lines {
    parts := s.SplitAfterN(line, ": ", 2)
    subparts := s.Split(parts[1], ", ")
    sue := Sue{}
    for _, str := range subparts {
      kv := s.Split(str, ": ")
      v, _ := strconv.Atoi(kv[1])
      sue[kv[0]] = v
    }
    sues = append(sues, sue)
  }
  return sues
}

func part1(sues []Sue) int {
  expected := Sue{
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
  }
  for i, sue := range sues {
    found := true
    for k, v := range expected {
      sv, ok := sue[k]
      if ok && v != sv {
        found = false
      }
    }
    if found {
      return i + 1
    }
  }
	return -1
}

func part2(sues []Sue) int {
  expected := Sue{
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
  }
  for i, sue := range sues {
    found := true
    for k, v := range expected {
      sv, ok := sue[k]
      if k == "cats" || k == "trees" {
        if ok && v >= sv {
          found = false
        }
      } else if k == "goldfish" || k == "pomeranians" {
        if ok && v <= sv {
          found = false
        }
      } else {
        if ok && v != sv {
          found = false
        }
      }
    }
    if found {
      return i + 1
    }
  }
	return -1
}
