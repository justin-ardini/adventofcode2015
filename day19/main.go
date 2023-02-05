package main

import (
	"fmt"
  "sort"
	s "strings"
	"../util"
)

type Replacements map[string][]string

type Entry struct {
  m string
  distance int
}

type byLength []string

func (s byLength) Len() int {
    return len(s)
}

func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
    return len(s[i]) > len(s[j])
}

func main() {
	path := "./day19/input.txt"
	lines, _ := util.ReadLines(path)
  repl := parseLines(lines)
  molecule := lines[len(lines) - 1]
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(repl, molecule))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(repl, molecule))
}

func parseLines(lines []string) Replacements {
  repl := Replacements{}
  for _, line := range lines {
    parts := s.Split(line, " => ")
    if len(parts) < 2 {
      continue
    }
    in := parts[0]
    out := parts[1]
    outs, ok := repl[in]
    if !ok {
      outs = []string{}
    }
    repl[in] = append(outs, out)
  }
  return repl
}

func part1(repl Replacements, molecule string) int {
  results := map[string]bool{}
  for i := 0; i < len(molecule); i++ {
    s := string(molecule[i])
    for _, r := range repl[s] {
      molS := molecule[:i] + r + molecule[i + 1:]
      results[molS] = true
    }
    if i < len(molecule) - 1 {
      d := string(molecule[i]) + string(molecule[i + 1])
      for _, r := range repl[d] {
        molD := molecule[:i] + r + molecule[i + 2:]
        results[molD] = true
      }
    }
  }
	return len(results)
}

func part2(repl Replacements, molecule string) int {
  rev := map[string]string{}
  for k, arr := range repl {
    for _, v := range arr {
      rev[v] = k
    }
  }
  keys := make([]string, 0, len(rev))
  for k := range rev {
    keys = append(keys, k)
  }
  sort.Sort(byLength(keys))

  m := molecule
  steps := 0
  for {
    m = simplify(m, rev, keys)
    steps++
    if m == "e" {
      return steps
    }
  }
}

func simplify(m string, repl map[string]string, keys []string) string {
  for _, k := range keys {
    i := s.Index(m, k)
    if i != -1 {
      r := repl[k]
      return m[:i] + r + m[i + len(k):]
    }
  }
  return ""
}
