package main

import (
	"fmt"
  "strconv"
  s "strings"
	"../util"
)

type Graph map[string]map[string]int

type Entry struct {
  visited []string
  happiness int
}

func main() {
	path := "./day13/input.txt"
	lines, _ := util.ReadLines(path)
  graph := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(graph))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(graph))
}

func parseLines(lines []string) Graph {
  graph := Graph{}
  for _, line := range lines {
    parts := s.Split(line, " happiness units by sitting next to ")
    first := s.Split(parts[0], " ")
    units, _ := strconv.Atoi(first[3])
    if first[2] == "lose" {
      units = -units
    }
    from := first[0]
    to := s.Split(parts[1], ".")[0]

    fromMap, ok := graph[from]
    if !ok {
      fromMap = map[string]int{}
      graph[from] = fromMap
    }
    fromMap[to] = units
  }
  return graph
}

func part1(graph Graph) int {
  q := []Entry{}
  max := -1
  for k := range graph {
    q = append(q, Entry{visited: []string{k}, happiness: 0})
  }
  for ; len(q) > 0; {
    var v Entry
    v, q = q[0], q[1:]
    if len(v.visited) == len(graph) {
      f := v.visited[len(graph) - 1]
      t := v.visited[0]
      total := v.happiness + graph[f][t] + graph[t][f]
      if total > max {
        max = total
      }
    }
    from := v.visited[len(v.visited) - 1]
    for to, happiness := range graph[from] {
      if indexOf(v.visited, to) == -1 {
        toHappiness := graph[to][from]
        q = append(q, Entry{visited: append(append(v.visited[:0:0], v.visited...), to), happiness: v.happiness + happiness + toHappiness})
      }
    }
  }
  return max
}

func part2(graph Graph) int {
  graph["You"] = map[string]int{}
  for k := range graph {
    graph[k]["You"] = 0
    graph["You"][k] = 0
  }
  return part1(graph)
}

func indexOf(arr []string, v string) int {
  for i, a := range arr {
    if a == v {
      return i
    }
  }
  return -1
}
