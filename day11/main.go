package main

import (
	"fmt"
	"../util"
)

func main() {
	path := "./day11/input.txt"
	lines, _ := util.ReadLines(path)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(lines[0]))
	fmt.Println("-- Part 2 --")
	fmt.Println(part1(part1(lines[0])))
}

func part1(pw string) string {
  for {
    pw = increment(pw)
    if (isValid(pw)) {
      return pw
    }
  }
}

func increment(pw string) string {
  bytes := []byte(pw)
  for i := len(pw) - 1; i >= 0; i-- {
    c := pw[i]
    if c != 122 {
      bytes[i] = pw[i] + 1
      break
    } else {
      bytes[i] = 97
    }
  }
  return string(bytes)
}

func isValid(pw string) bool {
  triple := false
  for i := 2; i < len(pw); i++ {
    a := pw[i - 2]
    b := pw[i - 1]
    c := pw[i]
    if (c - b == 1 && b - a == 1) {
      triple = true
      break
    }
  }
  if !triple {
    return false
  }
  for i := 0; i < len(pw); i++ {
    c := pw[i]
    if c == 'i' || c == 'o' || c == 'l' {
      return false
    }
  }
  return hasPairs(pw)
}

func hasPairs(s string) bool {
  pairs := 0
  for i := 1; i < len(s); i++ {
    c1 := s[i - 1]
    c2 := s[i]
    if c1 == c2 {
      pairs++
      i++  // No overlap
    }
  }
  return pairs >= 2
}
