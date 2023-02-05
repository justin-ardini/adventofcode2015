package main

import (
	"fmt"
  "strconv"
	"../util"
)

func main() {
	path := "./day24/input.txt"
	lines, _ := util.ReadLines(path)
  ns := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(ns))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(ns))
}

func parseLines(lines []string) []int {
  ns := []int{}
  for _, line := range lines {
    n, _ := strconv.Atoi(line)
    ns = append(ns, n)
  }
  return ns
}

func part1(ns []int) int {
  total := 0
  for _, n := range ns {
    total += n
  }
  target := total / 3
  _, qe := combos(0, target, 0, 1, ns, 6)
	return qe
}

func part2(ns []int) int {
  total := 0
  for _, n := range ns {
    total += n
  }
  target := total / 4
  _, qe := combos(0, target, 0, 1, ns, 7)
	return qe
}

func combos(total, target, l, qe int, ns []int, ml int) (bool, int) {
  if l > ml {
    return false, 0
  }
  if total == target {
    return true, qe
  }
  if total > target || len(ns) == 0 {
    return false, 0
  }
  has, qe1 := combos(total, target, l, qe, ns[1:], ml)
  has2, qe2 := combos(total + ns[0], target, l + 1, qe * ns[0], ns[1:], ml)
  if has {
    if has2 {
      if qe1 < qe2 {
        return has, qe1
      } else {
        return has2, qe2
      }
    } else {
      return has, qe1
    }
  } else {
    return has2, qe2
  }
}

