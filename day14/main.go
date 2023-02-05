package main

import (
	"fmt"
	"strconv"
  s "strings"
	"../util"
)

type Reindeer map[string]Entry

type Entry struct {
  speed int
  duration int
  rest int
}

func main() {
	path := "./day14/input.txt"
	lines, _ := util.ReadLines(path)
  reindeer := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(reindeer))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(reindeer))
}

func parseLines(lines []string) Reindeer {
  reindeer := Reindeer{}
  for _, line := range lines {
    parts := s.Split(line, " seconds, but then must rest for ")
    first := s.Split(parts[0], " ")
    second := s.Split(parts[1], " ")
    key := first[0]
    speed, _ := strconv.Atoi(first[3])
    duration, _ := strconv.Atoi(first[6])
    rest, _ := strconv.Atoi(second[0])
    reindeer[key] = Entry{speed: speed, duration: duration, rest: rest}
  }
  return reindeer
}

func part1(reindeer Reindeer) int {
  max := 0
  end := 2503
  for _, v := range reindeer {
    total := distance(v, end)
    if total > max {
      max = total
    }
  }
  return max
}

func distance(v Entry, end int) int {
  c := v.duration + v.rest
  n := end / c
  r := end - c * n
  if r > v.duration {
    r = v.duration
  }
  return v.speed * v.duration * n + v.speed * r
}

func part2(reindeer Reindeer) int {
  scores := map[string]int{}
  for k := range reindeer {
    scores[k] = 0
  }
  for i := 1; i <= 2503; i++ {
    step(reindeer, scores, i)
  }
  max := 0
  for _, v := range scores {
    if v > max {
      max = v
    }
  }
  return max
}

func step(reindeer Reindeer, scores map[string]int, i int) {
  var kMax []string
  max := 0
  for k, v := range reindeer {
    total := distance(v, i)
    if total > max {
      max = total
      kMax = []string{k}
    } else if total == max {
      kMax = append(kMax, k)
    }
  }
  for _, k := range kMax {
    scores[k]++
  }
}
