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
  distance int
}

func main() {
	path := "./day09/input.txt"
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
    parts := s.Split(line, " = ")
    distance, _ := strconv.Atoi(parts[1])
    fromTo := s.Split(parts[0], " to ")
    from := fromTo[0]
    to := fromTo[1]
    fromMap, fok := graph[from]
    if !fok {
      fromMap = map[string]int{}
      graph[from] = fromMap
    }
    toMap, tok := graph[to]
    if !tok {
      toMap = map[string]int{}
      graph[to] = toMap
    }
    fromMap[to] = distance
    toMap[from] = distance
  }
  return graph
}

func part1(graph Graph) int {
  q := []Entry{}
  min := 100000
  for k := range graph {
    q = append(q, Entry{visited: []string{k}, distance: 0})
  }
  for ; len(q) > 0; {
    var v Entry
    v, q = q[0], q[1:]
    if len(v.visited) == len(graph) {
      if v.distance < min {
        min = v.distance
      }
    }
    for to, distance := range graph[v.visited[len(v.visited) - 1]] {
      if indexOf(v.visited, to) == -1 {
        q = append(q, Entry{visited: append(append(v.visited[:0:0], v.visited...), to), distance: v.distance + distance})
      }
    }
  }
  return min
}

func part2(graph Graph) int {
  q := []Entry{}
  max := -1
  for k := range graph {
    q = append(q, Entry{visited: []string{k}, distance: 0})
  }
  for ; len(q) > 0; {
    var v Entry
    v, q = q[0], q[1:]
    if len(v.visited) == len(graph) {
      if v.distance > max {
        max = v.distance
      }
    }
    for to, distance := range graph[v.visited[len(v.visited) - 1]] {
      if indexOf(v.visited, to) == -1 {
        q = append(q, Entry{visited: append(append(v.visited[:0:0], v.visited...), to), distance: v.distance + distance})
      }
    }
  }
  return max
}

func indexOf(arr []string, v string) int {
  for i, a := range arr {
    if a == v {
      return i
    }
  }
  return -1
}
