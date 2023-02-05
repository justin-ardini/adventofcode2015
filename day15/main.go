package main

import (
	"fmt"
	"strconv"
  s "strings"
	"../util"
)

type Ingredients map[string]Entry

type Entry struct {
  capacity int
  durability int
  flavor int
  texture int
  calories int
}

func main() {
	path := "./day15/input.txt"
	lines, _ := util.ReadLines(path)
  ingredients := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(ingredients))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(ingredients))
}

func parseLines(lines []string) Ingredients {
  ingredients := Ingredients{}
  for _, line := range lines {
    parts := s.Split(line, ": ")
    key := parts[0]
    subparts := s.Split(parts[1], ", ")
    e := Entry{}
    for _, subpart := range subparts {
      p := s.Split(subpart, " ")
      v, _ := strconv.Atoi(p[1])
      switch p[0] {
      case "capacity":
        e.capacity = v
      case "durability":
        e.durability = v
      case "flavor":
        e.flavor = v
      case "texture":
        e.texture = v
      case "calories":
        e.calories = v
      }
    }
    ingredients[key] = e
  }
  return ingredients
}

func part1(ingredients Ingredients) int {
  butter := ingredients["Butterscotch"]
  candy := ingredients["Candy"]
  frosting := ingredients["Frosting"]
  sugar := ingredients["Sugar"]
  max := 0
  for b := 0; b <= 100; b++ {
    for c := 0; c <= 100 - b; c++ {
      for f := 0; f <= 100 - b - c; f++ {
        for s := 0; s <= 100 - b - c - f; s++ {
          capacity := orZero(butter.capacity * b + candy.capacity * c + frosting.capacity * f + sugar.capacity * s)
          durability := orZero(butter.durability * b + candy.durability * c + frosting.durability * f + sugar.durability * s)
          flavor := orZero(butter.flavor * b + candy.flavor * c + frosting.flavor * f + sugar.flavor * s)
          texture := orZero(butter.texture * b + candy.texture * c + frosting.texture * f + sugar.texture * s)
          score := capacity * durability * flavor * texture
          if score > max {
            max = score
          }
        }
      }
    }
  }
	return max
}

func orZero(i int) int {
  if i < 0 {
    return 0
  }
  return i
}

func part2(ingredients Ingredients) int {
  butter := ingredients["Butterscotch"]
  candy := ingredients["Candy"]
  frosting := ingredients["Frosting"]
  sugar := ingredients["Sugar"]
  max := 0
  for b := 0; b <= 100; b++ {
    for c := 0; c <= 100 - b; c++ {
      for f := 0; f <= 100 - b - c; f++ {
        for s := 0; s <= 100 - b - c - f; s++ {
          calories := butter.calories * b + candy.calories * c + frosting.calories * f + sugar.calories * s
          if calories == 500 {
            capacity := orZero(butter.capacity * b + candy.capacity * c + frosting.capacity * f + sugar.capacity * s)
            durability := orZero(butter.durability * b + candy.durability * c + frosting.durability * f + sugar.durability * s)
            flavor := orZero(butter.flavor * b + candy.flavor * c + frosting.flavor * f + sugar.flavor * s)
            texture := orZero(butter.texture * b + candy.texture * c + frosting.texture * f + sugar.texture * s)
            score := capacity * durability * flavor * texture
            if score > max {
              max = score
            }
          }
        }
      }
    }
  }
	return max
}
