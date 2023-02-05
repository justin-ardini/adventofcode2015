package main

import (
	"fmt"
	"../util"
)

func main() {
	path := "./day05/input.txt"
	lines, _ := util.ReadLines(path)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(lines))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
  n := 0
  for _, l := range lines {
    if isNice(l) {
      n += 1
    }
  }
  return n
}

func isNice(s string) bool {
  vowels := 0
  pair := false
  p := ""
  for _, b := range s {
    c := string(b)
    if p == c {
      pair = true
    }
    if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
      vowels++
    }
    if p == "a" && c == "b" || p == "c" && c == "d" || p == "p" && c == "q" || p == "x" && c == "y" {
      return false
    }
    p = c
  }
  return pair && vowels >= 3
}

func part2(lines []string) int {
  n := 0
  for _, l := range lines {
    if isNice2(l) {
      n += 1
    }
  }
  return n
}

func isNice2(s string) bool {
  pairs := make(map[string]bool)
  pair := false
  repeat := false
  p1 := ""
  p2 := ""
  for i := 0; i < len(s); i += 2 {
    c1 := string(s[i])
    c2 := string(s[i + 1])
    if p1 == c1 || p2 == c2 {
      repeat = true
    }
    if pairs[c1 + c2] {
      pair = true
    }
    if p2 != "" {
      if p1 + p2 != p2 + c1 && pairs[p2 + c1] {
        pair = true
      }
      pairs[p2 + c1] = true
    }
    pairs[c1 + c2] = true
    p1 = c1
    p2 = c2
  }
  return pair && repeat
}
